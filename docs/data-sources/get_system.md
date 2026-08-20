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

* `constellation` (String, computed) - The constellation that the system is part of.
* `factions` (List(Object({symbol})), computed) - Factions that control this system.
* `name` (String, computed) - The name of the system.
* `sector_symbol` (String, computed) - The symbol of the sector.
* `symbol` (String, computed) - The symbol of the system.
* `type` (String, computed) - The type of system.
* `waypoints` (List(Object({orbitals, orbits, symbol, type, x, y})), computed) - Waypoints in this system.
* `x` (Number, computed) - Relative position of the system in the sector in the x axis.
* `y` (Number, computed) - Relative position of the system in the sector in the y axis.

