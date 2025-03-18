//ftl:module ad
package ad

import (
	"context"
	_ "embed"
	"github.com/block/ftl/common/slices"

	"ftl/builtin"

	"github.com/block/ftl/go-runtime/ftl"
)

const maxAdsToServe = 2

type AdRequest struct {
	ContextKeys []string
}

type Ad struct {
	RedirectURL string
	Text        string
}

type AdResponse struct {
	Name string
	Ads  []Ad
}

//ftl:ingress GET /ad
func Get(ctx context.Context, req builtin.HttpRequest[ftl.Unit, ftl.Unit, AdRequest], getAds GetAdsClient, getAd GetAdClient) (builtin.HttpResponse[AdResponse, ftl.Unit], error) {
	var ads []Ad
	var err error
	if len(req.Query.ContextKeys) > 0 {
		ads, err = contextualAds(ctx, req.Query.ContextKeys, getAd)
		if err != nil {
			return builtin.HttpResponse[AdResponse, ftl.Unit]{}, err
		}
	} else {
		ads, err = randomAds(ctx, getAds)
	}

	return builtin.HttpResponse[AdResponse, ftl.Unit]{
		Body: ftl.Some(AdResponse{Name: "ad", Ads: ads}),
	}, nil
}

func contextualAds(ctx context.Context, contextKeys []string, client GetAdClient) (ads []Ad, err error) {
	for _, key := range contextKeys {
		ad, err := client(ctx, GetAdQuery{Name: key})
		if err != nil {
			return nil, err
		}
		ads = append(ads, Ad{RedirectURL: ad.Url, Text: ad.Text})
	}
	return ads, err
}

func randomAds(ctx context.Context, client GetAdsClient) (ads []Ad, err error) {
	ret, err := client(ctx)
	if err != nil {
		return nil, err
	}
	return slices.Map(ret, func(t GetAdsResult) Ad {
		return Ad{RedirectURL: t.Url, Text: t.Text}
	}), err
}
