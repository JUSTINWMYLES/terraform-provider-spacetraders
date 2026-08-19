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

* `system_symbol` (String, required)
* `waypoint_symbol` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `is_complete` (Bool, computed)
* `materials` (List(Object({fulfilled, required, trade_symbol})), computed)
* `symbol` (String, computed)

