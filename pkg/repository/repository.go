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
	errTickerNotFound      = errors.New("record not found")
	errSellMoreHave        = errors.New("you can't sell more than you have")
	errWrongFormatCurrency = errors.New("wrong format of currency")
)

// AddPosition Add new or update position
func AddPosition(ctx context.Context, db *gorm.DB, position *models.Position) (int, error) {
	var p models.Position
	r := db.Where("ticker = ?", position.Ticker).First(&p)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		position.Amount = toFixed(float64(position.Count)*position.Price, 2)
		db.Create(&position)
		return 201, nil
	}
	if r.Error != nil {
		return 501, r.Error
	}

	position.Count += p.Count
	position.Amount = toFixed(float64(position.Count)*position.Price, 2)

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}},
		DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount", "updated_at"}),
	}).Create(&position)

	return 200, nil
}

// DelPosition Del or update position
func DelPosition(ctx context.Context, db *gorm.DB, position *models.Position) (int, error) {
	var p models.Position
	r := db.Where("ticker = ?", position.Ticker).First(&p)
	if r.Error != nil {
		return 501, r.Error
	}

	position.Count = p.Count - position.Count
	if position.Count < 0 {
		return 400, errSellMoreHave
	}

	position.Amount = toFixed(float64(position.Count)*position.Price, 2)
	if position.Count == 0 {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ticker"}},
			DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount", "updated_at", "deleted_at"}),
		}).Create(&position)
		return 200, nil
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}},
		DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount", "updated_at"}),
	}).Create(&position)

	return 200, nil
}

// GetPosition get one position
func GetPosition(ctx context.Context, db *gorm.DB, position *models.Position) {
	var p models.Position
	result := db.WithContext(ctx).Find(&p)
	log.Println(db.Scan(&result))
}

// GetPositions all positions
func GetPositions(ctx context.Context, db *gorm.DB) ([]models.Position, int, error) {
	var userPositions []models.Position

	rows, err := db.Model(&models.Position{}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, 501, err
	}

	for rows.Next() {
		var p models.Position
		db.ScanRows(rows, &p)
		if p.Count != 0 {
			userPositions = append(userPositions, p)
		}
	}

	return userPositions, 200, nil
}
