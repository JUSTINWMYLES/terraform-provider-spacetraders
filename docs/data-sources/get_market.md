---
page_title: "spacetraders_get_market Data Source - spacetraders"
subcategory: ""
description: |-
  Retrieve imports, exports and exchange data from a marketplace. Requires a waypoint that has the \`Marketplace\` trait to use.  Send a ship to the waypoint to access trade good prices and recent transactions. Refer to the \[Market Overview page\](https://docs.spacetraders.io/game-concepts/markets) to gain better a understanding of the market in the game.
---

# spacetraders_get_market Data Source

Retrieve imports, exports and exchange data from a marketplace. Requires a waypoint that has the \`Marketplace\` trait to use.  Send a ship to the waypoint to access trade good prices and recent transactions. Refer to the \[Market Overview page\](https://docs.spacetraders.io/game-concepts/markets) to gain better a understanding of the market in the game.

## Example Usage

```terraform
data "spacetraders_get_market" "example" {
  system_symbol = null
  waypoint_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `system_symbol` (String, required)
* `waypoint_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `exchange` (List(Object({description, name, symbol})), computed)
* `exports` (List(Object({description, name, symbol})), computed)
* `imports` (List(Object({description, name, symbol})), computed)
* `symbol` (String, computed)
* `trade_goods` (List(Object({activity, purchase_price, sell_price, supply, symbol, trade_volume, type})), computed)
* `transactions` (List(Object({price_per_unit, ship_symbol, timestamp, total_price, trade_symbol, type, units, waypoint_symbol})), computed)

