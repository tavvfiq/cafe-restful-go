package helper

// DefaultResponse default response from controller
type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

// ErrorResponse helpers for returning error response from controllers
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"msg"`
}

// AuthResponse helpers for returning success response from controllers
type AuthResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

// MenuResponsePagination Menu response with pagination
// type MenuResponsePagination struct {
// 	Status   int         `json:"status"`
// 	Message  string      `json:"msg"`
// 	Menu     []tb.Menu   `json:"menu"`
// 	PageInfo tb.PageInfo `json:"pageInfo"`
// }
