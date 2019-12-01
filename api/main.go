package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"api/routes"
	"api/utils"

	"github.com/go-ini/ini"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	logger := initLogger()

	cfg, err := ini.Load("config/site.ini")

	utils.InitBoot()

	handleError(err, logger)

	db, err := utils.InitDb(cfg, logger)
	handleError(err, logger)

	keys, _ := cfg.GetSection("SERVER")

	origins := keys.Key("ORIGINS").Strings(",")

	e := echo.New()
	e.Static("/public", "public")
	e.Static("/static", "static")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	e.Use(middleware.Recover())

	e.Use(session.Middleware(sessions.NewFilesystemStore("./tmp", []byte("!ilovejesus123!"))))
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}

	startHandlers(utils.NewEnviroment(e, db, cfg, logger))

}

// handleError ...
func handleError(err error, log *zap.Logger) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func startHandlers(env *utils.Enviroment) {
	routes := []utils.Ihandler{
		&routes.Authentication{},
		&routes.Register{},
	}

	for _, rt := range routes {
		rt.Init("/api", env)
	}

	env.Log.Panic(env.Echo.Start(":4000").Error())
}

func initLogger() (logger *zap.Logger) {

	f, err := ioutil.ReadFile("config/zap.json")
	if err != nil {
		log.Fatal("could not get zap config file")
	}

	var cfg zap.Config
	if err := json.Unmarshal(f, &cfg); err != nil {
		log.Panicln(err)
	}

	logger, err = cfg.Build()
	if err != nil {
		log.Panicln(err)
	}

	defer logger.Sync()

	return
}
