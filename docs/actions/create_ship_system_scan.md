---
page_title: "spacetraders_create_ship_system_scan Action - spacetraders"
subcategory: ""
description: |-
  Scan for nearby systems, retrieving information on the systems' distance from the ship and their waypoints. Requires a ship to have the \`Sensor Array\` mount installed to use.  The ship will enter a cooldown after using this function, during which it cannot execute certain actions.
---

# spacetraders_create_ship_system_scan Action

Scan for nearby systems, retrieving information on the systems' distance from the ship and their waypoints. Requires a ship to have the \`Sensor Array\` mount installed to use.  The ship will enter a cooldown after using this function, during which it cannot execute certain actions.

## Example Usage

```terraform
action "spacetraders_create_ship_system_scan" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
