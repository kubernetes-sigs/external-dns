package infoblox

import (
	"fmt"
	"reflect"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func PagingGetObject[T any](
	c ibclient.IBConnector,
	obj ibclient.IBObject,
	ref string,
	queryParams map[string]string,
	res *[]T,
) (err error) {

	pagingResponse := pagingResponseStruct[T]{
		NextPageId: "",
		Result:     make([]T, 0),
	}

	//copy query params and update them
	queryParamsCopy := map[string]string{}
	for k, v := range queryParams {
		queryParamsCopy[k] = v
	}

	queryParamsCopy["_return_as_object"] = "1"
	queryParamsCopy["_paging"] = "1"
	queryParamsCopy["_max_results"] = "1000"

	err = c.GetObject(obj, "", ibclient.NewQueryParams(false, queryParamsCopy), &pagingResponse)
	if err != nil {
		return fmt.Errorf("could not fetch object: %s", err)
	} else {
		*res = append(*res, pagingResponse.Result...)
	}

	for {
		if pagingResponse.NextPageId == "" {
			return
		}
		queryParamsCopy["_page_id"] = pagingResponse.NextPageId
		pagingResponse.NextPageId = ""
		pagingResponse.Result = make([]T, 0)
		err = c.GetObject(obj, "", ibclient.NewQueryParams(false, queryParamsCopy), &pagingResponse)
		if err != nil {
			return fmt.Errorf("could not fetch object: %s", err)
		}

		*res = append(*res, pagingResponse.Result...)
		fmt.Print(fmt.Sprintln("Paging to retrieve", reflect.TypeOf(obj), len(*res)))
	}
}

type pagingResponseStruct[T any] struct {
	NextPageId string `json:"next_page_id,omitempty"`
	Result     []T    `json:"result,omitempty"`
}
