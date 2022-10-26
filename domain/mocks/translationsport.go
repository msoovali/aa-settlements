package mocks

import (
	"github.com/msoovali/aa-settlements/domain/localizer"
	"github.com/msoovali/aa-settlements/domain/port"
)

type TranslationsPortMock struct {
	port.TranslationsPort
	FakeGet  func(locale string) (*localizer.Localizer, error)
	GetCalls int
}

func (m *TranslationsPortMock) Get(locale string) (*localizer.Localizer, error) {
	m.GetCalls++
	if m.FakeGet != nil {
		return m.FakeGet(locale)
	}

	return nil, nil
}
