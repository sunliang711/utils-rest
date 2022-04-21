package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"nft-studio-backend/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const AppName = "nft"

func Init() {
	logrus.Infof("config init()...")
	configPath := pflag.StringP("config", "c", "", "config file")
	pflag.Parse()

	if *configPath != "" {
		viper.SetConfigFile(*configPath)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(fmt.Sprintf("/etc/%s", AppName))
		viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", AppName))
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Read config file: %v error: %v", *configPath, err)
	}
	logrus.Infof("using config file: %s", viper.ConfigFileUsed())

	viper.AutomaticEnv()
	viper.SetEnvPrefix(AppName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: viper.GetBool("log.show_fulltime")})
	if viper.GetBool("log.report_caller") {
		logrus.Info("logrus: enable report caller")
		logrus.SetReportCaller(true)
	}
	loglevel := viper.GetString("log.level")
	logrus.Infoln("Log level: ", loglevel)
	logrus.SetLevel(utils.LogLevel(loglevel))

	var output io.Writer
	logfilePath := viper.GetString("log.logfile")
	if logfilePath != "" {
		handler, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logrus.Fatalf("Open logfile: %v error: %v", logfilePath, err)
		}
		logrus.Infof("Logfile path: %v", logfilePath)
		output = handler
	} else {
		output = os.Stderr
	}
	logrus.SetOutput(output)
}
