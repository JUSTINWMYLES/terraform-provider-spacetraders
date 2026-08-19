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

* `ship_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `expiration` (String, computed)
* `remaining_seconds` (Number, computed)
* `total_seconds` (Number, computed)

