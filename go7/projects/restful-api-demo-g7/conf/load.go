package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 如何把配置映射成Config对象

// 从 Toml 格式的配置文件加载配置
func LoadConfigFromToml(filepath string) error {
	config = NewDefaultConfig()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return fmt.Errorf("load config from file error %s  ,path: %s", err, filepath)
	}

	// 加载db的全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return err
	}

	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}

// 加载全局实例
func LoadGlobal() (err error) {
	// 加载db全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return
	}
	return
}
