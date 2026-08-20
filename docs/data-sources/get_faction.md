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

* `faction_symbol` (String, required) - The faction symbol

### Attributes

In addition to all arguments above, the following attributes are exported:

* `description` (String, computed) - Description of the faction.
* `headquarters` (String, computed) - The waypoint in which the faction's HQ is located in.
* `is_recruiting` (Bool, computed) - Whether or not the faction is currently recruiting new agents.
* `name` (String, computed) - Name of the faction.
* `symbol` (String, computed) - The symbol of the faction.
* `traits` (List(Object({description, name, symbol})), computed) - List of traits that define this faction.

