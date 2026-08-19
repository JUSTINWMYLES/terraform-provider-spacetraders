---
page_title: "spacetraders_get_agents Data Source - spacetraders"
subcategory: ""
description: |-
  List all public agent details.
---

# spacetraders_get_agents Data Source

List all public agent details.

## Example Usage

```terraform
data "spacetraders_get_agents" "example" {
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

* `items` (List(Object({credits, headquarters, ship_count, starting_faction, symbol})), computed)

