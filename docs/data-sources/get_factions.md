---
page_title: "spacetraders_get_factions Data Source - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all the factions in the game.
---

# spacetraders_get_factions Data Source

Return a paginated list of all the factions in the game.

## Example Usage

```terraform
data "spacetraders_get_factions" "example" {
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

* `items` (List(Object({description, headquarters, is_recruiting, name, symbol, traits})), computed)

