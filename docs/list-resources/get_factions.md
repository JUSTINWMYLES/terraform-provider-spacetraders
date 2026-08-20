---
page_title: "spacetraders_get_factions List Resource - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all the factions in the game.
---

# spacetraders_get_factions List Resource

Return a paginated list of all the factions in the game.

## Example Usage

```terraform
list "spacetraders_get_factions" "example" {
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

* `faction_symbol` (String, computed)
