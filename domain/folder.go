package domain

import (
	"fmt"
	"log"

	"github.com/msoovali/aa-settlements/config"
)

func (s *service) createFolders(folders []config.Folder, parentFolderId string) error {
	for _, folder := range folders {
		if folder.Name == "" {
			continue
		}
		folderName := folder.Name
		if folderName == "{{month}}" {
			folderName = s.getCurrentMonthFolderName()
		}
		log.Printf("Creating folder %s", folderName)
		id, err := s.ports.DrivePort.CreateFolder(folderName, parentFolderId)
		if err != nil {
			return fmt.Errorf("failed to create folder %s. Error: %v", folderName, err)
		}
		if folder.SubFolders != nil {
			err = s.createFolders(folder.SubFolders, id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *service) getCurrentMonthFolderName() string {
	return s.localizer.Translate(s.clock.Now().Month().String())
}
