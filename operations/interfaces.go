package operations

import (
	"github.com/DefinitelyNotAGoat/go-tezos/account"
	"github.com/DefinitelyNotAGoat/go-tezos/delegate"
)

type TezosOperationsService interface {
	CreateBatchPaymentForFirstSend(payments []delegate.Payment, wallet account.Wallet, paymentFee int, gasLimit int, batchSize int) ([]string, error)
	CreateBatchPayment(payments []delegate.Payment, wallet account.Wallet, paymentFee int, gaslimit int, batchSize int) ([]string, error)
	InjectOperation(op string) ([]byte, error)
	GetBlockOperationHashes(id interface{}) ([]string, error)
}
