# Kaiko Go SDK

Official Go client library for [Kaiko](https://www.kaiko.com/) gRPC streaming APIs. Provides real-time and historical cryptocurrency market data including trades, OHLCV, VWAP, order books, indices, derivatives metrics, and more.

## Installation

```bash
go get github.com/kaikodata/kaiko-go-sdk
```

Requires Go 1.22+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	kaikosdk "github.com/kaikodata/kaiko-go-sdk"
	"github.com/kaikodata/kaiko-go-sdk/core"
	"github.com/kaikodata/kaiko-go-sdk/stream/trades_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Connect to Kaiko gateway
	conn, err := grpc.NewClient(
		"gateway-v0-grpc.kaiko.ovh:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	// Authenticate with your API key
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer <YOUR_API_KEY>")

	// Subscribe to BTC/USD trades on Coinbase
	client := kaikosdk.NewStreamTradesServiceV1Client(conn)
	stream, err := client.Subscribe(ctx, &trades_v1.StreamTradesRequestV1{
		InstrumentCriteria: &core.InstrumentCriteria{
			Exchange:       "cbse",
			InstrumentClass: "spot",
			Code:           "btc-usd",
		},
	})
	if err != nil {
		log.Fatalf("subscribe failed: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		fmt.Printf("Trade: %+v\n", resp)
	}
}
```

## Available Streaming Services

| Service | Description |
|---------|-------------|
| `StreamTradesServiceV1` | Real-time and historical trades |
| `StreamAggregatesOHLCVServiceV1` | OHLCV candles |
| `StreamAggregatesVWAPServiceV1` | Volume-weighted average price |
| `StreamAggregatedQuoteServiceV2` | Aggregated quotes |
| `StreamAggregatedPriceServiceV1` | Aggregated prices |
| `StreamAggregatedStatePriceServiceV1` | Aggregated state prices |
| `StreamAggregatesSpotExchangeRateV2ServiceV1` | Spot exchange rates |
| `StreamAggregatesSpotDirectExchangeRateV2ServiceV1` | Direct exchange rates |
| `StreamMarketUpdateServiceV1` | Market updates (trades, top-of-book) |
| `StreamOrderbookL2ServiceV1` | Level 2 order book snapshots |
| `StreamOrderbookL2ReplayServiceV1` | Level 2 order book replay |
| `StreamIndexServiceV1` | Kaiko indices |
| `StreamIndexMultiAssetsServiceV1` | Multi-asset indices |
| `StreamIndexForexRateServiceV1` | Forex rate indices |
| `StreamExoticIndicesServiceV1` | Exotic indices |
| `StreamCustomDurationIndicesServiceV1` | Custom duration indices |
| `StreamDerivativesInstrumentMetricsServiceV1` | Derivatives instrument metrics |
| `StreamIvSviParametersServiceV1` | Implied volatility SVI parameters |

All services follow the same `Subscribe` pattern shown in the Quick Start example.

## Package Structure

```
kaikosdk               # Root package — gRPC service clients and servers
├── core/              # Shared types (InstrumentCriteria, DataInterval, Assets, etc.)
└── stream/            # Request/response types per service
    ├── trades_v1/
    ├── aggregates_ohlcv_v1/
    ├── aggregates_vwap_v1/
    ├── orderbookl2_v1/
    └── ...
```

## Authentication

All API calls require a valid Kaiko API key passed via gRPC metadata:

```go
ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer <YOUR_API_KEY>")
```

## Code Generation

This SDK is generated from Kaiko's protobuf definitions using `protoc` with `protoc-gen-go` and `protoc-gen-go-grpc`. The `.pb.go` files should not be edited manually.

## License

See [LICENSE](LICENSE) for details.
