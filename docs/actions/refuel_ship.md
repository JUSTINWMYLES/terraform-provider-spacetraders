---
page_title: "spacetraders_refuel_ship Action - spacetraders"
subcategory: ""
description: |-
  Refuel your ship by buying fuel from the local market.  Requires the ship to be docking in a waypoint that has the \`Marketplace\` trait, and the market must be selling fuel in order to refuel.  Each fuel bought from the market replenishes 100 units in your ship's fuel.  Ships will always be refuel to their frame's maximum fuel capacity when using this action.
---

# spacetraders_refuel_ship Action

Refuel your ship by buying fuel from the local market.  Requires the ship to be docking in a waypoint that has the \`Marketplace\` trait, and the market must be selling fuel in order to refuel.  Each fuel bought from the market replenishes 100 units in your ship's fuel.  Ships will always be refuel to their frame's maximum fuel capacity when using this action.

## Example Usage

```terraform
action "spacetraders_refuel_ship" "example" {
  config {
    from_cargo = null
    ship_symbol = "example"
    units = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `from_cargo` (Dynamic, optional)
* `ship_symbol` (String, required)
* `units` (Number, optional)
