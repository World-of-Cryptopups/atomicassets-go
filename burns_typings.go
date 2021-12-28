package atomicassets

type BurnsProps struct {
	Success   bool             `json:"success"`
	Data      []BurnsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64            `json:"query_time"`
	Message   string           `json:"message,omitempty"` // error message will be here
}

type BurnsDataProps AccountsDataProps

type BurnsAccountProps struct {
	Success   bool                    `json:"success"`
	Data      []BurnsAccountDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                   `json:"query_time"`
	Message   string                  `json:"message,omitempty"` // error message will be here
}

type BurnsAccountDataProps struct {
	Collections []BurnsAccountDataCollectionsProps `json:"collections"`
	Templates   []BurnsAccountDataTemplatesProps   `json:"templates"`
	Assets      string                             `json:"assets"`
}

type BurnsAccountDataCollectionsProps struct {
	CollectionName string `json:"collection_name"`
	Assets         string `json:"assets"`
}

type BurnsAccountDataTemplatesProps struct {
	CollectionName string `json:"collection_name"`
	TemplateID     string `json:"template_id"`
	Assets         string `json:"assets"`
}
