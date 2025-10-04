package main

import (
	"github.com/CubeShield/CubeHopper/cmd"
	"github.com/CubeShield/CubeHopper/internal/api"
	"github.com/CubeShield/CubeHopper/internal/config"
)

func main() {
	configManager, _ := config.NewConfigManager()

	apiClient := api.NewApiClient(configManager.Config().ApiBaseUrl)

	updater := cmd.NewUpdater(apiClient, configManager)
	updater.RunUpdate()
}