package handlers

import (
	"context"
	"errors"

	"github.com/bongnv/expenses"
	"github.com/bongnv/expenses/internal/storage"
	"github.com/bongnv/gokit/util/entity/query"
	"github.com/bongnv/gokit/util/log"
)

type implExpenseHandler struct {
	ExpenseDAO storage.ExpenseDAO
}

func (i *implExpenseHandler) Get(ctx context.Context, req *expenses.GetExpenseRequest) (*expenses.GetExpenseResponse, error) {
	log.Info("ID", req.ID)
	entities, err := i.ExpenseDAO.Query(ctx,
		query.Equal("id", req.ID),
		query.Limit(1),
	)

	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, errors.New("handlers: resource not found")
	}

	return &expenses.GetExpenseResponse{
		Expense: entityToExpense(entities[0]),
	}, nil
}

func (i *implExpenseHandler) List(ctx context.Context, req *expenses.ListExpenseRequest) (*expenses.ListExpenseResponse, error) {
	filters := []query.Query{
		query.OrderBy("id", false),
	}
	if req.Offset > 0 {
		filters = append(filters, query.Offset(req.Offset))
	}

	if req.Limit > 0 {
		filters = append(filters, query.Limit(req.Limit))
	}
	entities, err := i.ExpenseDAO.Query(ctx, filters...)

	if err != nil {
		return nil, err
	}

	items := make([]*expenses.Expense, len(entities))
	for i, entity := range entities {
		items[i] = entityToExpense(entity)
	}

	return &expenses.ListExpenseResponse{
		Expenses: items,
	}, nil
}

func (i *implExpenseHandler) Create(ctx context.Context, req *expenses.CreateExpenseRequest) (*expenses.CreateExpenseResponse, error) {
	entity := entityFromExpense(req.Expense)
	err := i.ExpenseDAO.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &expenses.CreateExpenseResponse{
		Expense: entityToExpense(entity),
	}, nil
}

func (i *implExpenseHandler) Update(ctx context.Context, req *expenses.UpdateExpenseRequest) (*expenses.UpdateExpenseResponse, error) {
	entity := entityFromExpense(req.Expense)
	err := i.ExpenseDAO.Update(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &expenses.UpdateExpenseResponse{
		Expense: entityToExpense(entity),
	}, nil
}

func (i *implExpenseHandler) Delete(ctx context.Context, req *expenses.DeleteExpenseRequest) error {
	entities, err := i.ExpenseDAO.Query(ctx,
		query.Equal("id", req.ID),
		query.Limit(1),
	)

	if err != nil {
		return err
	}

	if len(entities) == 0 {
		return nil
	}

	return i.ExpenseDAO.Delete(ctx, entities[0])
}

func entityToExpense(entity *storage.Expense) *expenses.Expense {
	if entity == nil {
		return nil
	}

	return &expenses.Expense{
		ID:       entity.ID,
		Account:  entity.Account,
		Category: entity.Category,
		Currency: entity.Currency,
		Amount:   entity.Amount,
		Date:     entity.Date,
	}
}

func entityFromExpense(item *expenses.Expense) *storage.Expense {
	if item == nil {
		return nil
	}

	return &storage.Expense{
		ID:       item.ID,
		Account:  item.Account,
		Category: item.Category,
		Currency: item.Currency,
		Amount:   item.Amount,
		Date:     item.Date,
	}
}
