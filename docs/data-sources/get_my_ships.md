---
page_title: "spacetraders_get_my_ships Data Source - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all of ships under your agent's ownership.
---

# spacetraders_get_my_ships Data Source

Return a paginated list of all of ships under your agent's ownership.

## Example Usage

```terraform
data "spacetraders_get_my_ships" "example" {
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

* `items` (List(Object({cargo, cooldown, crew, engine, frame, fuel, modules, mounts, nav, reactor, registration, symbol})), computed)

