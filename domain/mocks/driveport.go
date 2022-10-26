package mocks

import "github.com/msoovali/aa-settlements/domain/port"

type DrivePortMock struct {
	port.GoogleDrivePort
	FakeCopySpreadsheet                 func(targetFile string, destinationFile string) (string, error)
	CopySpreadsheetCalls                int
	FakeCopySpreadsheetFromId           func(targetId string, destinationName string) (string, error)
	CopySpreadsheetFromIdCalls          int
	FakeCreateFolder                    func(name string, parentFolderId string) (string, error)
	CreateFolderCalls                   int
	FakeMoveSpreadsheetToAnotherFolder  func(fileName, sourceFolderId, targetFolderId string) (string, error)
	MoveSpreadsheetToAnotherFolderCalls int
	FakeMoveFolderToAnotherFolder       func(fileName, sourceFolderId, targetFolderId string) (string, error)
	MoveFolderToAnotherFolderCalls      int
}

func (m *DrivePortMock) CopySpreadsheet(targetFile string, destinationFile string) (string, error) {
	m.CopySpreadsheetCalls++
	if m.FakeCopySpreadsheet != nil {
		return m.FakeCopySpreadsheet(targetFile, destinationFile)
	}

	return "", nil
}

func (m *DrivePortMock) CopySpreadsheetFromId(targetId string, destinationName string) (string, error) {
	m.CopySpreadsheetFromIdCalls++
	if m.FakeCopySpreadsheetFromId != nil {
		return m.FakeCopySpreadsheetFromId(targetId, destinationName)
	}

	return "", nil
}

func (m *DrivePortMock) CreateFolder(name string, parentFolderId string) (string, error) {
	m.CreateFolderCalls++
	if m.FakeCreateFolder != nil {
		return m.FakeCreateFolder(name, parentFolderId)
	}

	return "", nil
}

func (m *DrivePortMock) MoveSpreadsheetToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error) {
	m.MoveSpreadsheetToAnotherFolderCalls++
	if m.FakeMoveSpreadsheetToAnotherFolder != nil {
		return m.FakeMoveSpreadsheetToAnotherFolder(fileName, sourceFolderId, targetFolderId)
	}

	return "", nil
}

func (m *DrivePortMock) MoveFolderToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error) {
	m.MoveFolderToAnotherFolderCalls++
	if m.FakeMoveFolderToAnotherFolder != nil {
		return m.FakeMoveFolderToAnotherFolder(fileName, sourceFolderId, targetFolderId)
	}

	return "", nil
}
