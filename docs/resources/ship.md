---
page_title: "spacetraders_ship Resource - spacetraders"
subcategory: ""
description: |-
  Retrieve the details of a ship under your agent's ownership.
---

# spacetraders_ship Resource

Retrieve the details of a ship under your agent's ownership.

## Example Usage

```terraform
resource "spacetraders_ship" "example" {
  ship_type = null
  waypoint_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_type` (String, required)
* `waypoint_symbol` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cargo` (Object({capacity, inventory, units}), computed)
  * `capacity` (Number, computed)
  * `inventory` (List(Object({description, name, symbol, units})), computed)
  * `units` (Number, computed)
* `cooldown` (Object({expiration, remaining_seconds, ship_symbol, total_seconds}), computed)
  * `expiration` (String, computed)
  * `remaining_seconds` (Number, computed)
  * `ship_symbol` (String, computed)
  * `total_seconds` (Number, computed)
* `crew` (Object({capacity, current, morale, required, rotation, wages}), computed)
  * `capacity` (Number, computed)
  * `current` (Number, computed)
  * `morale` (Number, computed)
  * `required` (Number, computed)
  * `rotation` (String, computed)
  * `wages` (Number, computed)
* `engine` (Object({condition, description, integrity, name, quality, requirements, speed, symbol}), computed)
  * `condition` (Number, computed)
  * `description` (String, computed)
  * `integrity` (Number, computed)
  * `name` (String, computed)
  * `quality` (Number, computed)
  * `requirements` (Object({crew, power, slots}), computed)
    * `crew` (Number, computed)
    * `power` (Number, computed)
    * `slots` (Number, computed)
  * `speed` (Number, computed)
  * `symbol` (String, computed)
* `frame` (Object({condition, description, fuel_capacity, integrity, module_slots, mounting_points, name, quality, requirements, symbol}), computed)
  * `condition` (Number, computed)
  * `description` (String, computed)
  * `fuel_capacity` (Number, computed)
  * `integrity` (Number, computed)
  * `module_slots` (Number, computed)
  * `mounting_points` (Number, computed)
  * `name` (String, computed)
  * `quality` (Number, computed)
  * `requirements` (Object({crew, power, slots}), computed)
    * `crew` (Number, computed)
    * `power` (Number, computed)
    * `slots` (Number, computed)
  * `symbol` (String, computed)
* `fuel` (Object({capacity, consumed, current}), computed)
  * `capacity` (Number, computed)
  * `consumed` (Object({amount, timestamp}), computed)
    * `amount` (Number, computed)
    * `timestamp` (String, computed)
  * `current` (Number, computed)
* `id` (String, computed)
* `modules` (List(Object({capacity, description, name, range, requirements, symbol})), computed)
* `mounts` (List(Object({deposits, description, name, requirements, strength, symbol})), computed)
* `nav` (Object({flight_mode, route, status, system_symbol, waypoint_symbol}), computed)
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
* `reactor` (Object({condition, description, integrity, name, power_output, quality, requirements, symbol}), computed)
  * `condition` (Number, computed)
  * `description` (String, computed)
  * `integrity` (Number, computed)
  * `name` (String, computed)
  * `power_output` (Number, computed)
  * `quality` (Number, computed)
  * `requirements` (Object({crew, power, slots}), computed)
    * `crew` (Number, computed)
    * `power` (Number, computed)
    * `slots` (Number, computed)
  * `symbol` (String, computed)
* `registration` (Object({faction_symbol, name, role}), computed)
  * `faction_symbol` (String, computed)
  * `name` (String, computed)
  * `role` (String, computed)
* `symbol` (String, computed)

