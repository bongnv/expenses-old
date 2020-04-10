package expenses

import (
	"context"
	"time"
)

// Service is a simple interface for a service.
type Service interface {
	Hello(ctx context.Context, p *Request) (*Response, error)
	Bye(ctx context.Context, req *ByeRequest) (*ByeResponse, error)
}

// Request presents a request.
type Request struct {
	Name string `json:"name"`
}

// Response ...
type Response struct {
	Text string `json:"text"`
}

// ByeRequest presents a request.
type ByeRequest struct {
	Name string `json:"name"`
}

// ByeResponse ...
type ByeResponse struct {
	Text string `json:"text"`
}

// Expense ...
type Expense struct {
	ID       int64     `json:"id"`
	Category string    `json:"category"`
	Currency string    `json:"currency"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
	Account  string    `json:"account"`
}
