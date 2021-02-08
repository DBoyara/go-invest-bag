package repository

import (
	"context"
	"errors"
	"log"

	"github.com/DBoyara/go-invest-bag/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrTickerNotFound = errors.New("record not found")
	ErrSellMoreHave   = errors.New("you can't sell more than you have")
)

func AddPosition(db *gorm.DB, ctx context.Context, position *models.Position) (int, error) {
	var p models.Position
	r := db.Where("ticker = ?", position.Ticker).First(&p)
	if r.Error == ErrTickerNotFound {
		db.Create(&position)
		return 201, nil
	}
	if r.Error != nil {
		return 501, r.Error
	}

	position.Count += p.Count
	position.Amount = float64(position.Count) * position.Price

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}},
		DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount"}),
	}).Create(&position)

	return 200, nil
}

func DelPosition(db *gorm.DB, ctx context.Context, position *models.Position) (int, error) {
	var p models.Position
	r := db.Where("ticker = ?", position.Ticker).First(&p)
	if r.Error != nil {
		return 501, r.Error
	}

	position.Count = p.Count - position.Count
	if position.Count < 0 {
		return 400, ErrSellMoreHave
	}
	if position.Count == 0 {
		db.Where("ticker = ?", position.Ticker).Delete(&position)
		return 200, nil
	}
	position.Amount = float64(position.Count) * position.Price

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}},
		DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount"}),
	}).Create(&position)

	return 200, nil
}

func GetPosition(db *gorm.DB, ctx context.Context, position *models.Position) {
	var p *models.Position
	result := db.WithContext(ctx).Find(&p)
	log.Println(db.Scan(&result))
}

func GetPositions() {

}
