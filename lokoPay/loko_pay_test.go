package lokoPay_test

import (
	"encoding/json"
	"fmt"
	"loko-golang/lokoPay"
	"loko-golang/lokoPay/constants"
	"loko-golang/lokoPay/payloads"
	"net/http"
	"testing"
	"time"
)

var lokoPayServ *lokoPay.LokoPay

func init() {
	publishableKey := "xxx"
	secretKey := "xxx"
	liveMode := false
	lokoPayServ = lokoPay.NewLokoPay(
		publishableKey,
		secretKey,
		liveMode,
	)
}

func TestLokoPay_PaymentProcess(t *testing.T) {
	t.Log("payment process")
	customer := payloads.NewCustomer("test-xx-1")
	createPaymentParams := payloads.NewCreatePaymentRequest("10000", "USDC")
	createPaymentParams.SetCustomer(customer)
	//1,creat a payment
	fmt.Println("create payment")
	payment, err := lokoPayServ.Payment().Create(createPaymentParams)
	if err != nil {
		t.Errorf("create payment error=%v", err)
		return
	}
	//2,retrieve the payment for prices
	fmt.Println("retrieve payment for prices")
	startTime := time.Now().Unix()
	for {
		time.Sleep(time.Second * 3)
		payment, err = lokoPayServ.Payment().Retrieve(payment.ID)
		if err != nil {
			t.Errorf("retrieve payment error=%v", err)
			return
		}
		if len(payment.SupportedCryptocurrencies) > 0 {
			break
		}
		if time.Now().Unix()-startTime > 30 {
			t.Error("get price over time")
			return
		}
	}
	var pickedCrypto *payloads.CryptoCurrency
	for _, cryptoCurrency := range payment.SupportedCryptocurrencies {
		if cryptoCurrency.Network == "Immutable zkEVM" && cryptoCurrency.PricePair == "USDC-USDC" {
			pickedCrypto = cryptoCurrency
			break
		}
	}
	if pickedCrypto == nil {
		t.Error("no supported crypto currency")
		return
	}
	//3,confirm payment
	fmt.Println("confirm payment")
	confirmPaymentParam := payloads.NewConfirmPaymentRequest(pickedCrypto)
	payment, err = lokoPayServ.Payment().Confirm(payment.ID, confirmPaymentParam)
	if err != nil {
		t.Errorf("confirm payment error=%v", err)
		return
	}
	//4,retrieve the payment for deposit address
	fmt.Println("retrieve payment for deposit address")
	startTime = time.Now().Unix()
	for {
		time.Sleep(time.Second * 3)
		payment, err = lokoPayServ.Payment().Retrieve(payment.ID)
		if err != nil {
			t.Errorf("retrieve payment error=%v", err)
			return
		}
		if payment.CurrencyDueAddress != "" {
			break
		}
		if time.Now().Unix()-startTime > 30 {
			t.Error("get deposited address over time")
			return
		}
	}
	t.Logf("deposited address : %v", payment.CurrencyDueAddress)
}

func TestLokoPay_PaymentList(t *testing.T) {
	t.Log("payment list")
	queryParams := payloads.NewQueryParam()
	queryParams.SetLimit(3)
	payments, err := lokoPayServ.Payment().List(queryParams)
	if err != nil {
		t.Errorf("retrieve payment error=%v", err)
		return
	}
	paymentsBytes, _ := json.MarshalIndent(payments, "", "  ")
	t.Logf("payments: %v", string(paymentsBytes))
}

func TestLokoPay_PayoutProcess(t *testing.T) {
	t.Log("payout process")
	//1,create a payout
	fmt.Println("create payout")
	customer := payloads.NewCustomer("test-xx-1")
	customer.SetDestinationAddress("0xD15AA5E00971Bf47877145088d3F89b848fA24dA")
	customer.SetDestinationCurrency("USDC")
	customer.SetDestinationNetwork("Immutable zkEVM")
	createPayoutParams := payloads.NewCreatePayoutRequest("10000", "USDC")
	createPayoutParams.SetCustomer(customer)
	createPayoutParams.SetTransferWithNativeToken(true)
	payout, err := lokoPayServ.Payout().Create(createPayoutParams)
	if err != nil {
		t.Errorf("create payout error=%v", err)
		return
	}
	//2,retrieve the payout for fees
	fmt.Println("retrieve payout for fees")
	startTime := time.Now().Unix()
	for {
		time.Sleep(time.Second * 3)
		payout, err = lokoPayServ.Payout().Retrieve(payout.ID)
		if err != nil {
			t.Errorf("retrieve payout error=%v", err)
			return
		}
		if len(payout.DestinationNetworkDetails) > 0 {
			break
		}
		if time.Now().Unix()-startTime > 30 {
			t.Error("get fees over time")
			return
		}
	}
	//3,confirm payout
	fmt.Println("confirm payout")
	confirmPayoutParam := payloads.NewConfirmPayoutRequest(payout.DestinationNetworkDetails)
	payout, err = lokoPayServ.Payout().Confirm(payout.ID, confirmPayoutParam)
	if err != nil {
		t.Errorf("confirm payout error=%v", err)
		return
	}
	//4,retrieve the payout for blockchain transaction details
	fmt.Println("retrieve payout for blockchain transaction details")
	startTime = time.Now().Unix()
	for {
		time.Sleep(time.Second * 3)
		payout, err = lokoPayServ.Payout().Retrieve(payout.ID)
		if err != nil {
			t.Errorf("retrieve payout error=%v", err)
			return
		}
		if len(payout.BlockchainTransactionDetails) > 0 {
			break
		}
		if time.Now().Unix()-startTime > 30 {
			t.Error("get blockchain transaction details over time")
			return
		}
	}
	blockchainTransactionDetailsBytes, _ := json.MarshalIndent(payout.BlockchainTransactionDetails, "", "  ")
	t.Logf("blockchian transaction details : %v", string(blockchainTransactionDetailsBytes))
}

func TestLokoPay_PayoutList(t *testing.T) {
	t.Log("payout list")
	queryParams := payloads.NewQueryParam()
	queryParams.SetLimit(3)
	payouts, err := lokoPayServ.Payout().List(queryParams)
	if err != nil {
		t.Errorf("retrieve payout error=%v", err)
		return
	}
	payoutsBytes, _ := json.MarshalIndent(payouts, "", "  ")
	t.Logf("payouts: %v", string(payoutsBytes))
}

// 查看不同网络的费用
func TestLokoPay_Networkfees(t *testing.T) {
	t.Log("network fees")
	networkfees, err := lokoPayServ.Payout().Networkfees()
	if err != nil {
		t.Errorf("Networkfees() error=%v", err)
		return
	}
	networkFeesBytes, _ := json.MarshalIndent(networkfees, "", "  ")
	t.Logf("Network fees: %v", string(networkFeesBytes))
}

// 查看支持的加密货币
func TestLokoPay_CustomerWallet_GetSupportedCryptocurrencies(t *testing.T) {
	t.Log("customer wallet get supported cryptocurrencies")
	customerWallet, err := lokoPayServ.CustomerWallet().Create(nil)
	if err != nil {
		t.Errorf("CustomerWallet() error=%v", err)
	}
	cryptocurrenciesBytes, _ := json.MarshalIndent(customerWallet.SupportedCryptocurrencies, "", "  ")
	t.Logf("Customer wallet: %v", string(cryptocurrenciesBytes))
}

// 查看/创建customer的钱包地址
func TestLokoPay_CustomerWallet_CreateWalletAddress(t *testing.T) {
	t.Log("customer wallet create wallet address")
	params := &payloads.CreateCustomerWalletRequest{
		Customer: &payloads.Customer{
			ID: "test-xx-1",
		},
		Network:  "Ethereum",
		Currency: "USDC",
	}
	customerWallet, err := lokoPayServ.CustomerWallet().Create(params)
	if err != nil {
		t.Errorf("CustomerWallet() error=%v", err)
	}
	walletAddressesBytes, _ := json.MarshalIndent(customerWallet.WalletAddresses, "", "  ")
	t.Logf("Customer wallet: %v", string(walletAddressesBytes))
}

func TestLokoPay_WebhookEvent(t *testing.T) {
	t.Log("webhook event")
	baseUrl := "https://a966-116-30-101-0.ngrok-free.app"
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		webhookEvent, err := lokoPayServ.WebhookEvent().Retrieve(baseUrl+"/webhook", r)
		if err != nil {
			t.Errorf("Retrieve() error=%v", err)
			return
		}
		switch webhookEvent.Type {
		case constants.WebhookEventTypePaymentDeposited.String(),
			constants.WebhookEventTypePaymentFailed.String(),
			constants.WebhookEventTypePaymentExpired.String():
			payment, err := lokoPayServ.Payment().DecodePayment(string(webhookEvent.Data))
			if err != nil {
				t.Errorf("DecodePayment() error=%v", err)
				return
			}
			paymentBytes, _ := json.MarshalIndent(payment, "", "  ")
			fmt.Println("payment:", string(paymentBytes))
		case constants.WebhookEventTypePayoutPending.String(),
			constants.WebhookEventTypePayoutSucceeded.String(),
			constants.WebhookEventTypePayoutFailed.String():
			payout, err := lokoPayServ.Payout().DecodePayout(string(webhookEvent.Data))
			if err != nil {
				t.Errorf("DecodePayout() error=%v", err)
				return
			}
			payoutBytes, _ := json.MarshalIndent(payout, "", "  ")
			fmt.Println("payout:", string(payoutBytes))
		default:
			t.Errorf("unknown webhook event type")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":0}`))
	})
	port := ":8888"
	if err := http.ListenAndServe(port, nil); err != nil {
		t.Errorf("ListenAndServe() error=%v", err)
	}
}
