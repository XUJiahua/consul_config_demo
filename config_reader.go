package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type ConfigReader interface {
	ReadConfig() error
}

type localConfigReader struct {
}

func NewLocalConfigReader() ConfigReader {
	return &localConfigReader{}
}

func (localConfigReader) ReadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "unable to read config")
	}

	showConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		showConfig()

		logrus.Info("app reload...")
	})

	return nil
}

type remoteConfigReader struct {
	Provider string
	Endpoint string
	Key string
}

func NewRemoteConfigReader(provider, endpoint, key string) ConfigReader{
	return &remoteConfigReader{
		Provider:provider,
		Endpoint:endpoint,
		Key:key,
	}
}

func (rcr remoteConfigReader) ReadConfig() error {
	err := viper.AddRemoteProvider(rcr.Provider, rcr.Endpoint, rcr.Key)
	if err != nil {
		return errors.Wrap(err, "unable to add remote provider")
	}

	viper.SetConfigType("json") // Need to explicitly set this to json

	err = viper.ReadRemoteConfig()
	if err != nil {
		return errors.Wrap(err, "unable to read remote config")
	}

	err = viper.WatchRemoteConfigOnChannel()
	if err != nil {
		return errors.Wrap(err, "unable to watch")
	}

	viper.OnRemoteConfigChange(func() {
		showConfig()
		logrus.Info("app start/reload...")
	})

	return nil
}

func showConfig() {
	hostname := viper.GetString("hostname")
	port := viper.GetInt("port")
	fmt.Println(hostname, port)
}
