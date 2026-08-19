---
page_title: "spacetraders_get_status Data Source - spacetraders"
subcategory: ""
description: |-
  Return the status of the game server. This also includes a few global elements, such as announcements, server reset dates and leaderboards.
---

# spacetraders_get_status Data Source

Return the status of the game server. This also includes a few global elements, such as announcements, server reset dates and leaderboards.

## Example Usage

```terraform
data "spacetraders_get_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `announcements` (List(Object({body, title})), computed)
* `description` (String, computed)
* `health` (Object({last_market_update}), computed)
  * `last_market_update` (String, computed)
* `leaderboards` (Object({most_credits, most_submitted_charts}), computed)
  * `most_credits` (List(Object({agent_symbol, credits})), computed)
  * `most_submitted_charts` (List(Object({agent_symbol, chart_count})), computed)
* `links` (List(Object({name, url})), computed)
* `reset_date` (String, computed)
* `server_resets` (Object({frequency, next}), computed)
  * `frequency` (String, computed)
  * `next` (String, computed)
* `stats` (Object({accounts, agents, ships, systems, waypoints}), computed)
  * `accounts` (Number, computed)
  * `agents` (Number, computed)
  * `ships` (Number, computed)
  * `systems` (Number, computed)
  * `waypoints` (Number, computed)
* `status` (String, computed)
* `version` (String, computed)

