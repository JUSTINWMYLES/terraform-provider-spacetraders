---
page_title: "spacetraders_get_repair_ship Data Source - spacetraders"
subcategory: ""
description: |-
  Get the cost of repairing a ship. Requires the ship to be docked at a waypoint that has the \`Shipyard\` trait.
---

# spacetraders_get_repair_ship Data Source

Get the cost of repairing a ship. Requires the ship to be docked at a waypoint that has the \`Shipyard\` trait.

## Example Usage

```terraform
data "spacetraders_get_repair_ship" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `transaction` (Object({ship_symbol, timestamp, total_price, waypoint_symbol}), computed) - Result of a repair transaction.
  * `ship_symbol` (String, computed) - The symbol of the ship.
  * `timestamp` (String, computed) - The timestamp of the transaction.
  * `total_price` (Number, computed) - The total price of the transaction.
  * `waypoint_symbol` (String, computed) - The symbol of the waypoint.

