package prices

import (
	"encoding/json"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/stretchr/testify/assert"
)

type priceTest struct {
	Input        float64
	InputInt     int
	OutputInt    int
	OutputString string
	OutputUSD    string
	OutputAbs    Price
}

var priceTests = []priceTest{
	{
		Input:        0,
		InputInt:     0,
		OutputInt:    0,
		OutputString: "0.00",
		OutputUSD:    "$0.00",
		OutputAbs:    0,
	},
	{
		Input:        1.235235,
		InputInt:     124,
		OutputInt:    124,
		OutputString: "1.24",
		OutputUSD:    "$1.24",
		OutputAbs:    1.24,
	},
	{
		Input:        -1.99,
		InputInt:     -199,
		OutputInt:    -199,
		OutputString: "-1.99",
		OutputUSD:    "$-1.99",
		OutputAbs:    1.99,
	},
}

func TestPrices(t *testing.T) {
	assert := assert.New(t)
	for _, pt := range priceTests {
		p := New(pt.Input)
		assert.Equal(pt.OutputInt, p.Int())
		assert.Equal(pt.OutputString, p.String())
		assert.Equal(pt.OutputUSD, p.USD())
		assert.Equal(pt.OutputAbs, p.Abs())

		//input int
		p = NewFromInt(pt.InputInt)
		assert.Equal(pt.OutputInt, p.Int())
		assert.Equal(pt.OutputString, p.String())
		assert.Equal(pt.OutputUSD, p.USD())
		assert.Equal(pt.OutputAbs, p.Abs())
	}
}

type jsonTest struct {
	Price    Price  `json:"price" bson:"price"`
	PtrPrice *Price `json:"ptr_price" bson:"ptr_price"`
}

func TestJSON(t *testing.T) {
	p := New(1.46)
	jss := jsonTest{
		Price:    p,
		PtrPrice: &p,
	}

	b, err := json.Marshal(jss)
	assert.Nil(t, err)
	str := `{"price":1.46,"ptr_price":1.46}`
	assert.Equal(t, str, string(b))

	var unm jsonTest
	json.Unmarshal(b, &unm)

	assert.Equal(t, jss.Price, unm.Price)
	assert.Equal(t, jss.PtrPrice, unm.PtrPrice)
}

func TestBSON(t *testing.T) {
	p := New(1.46)
	jss := jsonTest{
		Price:    p,
		PtrPrice: &p,
	}

	b, err := bson.Marshal(jss)
	assert.Nil(t, err)

	var unm jsonTest
	bson.Unmarshal(b, &unm)

	assert.Equal(t, jss.Price, unm.Price)
	assert.Equal(t, jss.PtrPrice, unm.PtrPrice)
}
