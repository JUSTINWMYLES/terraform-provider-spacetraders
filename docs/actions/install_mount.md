---
page_title: "spacetraders_install_mount Action - spacetraders"
subcategory: ""
description: |-
  Install a mount on a ship.  In order to install a mount, the ship must be docked and located in a waypoint that has a \`Shipyard\` trait. The ship also must have the mount to install in its cargo hold.  An installation fee will be deduced by the Shipyard for installing the mount on the ship.
---

# spacetraders_install_mount Action

Install a mount on a ship.  In order to install a mount, the ship must be docked and located in a waypoint that has a \`Shipyard\` trait. The ship also must have the mount to install in its cargo hold.  An installation fee will be deduced by the Shipyard for installing the mount on the ship.

## Example Usage

```terraform
action "spacetraders_install_mount" "example" {
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
