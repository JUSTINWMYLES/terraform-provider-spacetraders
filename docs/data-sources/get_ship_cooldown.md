---
page_title: "spacetraders_get_ship_cooldown Data Source - spacetraders"
subcategory: ""
description: |-
  Retrieve the details of your ship's reactor cooldown. Some actions such as activating your jump drive, scanning, or extracting resources taxes your reactor and results in a cooldown.  Your ship cannot perform additional actions until your cooldown has expired. The duration of your cooldown is relative to the power consumption of the related modules or mounts for the action taken.  Response returns a 204 status code (no-content) when the ship has no cooldown.
---

# spacetraders_get_ship_cooldown Data Source

Retrieve the details of your ship's reactor cooldown. Some actions such as activating your jump drive, scanning, or extracting resources taxes your reactor and results in a cooldown.  Your ship cannot perform additional actions until your cooldown has expired. The duration of your cooldown is relative to the power consumption of the related modules or mounts for the action taken.  Response returns a 204 status code (no-content) when the ship has no cooldown.

## Example Usage

```terraform
data "spacetraders_get_ship_cooldown" "example" {
  ship_symbol = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship that is on cooldown

### Attributes

In addition to all arguments above, the following attributes are exported:

* `expiration` (String, computed) - The date and time when the cooldown expires in ISO 8601 format
* `remaining_seconds` (Number, computed) - The remaining duration of the cooldown in seconds
* `total_seconds` (Number, computed) - The total duration of the cooldown in seconds

