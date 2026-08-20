---
page_title: "spacetraders_siphon_resources Action - spacetraders"
subcategory: ""
description: |-
  Siphon gases or other resources from gas giants.  The ship must be in orbit to be able to siphon and must have siphon mounts and a gas processor installed.
---

# spacetraders_siphon_resources Action

Siphon gases or other resources from gas giants.  The ship must be in orbit to be able to siphon and must have siphon mounts and a gas processor installed.

## Example Usage

```terraform
action "spacetraders_siphon_resources" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required) - The symbol of the ship.
