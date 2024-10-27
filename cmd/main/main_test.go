package main

import (
	"net/http"
	"testing"

	"github.com/sshaparenko/validgate/internal/domain"
	"github.com/sshaparenko/validgate/internal/handlers"
	"github.com/steinfletcher/apitest"
)

func Test_ValidCardNumber(t *testing.T) {
	card := domain.Card{
		CardNumber: "4111111111111111",
		ExpMonth:   11,
		ExpYear:    2024,
	}

	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Post("api/v1/valid").
		JSON(card).
		Expect(t).
		Status(http.StatusOK).End()
}

func Test_InvalidCardNumber(t *testing.T) {
	card := domain.Card{
		CardNumber: "4111111111111110",
		ExpMonth:   11,
		ExpYear:    2024,
	}

	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Post("api/v1/valid").
		JSON(card).
		Expect(t).
		Status(http.StatusBadRequest).End()
}

func Test_InvalidMonth(t *testing.T) {
	card := domain.Card{
		CardNumber: "4111111111111111",
		ExpMonth:   10,
		ExpYear:    2024,
	}

	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Post("api/v1/valid").
		JSON(card).
		Expect(t).
		Status(http.StatusBadRequest).End()
}

func Test_InvalidYear(t *testing.T) {
	card := domain.Card{
		CardNumber: "4111111111111111",
		ExpMonth:   11,
		ExpYear:    2023,
	}

	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Post("api/v1/valid").
		JSON(card).
		Expect(t).
		Status(http.StatusBadRequest).End()
}
