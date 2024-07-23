package redis

import (
	"context"
	"encoding/json"
	"lesson64/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type StockRepo struct {
	db *redis.Client
}

func NewStockRepo(db *redis.Client) *StockRepo {
	return &StockRepo{
		db: db,
	}
}
func (s *StockRepo) GetStockPrice(ctx context.Context, companyName string) (*models.Stock, error) {
	data, err := s.db.HGetAll(ctx, companyName).Result()
	if err != nil {
		return nil, err
	}

	res := models.Stock{
		CompanyName: companyName,
	}

	if val, ok := data["Price"]; ok {
		res.Price, _ = strconv.Atoi(val)
	}
	if val, ok := data["MaxPrice"]; ok {
		res.MaxPrice, _ = strconv.Atoi(val)
	}
	if val, ok := data["MinPrice"]; ok {
		res.MinPrice, _ = strconv.Atoi(val)
	}
	if val, ok := data["CurrentTime"]; ok {
		res.CurrentTime, _ = time.Parse(time.RFC3339, val)
	}
	return &res, nil
}

func (s *StockRepo) WritePrice(ctx context.Context, stockData string) error {
	stock := models.Stock{}
	err := json.Unmarshal([]byte(stockData), &stock)
	if err != nil {
		return err
	}

	currentStock, err := s.GetStockPrice(ctx, stock.CompanyName)
	if err != nil && err != redis.Nil {
		return err
	}

	if currentStock != nil {
		if stock.Price > currentStock.MaxPrice {
			stock.MaxPrice = stock.Price
		} else {
			stock.MaxPrice = currentStock.MaxPrice
		}

		if stock.Price < currentStock.MinPrice || currentStock.MinPrice == 0 {
			stock.MinPrice = stock.Price
		} else {
			stock.MinPrice = currentStock.MinPrice
		}
	} else {
		stock.MaxPrice = stock.Price
		stock.MinPrice = stock.Price
	}

	data := map[string]interface{}{
		"CompanyName": stock.CompanyName,
		"Price":       strconv.Itoa(stock.Price),
		"MaxPrice":    strconv.Itoa(stock.MaxPrice),
		"MinPrice":    strconv.Itoa(stock.MinPrice),
		"CurrentTime": stock.CurrentTime.Format(time.RFC3339),
	}

	err = s.db.HSet(ctx, stock.CompanyName, data).Err()
	return err
}
