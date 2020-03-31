package expenses

import "context"

// ExpenseService is a CRUD interface for a Expense service.
//go:generate gokit service -interface ExpenseService
type ExpenseService interface {
	//gokit: method:"GET" path:"/expenses/{ID}"
	Get(ctx context.Context, req *GetExpenseRequest) (*GetExpenseResponse, error)

	//gokit: method:"GET" path:"/expenses"
	List(ctx context.Context, req *ListExpenseRequest) (*ListExpenseResponse, error)

	//gokit: method:"POST" path:"/expenses"
	Create(ctx context.Context, req *CreateExpenseRequest) (*CreateExpenseResponse, error)

	//gokit: method:"PUT" path:"/expenses/{ID}"
	Update(ctx context.Context, req *UpdateExpenseRequest) (*UpdateExpenseResponse, error)

	//gokit: method:"DELETE" path:"/expenses/{ID}"
	Delete(ctx context.Context, req *DeleteExpenseRequest) error
}

// GetExpenseRequest ...
type GetExpenseRequest struct {
	ID string `json:"id"`
}

// GetExpenseResponse ...
type GetExpenseResponse struct {
	Expense *Expense `json:"expense"`
}

// ListExpenseRequest ...
type ListExpenseRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

// ListExpenseResponse ...
type ListExpenseResponse struct {
	Expenses []*Expense `json:"expenses"`
}

// CreateExpenseRequest ...
type CreateExpenseRequest struct {
	Expense *Expense `json:"expense"`
}

// CreateExpenseResponse ...
type CreateExpenseResponse struct {
	Expense *Expense `json:"expense"`
}

// UpdateExpenseRequest ...
type UpdateExpenseRequest struct {
	Expense *Expense `json:"expense"`
}

// UpdateExpenseResponse ...
type UpdateExpenseResponse struct {
	Expense *Expense `json:"expense"`
}

// DeleteExpenseRequest ...
type DeleteExpenseRequest struct {
	ID string `json:"id"`
}
