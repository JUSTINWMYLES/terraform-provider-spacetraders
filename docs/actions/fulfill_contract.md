---
page_title: "spacetraders_fulfill_contract Action - spacetraders"
subcategory: ""
description: |-
  Fulfill a contract. Can only be used on contracts that have all of their delivery terms fulfilled.
---

# spacetraders_fulfill_contract Action

Fulfill a contract. Can only be used on contracts that have all of their delivery terms fulfilled.

## Example Usage

```terraform
action "spacetraders_fulfill_contract" "example" {
  config {
    contract_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `contract_id` (String, required)
