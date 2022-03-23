package atomicassets

import (
	"fmt"
	"strconv"

	urljoin "github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const TEMPLATES_API = "templates"

type GetTemplatesQuery struct {
	CollectionName      string   `url:"collection_name,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	IssuedSupply        int      `url:"issued_supply,omitempty"`
	MinIssuedSupply     int      `url:"min_issued_supply,omitempty"`
	MaxIssuedSupply     int      `url:"max_issued_supply,omitempty"`
	HasAssets           bool     `url:"has_assets,omitempty"`
	MaxSupply           int      `url:"max_supply,omitempty"`
	IsBurnable          bool     `url:"is_burnable,omitempty"`
	IsTransferable      bool     `url:"is_transferable,omitempty"`
	AuthorizedAccount   string   `url:"authorized_account,omitempty"`
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

// GetTemplates fetches templates.
func (a *AtomicAssets) GetTemplates(options *GetTemplatesQuery) (*TemplatesProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, TEMPLATES_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, TEMPLATES_API)
	}

	fmt.Println(u)

	var result = &TemplatesProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTemplateID gets the template by its id.
func (a *AtomicAssets) GetTemplateID(collectionName string, templateID int) (*TemplateIDProps, error) {
	u := urljoin.UrlJoin(a.apiUrl, TEMPLATES_API, collectionName, strconv.Itoa(templateID))

	var result = &TemplateIDProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTemplateIDStats gets the collection's stats.
func (a *AtomicAssets) GetTemplateIDStats(collectionName string, templateID int) (*TemplateIDStatsProps, error) {
	u := urljoin.UrlJoin(a.apiUrl, TEMPLATES_API, collectionName, strconv.Itoa(templateID), "stats")

	var result = &TemplateIDStatsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetTemplateIDLogs GetAssetLogsQuery

// GetTemplateIDLogs gets the collection's logs.
func (a *AtomicAssets) GetTemplateIDLogs(collectionName string, templateID int, options *GetTemplateIDLogs) (*TemplateIDLogsProps, error) {
	var opts = &GetTemplateIDLogs{
		Limit: 100,
		Page:  1,
		Order: "desc",
	}

	if options != nil {
		opts = &GetTemplateIDLogs{
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
	u := urljoin.UrlJoin(a.apiUrl, SCHEMAS_API, collectionName, strconv.Itoa(templateID), "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &TemplateIDLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
