package redis

import (
	"context"
	"encoding/json"
	"lesson62/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ItemRepo struct {
	Db *redis.Client
}

func NewItemRepo(db *redis.Client) *ItemRepo {
	return &ItemRepo{Db: db}
}

func (i *ItemRepo) CreateItem(ctx context.Context, item *models.Item) (*models.ItemInfo, error) {

	currentTime := time.Now().String()
	res := models.ItemInfo{
		Id:          uuid.NewString(),
		Title:       item.Title,
		Description: item.Description,
		Price:       item.Price,
		// CreatedAt:   currentTime,
		// UpdatedAt:   currentTime,
	}
	res.CreatedAt = currentTime
	res.UpdatedAt = currentTime
	
	itemData, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = i.Db.HSet(ctx, "items", res.Id, string(itemData)).Err()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (i *ItemRepo) GetItemById(ctx context.Context, id string) (*models.ItemInfo, error) {

	data, err := i.Db.HGet(ctx, "items", id).Result()
	if err != nil {
		return nil, err
	}
	res := models.ItemInfo{}

	err = json.Unmarshal([]byte(data), &res)

	return &res, err
}

func (i *ItemRepo) GetItems(ctx context.Context, id string) (*models.ItemInfo, error) {

	data, err := i.Db.HGetAll(ctx, "items").Result()
	if err != nil {
		return nil, err
	}
	res := models.ItemInfo{}

	err = json.Unmarshal([]byte(data), &res)

	return &res, err
}

