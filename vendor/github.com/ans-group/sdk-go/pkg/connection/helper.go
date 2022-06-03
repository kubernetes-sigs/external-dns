package connection

func Get[T any](conn Connection, resource string, parameters APIRequestParameters, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	response, err := conn.Get(resource, parameters)
	return handleResponse[T](response, err, handlers)
}

func Post[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	response, err := conn.Post(resource, body)
	return handleResponse[T](response, err, handlers)
}

func Put[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	response, err := conn.Put(resource, body)
	return handleResponse[T](response, err, handlers)
}

func Patch[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	response, err := conn.Patch(resource, body)
	return handleResponse[T](response, err, handlers)
}

func Delete[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	response, err := conn.Delete(resource, body)
	return handleResponse[T](response, err, handlers)
}

func handleResponse[T any](response *APIResponse, err error, handlers []ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	if err != nil {
		return responseBody, err
	}

	return responseBody, response.HandleResponse(responseBody, handlers...)
}
