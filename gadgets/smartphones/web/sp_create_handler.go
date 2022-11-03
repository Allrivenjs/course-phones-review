package web

import (
	"encoding/json"
	"github.com/allrivenjs/course-phones-review/gadgets/smartphones/gateway"
	"github.com/allrivenjs/course-phones-review/internal/database"
	"net/http"
)

func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := h.Create(nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "error in create smartphone"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateSmartphoneHandler struct {
	gateway.SmartphoneStorageGateway
}

func NewCreateSmartphoneHandler(client *database.MySqlClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{
		gateway.NewSmartphoneCreateGateway(client),
	}
}
