package model

import "encoding/json"

func UnmarshalCryptoCurrency(data []byte) (CryptoCurrency, error) {
	var r CryptoCurrency
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CryptoCurrency) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CryptoCurrency struct {
	ID                int64         `json:"id"`
	Name              string        `json:"name"`
	Symbol            string        `json:"symbol"`
	Slug              string        `json:"slug"`
	CirculatingSupply float64       `json:"circulating_supply"`
	TotalSupply       float64       `json:"total_supply"`
	MaxSupply         interface{}   `json:"max_supply"`
	DateAdded         string        `json:"date_added"`
	NumMarketPairs    int64         `json:"num_market_pairs"`
	Tags              []interface{} `json:"tags"`
	Platform          interface{}   `json:"platform"`
	CmcRank           int64         `json:"cmc_rank"`
	LastUpdated       string        `json:"last_updated"`
	Quote             Quote         `json:"quote"`
}

type Quote struct {
	Usd Usd `json:"USD"`
}

type Usd struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}
