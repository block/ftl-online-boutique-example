// Code generated by FTL. DO NOT EDIT.
package checkout

import (
	"context"
	ftlbuiltin "ftl/builtin"
	ftlcart "ftl/cart"
	ftlcurrency "ftl/currency"
	ftlpayment "ftl/payment"
	ftlproductcatalog "ftl/productcatalog"
	ftlshipping "ftl/shipping"
	"github.com/block/ftl/common/reflection"
	"github.com/block/ftl/go-runtime/ftl"
	"github.com/block/ftl/go-runtime/server"
)

type PlaceOrderClient func(context.Context, ftlbuiltin.HttpRequest[PlaceOrderRequest, string, ftl.Unit]) (ftlbuiltin.HttpResponse[Order, ErrorResponse], error)

func init() {
	reflection.Register(

		reflection.ProvideResourcesForVerb(
			PlaceOrder,
			server.VerbClient[ftlcart.GetCartClient, ftlbuiltin.HttpRequest[ftl.Unit, ftl.Unit, ftlcart.GetCartRequest], ftlbuiltin.HttpResponse[ftlcart.Cart, ftl.Unit]](),
			server.VerbClient[ftlproductcatalog.GetClient, ftlbuiltin.HttpRequest[ftl.Unit, ftlproductcatalog.GetRequest, ftl.Unit], ftlbuiltin.HttpResponse[ftlproductcatalog.Product, ftlproductcatalog.ErrorResponse]](),
			server.VerbClient[ftlcurrency.ConvertClient, ftlbuiltin.HttpRequest[ftlcurrency.ConvertRequest, ftl.Unit, ftl.Unit], ftlbuiltin.HttpResponse[ftlcurrency.Money, string]](),
			server.VerbClient[ftlshipping.GetQuoteClient, ftlbuiltin.HttpRequest[ftlshipping.ShippingRequest, ftl.Unit, ftl.Unit], ftlbuiltin.HttpResponse[ftlcurrency.Money, ftl.Unit]](),
			server.VerbClient[ftlpayment.ChargeClient, ftlbuiltin.HttpRequest[ftlpayment.ChargeRequest, ftl.Unit, ftl.Unit], ftlbuiltin.HttpResponse[ftlpayment.ChargeResponse, ftlpayment.ErrorResponse]](),
			server.VerbClient[ftlshipping.ShipOrderClient, ftlbuiltin.HttpRequest[ftlshipping.ShippingRequest, ftl.Unit, ftl.Unit], ftlbuiltin.HttpResponse[ftlshipping.ShipOrderResponse, ftl.Unit]](),
			server.VerbClient[ftlcart.EmptyCartClient, ftlbuiltin.HttpRequest[ftlcart.EmptyCartRequest, ftl.Unit, ftl.Unit], ftlbuiltin.HttpResponse[ftlcart.EmptyCartResponse, ftl.Unit]](),
		),
	)
}
