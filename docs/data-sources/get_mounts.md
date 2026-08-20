---
page_title: "spacetraders_get_mounts Data Source - spacetraders"
subcategory: ""
description: |-
  Get the mounts installed on a ship.
---

# spacetraders_get_mounts Data Source

Get the mounts installed on a ship.

## Example Usage

```terraform
data "spacetraders_get_mounts" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({deposits, description, name, requirements, strength, symbol})), computed)

