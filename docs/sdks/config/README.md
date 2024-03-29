# Config
(*Config*)

### Available Operations

* [SubscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## SubscribeToWebhooks

Subscribe to webhooks.

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

    ctx := context.Background()
    res, err := s.Config.SubscribeToWebhooks(ctx, []bar.RequestBody{
        bar.RequestBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]RequestBody](../../.md)                            | :heavy_check_mark:                                    | The request object to use for the request.            |
| `opts`                                                | [][bar.Option](../../option.md)                       | :heavy_minus_sign:                                    | The options for this request.                         |


### Response

**[*SubscribeToWebhooksResponse](../../subscribetowebhooksresponse.md), error**
| Error Object     | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| bar.BadRequest   | 400              | application/json |
| bar.APIError     | 5XX              | application/json |
| bar.SDKError     | 4xx-5xx          | */*              |
