package order

import "context"

type OrderUsecase interface{
	Insert(c context.Context, payload OrderPayload) (id string, err error)
	GetOrder(c context.Context, id string) (order Order, err error)
	StatusUpdate(c context.Context, id string, status string) (err error)
}

type orderUsecase struct{
	orderRepository OrderRepository
}

func NewOrderUsecase(orderRepository OrderRepository) OrderUsecase{
	return &orderUsecase{orderRepository}
}

func (uc *orderUsecase) Insert(c context.Context, payload OrderPayload) (id string, err error){
	id, err = uc.orderRepository.Save(c, payload)

	return id, err
}

func (uc *orderUsecase) GetOrder(c context.Context, id string) (order Order, err error){
	order, err = uc.orderRepository.GetById(c, id)

	return order, err
}

func (uc *orderUsecase) StatusUpdate(c context.Context, id string, status string) (err error){
	err = uc.orderRepository.StatusUpdate(c, id, status)

	return err
}