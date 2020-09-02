package server

import (
	"fmt"
	"trouble-shooting/config"
	"trouble-shooting/router"
	"trouble-shooting/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func Start() {
	cfg := config.NewConfig()
	ginServer := router.InitRouter()
	//ginpprof.Wrapper(ginServer)
	utils.Logger.WithFields(logrus.Fields{
		"env": os.Getenv("env"),
	}).Info("Current ENV")
	utils.Logger.WithFields(logrus.Fields{
		"port": cfg.HTTP.Port,
	}).Info("Start gin-admin on Port")
	utils.Logger.WithFields(logrus.Fields{
		"logPath": cfg.Log.LogPath,
		"logName": cfg.Log.LogName,
	}).Info("log info")
	//err := ginServer.RunTLS(":"+ string(cfg.HTTP.Port), cfg.HTTP.Certificate, cfg.HTTP.CertificateKey)
	err := ginServer.Run(cfg.HTTP.Host + ":" + fmt.Sprintf("%d", cfg.HTTP.Port))

	if err != nil {
		utils.Logger.Fatalf("Gin  Start err: %s", err.Error())
	}
}
