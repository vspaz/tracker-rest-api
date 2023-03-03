package app

import (
	"github.com/vspaz/simplelogger/pkg/logging"
	"github.com/vspaz/tracker-rest-api/config"
	"github.com/vspaz/tracker-rest-api/web/handlers"
)

func Run() {
	conf := config.GetConfig().Config
	logger := logging.GetTextLogger(conf.Logging.Level).Logger
	router := handlers.NewRouter(conf, logger)
	router.ConfigureMiddleware()
	router.RegisterHandlers()
	router.StartServer()
}
