package briq

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type (
	UsersResponse struct {
		Response
		Users []User
	}

	TransactionResponse struct {
		Response
		Transaction *Transaction
	}

	TransactionsResponse struct {
		Response
		Transactions []Transaction
	}
)

func (b Briq) Users(p Pagination) UsersResponse {
	var users = make([]User, 0)

	res := b.request(http.MethodGet, b.uri(endpointUsers, b.key()), &p, nil, &users)

	return UsersResponse{Response: res, Users: users}
}

func (b Briq) Transactions(p Pagination) TransactionsResponse {
	var transactions = make([]Transaction, 0)

	res := b.request(http.MethodGet, b.uri(endpointTransactions, b.Key), &p, nil, &transactions)

	return TransactionsResponse{Response: res, Transactions: transactions}
}

func (b Briq) DoTransaction(t TransactionInput) TransactionResponse {
	var transaction = new(Transaction)

	bs, err := json.Marshal(t)
	if err != nil {
		return TransactionResponse{Response: Response{Error: err}}
	}

	res := b.request(http.MethodPost, b.uri(endpointTransactions, b.Key), nil, bytes.NewBuffer(bs), transaction)

	return TransactionResponse{Response: res, Transaction: transaction}
}
