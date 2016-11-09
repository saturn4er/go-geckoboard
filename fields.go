package geckoboard

import "encoding/json"

type Field interface {
	MarshalJSON() ([]byte, error)
}
type Fields map[string]Field

type DateField struct {
	Name string
}

func (f *DateField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type: "date",
		Name: f.Name,
	})
}
func NewDateField(name string) *DateField {
	result := new(DateField)
	result.Name = name
	return result
}

type DateTimeField struct {
	Name string
}

func (f *DateTimeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type: "datetime",
		Name: f.Name,
	})
}

func NewDateTimeField(name string) *DateTimeField {
	result := new(DateTimeField)
	result.Name = name
	return result
}

type NumberField struct {
	Name string
}
func (f *NumberField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type: "number",
		Name: f.Name,
	})
}
func NewNumberField(name string) *NumberField {
	result := new(NumberField)
	result.Name = name
	return result
}

type PercentageField struct {
	Name string
}

func (f *PercentageField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type: "percentage",
		Name: f.Name,
	})
}
func NewPercentageField(name string) *PercentageField {
	result := new(PercentageField)
	result.Name = name
	return result
}

type MoneyField struct {
	Name     string
	Currency string
}

func (f *MoneyField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Currency string `json:"currency_code"`
		Type     string `json:"type"`
		Name     string `json:"name"`
	}{
		Currency: f.Currency,
		Type:     "money",
		Name:     f.Name,
	})
}
func NewMoneyField(name, currency string) *MoneyField {
	result := new(MoneyField)
	result.Name = name
	result.Currency = currency
	return result
}

type StringField struct {
	Name string
}

func (f *StringField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{
		Type: "string",
		Name: f.Name,
	})
}
func NewStringField(name string) *StringField {
	result := new(StringField)
	result.Name = name
	return result
}
