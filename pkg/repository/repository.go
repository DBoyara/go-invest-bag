package repository

import (
	"context"
	"log"

	"github.com/DBoyara/go-invest-bag/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddPosition(db *gorm.DB, ctx context.Context, position *models.Position) (*models.Position, int) {

	log.Println(position.Price)
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}},
		DoUpdates: clause.AssignmentColumns([]string{"count", "price", "amount"}),
	}).Create(&position)

	return position, 201
}

func GetPosition(db *gorm.DB, ctx context.Context) {
	var p *models.Position
	result := db.WithContext(ctx).Find(&p)
	log.Println(db.Scan(&result))
}

func UpdatePosition() {

}

func DelPosition() {

}

func GetPositions() {

}
