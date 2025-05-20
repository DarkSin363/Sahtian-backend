package auth

import (
	"context"
	in_errors "github.com/BigDwarf/sahtian/internal/errors"
	"github.com/labstack/echo/v4"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"strings"
	"time"

	"github.com/BigDwarf/sahtian/internal/model"
)

type Repository interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

type Service struct {
	rep         Repository
	authToken   string
	debug       bool
	debugUserId int64
}

func NewService(rep Repository, token string, debug bool, debugUserId int64) *Service {
	return &Service{rep: rep, authToken: token, debug: debug, debugUserId: debugUserId}
}

func (s *Service) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "swagger") {
			return next(c)
		}

		if s.debug {
			user := &model.User{
				ID: s.debugUserId,
			}

			c.Set("user", user)
			c.Set("userId", user.ID)
			c.Set("startParam", "")
			return next(c)
		}

		initData := c.Request().Header.Get("X-sahtian-Init-Data")
		err := initdata.Validate(initData, s.authToken, 24*time.Hour)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, in_errors.ErrUnauthorized)
		}

		data, err := initdata.Parse(initData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		user := &model.User{
			AddedToAttachmentMenu: data.User.AddedToAttachmentMenu,
			AllowsWriteToPm:       data.User.AllowsWriteToPm,
			FirstName:             data.User.FirstName,
			ID:                    data.User.ID,
			IsBot:                 data.User.IsBot,
			IsPremium:             data.User.IsPremium,
			LastName:              data.User.LastName,
			Username:              data.User.Username,
			LanguageCode:          data.User.LanguageCode,
		}

		c.Set("user", user)
		c.Set("userId", user.ID)
		c.Set("startParam", data.StartParam)

		return next(c)
	}
}

func (s *Service) AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().URL.Path, "/api/v1/admin") {
			id := c.Get("userId").(int64)

			user, err := s.rep.GetUser(c.Request().Context(), id)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, in_errors.ErrUnauthorized)
			}

			if !user.IsAdmin {
				return c.JSON(http.StatusUnauthorized, in_errors.ErrUnauthorized)
			}
		}
		return next(c)
	}
}
