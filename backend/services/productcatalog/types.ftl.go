// Code generated by FTL. DO NOT EDIT.
package productcatalog

import (
	"context"
	ftlbuiltin "ftl/builtin"
	"github.com/block/ftl/common/reflection"
	"github.com/block/ftl/go-runtime/ftl"
)

type GetClient func(context.Context, ftlbuiltin.HttpRequest[ftl.Unit, GetRequest, ftl.Unit]) (ftlbuiltin.HttpResponse[Product, ErrorResponse], error)

type ListClient func(context.Context, ftlbuiltin.HttpRequest[ftl.Unit, ftl.Unit, ListRequest]) (ftlbuiltin.HttpResponse[ListResponse, ftl.Unit], error)

type SearchClient func(context.Context, SearchRequest) (SearchResponse, error)

func init() {
	reflection.Register(

		reflection.ProvideResourcesForVerb(
			Get,
		),

		reflection.ProvideResourcesForVerb(
			List,
		),

		reflection.ProvideResourcesForVerb(
			Search,
		),
	)
}
