package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sshaparenko/validgate/internal/domain"
	"github.com/sshaparenko/validgate/internal/service"
)

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	card := domain.Card{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		renderErr := render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "PARS-1001",
				Message: "Failed to read request body",
			},
		})
		if renderErr != nil {
			log.Println(fmt.Errorf("in service.ValidateCard: %w", err))
		}
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(b, &card); err != nil {
		renderErr := render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "PARS-1002",
				Message: "Failed to unmarshal request body",
			},
		})
		if renderErr != nil {
			log.Println(fmt.Errorf("in service.ValidateCard: %w", err))
		}
		return
	}

	result, err := card.Validate()

	if errors.Is(err, service.ErrLuhnCheck) {
		renderErr := render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "CARD-VAL-1001",
				Message: "Card number fails Luhn check",
			},
		})
		if renderErr != nil {
			log.Println(fmt.Errorf("in service.ValidateCard: %w", err))
		}
		return
	}

	if errors.Is(err, service.ErrExpDate) {
		renderErr := render.Render(w, r, &domain.ValidationResponse{
			Valid: false,
			Error: domain.ValidationError{
				Code:    "CARD-VAL-1002",
				Message: "card is expired",
			},
		})
		if renderErr != nil {
			log.Println(fmt.Errorf("in service.ValidateCard: %w", err))
		}
		return
	}

	renderErr := render.Render(w, r, &domain.ValidationResponse{
		Valid: result,
	})
	if renderErr != nil {
		log.Println(fmt.Errorf("in service.ValidateCard: %w", err))
	}
}
