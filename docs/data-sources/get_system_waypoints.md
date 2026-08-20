---
page_title: "spacetraders_get_system_waypoints Data Source - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all of the waypoints for a given system.  If a waypoint is uncharted, it will return the \`Uncharted\` trait instead of its actual traits.
---

# spacetraders_get_system_waypoints Data Source

Return a paginated list of all of the waypoints for a given system.  If a waypoint is uncharted, it will return the \`Uncharted\` trait instead of its actual traits.

## Example Usage

```terraform
data "spacetraders_get_system_waypoints" "example" {
  limit = null
  page = null
  system_symbol = null
  traits = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional) - How many entries to return per page
* `page` (Number, optional) - What entry offset to request
* `system_symbol` (String, required)
* `traits` (String, optional) - Filter waypoints by one or more traits.
* `type` (String, optional) - Filter waypoints by type.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({chart, faction, is_under_construction, modifiers, orbitals, orbits, symbol, system_symbol, traits, type, x, y})), computed)

