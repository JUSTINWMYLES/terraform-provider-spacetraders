---
page_title: "spacetraders_get_my_account Data Source - spacetraders"
subcategory: ""
description: |-
  Fetch your account details.
---

# spacetraders_get_my_account Data Source

Fetch your account details.

## Example Usage

```terraform
data "spacetraders_get_my_account" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `account` (Object({created_at, email, id, token}), computed)
  * `created_at` (String, computed)
  * `email` (String, computed)
  * `id` (String, computed)
  * `token` (String, computed)

