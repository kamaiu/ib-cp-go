package gen

import (
	"encoding/json"
	"fmt"
	_ "github.com/valyala/fasthttp"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	s := `
LastPrice               31
	Symbol                  55
	Text                    58
	High                    70
	Low                     71
	Position                72
	MarketValue             73
	AveragePrice            74
	UnrealizedPnL           75
	FormattedPosition       76
	FormattedUnrealizedPnL  77
	DailyPnL                78
	ChangePrice             82
	ChangePercent           83
	BidPrice                84
	AskSize                 85
	AskPrice                86
	Volume                  87
	BidSize                 88

	Exchange         6004
	Conid6008        6008
	SecurityType     6070
	Months           6072
	RegularExpiry    6073
	Marker           6119
	UnderlyingConid  6457
	MarketDataAvailability  6509
	CompanyName             7051
	LastSize                7059
	ConidAndExchange        7094
	ContractDescription     7219
	ContractDescription2    7220
	ListingExchange         7221
	Industry                7280
	Category                7281
	AverageDailyVolume      7282
	HistoricVolume30d       7284
	PutCallRatio            7285
	DividendAmount          7286
	DividendYieldPercent    7287
	ExDateOfTheDividend     7288
	MarketCap               7289
	PE                      7290
	EPS                     7291
	CostBasis               7292
	High52Week              7293
	Low52Week               7294
	OpenPrice               7295
	ClosePrice              7296

	Delta              7308
	Gamma              7309
	Theta              7310
	Vega               7311
	ImpliedVolatility  7633

`
	parts := strings.Split(s, "\n")
	type pair struct {
		num  string
		name string
	}
	pairs := make([]pair, 0)
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if len(p) == 0 {
			continue
		}
		i := strings.IndexByte(p, ' ')
		name := strings.TrimSpace(p[0:i])
		pairs = append(pairs, pair{
			num:  strings.TrimSpace(p[i:]),
			name: name,
		})
	}

	for _, p := range pairs {
		fmt.Println("\"" + p.num + "\":\"" + p.name + "\",")
	}
}

func TestGen(t *testing.T) {
	f, err := ioutil.ReadFile("doc.json")
	if err != nil {
		t.Fatal(err)
	}
	spec := &Spec{}
	err = json.Unmarshal(f, spec)
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewResolver(spec)
	if err != nil {
		t.Fatal(err)
	}

	g := &Generator{resolver: r}
	err = g.generate()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)
}
