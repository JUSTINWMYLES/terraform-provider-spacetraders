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

* `ship_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `flight_mode` (String, computed)
* `route` (Object({arrival, departure_time, destination, origin}), computed)
  * `arrival` (String, computed)
  * `departure_time` (String, computed)
  * `destination` (Object({symbol, system_symbol, type, x, y}), computed)
    * `symbol` (String, computed)
    * `system_symbol` (String, computed)
    * `type` (String, computed)
    * `x` (Number, computed)
    * `y` (Number, computed)
  * `origin` (Object({symbol, system_symbol, type, x, y}), computed)
    * `symbol` (String, computed)
    * `system_symbol` (String, computed)
    * `type` (String, computed)
    * `x` (Number, computed)
    * `y` (Number, computed)
* `status` (String, computed)
* `system_symbol` (String, computed)
* `waypoint_symbol` (String, computed)

