---
page_title: "spacetraders_navigate_ship Action - spacetraders"
subcategory: ""
description: |-
  Navigate to a target destination. The ship must be in orbit to use this function. The destination waypoint must be within the same system as the ship's current location. Navigating will consume the necessary fuel from the ship's manifest based on the distance to the target waypoint.  The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.  To travel between systems, see the ship's Warp or Jump actions.
---

# spacetraders_navigate_ship Action

Navigate to a target destination. The ship must be in orbit to use this function. The destination waypoint must be within the same system as the ship's current location. Navigating will consume the necessary fuel from the ship's manifest based on the distance to the target waypoint.  The returned response will detail the route information including the expected time of arrival. Most ship actions are unavailable until the ship has arrived at it's destination.  To travel between systems, see the ship's Warp or Jump actions.

## Example Usage

```terraform
action "spacetraders_navigate_ship" "example" {
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
