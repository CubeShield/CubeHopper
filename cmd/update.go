package cmd

import (
	"fmt"

	"github.com/CubeShield/CubeHopper/internal/api"
	"github.com/CubeShield/CubeHopper/internal/config"
	"github.com/CubeShield/CubeHopper/internal/filesystem"
	"github.com/CubeShield/CubeHopper/internal/types"
	"github.com/CubeShield/CubeHopper/internal/updater"
)

type Updater struct {
	apiClient *api.ApiClient
	cm *config.ConfigManager
	fm *filesystem.FileManager
}

func NewUpdater(apiClient *api.ApiClient, cm *config.ConfigManager) *Updater {
	return &Updater{
		apiClient: apiClient,
		cm: cm,
		fm: &filesystem.FileManager{BasePath: cm.Config().MinecraftPath},
	}
}

func (u *Updater) RunUpdate() error {
	fmt.Println("Последняя версия")
	latestInstance, err := u.apiClient.GetInstance()
	if err != nil {
		return fmt.Errorf("не удалось получить последнюю версию сборки: %w", err)
	}
	fmt.Printf("Версия: %s\n", latestInstance.Version)
	fmt.Printf("\n")
	for _, instanceContainer := range latestInstance.Containers {
		var installedContainer *types.Container

		for i, container := range u.cm.Config().InstalledContainers {
			if instanceContainer.ContentType == container.ContentType {
				installedContainer = &u.cm.Config().InstalledContainers[i]
				break

			}
		}

		if installedContainer == nil {
			installedContainer = &types.Container{
				ContentType: instanceContainer.ContentType,
				Content:     make([]types.Content, 0),
			}
		}

		contentProcessor := updater.NewContentProcessor(instanceContainer, *installedContainer, u.fm, u.apiClient)
		err := contentProcessor.Process()
		if err != nil {
			return err
		}
	}

	cfg := u.cm.Config()
	cfg.InstalledContainers = latestInstance.Containers
	u.cm.Save()
	
	return nil
}