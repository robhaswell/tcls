# tcls
Thera Content Location Service

## What is it?

TCLS is a service for the videogame EVE Online.
It is designed to search for interesting connections to Thera and notify a Slack channel about those which appear.

## How does it work?

TCLS scans for Thera connections in Eve-Scout and your corporation's bookmarks then compares them for proximity to content hotspots, including:

* Trade hubs
* POS/sov timers
* High NPC kills

Any connections which are within a threshold number of jumps will be announced on a configurable Slack channel.

## Contributing

TCLS is a microservices application, mostly written in Go and deployed using Kubernetes.
For more information see CONTRIBUTING.md
