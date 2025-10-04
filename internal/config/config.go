package config

import "github.com/CubeShield/CubeHopper/internal/types"


type Config struct {
	MinecraftPath string `mapstructure:"minecraft_path"`
	ApiBaseUrl string `mapstructure:"api_base_url"`
	InstalledContainers []types.Container `mapstructure:"installed_containers"`
}
