package memory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
)

func TestNewBlockRepository(t *testing.T) {
	tests := []struct {
		name   string
		expRes block.Repository
	}{
		{
			name:   "should return a block repository",
			expRes: &BlockRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewBlockRepository()
			assert.Equal(t, tt.expRes, mr)
		})
	}
}

func TestBlockRepository_Get(t *testing.T) {
	type fields struct {
		block block.Block
	}
	tests := []struct {
		name   string
		fields fields
		expRes block.Block
		expErr error
	}{
		{
			name: "should return proper block when a block exist",
			fields: fields{
				block: block.Block{Number: "x0253454"},
			},
			expRes: block.Block{Number: "x0253454"},
			expErr: nil,
		},
		{
			name:   "should return proper error when no block exist",
			fields: fields{block: block.Block{}},
			expRes: block.Block{},
			expErr: fmt.Errorf("no block exists in memory"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := BlockRepository{block: tt.fields.block}
			res, err := mr.Get()
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}

func TestBlockRepository_Add(t *testing.T) {
	type fields struct {
		block block.Block
	}
	type args struct {
		block block.Block
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name:   "should add block in memory",
			fields: fields{block: block.Block{}},
			args:   args{block: block.Block{Number: "0x9922"}},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := BlockRepository{block: tt.fields.block}
			err := mr.Add(tt.args.block)
			assert.Equal(t, tt.expErr, err)
			b, _ := mr.Get()
			assert.Equal(t, tt.args.block, b)
		})
	}
}
