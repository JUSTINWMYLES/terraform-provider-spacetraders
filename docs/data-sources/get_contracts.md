---
page_title: "spacetraders_get_contracts Data Source - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all your contracts.
---

# spacetraders_get_contracts Data Source

Return a paginated list of all your contracts.

## Example Usage

```terraform
data "spacetraders_get_contracts" "example" {
  limit = null
  page = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional)
* `page` (Number, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({accepted, deadline_to_accept, expiration, faction_symbol, fulfilled, id, terms, type})), computed)

