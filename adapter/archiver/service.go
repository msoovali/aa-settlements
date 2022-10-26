package archiver

import (
	"fmt"
	"time"

	"github.com/msoovali/aa-settlements/config"
	"github.com/msoovali/aa-settlements/domain/localizer"
	"github.com/msoovali/aa-settlements/domain/port"
)

type service struct {
	drivePort port.GoogleDrivePort
	config    config.Config
	localizer *localizer.Localizer
}

func New(drivePort port.GoogleDrivePort, config config.Config, localizer *localizer.Localizer) *service {
	return &service{
		drivePort: drivePort,
		config:    config,
		localizer: localizer,
	}
}

func (s *service) ArchiveLastYearMonthlyFolders(lastYear int, rootFolderId string) error {
	archiveFolderId, err := s.drivePort.CreateFolder(fmt.Sprintf("%d", lastYear), rootFolderId)
	if err != nil {
		return err
	}
	for month := 1; month <= 12; month++ {
		monthString := s.localizer.Translate(time.Month(month).String())
		_, err = s.drivePort.MoveFolderToAnotherFolder(monthString, rootFolderId, archiveFolderId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) ArchiveLastYearSpreadsheets(lastYear int, rootFolderId string) (lastMonthSpreadSheetId string, err error) {
	archiveFolderId, err := s.drivePort.CreateFolder(fmt.Sprintf("%d", lastYear), rootFolderId)
	if err != nil {
		return
	}
	for month := 1; month <= 12; month++ {
		fileName := fmt.Sprintf("%d_%02d", lastYear, month)
		lastMonthSpreadSheetId, err = s.drivePort.MoveSpreadsheetToAnotherFolder(fileName, rootFolderId, archiveFolderId)
		if err != nil {
			return
		}
	}

	return
}
