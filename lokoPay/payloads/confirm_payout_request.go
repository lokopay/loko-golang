package payloads

type ConfirmPayoutRequest struct {
	DestinationNetworkDetails []*DestinationNetworkDetail `json:"destination_network_details"`
}

func NewConfirmPayoutRequest(destinationNetworkDetails []*DestinationNetworkDetail) *ConfirmPayoutRequest {
	return &ConfirmPayoutRequest{
		DestinationNetworkDetails: destinationNetworkDetails,
	}
}
