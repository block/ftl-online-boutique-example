// Code generated by FTL. DO NOT EDIT.
package shipping

import (
	"context"
	ftlbuiltin "ftl/builtin"
	ftlcurrency "ftl/currency"
	"github.com/block/ftl/common/reflection"
	"github.com/block/ftl/go-runtime/ftl"
)

type GetQuoteClient func(context.Context, ftlbuiltin.HttpRequest[ShippingRequest, ftl.Unit, ftl.Unit]) (ftlbuiltin.HttpResponse[ftlcurrency.Money, ftl.Unit], error)

type ShipOrderClient func(context.Context, ftlbuiltin.HttpRequest[ShippingRequest, ftl.Unit, ftl.Unit]) (ftlbuiltin.HttpResponse[ShipOrderResponse, ftl.Unit], error)

func init() {
	reflection.Register(

		reflection.ProvideResourcesForVerb(
			GetQuote,
		),

		reflection.ProvideResourcesForVerb(
			ShipOrder,
		),
	)
}
