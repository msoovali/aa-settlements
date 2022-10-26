package mocks

import "github.com/msoovali/aa-settlements/domain/port"

type ArchiverPortMock struct {
	port.ArchiverPort
	FakeArchiveLastYearMonthlyFolders  func(lastYear int, rootFolderId string) error
	ArchiveLastYearMonthlyFoldersCalls int
	FakeArchiveLastYearSpreadsheets    func(lastYear int, rootFolderId string) (lastMonthSpreadSheetId string, err error)
	ArchiveLastYearSpreadsheetsCalls   int
}

func (m *ArchiverPortMock) ArchiveLastYearMonthlyFolders(lastYear int, rootFolderId string) error {
	m.ArchiveLastYearMonthlyFoldersCalls++
	if m.FakeArchiveLastYearMonthlyFolders != nil {
		return m.FakeArchiveLastYearMonthlyFolders(lastYear, rootFolderId)
	}

	return nil
}

func (m *ArchiverPortMock) ArchiveLastYearSpreadsheets(lastYear int, rootFolderId string) (lastMonthSpreadSheetId string, err error) {
	m.ArchiveLastYearSpreadsheetsCalls++
	if m.FakeArchiveLastYearSpreadsheets != nil {
		return m.FakeArchiveLastYearSpreadsheets(lastYear, rootFolderId)
	}

	return "", nil
}
