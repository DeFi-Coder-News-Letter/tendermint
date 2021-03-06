package conn

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/torusresearch/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
