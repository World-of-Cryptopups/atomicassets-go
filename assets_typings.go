package atomicassets

// GetAssets response
type AssetsProps struct {
	Success   bool              `json:"success"`
	Data      []AssetsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64             `json:"query_time"`
	Message   string            `json:"message,omitempty"` // error message will be here
}

// GetAssetID response
type AssetsIDProps struct {
	Success   bool            `json:"success"`
	Data      AssetsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64           `json:"query_time"`
	Message   string          `json:"message,omitempty"` // error message will be here
}

// GetAssetIDStats response
type AssetsIDStatsProps struct {
	Success   bool                 `json:"success"`
	Data      AssetsStatsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                `json:"query_time"`
	Message   string               `json:"message,omitempty"` // error message will be here

}

// GetAssetIDLogs response
type AssetsIDLogsProps struct {
	Success   bool           `json:"success"`
	Data      []LogDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64          `json:"query_time"`
	Message   string         `json:"message,omitempty"` // error message will be here
}

type LogDataProps struct {
	LogID          string `json:"log_id"`
	Name           string `json:"name"`
	Data           Data   `json:"data"`
	TXID           string `json:"txid"`
	CreatedAtBlock string `json:"created_at_block"`
	CreatedAtTime  string `json:"created_at_time"`
}

type AssetsStatsDataProps struct {
	TemplateMint string `json:"template_mint"`
}

type AssetsDataProps struct {
	Contract           string        `json:"contract"`
	AssetID            string        `json:"asset_id"`
	Owner              string        `json:"owner"`
	IsTransferable     bool          `json:"is_transferable"`
	IsBurnable         bool          `json:"is_burnable"`
	Collection         Collection    `json:"collection"`
	Schema             Schema        `json:"schema"`
	Template           Template      `json:"template"`
	MutableData        MutableData   `json:"mutable_data"`
	ImmutableData      MutableData   `json:"immutable_data"`
	TemplateMint       string        `json:"template_mint"`
	BackedTokens       []interface{} `json:"backed_tokens"`
	BurnedByAccount    interface{}   `json:"burned_by_account"`
	BurnedAtBlock      interface{}   `json:"burned_at_block"`
	BurnedAtTime       interface{}   `json:"burned_at_time"`
	UpdatedAtBlock     string        `json:"updated_at_block"`
	UpdatedAtTime      string        `json:"updated_at_time"`
	TransferredAtBlock string        `json:"transferred_at_block"`
	TransferredAtTime  string        `json:"transferred_at_time"`
	MintedAtBlock      string        `json:"minted_at_block"`
	MintedAtTime       string        `json:"minted_at_time"`
	Data               Data          `json:"data"`
	Name               string        `json:"name"`
}

type Collection struct {
	CollectionName     string        `json:"collection_name"`
	Name               string        `json:"name"`
	Img                string        `json:"img"`
	Author             string        `json:"author"`
	AllowNotify        bool          `json:"allow_notify"`
	AuthorizedAccounts []string      `json:"authorized_accounts"`
	NotifyAccounts     []interface{} `json:"notify_accounts"`
	MarketFee          float64       `json:"market_fee"`
	CreatedAtBlock     string        `json:"created_at_block"`
	CreatedAtTime      string        `json:"created_at_time"`
}

type Data map[string]interface{}

type MutableData map[string]interface{}

type Schema struct {
	SchemaName     string   `json:"schema_name"`
	Format         []Format `json:"format"`
	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  string   `json:"created_at_time"`
}

type Format struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}

type Template struct {
	TemplateID     string `json:"template_id"`
	MaxSupply      string `json:"max_supply"`
	IsTransferable bool   `json:"is_transferable"`
	IsBurnable     bool   `json:"is_burnable"`
	IssuedSupply   string `json:"issued_supply"`
	ImmutableData  Data   `json:"immutable_data"`
	CreatedAtTime  string `json:"created_at_time"`
	CreatedAtBlock string `json:"created_at_block"`
}

type Type string
