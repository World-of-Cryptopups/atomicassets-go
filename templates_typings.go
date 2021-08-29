package atomicassets

// GetTemplates response
type TemplatesProps struct {
	Success   bool            `json:"success"`
	Data      []TemplatesData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64           `json:"query_time"`
	Message   string          `json:"message,omitempty"` // error message will be here
}

// GetTemplateID response
type TemplateIDProps struct {
	Success   bool          `json:"success"`
	Data      TemplatesData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64         `json:"query_time"`
	Message   string        `json:"message,omitempty"` // error message will be here
}

// GetTemplateIDStats response
type TemplateIDStatsProps struct {
	Success   bool          `json:"success"`
	Data      TemplateStats `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64         `json:"query_time"`
	Message   string        `json:"message,omitempty"` // error message will be here
}

// GetTemplateIDLogs response
type TemplateIDLogsProps struct {
	Success   bool           `json:"success"`
	Data      []TemplateLogs `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64          `json:"query_time"`
	Message   string         `json:"message,omitempty"` // error message will be here
}

type TemplateStats struct {
	Assets string `json:"assets"`
	Burned string `json:"burned"`
}

type TemplateLogs AssetsIDLogsProps

type TemplatesData struct {
	Contract       string                 `json:"contract"`
	TemplateID     string                 `json:"template_id"`
	MaxSupply      string                 `json:"max_supply"`
	IsTransferable bool                   `json:"is_transferable"`
	IsBurnable     bool                   `json:"is_burnable"`
	ImmutableData  MutableData            `json:"immutable_data"`
	CreatedAtTime  string                 `json:"created_at_time"`
	CreatedAtBlock string                 `json:"created_at_block"`
	Schema         TemplateSchemaData     `json:"schema"`
	Collection     TemplateCollectionData `json:"collection"`
}

type TemplateSchemaData struct {
	SchemaName     string         `json:"schema_name"`
	Format         []SchemaFormat `json:"format"`
	CreatedAtBlock string         `json:"created_at_block"`
	CreatedAtTime  string         `json:"created_at_time"`
}

type TemplateCollectionData struct {
	CollectionName     string   `json:"collection_name"`
	Name               string   `json:"name"`
	Author             string   `json:"author"`
	AllowNotify        bool     `json:"allow_notify"`
	AuthorizedAccounts []string `json:"authorized_accounts"`
	NotifyAccounts     []string `json:"notify_accounts"`
	MarketFee          float64  `json:"market_fee"`
	CreatedAtBlock     string   `json:"created_at_block"`
	CreatedAtTime      string   `json:"created_at_time"`
}
