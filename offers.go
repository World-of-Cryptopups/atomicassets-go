package atomicassets

import (
	"fmt"
	"strconv"

	"github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const OFFERS_API = "offers"

type GetOffersQuery struct {
	Account                 string   `url:"account,omitempty"`
	Sender                  string   `url:"sender,omitempty"`
	Recipient               string   `url:"recipient,omitempty"`
	State                   string   `url:"state,omitempty"`
	IsRecipientContract     bool     `url:"is_recipient_contact,omitempty"`
	AssetID                 []string `url:"asset_id,omitempty" del:","`
	TemplateID              []string `url:"template_id,omitempty" del:","`
	SchemaName              []string `url:"schema_name,omitempty" del:","`
	CollectionName          []string `url:"collection_name,omitempty" del:","`
	AccountWhitelist        string   `url:"account_whitelist,omitempty"`
	AccountBlacklist        string   `url:"account_blacklist,omitempty"`
	SenderAssetWhitelist    string   `url:"sender_asset_whitelist,omitempty"`
	SenderAssetBlacklist    string   `url:"sender_asset_blacklist,omitempty"`
	RecipientAssetWhitelist string   `url:"recipient_asset_whitelist,omitempty"`
	RecipientAssetBlacklist string   `url:"recipient_asset_blacklist,omitempty"`
	HideContracts           bool     `url:"hide_contracts,omitempty"`
	Ids                     []string `url:"ids,omitempty" del:","`
	LowerBound              string   `url:"lower_bound,omitempty"`
	UpperBound              string   `url:"upper_bound,omitempty"`
	Before                  int64    `url:"before,omitempty"`
	After                   int64    `url:"after,omitempty"`
	Page                    int      `url:"page,omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Order                   string   `url:"order,omitempty"`
	Sort                    string   `url:"sort,omitempty"`
}

// GetOffers fetches offers.
func (a *AtomicAssets) GetOffers(options *GetOffersQuery) (*OffersProps, error) {
	var u string
	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, OFFERS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, OFFERS_API)
	}

	var result = &OffersProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOfferID gets the offer by its id.
func (a *AtomicAssets) GetOfferID(offerID int) (*OffersData, error) {
	u := urljoin.UrlJoin(a.apiUrl, OFFERS_API, strconv.Itoa(offerID))

	var result = &OffersData{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetOfferIDLogsQuery GetAssetLogsQuery

// GetOfferIDLogs gets the offer id's logs.
func (a *AtomicAssets) GetOfferIDLogs(offerID int, options *GetOfferIDLogsQuery) (*OfferIDLogsProps, error) {
	var opts = &GetOfferIDLogsQuery{
		Limit: 100,
		Page:  1,
		Order: "desc",
	}

	if options != nil {
		opts = &GetOfferIDLogsQuery{
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
	u := urljoin.UrlJoin(a.apiUrl, OFFERS_API, strconv.Itoa(offerID), "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &OfferIDLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
