---
page_title: "spacetraders_get_supply_chain Data Source - spacetraders"
subcategory: ""
description: |-
  Describes which import and exports map to each other.
---

# spacetraders_get_supply_chain Data Source

Describes which import and exports map to each other.

## Example Usage

```terraform
data "spacetraders_get_supply_chain" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `export_to_import_map` (Map(List(String)), computed)

