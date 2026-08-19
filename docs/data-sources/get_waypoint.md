---
page_title: "spacetraders_get_waypoint Data Source - spacetraders"
subcategory: ""
description: |-
  View the details of a waypoint.  If the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.
---

# spacetraders_get_waypoint Data Source

View the details of a waypoint.  If the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.

## Example Usage

```terraform
data "spacetraders_get_waypoint" "example" {
  system_symbol = null
  waypoint_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `system_symbol` (String, required)
* `waypoint_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `chart` (Object({submitted_by, submitted_on, waypoint_symbol}), computed)
  * `submitted_by` (String, computed)
  * `submitted_on` (String, computed)
  * `waypoint_symbol` (String, computed)
* `faction` (Object({symbol}), computed)
  * `symbol` (String, computed)
* `is_under_construction` (Bool, computed)
* `modifiers` (List(Object({description, name, symbol})), computed)
* `orbitals` (List(Object({symbol})), computed)
* `orbits` (String, computed)
* `symbol` (String, computed)
* `traits` (List(Object({description, name, symbol})), computed)
* `type` (String, computed)
* `x` (Number, computed)
* `y` (Number, computed)

