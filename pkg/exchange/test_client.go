package exchange

import (
	"context"
	"math/rand"

	"github.com/thrasher-corp/gocryptotrader/gctrpc"
	"google.golang.org/grpc"
)

const (
	testExchangeName = "test-exchange"
	testPriceValue   = 12.
	testAmountValue  = 0.001
)

type testClient struct{}

// GetTicker ...
func (m *testClient) GetTicker(ctx context.Context, in *gctrpc.GetTickerRequest, opts ...grpc.CallOption) (*gctrpc.TickerResponse, error) {
	return &gctrpc.TickerResponse{
		Pair: &gctrpc.CurrencyPair{
			Delimiter: "-",
			Base:      "BTC",
			Quote:     "USD",
		},
		Bid: testPriceValue,
		Ask: testPriceValue,
	}, nil
}

// UpdateAccountInfo ...
func (m *testClient) UpdateAccountInfo(ctx context.Context, in *gctrpc.GetAccountInfoRequest, opts ...grpc.CallOption) (*gctrpc.GetAccountInfoResponse, error) {
	return &gctrpc.GetAccountInfoResponse{
		Exchange: testExchangeName,
		Accounts: []*gctrpc.Account{
			{
				Currencies: []*gctrpc.AccountCurrencyInfo{
					{
						Currency:   "BTC",
						TotalValue: testPriceValue,
					},
					{
						Currency:   "USD",
						TotalValue: testPriceValue,
					},
				},
			},
		},
	}, nil
}

// SubmitOrder ...
func (m *testClient) SubmitOrder(ctx context.Context, in *gctrpc.SubmitOrderRequest, opts ...grpc.CallOption) (*gctrpc.SubmitOrderResponse, error) {
	return &gctrpc.SubmitOrderResponse{
		OrderId:     "1",
		OrderPlaced: true,
	}, nil
}

// GetOrderbook ...
func (m *testClient) GetOrderbook(ctx context.Context, in *gctrpc.GetOrderbookRequest, opts ...grpc.CallOption) (*gctrpc.OrderbookResponse, error) {
	bidPrice := 20000 + rand.Float64()*(30000-20000)
	askPrice := 20000 + rand.Float64()*(30000-20000)
	return &gctrpc.OrderbookResponse{
		Asks: []*gctrpc.OrderbookItem{
			{
				Price:  askPrice - 100,
				Amount: rand.Float64(),
			},
			{
				Price:  askPrice,
				Amount: rand.Float64(),
			},
		},
		Bids: []*gctrpc.OrderbookItem{
			{
				Price:  bidPrice - 100,
				Amount: rand.Float64(),
			},
			{
				Price:  bidPrice,
				Amount: rand.Float64(),
			},
		},
	}, nil
}

// Close ...
func (m *testClient) Close() error { return nil }

// SetupTestGRPCClient ...
func SetupTestGRPCClient() {
	setupClient = func() (grpcClient, error) {
		return &testClient{}, nil
	}
}
