---
page_title: "spacetraders_jettison Action - spacetraders"
subcategory: ""
description: |-
  Jettison cargo from your ship's cargo hold.
---

# spacetraders_jettison Action

Jettison cargo from your ship's cargo hold.

## Example Usage

```terraform
action "spacetraders_jettison" "example" {
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

* `ship_symbol` (String, required) - The symbol of the ship.
* `symbol` (String, required) - The good's symbol.
* `units` (Number, required) - Amount of units to jettison of this good.
