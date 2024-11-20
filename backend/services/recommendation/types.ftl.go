// Code generated by FTL. DO NOT EDIT.
package recommendation

import (
	"context"
	ftlbuiltin "ftl/builtin"
	ftlproductcatalog "ftl/productcatalog"
	"github.com/TBD54566975/ftl/go-runtime/ftl"
	"github.com/TBD54566975/ftl/go-runtime/ftl/reflection"
	"github.com/TBD54566975/ftl/go-runtime/server"
)

type ListClient func(context.Context, ftlbuiltin.HttpRequest[ftl.Unit, ftl.Unit, ListRequest]) (ftlbuiltin.HttpResponse[ListResponse, ErrorResponse], error)

func init() {
	reflection.Register(
		reflection.ProvideResourcesForVerb(
			List,
			server.VerbClient[ftlproductcatalog.ListClient, ftlbuiltin.HttpRequest[ftl.Unit, ftl.Unit, ftlproductcatalog.ListRequest], ftlbuiltin.HttpResponse[ftlproductcatalog.ListResponse, ftl.Unit]](),
		),
	)
}