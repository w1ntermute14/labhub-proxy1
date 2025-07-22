package models

type LdlRequest struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	RDW    float64 `json:"rdw"`
	WBC    float64 `json:"wbc"`
	RBC    float64 `json:"rbc"`
	HGB    float64 `json:"hgb"`
	HCT    float64 `json:"hct"`
	MCV    float64 `json:"mcv"`
	MCH    float64 `json:"mch"`
	MCHC   float64 `json:"mchc"`
	PLT    float64 `json:"plt"`
	NEU    float64 `json:"neu"`
	EOS    float64 `json:"eos"`
	BAS    float64 `json:"bas"`
	LYM    float64 `json:"lym"`
	MON    float64 `json:"mon"`
	SOE    float64 `json:"soe"`
	Chol   float64 `json:"chol"`
	GLU    float64 `json:"glu"`
}

type LdlResponse struct {
	UID        string  `json:"uid"`
	Prediction float64 `json:"prediction"`
	Model      string  `json:"model"`
}
