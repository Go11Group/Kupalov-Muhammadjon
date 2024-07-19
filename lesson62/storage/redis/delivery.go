package redis

import (
	"context"
	"encoding/json"
	"lesson62/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type DeliveryRepo struct {
	Db *redis.Client
}

func NewDeliveryRepo(db *redis.Client) *DeliveryRepo {
	return &DeliveryRepo{Db: db}
}

func (d *DeliveryRepo) CreateDelivery(ctx context.Context, Delivery *models.Delivery) (*models.DeliveryInfo, error) {

	currentTime := time.Now().String()
	res := models.DeliveryInfo{
		Id:           uuid.NewString(),
		StartPoint:   Delivery.StartPoint,
		EndPoint:     Delivery.EndPoint,
		Status:       Delivery.Status,
		DeliveryType: Delivery.DeliveryType,
	}
	res.CreatedAt = currentTime
	res.UpdatedAt = currentTime

	DeliveryData, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = d.Db.HSet(ctx, "deliveries", res.Id, string(DeliveryData)).Err()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (d *DeliveryRepo) GetDeliveryById(ctx context.Context, id string) (*models.DeliveryInfo, error) {

	data, err := d.Db.HGet(ctx, "deliveries", id).Result()
	if err != nil {
		return nil, err
	}
	res := models.DeliveryInfo{}

	err = json.Unmarshal([]byte(data), &res)

	return &res, err
}

func (d *DeliveryRepo) GetDeliverys(ctx context.Context) (*[]models.DeliveryInfo, error) {

	data, err := d.Db.HGetAll(ctx, "deliveries").Result()
	if err != nil {
		return nil, err
	}
	res := []models.DeliveryInfo{}
	for id := range data {
		var Delivery models.DeliveryInfo

		err = json.Unmarshal([]byte(data[id]), &Delivery)
		if err != nil {
			return nil, err
		}
		res = append(res, Delivery)
	}

	return &res, nil
}

func (d *DeliveryRepo) UpdateDelivery(ctx context.Context, Delivery *models.DeliveryUpdate) (*models.DeliveryInfo, error) {

	DeliveryInfo, err := d.GetDeliveryById(ctx, Delivery.Id)
	if err != nil {
		return nil, err
	}
	DeliveryInfo.Id = Delivery.Id
	DeliveryInfo.StartPoint = Delivery.StartPoint
	DeliveryInfo.EndPoint = Delivery.EndPoint
	DeliveryInfo.Status = Delivery.Status
	DeliveryInfo.DeliveryType = Delivery.DeliveryType
	DeliveryInfo.UpdatedAt = time.Now().Format(time.RFC3339)

	DeliveryData, err := json.Marshal(DeliveryInfo)
	if err != nil {
		return nil, err
	}

	err = d.Db.HSet(ctx, "deliveries", Delivery.Id, string(DeliveryData)).Err()

	return DeliveryInfo, err
}

func (d *DeliveryRepo) DeleteDelivery(ctx context.Context, id string) error {

	err := d.Db.HDel(ctx, "deliveries", id).Err()

	return err
}
