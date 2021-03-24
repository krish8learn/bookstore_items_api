package services

import (
	"net/http"

	"github.com/krish8learn/bookstore_items_api/src/domain/items"
	"github.com/krish8learn/booktstore_common/rest_errors"
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	//Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}
