---
page_title: "spacetraders_install_ship_module Action - spacetraders"
subcategory: ""
description: |-
  Install a module on a ship. The module must be in your cargo.
---

# spacetraders_install_ship_module Action

Install a module on a ship. The module must be in your cargo.

## Example Usage

```terraform
action "spacetraders_install_ship_module" "example" {
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
* `symbol` (String, required) - The symbol of the module to install.
