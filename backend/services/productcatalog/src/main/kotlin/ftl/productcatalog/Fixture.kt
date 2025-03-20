package ftl.productcatalog

import ftl.currency.Money

fun loadProducts(insert: InsertProductClient) {
    insert.insertProduct(
        InsertProductQuery(
            id = "OLJCESPC7Z",
            name = "Sunglasses",
            description = "Add a modern touch to your outfits with these sleek aviator sunglasses.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/sunglasses.jpg?raw=true",
            priceUsd = "19",
            categories = "accessories"
        )
    )

    insert.insertProduct(
        InsertProductQuery(
            id = "66VCHSJNUP",
            name = "Tank Top",
            description = "Perfectly cropped cotton tank, with a scooped neckline.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/tank-top.jpg?raw=true",
            priceUsd = "18",
            categories = "clothing,tops"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "1YMWWN1N4O",
            name = "Watch",
            description = "This gold-tone stainless steel watch will work with most of your outfits.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/watch.jpg?raw=true",
            priceUsd = "109",
            categories = "accessories"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "L9ECAV7KIM",
            name = "Loafers",
            description = "A neat addition to your summer wardrobe.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/loafers.jpg?raw=true",
            priceUsd = "89",
            categories = "footwear"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "2ZYFJ3GM2N",
            name = "Hairdryer",
            description = "This lightweight hairdryer has 3 heat and speed settings. It's perfect for travel.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/hairdryer.jpg?raw=true",
            priceUsd = "24",
            categories = "hair,beauty"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "0PUK6V6EV0",
            name = "Candle Holder",
            description = "This small but intricate candle holder is an excellent gift.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/candle-holder.jpg?raw=true",
            priceUsd = "18",
            categories = "decor,home"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "LS4PSXUNUM",
            name = "Salt & Pepper Shakers",
            description = "Add some flavor to your kitchen.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/salt-and-pepper-shakers.jpg?raw=true",
            priceUsd = "18",
            categories = "kitchen"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "9SIQT8TOJO",
            name = "Bamboo Glass Jar",
            description = "This bamboo glass jar can hold 57 oz (1.7 l) and is perfect for any kitchen.",
            picture = "https://raw.githubusercontent.com/GoogleCloudPlatform/microservices-demo/main/src/frontend/static/img/products/bamboo-glass-jar.jpg",
            priceUsd = "5",
            categories = "kitchen"
        )
    )
    insert.insertProduct(
        InsertProductQuery(
            id = "6E92ZMYYFZ",
            name = "Mug",
            description = "A simple mug with a mustard interior.",
            picture = "https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/src/frontend/static/img/products/mug.jpg?raw=true",
            priceUsd = "8",
            categories = "kitchen"
        )
    )

}