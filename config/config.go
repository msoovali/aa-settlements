package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	RootFolderId               string
	SubFolders                 []Folder
	ClearRanges                []string
	CutAndPasteRanges          []CutAndPasteRange
	CurrentPeriodMetadataRange string
	Locale                     string
	YearlyArchivedFolderIds    []string
}

type CutAndPasteRange struct {
	CutRange   string
	PasteRange string
}

type Folder struct {
	Name       string
	SubFolders []Folder
}

func LoadConfiguration(file string) (config *Config, err error) {
	configFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return
}
