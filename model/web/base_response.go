package web

type BaseResponse struct {
	Status  Status      `json:"status" swaggertype:"primitive,string" enums:"success,failed"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
