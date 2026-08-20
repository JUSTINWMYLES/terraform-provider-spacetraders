---
page_title: "spacetraders_create_ship_waypoint_scan Action - spacetraders"
subcategory: ""
description: |-
  Scan for nearby waypoints, retrieving detailed information on each waypoint in range. Scanning uncharted waypoints will allow you to ignore their uncharted state and will list the waypoints' traits.  Requires a ship to have the \`Sensor Array\` mount installed to use.  The ship will enter a cooldown after using this function, during which it cannot execute certain actions.
---

# spacetraders_create_ship_waypoint_scan Action

Scan for nearby waypoints, retrieving detailed information on each waypoint in range. Scanning uncharted waypoints will allow you to ignore their uncharted state and will list the waypoints' traits.  Requires a ship to have the \`Sensor Array\` mount installed to use.  The ship will enter a cooldown after using this function, during which it cannot execute certain actions.

## Example Usage

```terraform
action "spacetraders_create_ship_waypoint_scan" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
