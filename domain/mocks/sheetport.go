package mocks

import "github.com/msoovali/aa-settlements/domain/port"

type SheetPortMock struct {
	port.GoogleSheetPort
	FakeCutAndPasteRange  func(spreadSheetId string, cutRange string, pasteRange string) error
	CutAndPasteRangeCalls int
	FakeClearRange        func(spreadSheetId string, _range string) error
	ClearRangeCalls       int
	FakeSetValue          func(spreadSheetId string, value string, column string) error
	SetValueCalls         int
}

func (m *SheetPortMock) CutAndPasteRange(spreadSheetId string, cutRange string, pasteRange string) error {
	m.CutAndPasteRangeCalls++
	if m.FakeCutAndPasteRange != nil {
		return m.FakeCutAndPasteRange(spreadSheetId, cutRange, pasteRange)
	}

	return nil
}

func (m *SheetPortMock) ClearRange(spreadSheetId string, _range string) error {
	m.ClearRangeCalls++
	if m.FakeClearRange != nil {
		return m.FakeClearRange(spreadSheetId, _range)
	}

	return nil
}

func (m *SheetPortMock) SetValue(spreadSheetId string, value string, column string) error {
	m.SetValueCalls++
	if m.FakeSetValue != nil {
		return m.FakeSetValue(spreadSheetId, value, column)
	}

	return nil
}
