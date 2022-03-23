package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const BURNS_API = "burns"

type GetBurnsQuery struct {
	CollectionName      string   `url:"collection_name,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	TemplateID          string   `url:"template_id,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
	IDs                 []string `url:"ids,omitempty" del:","`
	LowerBound          string   `url:"lower_bound,omitempty"`
	UpperBound          string   `url:"upper_bound,omitempty"`
	Page                int64    `url:"page,omitempty"`
	Limit               int64    `url:"limit,omitempty"`
	Order               string   `url:"order,omitempty"`
}

// GetBurns gets the burns.
func (a *AtomicAssets) GetBurns(options *GetBurnsQuery) (*BurnsProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, BURNS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, BURNS_API)
	}

	var result = &BurnsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetBurnsAccount gets the burns of a specific account.
func (a *AtomicAssets) GetBurnsAccount(account string, options *GetBurnsQuery) (*BurnsAccountProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, BURNS_API, account, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, BURNS_API, account)
	}

	var result = &BurnsAccountProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
