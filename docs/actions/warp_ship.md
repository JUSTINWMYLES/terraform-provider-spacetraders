---
page_title: "spacetraders_warp_ship Action - spacetraders"
subcategory: ""
description: |-
  Warp your ship to a target destination in another system. The ship must be in orbit to use this function and must have the \`Warp Drive\` module installed. Warping will consume the necessary fuel from the ship's manifest.  The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.  To travel between systems, see the ship's Warp or Jump actions.
---

# spacetraders_warp_ship Action

Warp your ship to a target destination in another system. The ship must be in orbit to use this function and must have the \`Warp Drive\` module installed. Warping will consume the necessary fuel from the ship's manifest.  The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.  To travel between systems, see the ship's Warp or Jump actions.

## Example Usage

```terraform
action "spacetraders_warp_ship" "example" {
  config {
    ship_symbol = "example"
    waypoint_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
* `waypoint_symbol` (String, required)
