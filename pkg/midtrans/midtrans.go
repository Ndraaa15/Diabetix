package midtrans

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Midtrans struct {
	midtrans snap.Client
}

func NewMidtrans(conf *env.Env) *Midtrans {
	snapClient := snap.Client{}
	snapClient.New(conf.MidtransAPIKey, midtrans.Sandbox)

	return &Midtrans{
		midtrans: snapClient,
	}
}

func (m *Midtrans) CreateTransaction(ctx context.Context, orderID string, amount float64) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := m.midtrans.CreateTransaction(req)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to create midtrans transaction").
			WithError(err)
	}

	return snapResp.RedirectURL, nil
}
