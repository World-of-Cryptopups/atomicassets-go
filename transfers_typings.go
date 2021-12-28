package atomicassets

type TransfersProps struct {
	Success   bool            `json:"success"`
	Data      []TransfersData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64           `json:"query_time"`
	Message   string          `json:"message,omitempty"` // error message will be here
}

type TransfersData struct {
	Contract       string            `json:"contract"`
	TransferID     string            `json:"transfer_id"`
	SenderName     string            `json:"sender_name"`
	RecipientName  string            `json:"recipient_name"`
	Memo           string            `json:"memo"`
	Assets         []AssetsDataProps `json:"assets"`
	CreatedAtBlock string            `json:"created_at_block"`
	CreatedAtTime  string            `json:"created_at_time"`
}
