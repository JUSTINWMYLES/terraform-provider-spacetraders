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

* `agent_symbol` (String, required) - The agent symbol

### Attributes

In addition to all arguments above, the following attributes are exported:

* `credits` (Number, computed) - The number of credits the agent has available. Credits can be negative if funds have been overdrawn.
* `headquarters` (String, computed) - The headquarters of the agent.
* `ship_count` (Number, computed) - How many ships are owned by the agent.
* `starting_faction` (String, computed) - The faction the agent started with.
* `symbol` (String, computed) - Symbol of the agent.

