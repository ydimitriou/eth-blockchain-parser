package queries

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"
)

func TestGetSubscriberHandler(t *testing.T) {
	type args struct {
		repo subscriber.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes GetSubscriberHandler
	}{
		{
			name:   "should return a new GetSubscriberHandler",
			args:   args{repo: &subscriber.MockRepository{}},
			expRes: getSubscriberHandler{repo: &subscriber.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewGetSubscriberHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestGetSubscriberHandler_Handle(t *testing.T) {
	subAddress := "0x005"
	toAddress := "0x006"
	type fields struct {
		repo subscriber.Repository
	}
	type args struct {
		req GetSubscriberRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expRes *GetSubscriberResult
		expErr error
	}{
		{
			name: "should return subscriber when address exists",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					sub := subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{{From: subAddress, To: toAddress}},
					}
					mr.On("GetByAddress", subAddress).Return(&sub, nil)
					return &mr
				}(),
			},
			args: args{req: GetSubscriberRequest{Address: subAddress}},
			expRes: &GetSubscriberResult{
				Address:      subAddress,
				Transactions: []TransactionResult{{From: subAddress, To: toAddress}},
			},
			expErr: nil,
		},
		{
			name: "should return error when repo responds with error",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					mr.On("GetByAddress", subAddress).Return((*subscriber.Subscriber)(nil), errors.New("repository error"))
					return &mr
				}(),
			},
			args:   args{req: GetSubscriberRequest{Address: subAddress}},
			expRes: (*GetSubscriberResult)(nil),
			expErr: errors.New("repository error"),
		},
		{
			name: "should return nil result and nil response when no error and no subs exist",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					mr.On("GetByAddress", subAddress).Return((*subscriber.Subscriber)(nil), nil)
					return &mr
				}(),
			},
			args:   args{req: GetSubscriberRequest{Address: subAddress}},
			expRes: (*GetSubscriberResult)(nil),
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getSubscriberHandler{repo: tt.fields.repo}
			res, err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
