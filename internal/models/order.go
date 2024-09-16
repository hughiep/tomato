package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCanceled  OrderStatus = "canceled"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
)

type Order struct {
	ID     int
	Status OrderStatus
}

func NewOrder(id int) *Order {
	return &Order{
		ID:     id,
		Status: OrderStatusPending,
	}
}

func (o *Order) Pay() {
	o.Status = OrderStatusPaid
}

func (o *Order) Ship() {
	o.Status = OrderStatusShipped
}

func (o *Order) Deliver() {
	o.Status = OrderStatusDelivered
}

func (o *Order) IsPaid() bool {
	return o.Status == OrderStatusPaid
}

func (o *Order) IsShipped() bool {
	return o.Status == OrderStatusShipped
}

func (o *Order) IsDelivered() bool {
	return o.Status == OrderStatusDelivered
}

func (o *Order) IsPending() bool {
	return o.Status == OrderStatusPending
}

func (o *Order) StatusString() string {
	return string(o.Status)
}

func (o *Order) StatusIs(status OrderStatus) bool {
	return o.Status == status
}

func (o *Order) StatusIsAny(statuses ...OrderStatus) bool {
	for _, s := range statuses {
		if o.Status == s {
			return true
		}
	}
	return false
}

func (o *Order) StatusIsNot(status OrderStatus) bool {
	return o.Status != status
}

func (o *Order) StatusIsNone(statuses ...OrderStatus) bool {
	for _, s := range statuses {
		if o.Status == s {
			return false
		}
	}
	return true
}

func main() {
	order := NewOrder(1)
	order.Pay()
	order.Ship()
	order.Deliver()
	order.StatusString()
	order.StatusIs(OrderStatusDelivered)
	order.StatusIsAny(OrderStatusDelivered, OrderStatusPaid)
	order.StatusIsNot(OrderStatusPending)
	order.StatusIsNone(OrderStatusPending, OrderStatusPaid)
	order.IsPaid()
	order.IsShipped()
	order.IsDelivered()
	order.IsPending()
}
