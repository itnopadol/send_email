package ctrl


type Response struct {
	Status  string `json:"status"`
	Message string  `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
