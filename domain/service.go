package domain

import (
	"github.com/msoovali/aa-settlements/config"
	"github.com/msoovali/aa-settlements/domain/clock"
	"github.com/msoovali/aa-settlements/domain/localizer"
	"github.com/msoovali/aa-settlements/domain/port"
	"github.com/msoovali/aa-settlements/domain/sheet"
)

type service struct {
	ports     *port.Ports
	config    *config.Config
	localizer *localizer.Localizer
	clock     clock.Clock
}

func New(ports *port.Ports, config *config.Config, clock clock.Clock) (*service, error) {
	localizer, err := ports.TranslationsPort.Get(config.Locale)
	if err != nil {
		return nil, err
	}
	return &service{
		ports:     ports,
		config:    config,
		localizer: localizer,
		clock:     clock,
	}, nil
}

func (s *service) CreateNextMonthSettlements() error {
	newSheet, err := s.createSheet()
	if err != nil {
		return err
	}
	err = s.prepareNextMonthSheet(newSheet)
	if err != nil {
		return err
	}
	err = s.createFolders(s.config.SubFolders, s.config.RootFolderId)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) createSheet() (*sheet.Sheet, error) {
	now := s.clock.Now()
	if now.Month() == 1 {
		lastMonthSheetId, err := s.archiveYear(now.Year() - 1)
		if err != nil {
			return nil, err
		}
		return sheet.CopyNextMonthFromSheetId(s.ports.SheetPort, s.ports.DrivePort, lastMonthSheetId, s.clock)
	}

	return sheet.CopyNextMonthFromPrevious(s.ports.SheetPort, s.ports.DrivePort, s.clock)
}

func (s *service) archiveYear(lastYear int) (string, error) {
	for _, folderId := range s.config.YearlyArchivedFolderIds {
		err := s.ports.ArchiverPort.ArchiveLastYearMonthlyFolders(lastYear, folderId)
		if err != nil {
			return "", err
		}
	}
	return s.ports.ArchiverPort.ArchiveLastYearSpreadsheets(lastYear, s.config.RootFolderId)
}

func (s *service) prepareNextMonthSheet(sheet *sheet.Sheet) error {
	for _, _range := range s.config.CutAndPasteRanges {
		if err := sheet.CutAndPasteRange(_range.CutRange, _range.PasteRange); err != nil {
			return err
		}
	}
	for _, _range := range s.config.ClearRanges {
		if err := sheet.ClearRange(_range); err != nil {
			return err
		}
	}
	if err := sheet.SetCurrentMonthPeriodMetadata(s.config.CurrentPeriodMetadataRange); err != nil {
		return err
	}

	return nil
}
