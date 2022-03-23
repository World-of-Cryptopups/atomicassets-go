package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const ACCOUNTS_API = "accounts"

type GetAccountsQuery struct {
	Match               string   `url:"match,omitempty"`
	CollectionName      string   `url:"collection_name,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	TemplateID          string   `url:"template_id,omitempty"`
	HideOffers          bool     `url:"hide_offers,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
	IDs                 []string `url:"ids,omitempty" del:","`
	LowerBound          string   `url:"lower_bound,omitempty"`
	UpperBound          string   `url:"upper_bound,omitempty"`
	Page                int64    `url:"page,omitempty"`
	Limit               int64    `url:"limit,omitempty"`
	Order               string   `url:"order,omitempty"`
}

// GetAccounts gets the accounts which oen atomicassets NFTs.
func (a *AtomicAssets) GetAccounts(options *GetAccountsQuery) (*AccountsProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, ACCOUNTS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, ACCOUNTS_API)
	}

	var result = &AccountsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetAccountQuery struct {
	HideOffers          bool     `url:"hide_offers,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
}

// GetAccount gets a specific account.
func (a *AtomicAssets) GetAccount(account string, options *GetAccountQuery) (*AccountProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, ACCOUNTS_API, account, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, ACCOUNTS_API, account)
	}

	var result = &AccountProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAccountCollection gets the templates and schemas count by account.
func (a *AtomicAssets) GetAccountCollection(account, collection string) (*AccountCollectionProps, error) {
	u := urljoin.UrlJoin(a.apiUrl, ACCOUNTS_API, account, collection)

	var result = &AccountCollectionProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
