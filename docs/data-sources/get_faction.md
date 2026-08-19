---
page_title: "spacetraders_get_faction Data Source - spacetraders"
subcategory: ""
description: |-
  View the details of a faction.
---

# spacetraders_get_faction Data Source

View the details of a faction.

## Example Usage

```terraform
data "spacetraders_get_faction" "example" {
  faction_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `faction_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `description` (String, computed)
* `headquarters` (String, computed)
* `is_recruiting` (Bool, computed)
* `name` (String, computed)
* `symbol` (String, computed)
* `traits` (List(Object({description, name, symbol})), computed)

