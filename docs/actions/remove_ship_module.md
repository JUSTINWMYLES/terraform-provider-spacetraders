---
page_title: "spacetraders_remove_ship_module Action - spacetraders"
subcategory: ""
description: |-
  Remove a module from a ship. The module will be placed in cargo.
---

# spacetraders_remove_ship_module Action

Remove a module from a ship. The module will be placed in cargo.

## Example Usage

```terraform
action "spacetraders_remove_ship_module" "example" {
  config {
    ship_symbol = "example"
    symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
* `symbol` (String, required) - The symbol of the module to remove.
