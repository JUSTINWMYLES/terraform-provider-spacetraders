---
page_title: "spacetraders_get_my_ship_cargo Data Source - spacetraders"
subcategory: ""
description: |-
  Retrieve the cargo of a ship under your agent's ownership.
---

# spacetraders_get_my_ship_cargo Data Source

Retrieve the cargo of a ship under your agent's ownership.

## Example Usage

```terraform
data "spacetraders_get_my_ship_cargo" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `capacity` (Number, computed)
* `inventory` (List(Object({description, name, symbol, units})), computed)
* `units` (Number, computed)

