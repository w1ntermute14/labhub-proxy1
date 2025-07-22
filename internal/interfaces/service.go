package interfaces

import "net/http"

type ProxyService interface {
	HandleHba1c(w http.ResponseWriter, r *http.Request)
	HandleLdll(w http.ResponseWriter, r *http.Request)
	HandleFerr(w http.ResponseWriter, r *http.Request)
	HandleLdl(w http.ResponseWriter, r *http.Request)
	HandleTg(w http.ResponseWriter, r *http.Request)
	HandleHdl(w http.ResponseWriter, r *http.Request)
}
