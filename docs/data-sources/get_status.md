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
  * `last_market_update` (String, computed) - The date/time when the market was last updated.
* `leaderboards` (Object({most_credits, most_submitted_charts}), computed)
  * `most_credits` (List(Object({agent_symbol, credits})), computed) - Top agents with the most credits.
  * `most_submitted_charts` (List(Object({agent_symbol, chart_count})), computed) - Top agents with the most charted submitted.
* `links` (List(Object({name, url})), computed)
* `reset_date` (String, computed) - The date when the game server was last reset.
* `server_resets` (Object({frequency, next}), computed)
  * `frequency` (String, computed) - How often we intend to reset the game server.
  * `next` (String, computed) - The date and time when the game server will reset.
* `stats` (Object({accounts, agents, ships, systems, waypoints}), computed)
  * `accounts` (Number, computed) - Total number of accounts registered on the game server.
  * `agents` (Number, computed) - Number of registered agents in the game.
  * `ships` (Number, computed) - Total number of ships in the game.
  * `systems` (Number, computed) - Total number of systems in the game.
  * `waypoints` (Number, computed) - Total number of waypoints in the game.
* `status` (String, computed) - The current status of the game server.
* `version` (String, computed) - The current version of the API.

