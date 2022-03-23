package atomicassets

// GetAccounts response
type AccountsProps struct {
	Success   bool                   `json:"success"`
	Data      []CollectionsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                  `json:"query_time"`
	Message   string                 `json:"message,omitempty"` // error message will be here
}

type AccountsDataProps struct {
	Account string `json:"account"`
	Assets  string `json:"assets"`
}

type AccountProps struct {
	Success   bool               `json:"success"`
	Data      []AccountDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64              `json:"query_time"`
	Message   string             `json:"message,omitempty"` // error message will be here
}

type AccountDataProps struct {
	Collections []AccountDataCollectionsProps `json:"collections"`
	Templates   []AccountDataTemplatesProps
	Assets      string `json:"assets"`
}

type AccountDataCollectionsProps struct {
	Collection CollectionsDataProps `json:"collection"`
	Assets     string               `json:"assets"`
}

type AccountDataTemplatesProps struct {
	TemplateID string `json:"template_id"`
	Assets     string `json:"assets"`
}

type AccountCollectionProps struct {
	Success   bool                       `json:"success"`
	Data      AccountCollectionDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                      `json:"query_time"`
	Message   string                     `json:"message,omitempty"` // error message will be here
}

type AccountCollectionDataProps struct {
	Templates []AccountCollectionDataTemplatesProps `json:"templates"`
	Schemas   []AccountCollectionDataSchemasProps   `json:"schemas"`
}

type AccountCollectionDataTemplatesProps AccountDataTemplatesProps
type AccountCollectionDataSchemasProps struct {
	SchemaName string `json:"schema_name"`
	Assets     string `json:"assets"`
}
