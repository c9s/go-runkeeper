package main

import runkeeper "github.com/c9s/go-runkeeper"
import "github.com/davecgh/go-spew/spew"
import "log"
import "flag"

func main() {
	flag.Parse()

	args := flag.Args()

	client := runkeeper.NewClient(args[0])
	activities, err := client.GetFitnessActivityFeed(nil)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(activities)

	for _, activity := range activities.Items {
		log.Printf("%#v\n", activity)

		activityDetails, err := client.GetFitnessActivity(activity.Uri, nil)
		if err != nil {
			log.Println(err)
		}
		spew.Dump(activityDetails)
	}
}
