---
page_title: "spacetraders_get_system_waypoints List Resource - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all of the waypoints for a given system.  If a waypoint is uncharted, it will return the \`Uncharted\` trait instead of its actual traits.
---

# spacetraders_get_system_waypoints List Resource

Return a paginated list of all of the waypoints for a given system.  If a waypoint is uncharted, it will return the \`Uncharted\` trait instead of its actual traits.

## Example Usage

```terraform
list "spacetraders_get_system_waypoints" "example" {
  provider = spacetraders
  limit = 100
  config {
    limit = 1
    page = 1
    system_symbol = "example"
    traits = "example"
    type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional) - How many entries to return per page
* `page` (Number, optional) - What entry offset to request
* `system_symbol` (String, required)
* `traits` (String, optional) - Filter waypoints by one or more traits.
* `type` (String, optional) - Filter waypoints by type.

### Identity Attributes

The following identity attributes are exported for each matching result:

* `system_symbol` (String, computed)
* `waypoint_symbol` (String, computed)
