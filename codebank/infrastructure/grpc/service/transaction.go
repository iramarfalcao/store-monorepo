package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iramarfalcao/codebank/dto"
	"github.com/iramarfalcao/codebank/infrastructure/grpc/pb"
	"github.com/iramarfalcao/codebank/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            in.CreditCard.Name,
		Number:          in.CreditCard.Number,
		ExpirationMonth: in.CreditCard.ExpirationMonth,
		ExpirationYear:  in.CreditCard.ExpirationYear,
		CVV:             in.CreditCard.Cvv,
		Amount:          in.Amount,
		Store:           in.Store,
		Description:     in.Description,
	}

	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDto)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "transaction rejected by the bank")
	}
	return &empty.Empty{}, nil
}
