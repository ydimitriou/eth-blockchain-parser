package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"
)

func TestUpdateSubscriberHandler(t *testing.T) {
	type args struct {
		repo subscriber.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes UpdateSubscriberHandler
	}{
		{
			name:   "should return a new UpdateSubscriberHandler",
			args:   args{repo: &subscriber.MockRepository{}},
			expRes: updateSubscriberHandler{repo: &subscriber.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewUpdateSubscriberHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestUpdateSubscriberHandler_Handle(t *testing.T) {
	subAddress := "0x7444"
	toAddress := "0x910"
	type fields struct {
		repo subscriber.Repository
	}
	type args struct {
		req UpdateSubscriberRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should return nil on success",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					getRes := subscriber.Subscriber{Address: subAddress}
					updateReq := subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{{From: subAddress, To: toAddress}},
					}
					mr.On("GetByAddress", subAddress).Return(&getRes, nil)
					mr.On("Update", updateReq).Return(nil)
					return &mr
				}(),
			},
			args: args{req: UpdateSubscriberRequest{
				Address:     subAddress,
				Transaction: TransactionsRequest{From: subAddress, To: toAddress},
			}},
			expErr: nil,
		},
		{
			name: "should return error when repo GetByAddress returns error",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					sub := &subscriber.Subscriber{}
					mr.On("GetByAddress", subAddress).Return(sub, errors.New("repository error"))
					return &mr
				}(),
			},
			args:   args{req: UpdateSubscriberRequest{Address: subAddress}},
			expErr: errors.New("repository error"),
		},
		{
			name: "should return error when repo Update returns error",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					getRes := subscriber.Subscriber{Address: subAddress, Transactions: []transaction.Transaction{}}
					updateReq := subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{{From: subAddress, To: toAddress}},
					}
					mr.On("GetByAddress", subAddress).Return(&getRes, nil)
					mr.On("Update", updateReq).Return(errors.New("repository error"))
					return &mr
				}(),
			},
			args: args{req: UpdateSubscriberRequest{
				Address:     subAddress,
				Transaction: TransactionsRequest{From: subAddress, To: toAddress},
			}},
			expErr: errors.New("repository error"),
		},
		{
			name: "should return proper error when repo GetByAddress return nil subscriber",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					mr.On("GetByAddress", subAddress).Return((*subscriber.Subscriber)(nil), nil)

					return &mr
				}(),
			},
			args:   args{req: UpdateSubscriberRequest{Address: subAddress}},
			expErr: errors.New("update failed, subscriber with adress 0x7444 does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateSubscriberHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
