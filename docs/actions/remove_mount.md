---
page_title: "spacetraders_remove_mount Action - spacetraders"
subcategory: ""
description: |-
  Remove a mount from a ship.  The ship must be docked in a waypoint that has the \`Shipyard\` trait, and must have the desired mount that it wish to remove installed.  A removal fee will be deduced from the agent by the Shipyard.
---

# spacetraders_remove_mount Action

Remove a mount from a ship.  The ship must be docked in a waypoint that has the \`Shipyard\` trait, and must have the desired mount that it wish to remove installed.  A removal fee will be deduced from the agent by the Shipyard.

## Example Usage

```terraform
action "spacetraders_remove_mount" "example" {
  config {
    ship_symbol = "example"
    symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
* `symbol` (String, required)
