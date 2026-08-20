---
page_title: "spacetraders_negotiate_contract Action - spacetraders"
subcategory: ""
description: |-
  Negotiate a new contract with the HQ.  In order to negotiate a new contract, an agent must not have ongoing or offered contracts over the allowed maximum amount. Currently the maximum contracts an agent can have at a time is 1.  Once a contract is negotiated, it is added to the list of contracts offered to the agent, which the agent can then accept.   The ship must be present at any waypoint with a faction present to negotiate a contract with that faction.
---

# spacetraders_negotiate_contract Action

Negotiate a new contract with the HQ.  In order to negotiate a new contract, an agent must not have ongoing or offered contracts over the allowed maximum amount. Currently the maximum contracts an agent can have at a time is 1.  Once a contract is negotiated, it is added to the list of contracts offered to the agent, which the agent can then accept.   The ship must be present at any waypoint with a faction present to negotiate a contract with that faction.

## Example Usage

```terraform
action "spacetraders_negotiate_contract" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
