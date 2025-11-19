# Speedhive / Sporthive 

This package offers a Go implementation of the public API for [Speedhive](https://speedhive.mylaps.com) / [Sporthive](https://sporthive.com).

## Supported endpoints

From the list published in the [client settings](https://sporthive.com/api/clientSettings) we can find that there are a couple API's available. For some there is an openapi spec available, those are the ones that this library targets.

| Section                                                                          | Implemented | Note                            |
| -------------------------------------------------------------------------------- | ----------- | ------------------------------- |
| [Event results](https://eventresults-api.speedhive.com/swagger/docs/v1)          | Yes         |                                 |
| [Practice](https://practice-api.speedhive.com/swagger/docs/v1)                   | Yes         |                                 |
| [Users and products](https://usersandproducts-api.speedhive.com/swagger/docs/v1) | No          | Requires OAuth2                 |
| Live timing                                                                      | No          | Haven't found the specification |

## Usage

```go
client := eventresult.NewClient(nil) // Pass nil or an http.Client
event, err := client.Events.List()
```

```go
client := practice.NewClient(nil) // Pass nil or an http.Client
practice, err := client.Locations.List()
```