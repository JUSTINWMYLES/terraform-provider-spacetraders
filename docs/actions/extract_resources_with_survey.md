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

* `deposits` (List(Dynamic), required) - A list of deposits that can be found at this location. A ship will extract one of these deposits when using this survey in an extraction request. If multiple deposits of the same type are present, the chance of extracting that deposit is increased.
* `expiration` (String, required) - The date and time when the survey expires. After this date and time, the survey will no longer be available for extraction.
* `ship_symbol` (String, required) - The symbol of the ship.
* `signature` (String, required) - A unique signature for the location of this survey. This signature is verified when attempting an extraction using this survey.
* `size` (String, required) - The size of the deposit. This value indicates how much can be extracted from the survey before it is exhausted.
* `symbol` (String, required) - The symbol of the waypoint that this survey is for.
