---
page_title: "spacetraders_get_contracts List Resource - spacetraders"
subcategory: ""
description: |-
  Return a paginated list of all your contracts.
---

# spacetraders_get_contracts List Resource

Return a paginated list of all your contracts.

## Example Usage

```terraform
list "spacetraders_get_contracts" "example" {
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

* `contract_id` (String, computed)
