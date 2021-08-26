package atomicassets

// GetCollections response
type CollectionsProps struct {
	Success   bool                   `json:"success"`
	Data      []CollectionsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                  `json:"query_time"`
	Message   string                 `json:"message,omitempty"` // error message will be here
}

// GetCollectionName response
type CollectionNameProps struct {
	Success   bool                 `json:"success"`
	Data      CollectionsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                `json:"query_time"`
	Message   string               `json:"message,omitempty"` // error message will be here
}

// GetCollectionNameStats response
type CollectionNameStatsProps struct {
	Success   bool                      `json:"success"`
	Data      CollectionsStatsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                     `json:"query_time"`
	Message   string                    `json:"message,omitempty"` // error message will be here
}

// GetCollectionNameLogs response
type CollectionNameLogsProps AssetsIDLogsProps

type CollectionsDataProps struct {
	Contract           string   `json:"contract"`
	CollectionName     string   `json:"collection_name"`
	Name               string   `json:"name"`
	Author             string   `json:"author"`
	AllowNotify        bool     `json:"allow_notify"`
	AuthorizedAccounts []string `json:"authorized_accounts"`
	NotifyAccounts     []string `json:"notify_accounts"`
	MarketFee          float64  `json:"market_fee"`
	Data               Data     `json:"data"`
	CreatedAtBlock     string   `json:"created_at_block"`
	CreatedAtTime      string   `json:"created_at_time"`
}

type CollectionsStatsDataProps struct {
	Assets    string `json:"assets"`
	Burned    string `json:"burned"`
	Templates string `json:"templates"`
	Schemas   string `json:"schemas"`
}
