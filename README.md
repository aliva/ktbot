# ktbot

```sh
# Create twitter a app and get keys
# https://developer.twitter.com/en/apps
export KTBOT_CONSUMER_KEY=[twitter consumer key]
export KTBOT_CONSUMER_SECRET=[twitter consumer secret]

# generate tokens - run only once
go run login/main.go

export KTBOT_TOKEN=[twitter access token]
export KTBOT_TOKEN_SECRET=[twitter access secret]

go run .
```
