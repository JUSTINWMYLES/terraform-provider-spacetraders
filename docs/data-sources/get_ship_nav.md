---
page_title: "spacetraders_get_ship_nav Data Source - spacetraders"
subcategory: ""
description: |-
  Get the current nav status of a ship.
---

# spacetraders_get_ship_nav Data Source

Get the current nav status of a ship.

## Example Usage

```terraform
data "spacetraders_get_ship_nav" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `flight_mode` (String, computed) - The ship's set speed when traveling between waypoints or systems.
* `route` (Object({arrival, departure_time, destination, origin}), computed) - The routing information for the ship's most recent transit or current location.
  * `arrival` (String, computed) - The date time of the ship's arrival. If the ship is in-transit, this is the expected time of arrival.
  * `departure_time` (String, computed) - The date time of the ship's departure.
  * `destination` (Object({symbol, system_symbol, type, x, y}), computed) - The destination or departure of a ships nav route.
    * `symbol` (String, computed) - The symbol of the waypoint.
    * `system_symbol` (String, computed) - The symbol of the system.
    * `type` (String, computed) - The type of waypoint.
    * `x` (Number, computed) - Position in the universe in the x axis.
    * `y` (Number, computed) - Position in the universe in the y axis.
  * `origin` (Object({symbol, system_symbol, type, x, y}), computed) - The destination or departure of a ships nav route.
    * `symbol` (String, computed) - The symbol of the waypoint.
    * `system_symbol` (String, computed) - The symbol of the system.
    * `type` (String, computed) - The type of waypoint.
    * `x` (Number, computed) - Position in the universe in the x axis.
    * `y` (Number, computed) - Position in the universe in the y axis.
* `status` (String, computed) - The current status of the ship
* `system_symbol` (String, computed) - The symbol of the system.
* `waypoint_symbol` (String, computed) - The symbol of the waypoint.

