package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type RequestData struct {
	Uid    string  `json:"uid"`
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

func main() {
	http.HandleFunc("/proxy", proxyHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	
	q := r.URL.Query()
	parseFloat := func(key string) float64 {
		val, _ := strconv.ParseFloat(q.Get(key), 64)
		return val
	}
	parseInt := func(key string) int {
		val, _ := strconv.Atoi(q.Get(key))
		return val
	}

	data := RequestData{
		Uid:    "web-client",
		Age:    parseInt("age"),
		Gender: parseInt("gender"),
		Rdw:    parseFloat("rdw"),
		Wbc:    parseFloat("wbc"),
		Rbc:    parseFloat("rbc"),
		Hgb:    parseFloat("hgb"),
		Hct:    parseFloat("hct"),
		Mcv:    parseFloat("mcv"),
		Mch:    parseFloat("mch"),
		Mchc:   parseFloat("mchc"),
		Plt:    parseFloat("plt"),
		Neu:    parseFloat("neu"),
		Eos:    parseFloat("eos"),
		Bas:    parseFloat("bas"),
		Lym:    parseFloat("lym"),
		Mon:    parseFloat("mon"),
		Soe:    parseFloat("soe"),
		Chol:   parseFloat("chol"),
		Glu:    parseFloat("glu"),
	}


	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка при сериализации JSON", http.StatusInternalServerError)
		return
	}

		client := &http.Client{}
	req, err := http.NewRequest("POST", "https://apiml.labhub.online/api/v1/predict/hba1c", ioutil.NopCloser(bytes.NewReader(jsonData)))
	if err != nil {
		http.Error(w, "Ошибка создания запроса", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer 0l62<EJi/zJx]a?")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при обращении к внешнему API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
