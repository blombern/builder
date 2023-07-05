package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type KickbackArgs struct {
	FeeRecipient               *common.Address  `json:"feeRecipient"`
	FeeRecipientProof          *[]hexutil.Bytes `json:"feeRecipientProof"`
	FeePayer                   *common.Address  `json:"feePayer"`
	FeePayerProof              *[]hexutil.Bytes `json:"feePayerProof"`
	FeeTransactionIndex        hexutil.Bytes    `json:"feeTransactionIndex"`
	FeeTransactionProof        *[]hexutil.Bytes `json:"feeTransactionProof"`
	FeeTransactionReceiptProof *[]hexutil.Bytes `json:"receiptsProof"`
}
