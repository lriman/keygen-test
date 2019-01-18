package models

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Request struct {
	Key string `json:"key"`
}
