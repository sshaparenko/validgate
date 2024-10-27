package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sshaparenko/validgate/internal/domain"
	"github.com/sshaparenko/validgate/internal/handlers"
	"github.com/steinfletcher/apitest"
)

type CustomReader struct{}

func (cr *CustomReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated read error")
}

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

func Test_UnmurshalFailed(t *testing.T) {
	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Post("api/v1/valid").
		Body(`{  `).
		Expect(t).
		Status(http.StatusBadRequest).End()
}

func Test_InvalidBody(t *testing.T) {
	apitest.New().
		HandlerFunc(handlers.ValidateCard).
		Intercept(func(r *http.Request) {
			r.Body = io.NopCloser(&CustomReader{})
		}).
		Post("api/v1/valid").
		Expect(t).
		Status(http.StatusBadRequest).End()
}
