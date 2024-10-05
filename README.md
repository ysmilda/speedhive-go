# Speedhive / Sporthive 

This package offers a Go implementation of the public API for [Speedhive](https://speedhive.mylaps.com) / [Sporthive](https://sporthive.com).

## Supported endpoints

According to the [client settings](https://sporthive.com/api/clientSettings) the endpoints below exist. This table describes the status of the implementation.

| Section                                                                 | Implemented | Note                            |
| ----------------------------------------------------------------------- | ----------- | ------------------------------- |
| [Event results](https://eventresults-api.speedhive.com/swagger/docs/v1) | Yes         |                                 |
| [Practice](https://practice-api.speedhive.com/swagger/docs/v1)          | Yes         |                                 |
| Users and products                                                      | No          | Requires OAuth2                 |
| Live timing                                                             | No          | Haven't found the specification |

## Usage

```go
client := eventresult.NewClient(nil) // Pass nil or an http.Client
event, err := client.Events.List()
```

```go
client := practice.NewClient(nil) // Pass nil or an http.Client
practice, err := client.Locations.List()
```