package atomicassets

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const SCHEMAS_API = "schemas"

type GetSchemasQuery struct {
	CollectionName      string   `url:"collection_name,omitempty"`
	AuthorizedAccount   string   `url:"authorized_account,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	Match               string   `url:"match,omitempty"`
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

// GetSchemas fetchers collections.
func (a *AtomicAssets) GetSchemas(options *GetSchemasQuery) (*SchemaProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u, _ = urljoin.UrlJoin(a.API_URL, SCHEMAS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u, _ = urljoin.UrlJoin(a.API_URL, SCHEMAS_API)
	}

	var result = &SchemaProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCollectionName gets the collection with its name.
func (a *AtomicAssets) GetSchemaName(collectionName, schemaName string) (*SchemaNameProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, SCHEMAS_API, collectionName, schemaName)

	var result = &SchemaNameProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetCollectionNameStats gets the collection's stats.
func (a *AtomicAssets) GetSchemaNameStats(collectionName, schemaName string) (*SchemaNameStatsProps, error) {
	u, _ := urljoin.UrlJoin(a.API_URL, SCHEMAS_API, collectionName, schemaName, "stats")

	var result = &SchemaNameStatsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetSchemaNameLogs GetAssetLogsQuery

// GetCollectionNameStats gets the collection's logs.
func (a *AtomicAssets) GetSchemaNameLogs(collectionName, schemaName string, options *GetSchemaNameLogs) (*SchemaNameLogsProps, error) {
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
	u, _ := urljoin.UrlJoin(a.API_URL, SCHEMAS_API, collectionName, schemaName, "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &SchemaNameLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
