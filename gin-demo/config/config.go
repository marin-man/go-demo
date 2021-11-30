package config

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
)

func Init(output io.Writer, configFile string) error {
	if output == nil {
		output = ioutil.Discard
	}
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if nil != err {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		_, _ = fmt.Fprintf(output, "Config file change %s \n", e.Name)
	})
	return nil
}

func MustInit(output io.Writer, conf string) {
	if err := Init(output, conf); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
