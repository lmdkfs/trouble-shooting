package api

import (
	"trouble-shooting/config"
	"trouble-shooting/server"
	"trouble-shooting/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	cfgFile  string
	cfg      = config.NewConfig()
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "start server",
		Example:      "gin-admin server -c cfg.yaml",
		SilenceUsage: true,
		//PreRun: func(cmd *cobra.Command, args []string) {
		//	setup()
		//},
		Run: func(cmd *cobra.Command, args []string) {
			os.Setenv("TESTHTTP", "0")
			server.Start()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	StartCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cfg.yaml)")
	StartCmd.PersistentFlags().StringP("port", "p", "8888", "listen port")
	StartCmd.PersistentFlags().StringP("host", "s", "0.0.0.0", "listen host")
	StartCmd.PersistentFlags().StringP("logname", "l", "gin_admin.log", "logname")
	StartCmd.PersistentFlags().String("logpath", "", "logpath")
	viper.BindPFlag("server.port", StartCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("server.host", StartCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("server.logpath", StartCmd.PersistentFlags().Lookup("logpath"))
	viper.BindPFlag("server.logpname", StartCmd.PersistentFlags().Lookup("logpname"))

}

func initConfig() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Panicf("Get currrent dir Fail: %s", err.Error())
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Printf("Get $HOME Dir error: %s", err.Error())
			os.Exit(1)
		}
		viper.AddConfigPath(currentDir)
		viper.AddConfigPath(home)
		viper.SetConfigName(".cfg")

	}
	viper.SetEnvPrefix("ginadmin")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8888")
	viper.SetDefault("server.logpath", "/var/log/")
	viper.SetDefault("server.logname", "gin_admin.log")
	cfg.RunMode = viper.GetString("server.run_mode")
	cfg.HTTP.Host = viper.GetString("server.host")
	cfg.HTTP.Port = viper.GetInt("server.port")

	cfg.Log.LogName = viper.GetString("server.logname")
	cfg.Log.LogPath = viper.GetString("server.Logpath")

	exist, err := PathExists(cfg.Log.LogPath)
	//logrus.Printf("日志路径: %s", cfg.Log.LogPath)
	if err != nil {
		logrus.Errorf("Get dir error ![%v]\n", err)
	}
	if !exist {
		err := os.MkdirAll(cfg.Log.LogPath, 0755)
		if err != nil {
			logrus.Errorf("MkDirs Failed![%v]\n", err)
		} else {
			logrus.Info("MkDirs Success!\n")
		}
	}
	utils.NewLogger()

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
