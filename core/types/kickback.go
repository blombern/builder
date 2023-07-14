package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type KickbackArgs struct {
	Builder             *common.Address  `json:"builder"`
	BuilderProof        *[]hexutil.Bytes `json:"builderProof"`
	FeeRecipient        *common.Address  `json:"feeRecipient"`
	FeeRecipientProof   *[]hexutil.Bytes `json:"feeRecipientProof"`
	FeePayer            *common.Address  `json:"feePayer"`
	FeePayerProof       *[]hexutil.Bytes `json:"feePayerProof"`
	FeeTransactionIndex hexutil.Bytes    `json:"feeTransactionIndex"`
	FeeTransactionProof *[]hexutil.Bytes `json:"feeTransactionProof"`
	StateRoot           *common.Hash     `json:"stateRoot"`
	TransactionRoot     *common.Hash     `json:"transactionRoot"`
	// ReceiptsRoot               *common.Hash     `json:"receiptsRoot"`
	// FeeTransactionReceiptProof *[]hexutil.Bytes `json:"receiptsProof"`
}
