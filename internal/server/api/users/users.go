package users

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/BigDwarf/sahtian/internal/server/api"

	in_errors "github.com/BigDwarf/sahtian/internal/errors"
	"github.com/BigDwarf/sahtian/internal/model"
)

type GetUserRequest struct {
	UserId int64 `json:"user_id"`
}

type GetUserResponse struct {
	Id                int64    `json:"id"`
	FirstName         string   `json:"first_name"`
	AvatarUrl         string   `json:"avatar_url"`
	LastName          string   `json:"last_name"`
	Username          string   `json:"username"`
	DisplayName       string   `json:"display_name"`
	TelegramUsername  string   `json:"telegram_username"`
	InstagramUsername string   `json:"instagram_username"`
	XUsername         string   `json:"x_username"`
	Website           string   `json:"website"`
	Status            string   `json:"status"`
	Totalsahtians     int64    `json:"total_sahtians"`
	Todaysahtians     int64    `json:"today_sahtians"`
	Monthlysahtians   int64    `json:"monthly_sahtians"`
	Rank              int64    `json:"rank"`
	Features          Features `json:"features"`
}

type Features struct {
}

// handlerGetUser		 godoc
// @Summary      Get user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body body GetUserRequest true  "GetUserRequest"
// @Success      200  {object}  GetUserResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      404  {object} 	api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/getUser [post]
func (ctl *Controller) handlerGetUser(e echo.Context) error {
	var req GetUserRequest

	err := json.NewDecoder(e.Request().Body).Decode(&req)
	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	var responseUser *model.User
	id := e.Get("userId").(int64)

	responseUser, err = ctl.service.GetExistingUser(e.Request().Context(), id, req.UserId)
	if err != nil {
		if errors.Is(err, in_errors.ErrUserNotFound) {
			return e.JSON(http.StatusNotFound, &api.ErrorResponse{Error: err.Error()})
		}
		return e.JSON(http.StatusInternalServerError, &api.ErrorResponse{Error: err.Error()})
	}

	res := GetUserResponse{
		Id:          responseUser.ID,
		FirstName:   responseUser.FirstName,
		LastName:    responseUser.LastName,
		Username:    responseUser.Username,
		DisplayName: responseUser.Username,
		AvatarUrl:   responseUser.AvatarURL,
	}

	return e.JSON(http.StatusOK, res)
}

// handlerInit		 godoc
// @Summary      Initialize user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        referral_id   query     int    false  "Referral ID"
// @Param        utm_source    query     string false  "UTM Source"
// @Success      200  {object}  model.User
// @Failure      400  {object}  api.ErrorResponse
// @Failure      404  {object} 	api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/init [post]
func (ctl *Controller) handlerInit(e echo.Context) error {

	return e.JSON(http.StatusOK, nil)
}
