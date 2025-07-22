package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/w1ntermute14/labhub-proxy/config"
    	"github.com/w1ntermute14/labhub-proxy/internal/handler"
    	"github.com/w1ntermute14/labhub-proxy/internal/middleware"
    	"github.com/w1ntermute14/labhub-proxy/internal/services"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	service := services.NewProxyService(cfg)
	handler := handler.NewHandler(service)

	mux := http.NewServeMux()
	mux.Handle("/predict/hba1c", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleHba1c)))
	mux.Handle("/predict/ldll", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleLdll)))
	mux.Handle("/predict/ferr", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleFerr)))
	mux.Handle("/predict/ldl", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleLdl)))
	mux.Handle("/predict/tg", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleTg)))
	mux.Handle("/predict/hdl", middleware.AuthMiddleware(cfg.AuthToken, http.HandlerFunc(handler.HandleHdl)))


	fmt.Printf("Сервер запущен на порту :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
