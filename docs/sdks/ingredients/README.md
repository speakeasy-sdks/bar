# Ingredients
(*Ingredients*)

## Overview

The ingredients endpoints.

### Available Operations

* [ListIngredients](#listingredients) - Get a list of ingredients.

## ListIngredients

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bar"
	"context"
	"log"
)

func main() {
    s := bar.New(
        bar.WithSecurity(bar.Security{
            APIKey: bar.String("<YOUR_API_KEY>"),
        }),
    )


    ingredients := []string{
        "<value>",
    }

    ctx := context.Background()
    res, err := s.Ingredients.ListIngredients(ctx, ingredients)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `ingredients`                                                                         | []*string*                                                                            | :heavy_minus_sign:                                                                    | A list of ingredients to filter by. If not provided all ingredients will be returned. |


### Response

**[*ListIngredientsResponse](../../listingredientsresponse.md), error**
| Error Object     | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| bar.APIError     | 5XX              | application/json |
| bar.SDKError     | 4xx-5xx          | */*              |
