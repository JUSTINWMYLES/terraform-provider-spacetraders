---
page_title: "spacetraders_jump_ship Action - spacetraders"
subcategory: ""
description: |-
  Jump your ship instantly to a target connected waypoint. The ship must be in orbit to execute a jump.  A unit of antimatter is purchased and consumed from the market when jumping. The price of antimatter is determined by the market and is subject to change. A ship can only jump to connected waypoints
---

# spacetraders_jump_ship Action

Jump your ship instantly to a target connected waypoint. The ship must be in orbit to execute a jump.  A unit of antimatter is purchased and consumed from the market when jumping. The price of antimatter is determined by the market and is subject to change. A ship can only jump to connected waypoints

## Example Usage

```terraform
action "spacetraders_jump_ship" "example" {
  config {
    ship_symbol = "example"
    waypoint_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
* `waypoint_symbol` (String, required) - The symbol of the waypoint to jump to. The destination must be a connected waypoint.
