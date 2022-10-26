package googlesheet

import (
	"google.golang.org/api/sheets/v4"
)

type googleSheetGateway struct {
	service *sheets.Service
}

func NewSheetGateway(service *sheets.Service) *googleSheetGateway {
	return &googleSheetGateway{
		service: service,
	}
}

func (g *googleSheetGateway) CutAndPasteRange(spreadSheetId string, cutRange string, pasteRange string) error {
	values, err := g.service.Spreadsheets.Values.Get(spreadSheetId, cutRange).Do()
	if err != nil {
		return err
	}
	valueRange := &sheets.ValueRange{
		Values: values.Values,
	}
	if err = g.setValuesForRange(spreadSheetId, valueRange, pasteRange); err != nil {
		return err
	}
	g.ClearRange(spreadSheetId, cutRange)

	return err
}

func (g *googleSheetGateway) ClearRange(spreadSheetId string, _range string) error {
	_, err := g.service.Spreadsheets.Values.Clear(spreadSheetId, _range, &sheets.ClearValuesRequest{}).Do()

	return err
}

func (g *googleSheetGateway) SetValue(spreadSheetId string, value string, column string) error {
	values := make([][]interface{}, 1)
	values[0] = make([]interface{}, 1)
	values[0][0] = value
	return g.SetValuesForRange(spreadSheetId, values, column)
}

func (g *googleSheetGateway) SetValuesForRange(spreadSheetId string, values [][]interface{}, _range string) error {
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	return g.setValuesForRange(spreadSheetId, valueRange, _range)
}

func (g *googleSheetGateway) setValuesForRange(spreadSheetId string, valueRange *sheets.ValueRange, _range string) error {
	_, err := g.service.Spreadsheets.Values.Update(spreadSheetId, _range, valueRange).ValueInputOption("USER_ENTERED").Do()

	return err
}
