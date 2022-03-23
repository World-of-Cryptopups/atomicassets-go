package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const TRANSFERS_API = "transfers"

type GetTransfersQuery struct {
	Account             string   `url:"account,omitempty"`
	Sender              string   `url:"sender,omitempty"`
	Recipient           string   `url:"recipient,omitempty"`
	AssetID             string   `url:"asset_id,omitempty"`
	TemplateID          string   `url:"template_id,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	CollectionName      string   `url:"collection_name,omitempty"`
	HideContracts       bool     `url:"hide_contracts,omitempty"`
	IDs                 []string `url:"ids,omitempty" del:","`
	LowerBound          string   `url:"lower_bound,omitempty"`
	UpperBound          string   `url:"upper_bound,omitempty"`
	Before              int64    `url:"before,omitempty"`
	After               int64    `url:"after,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
	Page                int64    `url:"page,omitempty"`
	Limit               int64    `url:"limit,omitempty"`
	Order               string   `url:"order,omitempty"`
	Sort                string   `url:"sort,omitempty"`
}

// GetTransfers fetches transfers.
func (a *AtomicAssets) GetTransfers(options *GetTransfersQuery) (*TransfersProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, TRANSFERS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, TRANSFERS_API)
	}

	var result = &TransfersProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
