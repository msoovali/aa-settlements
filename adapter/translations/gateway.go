package translations

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/msoovali/aa-settlements/domain/localizer"
)

type translationsGateway struct{}

func NewTranslationsGateway() *translationsGateway {
	return &translationsGateway{}
}

func (g *translationsGateway) Get(locale string) (*localizer.Localizer, error) {
	values, err := loadLocale(locale)
	if err != nil {
		return nil, err
	}
	return localizer.Init(values), nil
}

func loadLocale(locale string) (values map[string]interface{}, err error) {
	file, err := os.Open(fmt.Sprintf("locales/%s.json", locale))
	if err != nil {
		return
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&values)
	return
}
