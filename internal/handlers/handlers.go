package handlers

import (
	"encoding/json"
	"errors"
	"exchanger_test/internal/models"
	"exchanger_test/internal/service"
	"exchanger_test/internal/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ExchangerHandler struct {
	service service.ExchangerLogic
}

func NewExchangerHandler(service service.ExchangerLogic) *ExchangerHandler {
	return &ExchangerHandler{service: service}
}

func (h *ExchangerHandler) GetExchanger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.WriteErrorResponse(w, http.StatusMethodNotAllowed, errors.New("method Not Allowed"))
			return
		}

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Error(err)
			utils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		var exchanger models.Exchanger

		if err := json.Unmarshal(reqBody, &exchanger); err != nil {
			logrus.Error(err)
			utils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		exchanges, err := h.service.ExchangeAmount(exchanger)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		utils.WriteResponse(w, http.StatusOK, exchanges)
	}
}
