// Code generated by FTL. DO NOT EDIT.
package productcatalog

import (
	"context"
	ftlbuiltin "ftl/builtin"
	"github.com/TBD54566975/ftl/go-runtime/ftl"
	"github.com/TBD54566975/ftl/go-runtime/ftl/reflection"
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
