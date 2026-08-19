---
page_title: "spacetraders_extract_resources_with_survey Action - spacetraders"
subcategory: ""
description: |-
  Use a survey when extracting resources from a waypoint. This endpoint requires a survey as the payload, which allows your ship to extract specific yields.  Send the full survey object as the payload which will be validated according to the signature. If the signature is invalid, or any properties of the survey are changed, the request will fail.
---

# spacetraders_extract_resources_with_survey Action

Use a survey when extracting resources from a waypoint. This endpoint requires a survey as the payload, which allows your ship to extract specific yields.  Send the full survey object as the payload which will be validated according to the signature. If the signature is invalid, or any properties of the survey are changed, the request will fail.

## Example Usage

```terraform
action "spacetraders_extract_resources_with_survey" "example" {
  config {
    deposits = "example"
    expiration = "example"
    ship_symbol = "example"
    signature = "example"
    size = "example"
    symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `deposits` (List(Dynamic), required)
* `expiration` (String, required)
* `ship_symbol` (String, required)
* `signature` (String, required)
* `size` (String, required)
* `symbol` (String, required)
