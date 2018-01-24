Bugtry
===

An experiment in receiving Bugsnag events from applications and forwarding them to Sentry.

## Goals

* Accept Bugsnag request
* Send Bugsnag request to Sentry
* Map Bugsnag requests to an appropriate Sentry endpoint based on Bugsnag credentials used
* Generate Sentry projects programatically
  * Via hardcoded YML
  * On the fly based on incoming Bugsnag events
    * Unlikely, but cool if this worked and incoming events had enough info to determine the name and flavour of Sentry projects to create
* (maybe) something to do with Sentry alerting on bugs
  * Show in a miminal web UI?
  * Send to the user via Slack (would require some configuration by the user)
  * Native OSX notifications would be cool
