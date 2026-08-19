---
page_title: "spacetraders_extract_resources Action - spacetraders"
subcategory: ""
description: |-
  Extract resources from a waypoint that can be extracted, such as asteroid fields, into your ship. Send an optional survey as the payload to target specific yields.  The ship must be in orbit to be able to extract and must have mining equipments installed that can extract goods, such as the \`Gas Siphon\` mount for gas-based goods or \`Mining Laser\` mount for ore-based goods.  The survey property is now deprecated. See the \`extract/survey\` endpoint for more details.
---

# spacetraders_extract_resources Action

Extract resources from a waypoint that can be extracted, such as asteroid fields, into your ship. Send an optional survey as the payload to target specific yields.  The ship must be in orbit to be able to extract and must have mining equipments installed that can extract goods, such as the \`Gas Siphon\` mount for gas-based goods or \`Mining Laser\` mount for ore-based goods.  The survey property is now deprecated. See the \`extract/survey\` endpoint for more details.

## Example Usage

```terraform
action "spacetraders_extract_resources" "example" {
  config {
    ship_symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ship_symbol` (String, required)
