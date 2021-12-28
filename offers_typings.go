package atomicassets

type OffersProps struct {
	Success   bool         `json:"success"`
	Data      []OffersData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64        `json:"query_time"`
	Message   string       `json:"message,omitempty"` // error message will be here
}

type OffersData struct {
	Contract            string            `json:"contract"`
	OfferID             string            `json:"offer_id"`
	SenderName          string            `json:"sender_name"`
	RecipientName       string            `json:"recipient_name"`
	Memo                string            `json:"memo"`
	State               int64             `json:"state"`
	IsSenderContract    bool              `json:"is_sender_contract"`
	IsRecipientContract bool              `json:"is_recipient_contract"`
	SenderAssets        []AssetsDataProps `json:"sender_assets"`
	RecipientAssets     []AssetsDataProps `json:"recipient_assets"`
	UpdatedAtBlock      string            `json:"updated_at_block"`
	UpdatedAtTime       string            `json:"updated_at_time"`
	CreatedAtBlock      string            `json:"created_at_block"`
	CreatedAtTime       string            `json:"created_at_time"`
}

type OfferIDLogsProps AssetsIDLogsProps
