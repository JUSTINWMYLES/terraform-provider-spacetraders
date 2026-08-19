---
page_title: "spacetraders_websocket_departure_events Data Source - spacetraders"
subcategory: ""
description: |-
  Subscribe to departure events for a system.            \#\# WebSocket Events            The following events are available:            - \`systems.{systemSymbol}.departure\`: A ship has departed from the system.            \#\# Subscribe using a message with the following format:            \`\`\`json           {             "action": "subscribe",             "systemSymbol": "{systemSymbol}"           }           \`\`\`            \#\# Unsubscribe using a message with the following format:            \`\`\`json           {             "action": "unsubscribe",             "systemSymbol": "{systemSymbol}"           }           \`\`\`
---

# spacetraders_websocket_departure_events Data Source

Subscribe to departure events for a system.            \#\# WebSocket Events            The following events are available:            - \`systems.{systemSymbol}.departure\`: A ship has departed from the system.            \#\# Subscribe using a message with the following format:            \`\`\`json           {             "action": "subscribe",             "systemSymbol": "{systemSymbol}"           }           \`\`\`            \#\# Unsubscribe using a message with the following format:            \`\`\`json           {             "action": "unsubscribe",             "systemSymbol": "{systemSymbol}"           }           \`\`\`

## Example Usage

```terraform
data "spacetraders_websocket_departure_events" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:


