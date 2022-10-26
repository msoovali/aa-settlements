package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/msoovali/aa-settlements/config"
	"github.com/msoovali/aa-settlements/domain/clock"
	"github.com/msoovali/aa-settlements/domain/localizer"
	"github.com/msoovali/aa-settlements/domain/mocks"
	"github.com/msoovali/aa-settlements/domain/port"
)

func TestCreateNextMonthSettlements(t *testing.T) {
	config := getConfig()
	t.Run("ifCurrentMonthIsJanary_shouldCallArchiver", func(t *testing.T) {
		translationsPort := getTranslationsPortMock(&localizer.Localizer{})
		drivePort := &mocks.DrivePortMock{}
		sheetPort := &mocks.SheetPortMock{}
		archiverPort := &mocks.ArchiverPortMock{}
		ports := port.Ports{
			DrivePort:        drivePort,
			TranslationsPort: translationsPort,
			SheetPort:        sheetPort,
			ArchiverPort:     archiverPort,
		}
		_time := time.Date(2023, time.January, 25, 0, 0, 0, 0, time.Now().Location())
		clock := &mocks.ClockMock{
			Time: &_time,
		}

		service, _ := New(&ports, config, clock, &localizer.Localizer{})
		err := service.CreateNextMonthSettlements()

		if err != nil {
			t.Errorf("Unexpected error received: %v", err)
		}
		expectedFolderCalls := len(config.YearlyArchivedFolderIds)
		if archiverPort.ArchiveLastYearMonthlyFoldersCalls != expectedFolderCalls {
			t.Errorf("Should call archiverPort.ArchiveLastYearMonthlyFolders %d times, but called %d times", expectedFolderCalls, archiverPort.ArchiveLastYearMonthlyFoldersCalls)
		}
		if archiverPort.ArchiveLastYearSpreadsheetsCalls != 1 {
			t.Errorf("Should call archiverPort.ArchiveLastYearSpreadsheets 1 times, but called %d times", archiverPort.ArchiveLastYearSpreadsheetsCalls)
		}
		if drivePort.CopySpreadsheetFromIdCalls != 1 {
			t.Errorf("Should call drivePort.CopySpreadsheetFromId 1 times, but called %d times", drivePort.CopySpreadsheetFromIdCalls)
		}
		if drivePort.CopySpreadsheetCalls != 0 {
			t.Errorf("Should call drivePort.CopySpreadsheet 0 times, but called %d times", drivePort.CopySpreadsheetFromIdCalls)
		}
	})
	for month := 2; month <= 12; month++ {
		t.Run(fmt.Sprintf("ifCurrentMonthIs%s_shouldNotCallArchiver", time.Month(month)), func(t *testing.T) {
			translationsPort := getTranslationsPortMock(&localizer.Localizer{})
			drivePort := &mocks.DrivePortMock{}
			sheetPort := &mocks.SheetPortMock{}
			archiverPort := &mocks.ArchiverPortMock{}
			ports := port.Ports{
				DrivePort:        drivePort,
				TranslationsPort: translationsPort,
				SheetPort:        sheetPort,
				ArchiverPort:     archiverPort,
			}
			_time := time.Date(2023, time.Month(month), 25, 0, 0, 0, 0, time.Now().Location())
			clock := &mocks.ClockMock{
				Time: &_time,
			}

			service, _ := New(&ports, config, clock, &localizer.Localizer{})
			err := service.CreateNextMonthSettlements()

			if err != nil {
				t.Errorf("Unexpected error received: %v", err)
			}
			if archiverPort.ArchiveLastYearMonthlyFoldersCalls != 0 {
				t.Errorf("Should call archiverPort.ArchiveLastYearMonthlyFolders 0 times, but called %d times", archiverPort.ArchiveLastYearMonthlyFoldersCalls)
			}
			if archiverPort.ArchiveLastYearSpreadsheetsCalls != 0 {
				t.Errorf("Should call archiverPort.ArchiveLastYearSpreadsheets 0 times, but called %d times", archiverPort.ArchiveLastYearSpreadsheetsCalls)
			}
			if drivePort.CopySpreadsheetFromIdCalls != 0 {
				t.Errorf("Should call drivePort.CopySpreadsheetFromId 0 times, but called %d times", drivePort.CopySpreadsheetFromIdCalls)
			}
			if drivePort.CopySpreadsheetCalls != 1 {
				t.Errorf("Should call drivePort.CopySpreadsheet 1 times, but called %d times", drivePort.CopySpreadsheetFromIdCalls)
			}
		})
	}
	t.Run("shouldPrepareNextMonthSheetAccordingToConfiguration", func(t *testing.T) {
		translationsPort := getTranslationsPortMock(&localizer.Localizer{})
		drivePort := &mocks.DrivePortMock{}
		sheetPort := &mocks.SheetPortMock{}
		archiverPort := &mocks.ArchiverPortMock{}
		ports := port.Ports{
			DrivePort:        drivePort,
			TranslationsPort: translationsPort,
			SheetPort:        sheetPort,
			ArchiverPort:     archiverPort,
		}

		service, _ := New(&ports, config, clock.RealClock{}, &localizer.Localizer{})
		err := service.CreateNextMonthSettlements()

		if err != nil {
			t.Errorf("Unexpected error received: %v", err)
		}
		if sheetPort.CutAndPasteRangeCalls != 1 {
			t.Errorf("Should call sheetPort.CutAndPasteRange 1 times, but called %d times", sheetPort.CutAndPasteRangeCalls)
		}
		if sheetPort.ClearRangeCalls != 2 {
			t.Errorf("Should call sheetPort.ClearRange 2 times, but called %d times", sheetPort.ClearRangeCalls)
		}
		if sheetPort.SetValueCalls != 1 {
			t.Errorf("Should call sheetPort.SetValue 1 times, but called %d times", sheetPort.SetValueCalls)
		}
	})
}

func getTranslationsPortMock(_localizer *localizer.Localizer) *mocks.TranslationsPortMock {
	return &mocks.TranslationsPortMock{
		FakeGet: func(locale string) (*localizer.Localizer, error) {
			return _localizer, nil
		},
	}
}

func getConfig() *config.Config {
	return &config.Config{
		SubFolders:              []config.Folder{{SubFolders: []config.Folder{{}}}, {}},
		ClearRanges:             []string{"Sheet1!A1:C3", "Sheet1!F12"},
		CutAndPasteRanges:       []config.CutAndPasteRange{{}},
		YearlyArchivedFolderIds: []string{"1", "2"},
	}
}
