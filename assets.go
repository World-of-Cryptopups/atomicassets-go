package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const ASSETS_API = "assets"

type GetAssetsQuery struct {
	CollectionName          string   `url:"collection_name,omitempty"`
	SchemaName              string   `url:"schema_name,omitempty"`
	TemplateID              int      `url:"template_id,omitempty"`
	IsTransferable          bool     `url:"is_transferable,omitempty"`
	IsBurnable              bool     `url:"is_burnable,omitempty"`
	Burned                  bool     `url:"burned,omitempty"`
	Owner                   string   `url:"owner,omitempty"`
	Match                   string   `url:"match,omitempty"`
	CollectionBlacklist     []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist     []string `url:"collection_whitelist,omitempty" del:","`
	OnlyDuplicateTemplates  bool     `url:"only_duplicate_templates,omitempty"`
	HasBackedTokens         bool     `url:"has_backed_tokens,omitempty"`
	AuthorizedAccount       bool     `url:"authorized_account,omitempty"`
	TemplateWhitelist       []string `url:"template_whitelist,omitempty" del:","`
	TemplateBlacklist       []string `url:"template_blacklist,omitempty" del:","`
	HideTemplatesByAccounts string   `url:"hide_templates_by_accounts,omitempty"`
	HideOffers              bool     `url:"hide_offers,omitempty"`
	IDs                     []string `url:"ids,omitempty" del:","`
	LowerBound              []string `url:"lower_bound,omitempty" del:","`
	UpperBound              []string `url:"upper_bound,omitempty" del:","`
	Before                  int64    `url:"before,omitempty"`
	After                   int64    `url:"after,omitempty"`
	Page                    int      `url:"page,omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Order                   string   `url:"order,omitempty"`
	Sort                    string   `url:"sort,omitempty"`
}

// GetAssets fetches assets.
func (a *AtomicAssets) GetAssets(options *GetAssetsQuery) (*AssetsProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u, _ = urljoin.UrlJoin(a.API_URL, ASSETS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u, _ = urljoin.UrlJoin(a.API_URL, ASSETS_API)
	}

	var result = &AssetsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAssetID fetches an assets by its id.
func (a *AtomicAssets) GetAssetID(id string) (*AssetsIDProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, ASSETS_API, id)

	var result = &AssetsIDProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAssetIDStats fetches the stats of an asset with its id.
func (a *AtomicAssets) GetAssetIDStats(id string) (*AssetsIDStatsProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, ASSETS_API, id, "stats")

	var result = &AssetsIDStatsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetAssetLogsQuery struct {
	Page            int    `url:"page,omitempty"`
	Limit           int    `url:"limit,omitempty"`
	Order           string `url:"order,omitempty"`
	ActionWhitelist string `url:"action_whitelist,omitempty"`
	ActionBlacklist string `url:"action_blacklist,omitempty"`
}

// GetAssetIDLogs fetches the logs of an asset with its id.
func (a *AtomicAssets) GetAssetIDLogs(id string, options *GetAssetLogsQuery) (*AssetsIDLogsProps, error) {
	var opts = &GetAssetLogsQuery{
		Limit: 100,
		Page:  1,
		Order: "desc",
	}

	if options != nil {
		opts = &GetAssetLogsQuery{
			ActionWhitelist: options.ActionWhitelist,
			ActionBlacklist: options.ActionBlacklist,
		}

		if options.Limit < 1 {
			opts.Limit = 100
		}
		if options.Page < 1 {
			opts.Page = 1
		}
		if options.Order == "desc" || options.Order == "asc" {
			opts.Order = options.Order
		} else {
			// default to `desc`
			opts.Order = "desc"
		}
	}

	v, _ := query.Values(opts)
	u, _ := urljoin.UrlJoin(a.API_URL, ASSETS_API, id, "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &AssetsIDLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
