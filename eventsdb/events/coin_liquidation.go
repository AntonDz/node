package events

import (
	"encoding/json"
	"github.com/PillarDevelopment/noax-go-node/core/types"
)

type CoinLiquidationEvent struct {
	Coin types.CoinSymbol
}

func (e CoinLiquidationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Coin string `json:"coin"`
	}{
		Coin: e.Coin.String(),
	})
}
