---
page_title: "spacetraders_deliver_contract Action - spacetraders"
subcategory: ""
description: |-
  Deliver cargo to a contract.  In order to use this API, a ship must be at the delivery location (denoted in the delivery terms as \`destinationSymbol\` of a contract) and must have a number of units of a good required by this contract in its cargo.  Cargo that was delivered will be removed from the ship's cargo.
---

# spacetraders_deliver_contract Action

Deliver cargo to a contract.  In order to use this API, a ship must be at the delivery location (denoted in the delivery terms as \`destinationSymbol\` of a contract) and must have a number of units of a good required by this contract in its cargo.  Cargo that was delivered will be removed from the ship's cargo.

## Example Usage

```terraform
action "spacetraders_deliver_contract" "example" {
  config {
    contract_id = "example"
    ship_symbol = "example"
    trade_symbol = "example"
    units = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `contract_id` (String, required) - The ID of the contract.
* `ship_symbol` (String, required) - Symbol of a ship located in the destination to deliver a contract and that has a good to deliver in its cargo.
* `trade_symbol` (String, required) - The symbol of the good to deliver.
* `units` (Number, required) - Amount of units to deliver.
