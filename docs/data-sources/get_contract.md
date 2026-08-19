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

* `contract_id` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `accepted` (Bool, computed)
* `deadline_to_accept` (String, computed)
* `expiration` (String, computed)
* `faction_symbol` (String, computed)
* `fulfilled` (Bool, computed)
* `id` (String, computed)
* `terms` (Object({deadline, deliver, payment}), computed)
  * `deadline` (String, computed)
  * `deliver` (List(Object({destination_symbol, trade_symbol, units_fulfilled, units_required})), computed)
  * `payment` (Object({on_accepted, on_fulfilled}), computed)
    * `on_accepted` (Number, computed)
    * `on_fulfilled` (Number, computed)
* `type` (String, computed)

