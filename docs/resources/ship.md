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

* `ship_type` (String, required) - Type of ship
* `waypoint_symbol` (String, required) - The symbol of the waypoint you want to purchase the ship at.

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cargo` (Object({capacity, inventory, units}), computed) - Ship cargo details.
  * `capacity` (Number, computed) - The max number of items that can be stored in the cargo hold.
  * `inventory` (List(Object({description, name, symbol, units})), computed) - The items currently in the cargo hold.
  * `units` (Number, computed) - The number of items currently stored in the cargo hold.
* `cooldown` (Object({expiration, remaining_seconds, ship_symbol, total_seconds}), computed) - A cooldown is a period of time in which a ship cannot perform certain actions.
  * `expiration` (String, computed) - The date and time when the cooldown expires in ISO 8601 format
  * `remaining_seconds` (Number, computed) - The remaining duration of the cooldown in seconds
  * `ship_symbol` (String, computed) - The symbol of the ship that is on cooldown
  * `total_seconds` (Number, computed) - The total duration of the cooldown in seconds
* `crew` (Object({capacity, current, morale, required, rotation, wages}), computed) - The ship's crew service and maintain the ship's systems and equipment.
  * `capacity` (Number, computed) - The maximum number of crew members the ship can support.
  * `current` (Number, computed) - The current number of crew members on the ship.
  * `morale` (Number, computed) - A rough measure of the crew's morale. A higher morale means the crew is happier and more productive. A lower morale means the ship is more prone to accidents.
  * `required` (Number, computed) - The minimum number of crew members required to maintain the ship.
  * `rotation` (String, computed) - The rotation of crew shifts. A stricter shift improves the ship's performance. A more relaxed shift improves the crew's morale.
  * `wages` (Number, computed) - The amount of credits per crew member paid per hour. Wages are paid when a ship docks at a civilized waypoint.
* `engine` (Object({condition, description, integrity, name, quality, requirements, speed, symbol}), computed) - The engine determines how quickly a ship travels between waypoints.
  * `condition` (Number, computed) - The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.
  * `description` (String, computed) - The description of the engine.
  * `integrity` (Number, computed) - The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.
  * `name` (String, computed) - The name of the engine.
  * `quality` (Number, computed) - The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.
  * `requirements` (Object({crew, power, slots}), computed) - The requirements for installation on a ship
    * `crew` (Number, computed) - The number of crew required for operation.
    * `power` (Number, computed) - The amount of power required from the reactor.
    * `slots` (Number, computed) - The number of module slots required for installation.
  * `speed` (Number, computed) - The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship.
  * `symbol` (String, computed) - The symbol of the engine.
* `frame` (Object({condition, description, fuel_capacity, integrity, module_slots, mounting_points, name, quality, requirements, symbol}), computed) - The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.
  * `condition` (Number, computed) - The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.
  * `description` (String, computed) - Description of the frame.
  * `fuel_capacity` (Number, computed) - The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount.
  * `integrity` (Number, computed) - The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.
  * `module_slots` (Number, computed) - The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed.
  * `mounting_points` (Number, computed) - The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed.
  * `name` (String, computed) - Name of the frame.
  * `quality` (Number, computed) - The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.
  * `requirements` (Object({crew, power, slots}), computed) - The requirements for installation on a ship
    * `crew` (Number, computed) - The number of crew required for operation.
    * `power` (Number, computed) - The amount of power required from the reactor.
    * `slots` (Number, computed) - The number of module slots required for installation.
  * `symbol` (String, computed) - Symbol of the frame.
* `fuel` (Object({capacity, consumed, current}), computed) - Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.
  * `capacity` (Number, computed) - The maximum amount of fuel the ship's tanks can hold.
  * `consumed` (Object({amount, timestamp}), computed) - An object that only shows up when an action has consumed fuel in the process. Shows the fuel consumption data.
    * `amount` (Number, computed) - The amount of fuel consumed by the most recent transit or action.
    * `timestamp` (String, computed) - The time at which the fuel was consumed.
  * `current` (Number, computed) - The current amount of fuel in the ship's tanks.
* `id` (String, computed)
* `modules` (List(Object({capacity, description, name, range, requirements, symbol})), computed) - Modules installed in this ship.
* `mounts` (List(Object({deposits, description, name, requirements, strength, symbol})), computed) - Mounts installed in this ship.
* `nav` (Object({flight_mode, route, status, system_symbol, waypoint_symbol}), computed) - The navigation information of the ship.
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
* `reactor` (Object({condition, description, integrity, name, power_output, quality, requirements, symbol}), computed) - The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.
  * `condition` (Number, computed) - The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.
  * `description` (String, computed) - Description of the reactor.
  * `integrity` (Number, computed) - The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.
  * `name` (String, computed) - Name of the reactor.
  * `power_output` (Number, computed) - The amount of power provided by this reactor. The more power a reactor provides to the ship, the lower the cooldown it gets when using a module or mount that taxes the ship's power.
  * `quality` (Number, computed) - The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.
  * `requirements` (Object({crew, power, slots}), computed) - The requirements for installation on a ship
    * `crew` (Number, computed) - The number of crew required for operation.
    * `power` (Number, computed) - The amount of power required from the reactor.
    * `slots` (Number, computed) - The number of module slots required for installation.
  * `symbol` (String, computed) - Symbol of the reactor.
* `registration` (Object({faction_symbol, name, role}), computed) - The public registration information of the ship
  * `faction_symbol` (String, computed) - The symbol of the faction the ship is registered with
  * `name` (String, computed) - The agent's registered name of the ship
  * `role` (String, computed) - The registered role of the ship
* `symbol` (String, computed) - The globally unique identifier of the ship in the following format: \`\[AGENT\_SYMBOL\]-\[HEX\_ID\]\`

