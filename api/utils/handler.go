package utils

import (
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// Dependencies required by the Handlers

type Enviroment struct {
	Echo  *echo.Echo
	Paths map[string]string
	DB    *gorm.DB
	Cfg   *ini.File
	Log   *zap.Logger
	Url   map[string]string
}

// New Enviroment
func NewEnviroment(echo *echo.Echo, db *gorm.DB, cfg *ini.File, log *zap.Logger) *Enviroment {
	return &Enviroment{
		Echo:  echo,
		Paths: GetPath(),
		Url:   BaseURL(),
		DB:    db,
		Cfg:   cfg,
		Log:   log,
	}
}

// Interface all Echo handlers must satisfy
type Ihandler interface {
	Init(prefix string, env *Enviroment) error
}
