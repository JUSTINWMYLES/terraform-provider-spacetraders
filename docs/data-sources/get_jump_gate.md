---
page_title: "spacetraders_get_jump_gate Data Source - spacetraders"
subcategory: ""
description: |-
  Get jump gate details for a waypoint. Requires a waypoint of type \`JUMP\_GATE\` to use.  Waypoints connected to this jump gate can be found by querying the waypoints in the system.
---

# spacetraders_get_jump_gate Data Source

Get jump gate details for a waypoint. Requires a waypoint of type \`JUMP\_GATE\` to use.  Waypoints connected to this jump gate can be found by querying the waypoints in the system.

## Example Usage

```terraform
data "spacetraders_get_jump_gate" "example" {
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

* `connections` (List(String), computed) - All the gates that are connected to this waypoint.
* `symbol` (String, computed) - The symbol of the waypoint.

