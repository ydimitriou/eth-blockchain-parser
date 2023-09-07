package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
)

func TestAddBlockHandler(t *testing.T) {
	type args struct {
		repo block.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes AddBlockHandler
	}{
		{
			name:   "should return a new AddBlockHandler",
			args:   args{repo: &block.MockRepository{}},
			expRes: addBlockHandler{repo: &block.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewAddBlockHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestAddBlockHandler_Handle(t *testing.T) {
	blockNumber := "0x1212"
	type fields struct {
		repo block.Repository
	}
	type args struct {
		req AddBlockRequest
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
				repo: func() *block.MockRepository {
					b := block.Block{Number: blockNumber}
					mr := block.MockRepository{}
					mr.On("Add", b).Return(nil)
					return &mr
				}(),
			},
			args:   args{req: AddBlockRequest{Number: blockNumber}},
			expErr: nil,
		},
		{
			name: "should return error when repo returns error",
			fields: fields{
				repo: func() *block.MockRepository {
					b := block.Block{Number: blockNumber}
					mr := block.MockRepository{}
					mr.On("Add", b).Return(errors.New("repository error"))
					return &mr
				}(),
			},
			args:   args{req: AddBlockRequest{Number: blockNumber}},
			expErr: errors.New("repository error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addBlockHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
