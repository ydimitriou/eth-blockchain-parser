package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

func TestAddSubscriberHandler(t *testing.T) {
	type args struct {
		repo subscriber.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes AddSubscriberHandler
	}{
		{
			name:   "should return a new AddSubscriberHandler",
			args:   args{repo: &subscriber.MockRepository{}},
			expRes: addSubscriberHandler{repo: &subscriber.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAddSubscriberHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestAddSubscriberHandler_Handle(t *testing.T) {
	subAdress := "0x100"
	type fields struct {
		repo subscriber.Repository
	}
	type args struct {
		req AddSubscriberRequest
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
					sub := subscriber.Subscriber{Address: subAdress}
					mr := subscriber.MockRepository{}
					mr.On("Add", sub).Return(nil)
					return &mr
				}(),
			},
			args:   args{req: AddSubscriberRequest{Address: subAdress}},
			expErr: nil,
		},
		{
			name: "should return error when repo returns error",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					sub := subscriber.Subscriber{Address: subAdress}
					mr := subscriber.MockRepository{}
					mr.On("Add", sub).Return(errors.New("repository error"))
					return &mr
				}(),
			},
			args:   args{req: AddSubscriberRequest{Address: subAdress}},
			expErr: errors.New("repository error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addSubscriberHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
