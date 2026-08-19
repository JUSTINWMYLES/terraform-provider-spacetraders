---
page_title: "spacetraders_get_shipyard Data Source - spacetraders"
subcategory: ""
description: |-
  Get the shipyard for a waypoint. Requires a waypoint that has the \`Shipyard\` trait to use. Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.
---

# spacetraders_get_shipyard Data Source

Get the shipyard for a waypoint. Requires a waypoint that has the \`Shipyard\` trait to use. Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.

## Example Usage

```terraform
data "spacetraders_get_shipyard" "example" {
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

* `modifications_fee` (Number, computed)
* `ship_types` (List(Object({type})), computed)
* `ships` (List(Object({activity, crew, description, engine, frame, modules, mounts, name, purchase_price, reactor, supply, type})), computed)
* `symbol` (String, computed)
* `transactions` (List(Object({agent_symbol, price, ship_symbol, ship_type, timestamp, waypoint_symbol})), computed)

