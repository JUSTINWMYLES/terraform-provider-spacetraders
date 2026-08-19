---
page_title: "spacetraders_sell_cargo Action - spacetraders"
subcategory: ""
description: |-
  Sell cargo in your ship to a market that trades this cargo. The ship must be docked in a waypoint that has the \`Marketplace\` trait in order to use this function.
---

# spacetraders_sell_cargo Action

Sell cargo in your ship to a market that trades this cargo. The ship must be docked in a waypoint that has the \`Marketplace\` trait in order to use this function.

## Example Usage

```terraform
action "spacetraders_sell_cargo" "example" {
  config {
    ship_symbol = "example"
    symbol = "example"
    units = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
* `symbol` (String, required)
* `units` (Number, required)
