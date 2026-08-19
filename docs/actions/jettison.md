---
page_title: "spacetraders_jettison Action - spacetraders"
subcategory: ""
description: |-
  Jettison cargo from your ship's cargo hold.
---

# spacetraders_jettison Action

Jettison cargo from your ship's cargo hold.

## Example Usage

```terraform
action "spacetraders_jettison" "example" {
  config {
    ship_symbol = "example"
    symbol = "example"
    units = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
* `symbol` (String, required)
* `units` (Number, required)
