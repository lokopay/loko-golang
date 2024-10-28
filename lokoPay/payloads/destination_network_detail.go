package payloads

type DestinationNetworkDetail struct {
	DestinationAmount                   string                   `json:"destination_amount,omitempty"`
	DestinationCurrency                 string                   `json:"destination_currency,omitempty"`
	DestinationNetwork                  string                   `json:"destination_network,omitempty"`
	DestinationNetworkDescription       string                   `json:"destination_network_description,omitempty"`
	DestinationNetworkFee               string                   `json:"destination_network_fee,omitempty"`
	DestinationNetworkFeeCurrency       string                   `json:"destination_network_fee_currency,omitempty"`
	DestinationNetworkFeeMonetary       string                   `json:"destination_network_fee_monetary,omitempty"`
	DestinationTransactionFeeCurrency   string                   `json:"destination_transaction_fee_currency,omitempty"`
	DestinationTransactionFeeFixed      string                   `json:"destination_transaction_fee_fixed,omitempty"`
	DestinationTransactionFeePercentage string                   `json:"destination_transaction_fee_percentage,omitempty"`
	DestinationTransactionFeeType       string                   `json:"destination_transaction_fee_type,omitempty"`
	ID                                  string                   `json:"id,omitempty"`
	TransferWithNativeToken             *TransferWithNativeToken `json:"transfer_with_native_token,omitempty"`
}
