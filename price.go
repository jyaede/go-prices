package prices

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

//Price ...
type Price float64

//New ...
func New(f float64) Price {
	v := toFixed(f, 2)
	return Price(v)
}

//NewFromInt ...
func NewFromInt(i int) Price {
	return New(toFixed(float64(float64(i)/100), 2))
}

//Float64 ...
func (p Price) Float64() float64 {
	return float64(p)
}

//Currency ...
func (p Price) Currency(symbol string) string {
	return fmt.Sprintf("%s%s", symbol, p)
}

//USD ...
func (p Price) USD() string {
	return p.Currency("$")
}

//Abs ...
func (p Price) Abs() Price {
	if p < 0 {
		return New(p.Float64() * -1)
	}
	return p
}

//Int ...
func (p Price) Int() int {
	return int(toFixed(p.Float64()*100, 0))
}

func (p Price) String() string {
	return strconv.FormatFloat(p.Float64(), 'f', 2, 64)
}

//SetBSON ...
func (p *Price) SetBSON(raw bson.Raw) error {
	var f float64
	if err := raw.Unmarshal(&f); err != nil {
		return err
	}
	*p = NewFromInt(round(f))
	return nil
}

//GetBSON ...
func (p Price) GetBSON() (interface{}, error) {
	return p.Int(), nil
}

//MarshalJSON ...
func (p Price) MarshalJSON() ([]byte, error) {
	return json.Marshal(toFixed(p.Float64(), 2))
}

//UnmarshalJSON ...
func (p *Price) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	*p = New(f)
	return nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
