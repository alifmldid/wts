package order

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

var baseURL = "http://localhost:8003/order"

type OrderRepository interface{
	Save(c context.Context, payload OrderPayload) (id string, err error)
	GetById(c context.Context, id string) (order Order, err error)
	StatusUpdate(c context.Context, id string, status string) (err error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository{
	return &orderRepository{}
}

func (repo *orderRepository) Save(c context.Context, payload OrderPayload) (id string, err error){
	client := &http.Client{}
	var responseData NewOrderResponse

	orderByte, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(orderByte))
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return "", err
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return "", err
	}

	return responseData.Data.Id, nil
}

func (repo *orderRepository) GetById(c context.Context, id string) (order Order, err error){
	client := &http.Client{}
	var responseData OrderResponse

	request, err := http.NewRequest("GET", baseURL+"/"+id, nil)
	if err != nil {
		return Order{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return Order{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return Order{}, err
	}

	order = responseData.Data

	return 
}

func (repo *orderRepository) StatusUpdate(c context.Context, id string, status string) (err error){
	client := &http.Client{}

	request, err := http.NewRequest("PUT", baseURL+"/"+id+"/"+status, nil)
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return err
}