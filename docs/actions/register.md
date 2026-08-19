---
page_title: "spacetraders_register Action - spacetraders"
subcategory: ""
description: |-
  Creates a new agent and ties it to an account.  The agent symbol must consist of a 3-14 character string, and will be used to represent your agent. This symbol will prefix the symbol of every ship you own. Agent symbols will be cast to all uppercase characters.  This new agent will be tied to a starting faction of your choice, which determines your starting location, and will be granted an authorization token, a contract with their starting faction, a command ship that can fly across space with advanced capabilities, a small probe ship that can be used for reconnaissance, and 175,000 credits.  > \#\#\#\# Keep your token safe and secure > > Keep careful track of where you store your token. You can generate a new token from our account dashboard, but if someone else gains access to your token they will be able to use it to make API requests on your behalf until the end of the reset.  If you are new to SpaceTraders, It is recommended to register with the COSMIC faction, a faction that is well connected to the rest of the universe. After registering, you should try our interactive \[quickstart guide\](https://docs.spacetraders.io/quickstart/new-game) which will walk you through a few basic API requests in just a few minutes.
---

# spacetraders_register Action

Creates a new agent and ties it to an account.  The agent symbol must consist of a 3-14 character string, and will be used to represent your agent. This symbol will prefix the symbol of every ship you own. Agent symbols will be cast to all uppercase characters.  This new agent will be tied to a starting faction of your choice, which determines your starting location, and will be granted an authorization token, a contract with their starting faction, a command ship that can fly across space with advanced capabilities, a small probe ship that can be used for reconnaissance, and 175,000 credits.  > \#\#\#\# Keep your token safe and secure > > Keep careful track of where you store your token. You can generate a new token from our account dashboard, but if someone else gains access to your token they will be able to use it to make API requests on your behalf until the end of the reset.  If you are new to SpaceTraders, It is recommended to register with the COSMIC faction, a faction that is well connected to the rest of the universe. After registering, you should try our interactive \[quickstart guide\](https://docs.spacetraders.io/quickstart/new-game) which will walk you through a few basic API requests in just a few minutes.

## Example Usage

```terraform
action "spacetraders_register" "example" {
  config {
    faction = "example"
    symbol = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `faction` (String, required)
* `symbol` (String, required)
