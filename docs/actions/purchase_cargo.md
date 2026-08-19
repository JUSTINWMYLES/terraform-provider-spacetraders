---
page_title: "spacetraders_purchase_cargo Action - spacetraders"
subcategory: ""
description: |-
  Purchase cargo from a market.  The ship must be docked in a waypoint that has \`Marketplace\` trait, and the market must be selling a good to be able to purchase it.  The maximum amount of units of a good that can be purchased in each transaction are denoted by the \`tradeVolume\` value of the good, which can be viewed by using the Get Market action.  Purchased goods are added to the ship's cargo hold.
---

# spacetraders_purchase_cargo Action

Purchase cargo from a market.  The ship must be docked in a waypoint that has \`Marketplace\` trait, and the market must be selling a good to be able to purchase it.  The maximum amount of units of a good that can be purchased in each transaction are denoted by the \`tradeVolume\` value of the good, which can be viewed by using the Get Market action.  Purchased goods are added to the ship's cargo hold.

## Example Usage

```terraform
action "spacetraders_purchase_cargo" "example" {
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
