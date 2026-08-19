---
page_title: "spacetraders_get_agents List Resource - spacetraders"
subcategory: ""
description: |-
  List all public agent details.
---

# spacetraders_get_agents List Resource

List all public agent details.

## Example Usage

```terraform
list "spacetraders_get_agents" "example" {
  provider = spacetraders
  limit = 100
  config {
    limit = 1
    page = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `limit` (Number, optional)
* `page` (Number, optional)

### Identity Attributes

The following identity attributes are exported for each matching result:

* `agent_symbol` (String, computed)
