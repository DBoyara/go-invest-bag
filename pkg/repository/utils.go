package repository

import (
	"math"
	"strconv"
	"strings"

	"github.com/DBoyara/go-invest-bag/pkg/models"
)

// const can export's to server
const (
	EURO = "euro"
	RUB  = "rub"
	USD  = "usd"
)

//CountingPackage counting
func CountingPackage(rubUsd string, rubEvro string, positions []models.Position) (models.Positions, error) {
	var proportionPositions models.Positions
	var stockAmount float64
	var obligationAmount float64

	rubUsdFloat, err := formatNumber(rubUsd)
	if err != nil {
		return proportionPositions, errWrongFormatCurrency
	}

	rubEvroFloat, err := formatNumber(rubEvro)
	if err != nil {
		return proportionPositions, errWrongFormatCurrency
	}

	for _, p := range positions {
		if p.Type == "Акция" {
			stock := models.Stock{
				Ticker: p.Ticker,
				Amount: convertAmount(rubUsdFloat, rubEvroFloat, p.Amount, p.Currency),
				Count:  p.Count,
			}
			proportionPositions.Stocks = append(proportionPositions.Stocks, stock)
			stockAmount += stock.Amount
		}
		if p.Type == "Облигация" {
			obligation := models.Obligation{
				Ticker: p.Ticker,
				Amount: convertAmount(rubUsdFloat, rubEvroFloat, p.Amount, p.Currency),
				Count:  p.Count,
			}
			proportionPositions.Obligations = append(proportionPositions.Obligations, obligation)
			obligationAmount += obligation.Amount
		}
	}

	proportionPositions.Relation.Total = toFixed(obligationAmount+stockAmount, 2)

	stockPercent := stockAmount * 100 / proportionPositions.Relation.Total
	proportionPositions.Relation.StockPercent = toFixed(stockPercent, 2)

	obligationPercent := obligationAmount * 100 / proportionPositions.Relation.Total
	proportionPositions.Relation.ObligationPercent = toFixed(obligationPercent, 2)

	return proportionPositions, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func convertAmount(rubUsdFloat float64, rubEvroFloat float64, amount float64, currency string) float64 {
	if currency == USD {
		return toFixed(amount*rubUsdFloat, 2)
	}
	if currency == EURO {
		return toFixed(amount*rubEvroFloat, 2)
	}
	return amount
}

func formatNumber(n string) (float64, error) {
	if strings.Contains(n, ".") {
		if number, err := strconv.ParseFloat(n, 64); err == nil {
			return number, err
		}
	}

	if number, err := strconv.ParseInt(n, 10, 64); err == nil {
		return float64(number), err
	}

	return 0, nil
}

func toFixedOnCurrency(num float64, c string) float64 {

	if c != RUB {
		output := math.Pow(10, 4.0)
		return float64(round(num*output)) / output
	}

	output := math.Pow(10, 2.0)
	return float64(round(num*output)) / output
}
