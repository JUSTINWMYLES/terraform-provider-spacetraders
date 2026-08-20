---
page_title: "spacetraders_transfer_cargo Action - spacetraders"
subcategory: ""
description: |-
  Transfer cargo between ships.  The receiving ship must be in the same waypoint as the transferring ship, and it must able to hold the additional cargo after the transfer is complete. Both ships also must be in the same state, either both are docked or both are orbiting.  The response body's cargo shows the cargo of the transferring ship after the transfer is complete.
---

# spacetraders_transfer_cargo Action

Transfer cargo between ships.  The receiving ship must be in the same waypoint as the transferring ship, and it must able to hold the additional cargo after the transfer is complete. Both ships also must be in the same state, either both are docked or both are orbiting.  The response body's cargo shows the cargo of the transferring ship after the transfer is complete.

## Example Usage

```terraform
action "spacetraders_transfer_cargo" "example" {
  config {
    body_ship_symbol = "example"
    ship_symbol = "example"
    trade_symbol = "example"
    units = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `body_ship_symbol` (String, required) - The symbol of the ship to transfer to.
* `ship_symbol` (String, required) - The symbol of the ship.
* `trade_symbol` (String, required) - The good's symbol.
* `units` (Number, required) - Amount of units to transfer.
