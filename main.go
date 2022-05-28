package main

import (
	"github.com/isther/backendTemplate/conf"
	"github.com/isther/backendTemplate/internal/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := routers.Init()
	logrus.Info("Server listen: ", conf.Server.Listen)
	r.Run(conf.Server.Listen)
}
