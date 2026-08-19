---
page_title: "spacetraders_get_my_agent_events Data Source - spacetraders"
subcategory: ""
description: |-
  Get recent events for your agent.
---

# spacetraders_get_my_agent_events Data Source

Get recent events for your agent.

## Example Usage

```terraform
data "spacetraders_get_my_agent_events" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({created_at, data, id, message, type})), computed)

