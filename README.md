# Webhook git cloner

This code has a single purpose; receive a push notification via webhook, match a ref, then clone the repo on a match.

## Platform Support

* Github
* Gitlab
* Gitea

## Development

* Start ngrok `./ngrok http 4567`
* Setup a webhook pointing to your ngrok url ensuring to add the github URI.
* Start the server `go run *.go <refs>`. Refs will be used to determine when to clone.

## Use cases

1. Mirroring a repo. You may wish to backup a repo when a change is made. For example, if master changes we would run the program with `refs/heads/master` and this service would clone master on every change.
2. Deployment. You may need an updated version to be deployed on a machine on every change to a release branch. In this case you would pass `refs/heads/release` and the service would clone the release branch on every change.

## Convention

When adding a webhook, the app will only respond to a url which contains the providers name as the uri. Example, `http://your-endpoint/github`

## Limitations

Currently, the app only clones a single repo per provider.