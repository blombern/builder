package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type KickbackArgs struct {
	Builder            *common.Address  `json:"builder_address"`
	BuilderProof       *[]hexutil.Bytes `json:"builder_proof"`
	FeeRecipient       *common.Address  `json:"fee_recipient_address"`
	FeeRecipientProof  *[]hexutil.Bytes `json:"fee_recipient_proof"`
	FeePayer           *common.Address  `json:"fee_payer_address"`
	FeePayerProof      *[]hexutil.Bytes `json:"fee_payer_proof"`
	PlaceholderTxProof *[]hexutil.Bytes `json:"placeholder_transaction_proof"`
	StateRoot          *common.Hash     `json:"state_root"`
	TransactionRoot    *common.Hash     `json:"transaction_root"`
}
