package models

type Delivery struct {
	StartPoint   string
	EndPoint     string
	ItemId       string
	Status       string
	DeliveryType string
}

type DeliveryInfo struct {
	Id           string
	StartPoint   string
	EndPoint     string
	Status       string
	DeliveryType string
	Time
}

type DeliveryUpdate struct {
	Id           string
	StartPoint   string
	EndPoint     string
	Status       string
	DeliveryType string
}
