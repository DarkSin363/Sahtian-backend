package app

import (
	"context"
	"crypto/rand"
	"net/http"

	"github.com/brpaz/echozap"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	_ "github.com/BigDwarf/sahtian/docs"
	"github.com/BigDwarf/sahtian/internal/config"
	"github.com/BigDwarf/sahtian/internal/log"
	"github.com/BigDwarf/sahtian/internal/repository"
	"github.com/BigDwarf/sahtian/internal/service/auth"
	"github.com/BigDwarf/sahtian/internal/service/telegram"
	"github.com/BigDwarf/sahtian/internal/service/users"
)

type ServerApplication struct {
	conf             config.Config
	srv              *echo.Echo
	usersService     *users.Service
	authService      *auth.Service
	telegramService  *telegram.Service
	dbClient         *mongo.Client
	usersRepository  *repository.UsersRepository
	checkTaskService *telegram.Service
}

func NewServerApplication(cfg *config.Config) *ServerApplication {
	app := &ServerApplication{
		conf: *cfg,
	}

	srv := echo.New()
	if cfg.EnableDocs {
		srv.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: app.conf.Server.CorsConfig.AllowOrigins,
		AllowMethods: app.conf.Server.CorsConfig.AllowMethods,
	}))
	srv.Use(echozap.ZapLogger(log.Logger()))
	srv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := make([]byte, 32)
			_, _ = rand.Read(token)
			requestId := uuid.NewSHA1(uuid.NameSpaceURL, token)
			c.Set("requestId", requestId.String())

			return next(c)
		}
	})
	srv.Use(middleware.Recover())
	srv.Use(app.AuthService().AuthMiddleware)

	app.srv = srv

	return app
}

func (app *ServerApplication) Run(_ context.Context) error {
	log.Info("sahtian Server RUN begin ... ")

	go app.TelegramService().Start()
	app.registerRoutes(app.srv)

	go func() {
		log.Sugar().Infof("http server listen : http://%s ", app.conf.Server.Listen)
		defer log.Sugar().Infof("http server closed : http://%s ", app.conf.Server.Listen)

		if err := app.srv.Start(app.conf.Server.Listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("http server closed", zap.Error(err))
		}
	}()

	log.Info("sahtian Server RUN finish ... ")

	return nil
}

func (app *ServerApplication) Shutdown() {
	log.Info("sahtian Server stopping...")

	ctx := context.Background()

	app.TelegramService().Stop()

	if err := app.srv.Shutdown(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error(err.Error())
		}
	}

	if err := app.dbClient.Disconnect(ctx); err != nil {
		log.Error(err.Error())
	}

	log.Info("sahtian Server stopped...")
}
