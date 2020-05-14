package handler

type ResposeMessage struct {
	Error *ErrorMessage `json:"error,omitempty"`
	Data  *DataMessage  `json:"data,omitempty"`
}

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type DataMessage struct {
	ID         string      `json:"id,omitempty"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes,omitempty"`
}
