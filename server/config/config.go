package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

type HttpServer struct {
	Addr string `yaml:"addr"`
}

type Mysql struct {
	HostName string `yaml:"host_name"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	DataBase string `yaml:"data_base"`
}

type Log struct {
	Level   string `yaml:"level"`
	Driver  string `yaml:"driver"`
	Path    string `yaml:"path"`
	MaxFile int    `yaml:"max_file"`
	MaxSize int    `yaml:"max_size"`
}

type ServerConfig struct {
	HttpServerConf *HttpServer     `yaml:"http_server"`
	LogConf        *logger.LogOpts `yaml:"log"`
	MysqlConf      *Mysql          `yaml:"mysql"`
}

var globalConfig *ServerConfig

func ConfigInit(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("open config file %s failed! err = %s\n", path, err.Error())
		return err
	}

	err = yaml.Unmarshal(bytes, globalConfig)
	if err != nil {
		fmt.Printf("config file Unmarshal %s failed!\n", string(bytes))
		return err
	}
	return nil
}

func Config() *ServerConfig {

	return globalConfig
}
