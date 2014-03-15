Nike+ API for Go
----------------

```go
import "github.com/c9s/go-nikeplus"
client := nikeplus.NewClient("{accessToken}")  // pass access token if you have one. if you don't, just pass an empty string

log.Println("Logining...")
client.Login("email", "password")

log.Println("Asking new access token")
accessToken, err := client.AskAccessToken()
if err != nil {
    log.Println(err)
}

activities , err := client.GetActivities()
for _, activity := range *activities.Data {
    activityDetails := client.GetActivityDetails(activity.Id)
    log.Println(activityDetails)
}
```

