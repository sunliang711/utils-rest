package main

import (
	"os"
	"os/signal"
	"syscall"

	"nft-studio-backend/config"
	"nft-studio-backend/server"
	"nft-studio-backend/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	// parse config file
	config.Init()

	utils.Init(123)

	srv := server.StartHttpServerWithConfig()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	signal.Notify(stopCh, syscall.SIGINT)

	<-stopCh
	logrus.Infof("got stop signal")
	srv.Stop()
}
