package handler

import "net/http"

type IndexHandler struct {
	RedirectPath string
}

func (m *IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Location", m.RedirectPath)
	res.WriteHeader(http.StatusFound)
}
