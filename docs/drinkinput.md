# DrinkInput


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | The name of the drink.                                            | Old Fashioned                                                     |
| `Type`                                                            | [*DrinkType](..//drinktype.md)                                    | :heavy_minus_sign:                                                | The type of drink.                                                |                                                                   |
| `Price`                                                           | *float64*                                                         | :heavy_check_mark:                                                | The price of one unit of the drink in US cents.                   | 1000                                                              |
| `ProductCode`                                                     | **string*                                                         | :heavy_minus_sign:                                                | The product code of the drink, only available when authenticated. | AC-A2DF3                                                          |