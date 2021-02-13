package models

type Ticker string
type Count uint16
type Percent float64

type ErrModel struct {
	Err string `json:"error"`
}

type Position struct {
	Ticker    Ticker  `gorm:"primaryKey" json:"tiker"`
	CreatedAt int64   `gorm:"autoCreateTime"`
	UpdatedAt int64   `gorm:"autoUpdateTime"`
	DeletedAt int64   `gorm:"autoUpdateTime"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Count     Count   `json:"count"`
	Price     float64 `json:"price"`
	Amount    float64
	Currency  string `json:"currency"`
}

type Positions struct {
	Obligations []Obligation
	Stocks      []Stock
	Relation    PositionRelation
}

type Obligation struct {
	Ticker Ticker
	Amount float64
	Count  Count
}

type Stock struct {
	Ticker Ticker
	Amount float64
	Count  Count
}

type PositionRelation struct {
	StockPercent      float64
	ObligationPercent float64
	Total             float64
}
