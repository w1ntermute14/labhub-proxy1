package models

type LdllRequest struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	Chol   float64 `json:"chol"`
	HDL    float64 `json:"hdl"`
	TG     float64 `json:"tg"`
}

type LdllResponse struct {
	UID        string  `json:"uid"`
	Prediction float64 `json:"prediction"`
	Model      string  `json:"model"`
}
