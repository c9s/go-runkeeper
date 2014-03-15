Nike+ API for Go
----------------

```go
import runkeeper "github.com/c9s/go-runkeeper"
client := runkeeper.NewClient("{accessToken}")  // pass access token if you have one. if you don't, just pass an empty string

log.Println("Logining...")
client.Login("email", "password")

log.Println("Asking new access token")
accessToken, err := client.AskAccessToken()
if err != nil {
    log.Println(err)
}

activities , err := client.GetActivities(nil)
for _, activity := range *activities.Data {
    activityDetails := client.GetActivityDetails(activity.Id)
    log.Println(activityDetails)
}

activities , err := client.GetActivities(runkeeper.Params{ "count": "10" })
for _, activity := range *activities.Data {
    activityDetails := client.GetActivityDetails(activity.Id)
    log.Println(activityDetails)
}
```

