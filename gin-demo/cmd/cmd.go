package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"helloworld/config"
	"helloworld/db"
	"helloworld/router"
	"log"
	"net/http"
	"os"
)

var (
	cfgFile string
	logger = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	// 配置初始化
	config.MustInit(os.Stdout, cfgFile)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启 debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := db.Mysql(
				viper.GetString("db.hostname"),
				viper.GetInt("db.port"),
				viper.GetString("db.username"),
				viper.GetString("db.password"),
				viper.GetString("db.dbname"),
			)
		if nil != err {
			return err
		}

		// 自动建表
		db.DB.AutoMigrate()

		defer db.DB.Close()

		r := router.SetupRouter()
		r.Run()

		port := viper.GetString("port")
		log.Println("port=", port)
		return http.ListenAndServe(port, nil)
	}
	return rootCmd.Execute()
}
