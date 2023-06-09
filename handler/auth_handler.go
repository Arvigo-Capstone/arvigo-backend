package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterAuthRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/login", loginHandler)
	authGroup.POST("/register-user", registerUserHandler)
	authGroup.POST("/update-user/:id", updateUserHandler)
	authGroup.POST("/register-partner", registerPartnerHandler)
}

func loginHandler(c echo.Context) error {
	var loginData datastruct.LoginUserInput
	if err := c.Bind(&loginData); err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	validationErrors := utils.ValidateStruct(loginData)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	token, statusCode, err := repository.Login(loginData)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Login success", token, http.StatusOK)
}

func updateUserHandler(c echo.Context) error {
	uID := utils.StrToUint64(c.Param("id"), 0)
	if uID == 0 {
		return utils.ResponseJSON(c, "Invalid user ID", nil, http.StatusBadRequest)
	}

	var userData datastruct.UserRegisterInput
	err := c.Bind(&userData)
	if err != nil {
		return err
	}

	validationErrors := utils.ValidateStruct(userData)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	statusCode, err := repository.UpdateUser(userData, uID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Updated", nil, statusCode)
}

func registerUserHandler(c echo.Context) error {
	var userData datastruct.UserRegisterInput
	err := c.Bind(&userData)
	if err != nil {
		return err
	}

	validationErrors := utils.ValidateStruct(userData)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	user, statusCode, err := repository.RegisterUser(userData)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Created", user, statusCode)
}

func registerPartnerHandler(c echo.Context) error {
	var userData datastruct.PartnerRegisterInput
	err := c.Bind(&userData)
	if err != nil {
		return err
	}

	validationErrors := utils.ValidateStruct(userData)
	if len(validationErrors) > 0 {
		return utils.ResponseJSON(c, "The data is not valid", validationErrors, http.StatusBadRequest)
	}

	user, statusCode, err := repository.RegisterPartner(userData)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), nil, statusCode)
	}

	return utils.ResponseJSON(c, "Created", user, statusCode)
}
