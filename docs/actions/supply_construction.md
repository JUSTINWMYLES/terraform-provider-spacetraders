---
page_title: "spacetraders_supply_construction Action - spacetraders"
subcategory: ""
description: |-
  Supply a construction site with the specified good. Requires a waypoint with a property of \`isUnderConstruction\` to be true.  The good must be in your ship's cargo. The good will be removed from your ship's cargo and added to the construction site's materials.
---

# spacetraders_supply_construction Action

Supply a construction site with the specified good. Requires a waypoint with a property of \`isUnderConstruction\` to be true.  The good must be in your ship's cargo. The good will be removed from your ship's cargo and added to the construction site's materials.

## Example Usage

```terraform
action "spacetraders_supply_construction" "example" {
  config {
    ship_symbol = "example"
    system_symbol = "example"
    trade_symbol = "example"
    units = 1
    waypoint_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
* `system_symbol` (String, required)
* `trade_symbol` (String, required)
* `units` (Number, required)
* `waypoint_symbol` (String, required)
