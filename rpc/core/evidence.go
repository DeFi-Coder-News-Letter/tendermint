package core

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

// ### Query Parameters
//
// | Parameter | Type   | Default | Required | Description                   |
// |-----------+--------+---------+----------+-------------------------------|
// | pubkey    | PubKey | nil     | true     | PubKey of the byzantine actor |
// | vote1     | Vote   | nil     | true     | First vote                    |
// | vote2     | Vote   | nil     | true     | Second vote                   |
func BroadcastDuplicateVote(pubkey crypto.PubKey, vote1 types.Vote, vote2 types.Vote) (*ctypes.ResultBroadcastDuplicateVote, error) {
	chainID := p2pTransport.NodeInfo().Network
	ev := &types.DuplicateVoteEvidence{pubkey, &vote1, &vote2}
	if err := vote1.Verify(chainID, pubkey); err != nil {
		return nil, fmt.Errorf("Error broadcasting evidence, invalid vote1: %v", err)
	}
	if err := vote2.Verify(chainID, pubkey); err != nil {
		return nil, fmt.Errorf("Error broadcasting evidence, invalid vote2: %v", err)
	}

	err := evidencePool.AddEvidence(ev)
	if err != nil {
		return nil, fmt.Errorf("Error broadcasting evidence, adding evidence: %v", err)
	}
	return &ctypes.ResultBroadcastDuplicateVote{ev.Hash()}, nil
}
