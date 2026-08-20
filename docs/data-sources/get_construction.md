---
page_title: "spacetraders_get_construction Data Source - spacetraders"
subcategory: ""
description: |-
  Get construction details for a waypoint. Requires a waypoint with a property of \`isUnderConstruction\` to be true.
---

# spacetraders_get_construction Data Source

Get construction details for a waypoint. Requires a waypoint with a property of \`isUnderConstruction\` to be true.

## Example Usage

```terraform
data "spacetraders_get_construction" "example" {
  system_symbol = null
  waypoint_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `system_symbol` (String, required) - The system symbol
* `waypoint_symbol` (String, required) - The waypoint symbol

### Attributes

In addition to all arguments above, the following attributes are exported:

* `is_complete` (Bool, computed) - Whether the waypoint has been constructed.
* `materials` (List(Object({fulfilled, required, trade_symbol})), computed) - The materials required to construct the waypoint.
* `symbol` (String, computed) - The symbol of the waypoint.

