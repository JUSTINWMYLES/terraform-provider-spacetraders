---
page_title: "spacetraders_get_agent Data Source - spacetraders"
subcategory: ""
description: |-
  Get public details for a specific agent.
---

# spacetraders_get_agent Data Source

Get public details for a specific agent.

## Example Usage

```terraform
data "spacetraders_get_agent" "example" {
  agent_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `agent_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `credits` (Number, computed)
* `headquarters` (String, computed)
* `ship_count` (Number, computed)
* `starting_faction` (String, computed)
* `symbol` (String, computed)

