---
page_title: "spacetraders_patch_ship_nav Action - spacetraders"
subcategory: ""
description: |-
  Update the nav configuration of a ship.  Currently only supports configuring the Flight Mode of the ship, which affects its speed and fuel consumption.
---

# spacetraders_patch_ship_nav Action

Update the nav configuration of a ship.  Currently only supports configuring the Flight Mode of the ship, which affects its speed and fuel consumption.

## Example Usage

```terraform
action "spacetraders_patch_ship_nav" "example" {
  config {
    flight_mode = "example"
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `flight_mode` (String, optional)
* `ship_symbol` (String, required)
