---
page_title: "spacetraders_get_my_ships List Resource - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all of ships under your agent's ownership.
---

# spacetraders_get_my_ships List Resource

Return a paginated list of all of ships under your agent's ownership.

## Example Usage

```terraform
list "spacetraders_get_my_ships" "example" {
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

* `limit` (Number, optional)
* `page` (Number, optional)

### Identity Attributes

The following identity attributes are exported for each matching result:

* `ship_symbol` (String, computed)
