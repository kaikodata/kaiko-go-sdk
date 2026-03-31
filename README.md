<p align="center">
  <img src="https://www.kaiko.com/assets/sites/1/2025/10/kaiko-logo-rgb_color.svg" alt="Kaiko" width="300">
</p>

# kaiko-go-sdk

Go client library for the Kaiko gRPC API. This SDK provides auto-generated gRPC stubs for streaming cryptocurrency market data.

## Installation

```bash
go get github.com/kaikodata/kaiko-go-sdk
```

## Requirements

- Go 1.24+
- A valid Kaiko API key

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	kaikosdk "github.com/kaikodata/kaiko-go-sdk"
	"github.com/kaikodata/kaiko-go-sdk/stream/trades_v1"
	"github.com/kaikodata/kaiko-go-sdk/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer <YOUR_API_KEY>")

	conn, err := grpc.NewClient(
		"gateway-v0-grpc.kaiko.ovh:443",
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

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
		fmt.Printf("Trade: %v\n", resp)
	}
}
```

## Available Services

All services expose a `Subscribe` method that returns a server-side streaming RPC.

| Service | Description |
|---------|-------------|
| `StreamTradesServiceV1` | Real-time trades |
| `StreamMarketUpdateServiceV1` | Market updates |
| `StreamAggregatedQuoteServiceV2` | Aggregated quotes |
| `StreamAggregatedPriceServiceV1` | Aggregated prices |
| `StreamAggregatedStatePriceServiceV1` | Aggregated state prices |
| `StreamAggregatesOHLCVServiceV1` | OHLCV aggregates |
| `StreamAggregatesVWAPServiceV1` | VWAP aggregates |
| `StreamAggregatesSpotExchangeRateV2ServiceV1` | Spot exchange rates |
| `StreamAggregatesSpotDirectExchangeRateV2ServiceV1` | Direct exchange rates |
| `StreamOrderbookL2ServiceV1` | Level 2 order book |
| `StreamOrderbookL2ReplayServiceV1` | Level 2 order book replay |
| `StreamIndexServiceV1` | Index values |
| `StreamIndexMultiAssetsServiceV1` | Multi-asset index values |
| `StreamIndexForexRateServiceV1` | Index forex rates |
| `StreamCompositeIndicesServiceV1` | Composite indices |
| `StreamConstantDurationIndicesServiceV1` | Constant duration indices |
| `StreamExoticIndicesServiceV1` | Exotic indices |
| `StreamDerivativesInstrumentMetricsServiceV1` | Derivatives instrument metrics |
| `StreamIvSviParametersServiceV1` | IV SVI parameters |

## Examples

For more complete examples, see the [kaiko-sdk-examples](https://github.com/kaikodata/kaiko-sdk-examples) repository.
