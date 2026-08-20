---
page_title: "spacetraders_get_my_factions Data Source - spacetraders"
subcategory: ""
description: |-
  Retrieve factions with which the agent has reputation.
---

# spacetraders_get_my_factions Data Source

Retrieve factions with which the agent has reputation.

## Example Usage

```terraform
data "spacetraders_get_my_factions" "example" {
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

* `items` (List(Object({reputation, symbol})), computed)

