package connection

<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
func GetRaw(conn Connection, resource string, parameters APIRequestParameters, responseBody interface{}, handlers ...ResponseHandler) error {
	response, err := conn.Get(resource, parameters)
	return handleResponse(response, err, responseBody, handlers)
}

func Get[T any](conn Connection, resource string, parameters APIRequestParameters, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	return responseBody, GetRaw(conn, resource, parameters, responseBody, handlers...)
}

func PostRaw(conn Connection, resource string, body interface{}, responseBody interface{}, handlers ...ResponseHandler) error {
	response, err := conn.Post(resource, body)
	return handleResponse(response, err, responseBody, handlers)
}

func Post[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	return responseBody, PostRaw(conn, resource, body, responseBody, handlers...)
}

func PutRaw(conn Connection, resource string, body interface{}, responseBody interface{}, handlers ...ResponseHandler) error {
	response, err := conn.Put(resource, body)
	return handleResponse(response, err, responseBody, handlers)
}

func Put[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	return responseBody, PutRaw(conn, resource, body, responseBody, handlers...)
}

func PatchRaw(conn Connection, resource string, body interface{}, responseBody interface{}, handlers ...ResponseHandler) error {
	response, err := conn.Patch(resource, body)
	return handleResponse(response, err, responseBody, handlers)
}

func Patch[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	return responseBody, PatchRaw(conn, resource, body, responseBody, handlers...)
}

func DeleteRaw(conn Connection, resource string, body interface{}, responseBody interface{}, handlers ...ResponseHandler) error {
	response, err := conn.Delete(resource, body)
	return handleResponse(response, err, responseBody, handlers)
}

func Delete[T any](conn Connection, resource string, body interface{}, handlers ...ResponseHandler) (*APIResponseBodyData[T], error) {
	responseBody := &APIResponseBodyData[T]{}
	return responseBody, DeleteRaw(conn, resource, body, responseBody, handlers...)
}

func handleResponse(response *APIResponse, err error, responseBody interface{}, handlers []ResponseHandler) error {
	if err != nil {
		return err
	}

	return response.HandleResponse(responseBody, handlers...)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
