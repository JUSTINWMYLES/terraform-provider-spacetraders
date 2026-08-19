---
page_title: "spacetraders_get_my_agent Data Source - spacetraders"
subcategory: ""
description: |-
  Fetch your agent's details.
---

# spacetraders_get_my_agent Data Source

Fetch your agent's details.

## Example Usage

```terraform
data "spacetraders_get_my_agent" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `account_id` (String, computed)
* `credits` (Number, computed)
* `headquarters` (String, computed)
* `ship_count` (Number, computed)
* `starting_faction` (String, computed)
* `symbol` (String, computed)

