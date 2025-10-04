package updater

import (
	"fmt"
	"path/filepath"

	"github.com/CubeShield/CubeHopper/internal/api"
	"github.com/CubeShield/CubeHopper/internal/filesystem"
	"github.com/CubeShield/CubeHopper/internal/types"
	"github.com/fatih/color"
)


type contentProcessor struct {
	contentType      string
	apiContent       []types.Content
	installedContent []types.Content

	fm        *filesystem.FileManager
	apiClient *api.ApiClient

	infoPrinter  *color.Color
}

func NewContentProcessor(
	container types.Container,
	installedContainer types.Container,
	fm *filesystem.FileManager,
	apiClient *api.ApiClient,
) *contentProcessor {
	return &contentProcessor{
		contentType:      container.ContentType,
		apiContent:       container.Content,
		installedContent: installedContainer.Content,
		fm:               fm,
		apiClient:        apiClient,
		infoPrinter:      color.New(color.FgWhite),
	}
}

func (p *contentProcessor) toSet(contentList []types.Content) map[string]types.Content {
	set := make(map[string]types.Content, len(contentList))
	for _, item := range contentList {
		set[item.File] = item
	}
	return set
}

func (p *contentProcessor) Process() error {
	fmt.Printf("● Обработка %s\n", p.contentType)

	apiSet := p.toSet(p.apiContent)
	installedSet := p.toSet(p.installedContent)

	for fileName := range installedSet {
		if _, exists := apiSet[fileName]; !exists {
			if err := p.delete(installedSet[fileName]); err != nil {
				fmt.Printf("Ошибка удаления %s: %v\n", fileName, err)
			}
		}
	}


	for fileName := range apiSet {
		if _, exists := installedSet[fileName]; !exists {
			if err := p.install(apiSet[fileName]); err != nil {
				fmt.Printf("Ошибка установки %s: %v\n", fileName, err)
			}
		}
	}
	fmt.Printf("✔ Завершнено\n\n")
	return nil
}

func (p *contentProcessor) delete(content types.Content) error {
	p.infoPrinter.Printf("· Удаление %s\n", content.File)
	fileNameWithPrefix := fmt.Sprintf("[CubeHopper] %s", content.File)
	relativePath := filepath.Join(p.contentType, fileNameWithPrefix)
	return p.fm.Delete(relativePath)
}

func (p *contentProcessor) install(content types.Content) error {
	p.infoPrinter.Printf("· Скачивание %s\n", content.File)
	
	stream, err := p.apiClient.DownloadFile(content.Url)
	if err != nil {
		return err
	}
	defer stream.Close()

	fileNameWithPrefix := fmt.Sprintf("[CubeHopper] %s", content.File)
	relativePath := filepath.Join(p.contentType, fileNameWithPrefix)

	return p.fm.Save(relativePath, stream)
}