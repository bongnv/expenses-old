package handlers

import (
	"context"

	"github.com/bongnv/expenses"
	"github.com/bongnv/expenses/internal/storage"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// HandlerImpl ...
type HandlerImpl struct {
	implExpenseHandler
}

func (h *HandlerImpl) Hello(ctx context.Context, req *expenses.Request) (*expenses.Response, error) {
	return nil, nil
}

func (h *HandlerImpl) Bye(ctx context.Context, req *expenses.ByeRequest) (*expenses.ByeResponse, error) {
	return nil, nil
}

// New creates a new service with business logic.
func New() *HandlerImpl {
	h := &HandlerImpl{}
	db := storage.InitDB()
	h.ExpenseDAO = storage.NewExpenseDAO(db)
	return h
}
