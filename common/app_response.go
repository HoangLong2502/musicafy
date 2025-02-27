package common

type successRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Paging  interface{} `json:"paging,omitempty"`
	Filter  interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{
		Code:    200,
		Message: "success",
		Data:    data,
		Paging:  paging,
		Filter:  filter,
	}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
