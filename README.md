# govtech-catalog-test-project

This is a simple guide to help you get started with our Go application for managing products. Follow the steps below to set up and use the service.

## Prerequisites

Before you can run the service, you need to have Go (Golang) installed on your system. This application is developed using Go version 1.21.x.

1. **Install Golang**: You can download and install Go version 1.21.x from the official website: [Go Downloads](https://golang.org/dl/)

## Getting Started

1. After installing Go, navigate to the project directory in your terminal.
2. Run the following command to ensure that all dependencies are downloaded and updated:

   ```shell
   go mod tidy
   ```

3. To start the service, run the following command:

   ```shell
   go run cmd/main.go
   ```

4. Next, copy the sample environment file to create your own environment file:

   ```shell
   cp .env_sample .env
   ```

The service should now be running locally on `http://localhost:8080`.

## Usage

(Continue with the usage instructions as provided in the original README)

Adding the step to copy the environment file ensures that you have the necessary environment variables set for your application to work correctly.

### Create a New Product

To create a new product, you can use the following `curl` command as an example:

```shell
curl --location --request POST 'http://localhost:8080/api/v1/products' \
--header 'Content-Type: application/json' \
--data-raw '{
  "sku": "SKU_0001",
  "title": "Cawan Ajaib",
  "description": "Ini cawan ajaib bisa memberikan anda makanan",
  "category": "Ghoin",
  "etalase": "Dunia Lain",
  "images": [
    {
      "image_url": "http://cawan_goib_1.png",
      "description": "gambar cawan goib saat baru lahir"
    },
    {
      "image_url": "http://cawan_goib_2.png",
      "description": "gambar cawan goib sakti"
    }
  ],
  "weight": 1,
  "price": 1000000000.99
}'
```

### Get a Product by ID

To retrieve a product by its ID, you can use the following `curl` command as an example:

```shell
curl --location --request GET 'http://localhost:8080/api/v1/products/1'
```

### Get a List of Products

#### Query Params:

| Query Parameter  | Example Value | Description                                       |
| ---------------- | ------------- | ------------------------------------------------- |
| `sku`            | `SKU_0001U`   | Search by SKU                                     |
| `title`          | `title`       | Search by Title                                   |
| `category`       | `Ghoin`       | Search by Category                                |
| `etalase`        | `Dunia Lain`  | Search by Showcase (Etalase)                     |
| `sort_created`   | `oldest`      | Sorting: Sort data by creation date (newest/oldest) |
| `sort_rating`    | `lowest`      | Sorting: Sort data by rating (lowest/highest)     |

These query parameters can be used to filter and sort the list of products based on the specified criteria.

To get a list of products with filtering and sorting options, use the following `curl` command as an example:

```shell
curl --location --request GET 'http://localhost:8080/api/v1/products?sku=SKU_0001U&title=title&category=Ghoin&etalase=Dunia Lain&sort_created=oldest&sort_rating=lowest'
```

### Update a Product

To update a product, you can use the following `curl` command as an example:

```shell
curl --location --request PUT 'http://localhost:8080/api/v1/products/1' \
--header 'Content-Type: application/json' \
--data-raw '{
  "sku": "SKU_0002U",
  "title": "Cawan Ajaib Update Bos 2",
  "description": "",
  "category": "",
  "images": [
    {
      "image_url": "http://cawan_goib_4.png",
      "description": "gambar cawan goib saat baru lahir updated"
    }
  ],
  "weight": 1.5,
  "price": 1000
}'
```

### Add a Product Review

To add a review to a product, use the following `curl` command as an example:

```shell
curl --location --request POST 'http://localhost:8080/api/v1/products/1/reviews' \
--header 'Content-Type: application/json' \
--data-raw '{
  "rating": 3,
  "review_comment": "barangnya b aja bosque 2"
}'
```

You can now start using the product management service with these commands. Be sure to replace the URLs and data as needed for your specific use case.