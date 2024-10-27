package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sshaparenko/validgate/internal/domain"
	"github.com/sshaparenko/validgate/internal/service"
)

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	card := domain.Card{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "PARS-1001",
				Message: "Failed to read request body",
			},
		})
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(b, &card); err != nil {
		render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "PARS-1002",
				Message: "Failed to unmarshal request body",
			},
		})
		return
	}

	result, err := card.Validate()

	if errors.Is(err, service.ErrLuhnCheck) {
		render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "CARD-VAL-1001",
				Message: "Card number fails Luhn check",
			},
		})
		return
	}

	if errors.Is(err, service.ErrExpDate) {
		render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "CARD-VAL-1002",
				Message: "card is expired",
			},
		})
		return
	}

	render.Render(w, r, &domain.ValidationResponse{
		Valid: result,
	})
}
