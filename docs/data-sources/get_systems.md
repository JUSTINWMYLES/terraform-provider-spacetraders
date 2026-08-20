---
page_title: "spacetraders_get_systems Data Source - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all systems.
---

# spacetraders_get_systems Data Source

Return a paginated list of all systems.

## Example Usage

```terraform
data "spacetraders_get_systems" "example" {
  limit = null
  page = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional) - How many entries to return per page
* `page` (Number, optional) - What entry offset to request

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({constellation, factions, name, sector_symbol, symbol, type, waypoints, x, y})), computed)

