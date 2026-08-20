---
page_title: "spacetraders_get_systems List Resource - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all systems.
---

# spacetraders_get_systems List Resource

Return a paginated list of all systems.

## Example Usage

```terraform
list "spacetraders_get_systems" "example" {
  provider = spacetraders
  limit = 100
  config {
    limit = 1
    page = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional) - How many entries to return per page
* `page` (Number, optional) - What entry offset to request

### Identity Attributes

The following identity attributes are exported for each matching result:

* `system_symbol` (String, computed)
