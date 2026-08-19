---
page_title: "spacetraders_get_ship_modules Data Source - spacetraders"
subcategory: ""
description: |-
  Get the modules installed on a ship.
---

# spacetraders_get_ship_modules Data Source

Get the modules installed on a ship.

## Example Usage

```terraform
data "spacetraders_get_ship_modules" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({capacity, description, name, range, requirements, symbol})), computed)

