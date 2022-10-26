package port

import "github.com/msoovali/aa-settlements/domain/localizer"

type Ports struct {
	DrivePort        GoogleDrivePort
	SheetPort        GoogleSheetPort
	TranslationsPort TranslationsPort
	ArchiverPort     ArchiverPort
}

type GoogleDrivePort interface {
	CopySpreadsheet(targetFile string, destinationFile string) (string, error)
	CopySpreadsheetFromId(targetId string, destinationName string) (string, error)
	CreateFolder(name string, parentFolderId string) (string, error)
	MoveSpreadsheetToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error)
	MoveFolderToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error)
}

type GoogleSheetPort interface {
	CutAndPasteRange(spreadSheetId string, cutRange string, pasteRange string) error
	ClearRange(spreadSheetId string, _range string) error
	SetValue(spreadSheetId string, value string, column string) error
}

type TranslationsPort interface {
	Get(locale string) (*localizer.Localizer, error)
}

type ArchiverPort interface {
	ArchiveLastYearMonthlyFolders(lastYear int, rootFolderId string) error
	ArchiveLastYearSpreadsheets(lastYear int, rootFolderId string) (lastMonthSpreadSheetId string, err error)
}
