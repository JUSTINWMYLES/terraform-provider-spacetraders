---
page_title: "spacetraders_get_system Data Source - spacetraders"
subcategory: ""
description: |-
  Get the details of a system. Requires the system to have been visited or charted.
---

# spacetraders_get_system Data Source

Get the details of a system. Requires the system to have been visited or charted.

## Example Usage

```terraform
data "spacetraders_get_system" "example" {
  system_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `system_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `constellation` (String, computed)
* `factions` (List(Object({symbol})), computed)
* `name` (String, computed)
* `sector_symbol` (String, computed)
* `symbol` (String, computed)
* `type` (String, computed)
* `waypoints` (List(Object({orbitals, orbits, symbol, type, x, y})), computed)
* `x` (Number, computed)
* `y` (Number, computed)

