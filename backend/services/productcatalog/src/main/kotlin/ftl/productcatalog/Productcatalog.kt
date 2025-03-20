package ftl.productcatalog

import ftl.currency.Money
import xyz.block.ftl.*
import jakarta.ws.rs.GET
import jakarta.ws.rs.NotFoundException
import jakarta.ws.rs.Path

@Export
data class Product(
    val id: String,
    val name: String,
    val description: String,
    val picture: String,
    val priceUsd: Money,
    val categories: List<String>
)


@Export
data class SearchRequest(
    val query: String
)

@Export
data class SearchResponse(
    val results: List<Product>
)

data class GetRequest(
    val id: String
)

data class GetResponse(
    val product: Product?
)

// Load database from JSON
private var database: List<Product> = listOf()

    @GET
    @Path("/productcatalog")
    fun listIngrees(): List<Product> {
        return list()
    }

    @Verb
    @Export
    fun list(): List<Product> {
        return database
    }

    @Verb
    @Export
    fun get(id: GetRequest): GetResponse {
        val product = database.find { it.id == id.id }
        return GetResponse(product)
    }

    @GET
    @Path("/productcatalog/{id}")
    fun getIngress(id: String): Product {
        val product = database.find { it.id == id }
        if (product != null) {
            return product
        } else {
            throw NotFoundException()
        }
    }

    @Export
    @Verb
    fun search(request: SearchRequest): SearchResponse {
        val query = request.query.lowercase()
        val results = database.filter { product ->
            product.name.lowercase().contains(query) || 
            product.description.lowercase().contains(query)
        }
        return SearchResponse(results = results)
    }


