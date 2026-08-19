---
page_title: "spacetraders_repair_ship Action - spacetraders"
subcategory: ""
description: |-
  Repair a ship, restoring the ship to maximum condition. The ship must be docked at a waypoint that has the \`Shipyard\` trait in order to use this function. To preview the cost of repairing the ship, use the Get action.
---

# spacetraders_repair_ship Action

Repair a ship, restoring the ship to maximum condition. The ship must be docked at a waypoint that has the \`Shipyard\` trait in order to use this function. To preview the cost of repairing the ship, use the Get action.

## Example Usage

```terraform
action "spacetraders_repair_ship" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
