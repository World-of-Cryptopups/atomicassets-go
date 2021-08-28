package atomicassets

// GetSchemas response
type SchemaProps struct {
	Success   bool          `json:"success"`
	Data      []SchemasData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64         `json:"query_time"`
	Message   string        `json:"message,omitempty"` // error message will be here
}

// GetSchemaName response
type SchemaNameProps struct {
	Success   bool        `json:"success"`
	Data      SchemasData `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64       `json:"query_time"`
	Message   string      `json:"message,omitempty"` // error message will be here
}

// GetSchemaNameStats response
type SchemaNameStatsProps struct {
	Success   bool        `json:"success"`
	Data      SchemaStats `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64       `json:"query_time"`
	Message   string      `json:"message,omitempty"` // error message will be here
}

// GetSchemaNameLogs response
type SchemaNameLogsProps AssetsIDLogsProps

type SchemaStats struct {
	Assets    string
	Burned    string
	Templates string
}

type SchemasData struct {
	Contract       string           `json:"contract"`
	SchemaName     string           `json:"schema_name"`
	Format         []SchemaFormat   `json:"format"`
	CreatedAtBlock string           `json:"created_at_block"`
	CreatedAtTime  string           `json:"created_at_time"`
	Collection     SchemaCollection `json:"collection"`
}

type SchemaFormat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SchemaCollection struct {
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
