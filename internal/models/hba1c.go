package models

type Hba1cRequest struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	Rdw    float64 `json:"rdw"`
	Wbc    float64 `json:"wbc"`
	Rbc    float64 `json:"rbc"`
	Hgb    float64 `json:"hgb"`
	Hct    float64 `json:"hct"`
	Mcv    float64 `json:"mcv"`
	Mch    float64 `json:"mch"`
	Mchc   float64 `json:"mchc"`
	Plt    float64 `json:"plt"`
	Neu    float64 `json:"neu"`
	Eos    float64 `json:"eos"`
	Bas    float64 `json:"bas"`
	Lym    float64 `json:"lym"`
	Mon    float64 `json:"mon"`
	Soe    float64 `json:"soe"`
	Chol   float64 `json:"chol"`
	Glu    float64 `json:"glu"`
}

type Hba1cResponse struct {
	UID        string  `json:"uid"`
	Prediction float64 `json:"prediction"`
	Model      string  `json:"model"`
}
