package queries

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
)

func TestNewGetBlockHandler(t *testing.T) {
	type args struct {
		repo block.Repository
		hp   hex.Provider
	}
	tests := []struct {
		name   string
		args   args
		expRes GetBlockHandler
	}{
		{
			name:   "should return a new GetBlockHandler",
			args:   args{repo: &block.MockRepository{}, hp: &hex.MockProvider{}},
			expRes: getBlockHandler{repo: &block.MockRepository{}, hexProvider: &hex.MockProvider{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewGetBlockHandler(tt.args.repo, tt.args.hp)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestNewGetBlockHandler_Handle(t *testing.T) {
	blockNumber := "0x3939"
	num := big.NewInt(125)
	type fields struct {
		hexProvider hex.Provider
		repo        block.Repository
	}
	tests := []struct {
		name   string
		fields fields
		expRes *GetBlockResult
		expErr error
	}{
		{
			name: "should return a GetBlockResult when repo returs a block",
			fields: fields{
				repo: func() *block.MockRepository {
					b := block.Block{Number: blockNumber}
					mr := block.MockRepository{}
					mr.On("Get").Return(b, nil)
					return &mr
				}(),
				hexProvider: func() *hex.MockProvider {
					hp := hex.MockProvider{}
					hp.On("HexToInt", blockNumber).Return(num)
					return &hp
				}(),
			},
			expRes: &GetBlockResult{Number: 125},
			expErr: nil,
		},
		{
			name: "should return error when repo return error",
			fields: fields{
				repo: func() *block.MockRepository {
					b := block.Block{}
					mr := block.MockRepository{}
					mr.On("Get").Return(b, errors.New("repository error"))
					return &mr
				}(),
				hexProvider: &hex.MockProvider{},
			},
			expRes: (*GetBlockResult)(nil),
			expErr: errors.New("repository error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getBlockHandler{
				repo:        tt.fields.repo,
				hexProvider: tt.fields.hexProvider,
			}
			res, err := h.Handle()
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
