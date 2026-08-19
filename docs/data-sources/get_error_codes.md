---
page_title: "spacetraders_get_error_codes Data Source - spacetraders"
subcategory: ""
description: |-
  Return a list of all possible error codes thrown by the game server.
---

# spacetraders_get_error_codes Data Source

Return a list of all possible error codes thrown by the game server.

## Example Usage

```terraform
data "spacetraders_get_error_codes" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({code, name})), computed)

