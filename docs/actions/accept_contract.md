---
page_title: "spacetraders_accept_contract Action - spacetraders"
subcategory: ""
description: |-
  Accept a contract by ID.   You can only accept contracts that were offered to you, were not accepted yet, and whose deadlines has not passed yet.
---

# spacetraders_accept_contract Action

Accept a contract by ID.   You can only accept contracts that were offered to you, were not accepted yet, and whose deadlines has not passed yet.

## Example Usage

```terraform
action "spacetraders_accept_contract" "example" {
  config {
    contract_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `contract_id` (String, required)
