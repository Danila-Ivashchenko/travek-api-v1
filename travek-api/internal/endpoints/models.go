package endpoints

// errors

type BaseError struct {
	mess string
}

func NewBaseError(mess string) *BaseError {
	return &BaseError{mess: mess}
}

func (b *BaseError) Error() string {
	return b.mess
}

// responses

type baseResponse struct {
	Exaption map[string]interface{} `json:"exaption"`
	Succsess bool                   `json:"success"`
	Data     map[string]interface{}
}

func GoodBaseResponse(name string, data interface{}) baseResponse {
	response := baseResponse{}
	response.Succsess = true
	response.Data = map[string]interface{}{
		name: data,
	}
	return response
}

func BadBaseResponse(err error) baseResponse {
	response := baseResponse{}
	response.Succsess = false
	response.Exaption = map[string]interface{}{
		"error": err.Error(),
	}
	return response
}
