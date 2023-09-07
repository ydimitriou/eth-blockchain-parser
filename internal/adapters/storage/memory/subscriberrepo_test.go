package memory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"
)

func TestNewSubscriberRepository(t *testing.T) {
	tests := []struct {
		name   string
		expRes subscriber.Repository
	}{
		{
			name:   "should return a subscriber repository",
			expRes: SubscriberRepository{subscribers: make(map[string]subscriber.Subscriber)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewSubscriberRepository()
			assert.Equal(t, tt.expRes, mr)
		})
	}
}

func TestSubscriberRepository_Add(t *testing.T) {
	type fields struct {
		subscribers map[string]subscriber.Subscriber
	}
	type args struct {
		subscriber subscriber.Subscriber
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should add subscriber in memory",
			fields: fields{
				subscribers: make(map[string]subscriber.Subscriber),
			},
			args: args{
				subscriber: subscriber.Subscriber{
					Address:      "0x1234",
					Transactions: []transaction.Transaction{{From: "0x1234", To: "0x987"}}},
			},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := SubscriberRepository{subscribers: tt.fields.subscribers}
			err := mr.Add(tt.args.subscriber)
			assert.Equal(t, tt.expErr, err)
			sub, _ := mr.GetByAddress(tt.args.subscriber.Address)
			assert.Equal(t, tt.args.subscriber, *sub)
		})
	}
}

func TestSubscriberRepository_Update(t *testing.T) {
	subAddress := "0x678"
	type fields struct {
		subscribers map[string]subscriber.Subscriber
	}
	type args struct {
		subscriber subscriber.Subscriber
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should update subscriber on success",
			fields: fields{
				subscribers: func() map[string]subscriber.Subscriber {
					mp := make(map[string]subscriber.Subscriber)
					mp[subAddress] = subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{},
					}
					return mp
				}(),
			},
			args: args{
				subscriber: subscriber.Subscriber{
					Address:      subAddress,
					Transactions: []transaction.Transaction{{From: subAddress, To: "0x333"}},
				},
			},
			expErr: nil,
		},
		{
			name: "should return error when address is not a subscriber already",
			fields: fields{
				subscribers: make(map[string]subscriber.Subscriber),
			},
			args: args{
				subscriber: subscriber.Subscriber{Address: subAddress},
			},
			expErr: fmt.Errorf("subscriber with address %v does not exist", subAddress),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := SubscriberRepository{subscribers: tt.fields.subscribers}
			err := mr.Update(tt.args.subscriber)
			assert.Equal(t, tt.expErr, err)
			if err == nil {
				sub, _ := mr.GetByAddress(tt.args.subscriber.Address)
				assert.Equal(t, tt.args.subscriber, *sub)
			}
		})
	}
}

func TestSubscriberRepository_Exist(t *testing.T) {
	subAddress := "0x555"
	type fields struct {
		subscribers map[string]subscriber.Subscriber
	}
	type args struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expRes bool
	}{
		{
			name: "should return true when address is a subscriber",
			fields: fields{
				subscribers: func() map[string]subscriber.Subscriber {
					mp := make(map[string]subscriber.Subscriber)
					mp[subAddress] = subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{},
					}
					return mp
				}(),
			},
			args:   args{address: subAddress},
			expRes: true,
		},
		{
			name:   "should return false when address is a subscriber",
			fields: fields{subscribers: make(map[string]subscriber.Subscriber)},
			args:   args{address: subAddress},
			expRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := SubscriberRepository{subscribers: tt.fields.subscribers}
			res := mr.Exist(tt.args.address)
			assert.Equal(t, tt.expRes, res)
		})
	}
}

func TestSubscriberRepository_GetByAddress(t *testing.T) {
	subAddress := "0x676"
	toAddress := "0x711"
	type fields struct {
		subscribers map[string]subscriber.Subscriber
	}
	type args struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expRes *subscriber.Subscriber
		expErr error
	}{
		{
			name: "should return proper subscriber when address exists in memory",
			fields: fields{
				subscribers: func() map[string]subscriber.Subscriber {
					mp := make(map[string]subscriber.Subscriber)
					mp[subAddress] = subscriber.Subscriber{
						Address:      subAddress,
						Transactions: []transaction.Transaction{{From: subAddress, To: toAddress}},
					}
					return mp
				}(),
			},
			args: args{address: subAddress},
			expRes: &subscriber.Subscriber{
				Address:      subAddress,
				Transactions: []transaction.Transaction{{From: subAddress, To: toAddress}},
			},
			expErr: nil,
		},
		{
			name:   "should return nil when address does not exist in memory",
			fields: fields{subscribers: make(map[string]subscriber.Subscriber)},
			args:   args{address: subAddress},
			expRes: (*subscriber.Subscriber)(nil),
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := SubscriberRepository{subscribers: tt.fields.subscribers}
			res, err := mr.GetByAddress(tt.args.address)
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
