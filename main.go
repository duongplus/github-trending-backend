package main

import (
	"demo-full-project/db"
	"demo-full-project/handler"
	"demo-full-project/helper"
	"demo-full-project/log"
	"demo-full-project/repository/repo_impl"
	"demo-full-project/router"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "renduong1",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.GET("/", welcome)
	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my app!")
}
