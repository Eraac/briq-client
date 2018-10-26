package briq

import "net/http"

type (
	UsersResponse struct {
		Page  Link
		Users []User
	}

	TransactionsResponse struct {
		Page         Link
		Transactions []Transaction
	}
)

func (b Briq) Users(p Pagination) (*UsersResponse, error) {
	var users = make([]User, 0)

	l, err := b.request(http.MethodGet, b.uri(endpointUsers, b.key()), p, nil, &users)
	if err != nil {
		return nil, err
	}

	return &UsersResponse{Page: *l, Users: users}, nil
}

func (b Briq) Transactions(p Pagination) (*TransactionsResponse, error) {
	var transactions = make([]Transaction, 0)

	l, err := b.request(http.MethodGet, b.uri(endpointTransactions, b.Key), p, nil, &transactions)
	if err != nil {
		return nil, err
	}

	return &TransactionsResponse{Page: *l, Transactions: transactions}, nil
}
