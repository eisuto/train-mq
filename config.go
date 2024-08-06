package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// Config 结构体表示服务器的配置项。
type Config struct {
	Port   int  `yaml:"port"`
	Banner bool `yaml:"banner"`
}

// DefaultConfig 提供默认的配置。
var DefaultConfig = Config{
	Port:   5757,
	Banner: true,
}

// LoadConfig 读取配置文件并返回 Config 结构体。
// 如果配置文件不存在，则创建一个默认配置文件。
func LoadConfig(filename string) (Config, error) {
	var config Config

	// 尝试读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// 配置文件不存在，创建默认配置文件
			data, _ = yaml.Marshal(DefaultConfig)
			err = os.WriteFile(filename, data, 0644)
			if err != nil {
				return config, fmt.Errorf("无法创建默认配置文件: %v", err)
			}
			return DefaultConfig, nil
		}
		return config, fmt.Errorf("无法读取配置文件: %v", err)
	}

	// 解析配置文件
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("配置文件格式错误: %v", err)
	}

	return config, nil
}
