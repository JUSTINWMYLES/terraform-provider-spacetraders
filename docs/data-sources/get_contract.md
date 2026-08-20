---
page_title: "spacetraders_get_contract Data Source - spacetraders"
subcategory: ""
description: |-
  Get the details of a specific contract.
---

# spacetraders_get_contract Data Source

Get the details of a specific contract.

## Example Usage

```terraform
data "spacetraders_get_contract" "example" {
  contract_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `contract_id` (String, required) - The contract ID to accept.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `accepted` (Bool, computed) - Whether the contract has been accepted by the agent
* `deadline_to_accept` (String, computed) - The time at which the contract is no longer available to be accepted
* `expiration` (String, computed) - Deprecated in favor of deadlineToAccept
* `faction_symbol` (String, computed) - The symbol of the faction that this contract is for.
* `fulfilled` (Bool, computed) - Whether the contract has been fulfilled
* `id` (String, computed) - ID of the contract.
* `terms` (Object({deadline, deliver, payment}), computed) - The terms to fulfill the contract.
  * `deadline` (String, computed) - The deadline for the contract.
  * `deliver` (List(Object({destination_symbol, trade_symbol, units_fulfilled, units_required})), computed) - The cargo that needs to be delivered to fulfill the contract.
  * `payment` (Object({on_accepted, on_fulfilled}), computed) - Payments for the contract.
    * `on_accepted` (Number, computed) - The amount of credits received up front for accepting the contract.
    * `on_fulfilled` (Number, computed) - The amount of credits received when the contract is fulfilled.
* `type` (String, computed) - Type of contract.

