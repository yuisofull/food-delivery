package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging" form:"paging,omitempty"`
	Filter interface{} `json:"filter" form:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleNewSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil, Filter: nil}
}
