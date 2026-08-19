---
page_title: "spacetraders_ship_refine Action - spacetraders"
subcategory: ""
description: |-
  Attempt to refine the raw materials on your ship. The request will only succeed if your ship is capable of refining at the time of the request. In order to be able to refine, a ship must have goods that can be refined and have installed a \`Refinery\` module that can refine it.  When refining, 100 basic goods will be converted into 10 processed goods.
---

# spacetraders_ship_refine Action

Attempt to refine the raw materials on your ship. The request will only succeed if your ship is capable of refining at the time of the request. In order to be able to refine, a ship must have goods that can be refined and have installed a \`Refinery\` module that can refine it.  When refining, 100 basic goods will be converted into 10 processed goods.

## Example Usage

```terraform
action "spacetraders_ship_refine" "example" {
  config {
    produce = "example"
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `produce` (String, required)
* `ship_symbol` (String, required)
