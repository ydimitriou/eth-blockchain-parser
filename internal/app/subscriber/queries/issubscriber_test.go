package queries

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

func TestIsSubscriberHandler(t *testing.T) {
	type args struct {
		repo subscriber.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes IsSubscriberHandler
	}{
		{
			name:   "should return a new IsSubscriberHandler",
			args:   args{repo: &subscriber.MockRepository{}},
			expRes: isSubscriberHandler{repo: &subscriber.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewIsSubscriberHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestIsSubscriberHandler_Handle(t *testing.T) {
	subAddress := "0x444"
	type fields struct {
		repo subscriber.Repository
	}
	type args struct {
		req IsSubscriberRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expRes bool
	}{
		{
			name: "should return true when subscriber exist",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					mr.On("Exist", subAddress).Return(true)
					return &mr
				}(),
			},
			args:   args{req: IsSubscriberRequest{Address: subAddress}},
			expRes: true,
		},
		{
			name: "should return false when subscriber not exist",
			fields: fields{
				repo: func() *subscriber.MockRepository {
					mr := subscriber.MockRepository{}
					mr.On("Exist", subAddress).Return(false)
					return &mr
				}(),
			},
			args:   args{req: IsSubscriberRequest{Address: subAddress}},
			expRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := isSubscriberHandler{repo: tt.fields.repo}
			res := h.Handle(tt.args.req)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
