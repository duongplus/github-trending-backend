package router

import (
	"demo-full-project/handler"
	"demo-full-project/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.HandleProfile)
	user.PUT("/profile/update", api.UserHandler.UpdateUserHandler)
}
