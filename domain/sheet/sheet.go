package sheet

import (
	"fmt"
	"time"

	"github.com/msoovali/aa-settlements/domain/clock"
	"github.com/msoovali/aa-settlements/domain/port"
)

type Sheet struct {
	id    string
	port  port.GoogleSheetPort
	clock clock.Clock
}

func CopyNextMonthFromPrevious(sheetPort port.GoogleSheetPort, drivePort port.GoogleDrivePort, clock clock.Clock) (*Sheet, error) {
	spreadsheetId, err := drivePort.CopySpreadsheet(getPreviousMonthFileName(clock), getCurrentMonthFileName(clock))
	if err != nil {
		return nil, err
	}

	return SheetFromId(spreadsheetId, sheetPort, clock), nil
}

func CopyNextMonthFromSheetId(sheetPort port.GoogleSheetPort, drivePort port.GoogleDrivePort, sheetId string, clock clock.Clock) (*Sheet, error) {
	spreadsheetId, err := drivePort.CopySpreadsheetFromId(sheetId, getCurrentMonthFileName(clock))
	if err != nil {
		return nil, err
	}

	return SheetFromId(spreadsheetId, sheetPort, clock), nil
}

func SheetFromId(id string, port port.GoogleSheetPort, clock clock.Clock) *Sheet {
	return &Sheet{
		id:    id,
		port:  port,
		clock: clock,
	}
}

func (s *Sheet) CutAndPasteRange(cutRange string, pasteRange string) error {
	return s.port.CutAndPasteRange(s.id, cutRange, pasteRange)
}

func (s *Sheet) ClearRange(_range string) error {
	return s.port.ClearRange(s.id, _range)
}

func (s *Sheet) SetValue(value string, column string) error {
	return s.port.SetValue(s.id, value, column)
}

func (s *Sheet) SetCurrentMonthPeriodMetadata(currentPeriodMetadataRange string) error {
	return s.SetValue(getCurrentMonthPeriodMetaData(s.clock), currentPeriodMetadataRange)
}

func getPreviousMonthFileName(clock clock.Clock) string {
	now := clock.Now()
	return fmt.Sprintf("%d_%02d", now.Year(), now.Month()-1)
}

func getCurrentMonthFileName(clock clock.Clock) string {
	now := clock.Now()
	return fmt.Sprintf("%d_%02d", now.Year(), now.Month())
}

func getCurrentMonthPeriodMetaData(clock clock.Clock) string {
	now := clock.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return fmt.Sprintf("Periood %02d.%02d - %02d.%02d", firstOfMonth.Day(), currentMonth, lastOfMonth.Day(), currentMonth)
}
