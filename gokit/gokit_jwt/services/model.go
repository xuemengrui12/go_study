package services

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

//接收http客户端的请求后，把请求参数转为请求模型对象，用于后续业务逻辑处理
type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

//用于向客户端响应结果
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}
