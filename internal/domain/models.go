package domain

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/sshaparenko/validgate/internal/service"
)

type Card struct {
	CardNumber string `json:"card_number"`
	ExpMonth   int    `json:"exp_month"`
	ExpYear    int    `json:"exp_year"`
}

func (c *Card) Validate() (bool, error) {
	result, err := service.LuhnCheck(c.CardNumber)
	if err != nil {
		return false, err
	}
	if !service.CheckExpDate(c.ExpMonth, c.ExpYear) {
		return false, service.ErrExpDate
	}

	return result, nil
}

type ValidationResponse struct {
	Valid bool            `json:"valid"`
	Error ValidationError `json:"error"`
}

func (vr *ValidationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	switch {
	case strings.HasPrefix(vr.Error.Code, "PARS"):
		render.Status(r, http.StatusBadRequest)
	case strings.HasPrefix(vr.Error.Code, "CARD-VAL"):
		render.Status(r, http.StatusBadRequest)
	default:
		render.Status(r, http.StatusOK)
	}
	return nil
}

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
