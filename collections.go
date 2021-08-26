package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const COLLECTIONS_API = "collections"

type GetCollectionsQuery struct {
	Author              string   `url:"author,omitempty"`
	Match               string   `url:"match,omitempty"`
	AuthorizedAccount   string   `url:"authorized_account,omitempty"`
	NotifyAccount       string   `url:"notify_account,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
	IDs                 []string `url:"ids,omitempty" del:","`
	LowerBound          string   `url:"lower_bound,omitempty"`
	UpperBound          string   `url:"upper_bound,omitempty"`
	Before              int64    `url:"before,omitempty"`
	After               int64    `url:"after,omitempty"`
	Page                int      `url:"page,omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Order               string   `url:"order,omitempty"`
	Sort                string   `url:"sort,omitempty"`
}

// GetCollections fetchers collections.
func (a *AtomicAssets) GetCollections(options *GetCollectionsQuery) (*CollectionsProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u, _ = urljoin.UrlJoin(a.API_URL, COLLECTIONS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u, _ = urljoin.UrlJoin(a.API_URL, COLLECTIONS_API)
	}

	var result = &CollectionsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCollectionName gets the collection with its name.
func (a *AtomicAssets) GetCollectionName(name string) (*CollectionNameProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, COLLECTIONS_API, name)

	var result = &CollectionNameProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCollectionNameStats gets the collection's stats.
func (a *AtomicAssets) GetCollectionNameStats(name string) (*CollectionNameStatsProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, COLLECTIONS_API, name, "stats")

	var result = &CollectionNameStatsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetCollectionNameLogs GetAssetLogsQuery

// GetCollectionNameStats gets the collection's logs.
func (a *AtomicAssets) GetCollectionNameLogs(name string, options *GetCollectionNameLogs) (*CollectionNameLogsProps, error) {
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
	u, _ := urljoin.UrlJoin(a.API_URL, COLLECTIONS_API, name, "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &CollectionNameLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
