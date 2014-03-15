RunKeeper API for Go
---------------------

```sh
go run -x bin/activity.go {access token}
```

```go
import runkeeper "github.com/c9s/go-runkeeper"
import "log"
import "flag"

func main() {
	flag.Parse()

	args := flag.Args()

	client := runkeeper.NewClient(args[0]) // pass access token
	activities, err := client.GetFitnessActivityFeed(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(activities)
	for _, activity := range activities.Items {
		log.Printf("%#v\n", activity)

		activityDetails, err := client.GetFitnessActivity(activity.Uri, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println(activityDetails)
	}
}
```

