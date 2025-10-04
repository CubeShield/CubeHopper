package config

import (
	"fmt"
	"os"

	"github.com/CubeShield/CubeHopper/internal/types"
	"github.com/spf13/viper"
)

type ConfigManager struct {
	v *viper.Viper
	c *Config
}

func NewConfigManager() (*ConfigManager, error) {
	v := viper.New()

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить текущую рабочую категорию: %w", err)
	}

	v.SetDefault("minecraft_path", currentDir)
	v.SetDefault("api_base_url", "http://api.cubeshield.ru:8000/api/v1")
	v.SetDefault("installed_containers", make([]types.Container, 0))

	v.SetConfigName("CubeHopperConfig")
	v.SetConfigType("json")
	v.AddConfigPath(".")


	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if writeErr := v.SafeWriteConfigAs("./CubeHopperConfig.json"); writeErr != nil {
				return nil, fmt.Errorf("не удалось создать стандартный конфигурационный файл: %w", writeErr)
			}
		} else {
			return nil, fmt.Errorf("не удалось прочитать конфигурационный файл: %w", err)
		}
	}

	cm := &ConfigManager{v: v}
	if err := cm.load(); err != nil {
		return nil, err
	}
	return cm, nil

}

func (cm *ConfigManager) load() error {
	var c Config
	if err := cm.v.Unmarshal(&c); err != nil {
		return fmt.Errorf("ошибка в обработке конфигурационного файла: %w", err)
	}
	cm.c = &c
	return nil
}

func (cm *ConfigManager) Save() error {
	cm.v.Set("minecraft_path", cm.c.MinecraftPath)
	cm.v.Set("api_base_url", cm.c.ApiBaseUrl)
	cm.v.Set("installed_containers", cm.c.InstalledContainers)

	if err := cm.v.WriteConfig(); err != nil {
		return fmt.Errorf("не удалось сохранить изменения в конфигурационный файл: %w", err)
	}

	return nil
}

func (cm *ConfigManager) Config() *Config {
	return cm.c
}