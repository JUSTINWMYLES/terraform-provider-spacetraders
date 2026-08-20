---
page_title: "spacetraders_get_waypoint Data Source - spacetraders"
subcategory: ""
description: |-
  View the details of a waypoint.  If the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.
---

# spacetraders_get_waypoint Data Source

View the details of a waypoint.  If the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.

## Example Usage

```terraform
data "spacetraders_get_waypoint" "example" {
  system_symbol = null
  waypoint_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `system_symbol` (String, required) - The symbol of the system.
* `waypoint_symbol` (String, required) - The waypoint symbol

### Attributes

In addition to all arguments above, the following attributes are exported:

* `chart` (Object({submitted_by, submitted_on, waypoint_symbol}), computed) - The chart of a system or waypoint, which makes the location visible to other agents.
  * `submitted_by` (String, computed) - The agent that submitted the chart for this waypoint.
  * `submitted_on` (String, computed) - The time the chart for this waypoint was submitted.
  * `waypoint_symbol` (String, computed) - The symbol of the waypoint.
* `faction` (Object({symbol}), computed) - The faction that controls the waypoint.
  * `symbol` (String, computed) - The symbol of the faction.
* `is_under_construction` (Bool, computed) - True if the waypoint is under construction.
* `modifiers` (List(Object({description, name, symbol})), computed) - The modifiers of the waypoint.
* `orbitals` (List(Object({symbol})), computed) - Waypoints that orbit this waypoint.
* `orbits` (String, computed) - The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined.
* `symbol` (String, computed) - The symbol of the waypoint.
* `traits` (List(Object({description, name, symbol})), computed) - The traits of the waypoint.
* `type` (String, computed) - The type of waypoint.
* `x` (Number, computed) - Relative position of the waypoint on the system's x axis. This is not an absolute position in the universe.
* `y` (Number, computed) - Relative position of the waypoint on the system's y axis. This is not an absolute position in the universe.

