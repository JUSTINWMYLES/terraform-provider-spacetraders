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

* `account_id` (String, computed) - Account ID that is tied to this agent. Only included on your own agent.
* `credits` (Number, computed) - The number of credits the agent has available. Credits can be negative if funds have been overdrawn.
* `headquarters` (String, computed) - The headquarters of the agent.
* `ship_count` (Number, computed) - How many ships are owned by the agent.
* `starting_faction` (String, computed) - The faction the agent started with.
* `symbol` (String, computed) - Symbol of the agent.

