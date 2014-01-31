# API documentation

Short documentation on available endpoints.

Request to the FutuLeague API must include the `Accept` header with including the value `application/vnd.futuleague.v1+json`, which corresponds to API version 1. Note though that as of now there is only only version of the API and it is still likely to change. Endpoints can also be appended with `.json` in order to get a response from the latest API version (and is easier to test in the browser).

## Endpoints

### GET requests - returns application/json

* `/` - Returns a list of games available.
* `/{game}` - Returns basic information about the club (currently only name) and a list of clubs in that game.
* `/players` - Returns a list of all players
* `/players/{player}` - Returns the information of a particular player, where {player} is the player's identifier.

### POST requests - returns application/json

* `/players` - Creates a new player. Takes two arguments `identifier` and `name`. Returns a redirect to /players/{identifier} with status 201 Created. Is identifier already exists returns 303 See other and does not change anything.

