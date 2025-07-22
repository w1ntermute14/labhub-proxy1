package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/w1ntermute14/labhub-proxy/config"
	"github.com/w1ntermute14/labhub-proxy/internal/models"
)

type ProxyService struct {
	cfg        *config.Config
	httpClient *http.Client
}

func NewProxyService(cfg *config.Config) *ProxyService {
	return &ProxyService{
		cfg:        cfg,
		httpClient: &http.Client{},
	}
}

func (s *ProxyService) callExternalAPI(url string, payload interface{}, result interface{}) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")


	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("Статус внешнего API: %d\nОтвет: %s\n", resp.StatusCode, string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("external API returned status: %s", resp.Status)
	}

	// После чтения тела и проверки статуса декодируем в result
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}

	return nil
}

// Методы для каждой модели:

func (s *ProxyService) PredictHba1c(req models.Hba1cRequest) (models.Hba1cResponse, error) {
	var res models.Hba1cResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/hba1c", req, &res)
	return res, err
}

func (s *ProxyService) PredictLdll(req models.LdllRequest) (models.LdllResponse, error) {
	var res models.LdllResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/ldll", req, &res)
	return res, err
}

func (s *ProxyService) PredictFerr(req models.FerrRequest) (models.FerrResponse, error) {
	var res models.FerrResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/ferr", req, &res)
	return res, err
}

func (s *ProxyService) PredictLdl(req models.LdlRequest) (models.LdlResponse, error) {
	var res models.LdlResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/ldl", req, &res)
	return res, err
}

func (s *ProxyService) PredictTg(req models.TgRequest) (models.TgResponse, error) {
	var res models.TgResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/tg", req, &res)
	return res, err
}

func (s *ProxyService) PredictHdl(req models.HdlRequest) (models.HdlResponse, error) {
	var res models.HdlResponse
	err := s.callExternalAPI("https://apiml.labhub.online/api/v1/predict/hdl", req, &res)
	return res, err
}
