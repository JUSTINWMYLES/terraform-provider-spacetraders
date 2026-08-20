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

* `system_symbol` (String, required) - The system symbol
* `waypoint_symbol` (String, required) - The waypoint symbol

### Attributes

In addition to all arguments above, the following attributes are exported:

* `modifications_fee` (Number, computed) - The fee to modify a ship at this shipyard. This includes installing or removing modules and mounts on a ship. In the case of mounts, the fee is a flat rate per mount. In the case of modules, the fee is per slot the module occupies.
* `ship_types` (List(Object({type})), computed) - The list of ship types available for purchase at this shipyard.
* `ships` (List(Object({activity, crew, description, engine, frame, modules, mounts, name, purchase_price, reactor, supply, type})), computed) - The ships that are currently available for purchase at the shipyard.
* `symbol` (String, computed) - The symbol of the shipyard. The symbol is the same as the waypoint where the shipyard is located.
* `transactions` (List(Object({agent_symbol, price, ship_symbol, ship_type, timestamp, waypoint_symbol})), computed) - The list of recent transactions at this shipyard.

