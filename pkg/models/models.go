package models

type ErrModel struct {
	Err string `json:"error"`
}

type Position struct {
	Ticker    string  `gorm:"primaryKey" json:"tiker"`
	CreatedAt int64   `gorm:"autoCreateTime"`
	UpdatedAt int64   `gorm:"autoUpdateTime"`
	DeletedAt int64   `gorm:"autoUpdateTime"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Count     uint16  `json:"count"`
	Price     float32 `json:"price"`
	Amount    uint32
	Currency  string `json:"currency"`
}
