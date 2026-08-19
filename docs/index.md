---
page_title: "spacetraders Provider"
subcategory: ""
description: |-
  SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using \[OpenAPI\](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  \`\`\`json http {   "method": "GET",   "url": "https://api.spacetraders.io/v2", } \`\`\`  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a \[Discord channel\](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.
---

# spacetraders Provider

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using \[OpenAPI\](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  \`\`\`json http {   "method": "GET",   "url": "https://api.spacetraders.io/v2", } \`\`\`  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a \[Discord channel\](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.

## Resources

- [spacetraders_ship](resources/ship.md)

## Data Sources

- [spacetraders_get_status](data-sources/get_status.md)
- [spacetraders_get_agents](data-sources/get_agents.md)
- [spacetraders_get_agent](data-sources/get_agent.md)
- [spacetraders_get_error_codes](data-sources/get_error_codes.md)
- [spacetraders_get_factions](data-sources/get_factions.md)
- [spacetraders_get_faction](data-sources/get_faction.md)
- [spacetraders_get_supply_chain](data-sources/get_supply_chain.md)
- [spacetraders_get_my_account](data-sources/get_my_account.md)
- [spacetraders_get_my_agent](data-sources/get_my_agent.md)
- [spacetraders_get_my_agent_events](data-sources/get_my_agent_events.md)
- [spacetraders_get_contracts](data-sources/get_contracts.md)
- [spacetraders_get_contract](data-sources/get_contract.md)
- [spacetraders_get_my_factions](data-sources/get_my_factions.md)
- [spacetraders_get_my_ships](data-sources/get_my_ships.md)
- [spacetraders_get_my_ship_cargo](data-sources/get_my_ship_cargo.md)
- [spacetraders_get_ship_cooldown](data-sources/get_ship_cooldown.md)
- [spacetraders_get_ship_modules](data-sources/get_ship_modules.md)
- [spacetraders_get_mounts](data-sources/get_mounts.md)
- [spacetraders_get_ship_nav](data-sources/get_ship_nav.md)
- [spacetraders_get_repair_ship](data-sources/get_repair_ship.md)
- [spacetraders_get_scrap_ship](data-sources/get_scrap_ship.md)
- [spacetraders_websocket_departure_events](data-sources/websocket_departure_events.md)
- [spacetraders_get_systems](data-sources/get_systems.md)
- [spacetraders_get_system](data-sources/get_system.md)
- [spacetraders_get_system_waypoints](data-sources/get_system_waypoints.md)
- [spacetraders_get_waypoint](data-sources/get_waypoint.md)
- [spacetraders_get_construction](data-sources/get_construction.md)
- [spacetraders_get_jump_gate](data-sources/get_jump_gate.md)
- [spacetraders_get_market](data-sources/get_market.md)
- [spacetraders_get_shipyard](data-sources/get_shipyard.md)

## Actions

- [spacetraders_accept_contract](actions/accept_contract.md)
- [spacetraders_deliver_contract](actions/deliver_contract.md)
- [spacetraders_fulfill_contract](actions/fulfill_contract.md)
- [spacetraders_create_chart](actions/create_chart.md)
- [spacetraders_dock_ship](actions/dock_ship.md)
- [spacetraders_extract_resources](actions/extract_resources.md)
- [spacetraders_extract_resources_with_survey](actions/extract_resources_with_survey.md)
- [spacetraders_jettison](actions/jettison.md)
- [spacetraders_jump_ship](actions/jump_ship.md)
- [spacetraders_install_ship_module](actions/install_ship_module.md)
- [spacetraders_remove_ship_module](actions/remove_ship_module.md)
- [spacetraders_install_mount](actions/install_mount.md)
- [spacetraders_remove_mount](actions/remove_mount.md)
- [spacetraders_patch_ship_nav](actions/patch_ship_nav.md)
- [spacetraders_navigate_ship](actions/navigate_ship.md)
- [spacetraders_negotiate_contract](actions/negotiate_contract.md)
- [spacetraders_orbit_ship](actions/orbit_ship.md)
- [spacetraders_purchase_cargo](actions/purchase_cargo.md)
- [spacetraders_ship_refine](actions/ship_refine.md)
- [spacetraders_refuel_ship](actions/refuel_ship.md)
- [spacetraders_repair_ship](actions/repair_ship.md)
- [spacetraders_create_ship_ship_scan](actions/create_ship_ship_scan.md)
- [spacetraders_create_ship_system_scan](actions/create_ship_system_scan.md)
- [spacetraders_create_ship_waypoint_scan](actions/create_ship_waypoint_scan.md)
- [spacetraders_sell_cargo](actions/sell_cargo.md)
- [spacetraders_siphon_resources](actions/siphon_resources.md)
- [spacetraders_create_survey](actions/create_survey.md)
- [spacetraders_transfer_cargo](actions/transfer_cargo.md)
- [spacetraders_warp_ship](actions/warp_ship.md)
- [spacetraders_register](actions/register.md)
- [spacetraders_supply_construction](actions/supply_construction.md)

## List Resources

- [spacetraders_get_agents](list-resources/get_agents.md)
- [spacetraders_get_factions](list-resources/get_factions.md)
- [spacetraders_get_contracts](list-resources/get_contracts.md)
- [spacetraders_get_my_ships](list-resources/get_my_ships.md)
- [spacetraders_get_systems](list-resources/get_systems.md)
- [spacetraders_get_system_waypoints](list-resources/get_system_waypoints.md)

