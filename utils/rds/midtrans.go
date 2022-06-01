package utils

import (
	"kost/configs"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func NewSnap(*configs.AppConfig) snap.Client {
	status, _ := strconv.Atoi(configs.Get().Payment.MidtransStatus)
	return snap.Client{
		ServerKey:  configs.Get().Payment.MidtransServerKey,
		Env:        midtrans.EnvironmentType(status),
		HttpClient: midtrans.GetHttpClient(midtrans.EnvironmentType(status)),
		Options: &midtrans.ConfigOptions{
			PaymentOverrideNotification: midtrans.PaymentOverrideNotification,
			PaymentAppendNotification:   midtrans.PaymentAppendNotification,
		},
	}
}
