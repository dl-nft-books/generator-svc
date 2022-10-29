package requests

import (
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
)

type ListTasksRequest struct {
	Account *string               `filter:"account"`
	Status  *resources.TaskStatus `filter:"status"`
}

func NewListTasksRequest(r *http.Request) (*ListTasksRequest, error) {
	var request ListTasksRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}