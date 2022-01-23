package scb

// SCB: QR Payment
// API Specification for Payment Confirmation (v1.0.2.2)
import (
	"encoding/json"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-proxy/pkg/datamodels"
	utils "github.com/zercle/thai-cross-bank-proxy/utils"
)

type Tag30Req struct {
	EventCode              string      `json:"eventCode"`
	TransactionType        string      `json:"transactionType"`
	ReverseFlag            string      `json:"reverseFlag"`
	PayeeProxyId           string      `json:"payeeProxyId"`
	PayeeProxyType         string      `json:"payeeProxyType"`
	PayeeAccountNumber     string      `json:"payeeAccountNumber"`
	PayeeName              string      `json:"payeeName"`
	PayerProxyId           string      `json:"payerProxyId"`
	PayerProxyType         string      `json:"payerProxyType"`
	PayerAccountNumber     string      `json:"payerAccountNumber"`
	PayerName              string      `json:"payerName"`
	SendingBankCode        string      `json:"sendingBankCode"`
	ReceivingBankCode      string      `json:"receivingBankCode"`
	Amount                 json.Number `json:"amount"`
	TransactionId          string      `json:"transactionId"`
	FastEasySlipNumber     string      `json:"fastEasySlipNumber"`
	TransactionDateAndTime string      `json:"transactionDateandTime"`
	BillPaymentRef1        string      `json:"billPaymentRef1"`
	BillPaymentRef2        string      `json:"billPaymentRef2"`
	BillPaymentRef3        string      `json:"billPaymentRef3"`
	CurrencyCode           string      `json:"currencyCode"`
	EquivalentAmount       json.Number `json:"equivalentAmount"`
	EquivalentCurrencyCode string      `json:"equivalentCurrencyCode"`
	ExchangeRate           string      `json:"exchangeRate"`
	ChannelCode            string      `json:"channelCode"`
	PartnerTransactionId   string      `json:"partnerTransactionId"`
	TepaCode               string      `json:"tepaCode"`
}

func (b Tag30Req) ToTransaction(result datamodels.Transaction, err error) {
	// convert SCB time format into RFC3339
	transactionTime, err := time.Parse(b.TransactionDateAndTime, utils.RFC3339Mili)

	if err != nil {
		log.Printf("%+v", err)
	}

	// check amount
	_, err = b.Amount.Float64()
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	result = datamodels.Transaction{
		PayeeBankCode: "014", // SCB
		PayeeAcc:      b.PayeeAccountNumber,
		PayeeName:     b.PayeeName,
		PayeePID:      b.PayeeProxyId,
		PayerBankCode: b.SendingBankCode,
		PayerAcc:      b.PayerAccountNumber,
		PayerName:     b.PayerName,
		PayerPID:      b.PayerProxyId,
		Reference1:    b.BillPaymentRef1,
		Reference2:    b.BillPaymentRef2,
		Reference3:    b.BillPaymentRef3,
		Amount:        b.Amount,
		CurrencyCode:  b.CurrencyCode,
		Channel:       b.ChannelCode,
		TxRef:         b.TransactionId,
		TxDateTime:    transactionTime,
	}

	return
}

// https://developer.scb/#/documents/api-reference-index/references/generic-response-codes.html
type Tag30Resp struct {
	ResCode       string `json:"resCode"`
	ResDesc       string `json:"resDesc"`
	TransactionId string `json:"transactionId"`
	ConfirmId     string `json:"confirmId"`
}
