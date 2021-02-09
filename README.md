<p align="center">
  <img src="https://github.com/nikunicke/ftapi/blob/master/resources/ftapi.png" width="500"/>
</p>

# :construction: WORK IN PROGRESS :construction:

```terminal
go get github.com/nikunicke/ftapi
```

The ftapi Go package lets you view and manage all available data from the [42 API](https://api.intra.42.fr/apidoc). The implementation of the package is based on a type of [Discovery Document](https://developers.google.com/discovery/v1/reference/apis) found in the 42 API documentation. You can view it [here](https://raw.githubusercontent.com/nikunicke/ftapi/master/ftapi.json). Unfortunately the document is missing crucial data, which results in some caveats on the package.

The 42 API uses [Oauth 2.0](https://oauth.net/2/) to handle authentication and authorization. To authenticate you need to have a registered application on the 42 Intra and a *credentials.json* file in the root of your project. The credentails file should include all necessary data for the Oauth2 authentication flow. If you are using [client credentials](https://oauth.net/2/grant-types/client-credentials/) grant type, you do not need redirect uris nor the auth uri.

```yaml
{
    "client_id": "YOU_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET",
    "auth_uri": "https://api.intra.42.fr/oauth/authorize",
    "token_uri": "https://api.intra.42.fr/oauth/token",
    "scopes": ["THE", "SCOPES", "YOU", "NEED"],
    "redirect_uris": ["urn:ietf:wg:oauth:2.0:oob","http://localhost"]
}
```
Please note that the authorization flow is *currently* only designed for command-line applications.

## Caveats
*   Some types do not have any members and cannot really be used for anything. :hankey:
*   Some struct members are of unknown type and therefore been defined as interface{}, requiring the user to use [type assertions](https://tour.golang.org/methods/15).
*   Broadcasts, commands, exams_users, squads and squads_users are currently unavailable

## Usage
To initiate a client service you need to specify the grant type. Currently supporting [Client Credentials](https://oauth.net/2/grant-types/client-credentials/) and [Authorization Code](https://oauth.net/2/grant-types/authorization-code/) grant types. Authorization Code will cache your token so that you do not need to login all the time.
```Go
// ftapi.AuthorizationCode or ftapi.ClientCredentials

s, err := ftapi.NewService(ftapi.AuthorizationCode)
if err != nil {
    log.Fatalf("%v", err)
}
```
Get your user data (auth code grant type required):
```Go
me, err := s.Me()
if err != nil {
    log.Fatalf("%v", err)
}
fmt.Println("Hello,", me.FirstName, me.LastName)
```
Initiate an events service:
```Go
events := ftapi.Events(s)
oneEvent, _ := events.Get("3647").Do()
eventsList, _ := events.List().P("campus_id", []string{"13"}).Do() // P() is optional
for _, event := range eventsList.Events {
    fmt.Println(event.Name)
}
```
## Coming up:
POST/PUT/DELETE features

