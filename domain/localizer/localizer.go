package localizer

import "fmt"

type Localizer struct {
	values map[string]interface{}
}

func Init(values map[string]interface{}) *Localizer {
	return &Localizer{
		values: values,
	}
}

func (l Localizer) Translate(key string) string {
	translation := l.values[key]

	if translation != nil {
		return fmt.Sprint(translation)
	}

	return key
}
