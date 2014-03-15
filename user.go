package runkeeper

/*
	{
		"userID": 1234567890,
		"profile": "/profile",
		"settings": "/settings",
		"fitness_activities": "/fitnessActivities",
		"strength_training_activities": "/strengthTrainingActivities",
		"background_activities": "/backgroundActivities",
		"sleep": "/sleep",
		"nutrition": "/nutrition",
		"weight": "/weight",
		"general_measurements": "/generalMeasurements",
		"diabetes": "/diabetes",
		"records": "/records",
		"team": "/team"
	}
*/

type User struct {
	Id                         int64  `json:"userID"`
	Profile                    string `json:"profile"`
	Settings                   string `json:"settings"`
	FitnessActivities          string `json:"fitness_activities"`
	StrengthTrainingActivities string `json:"strength_training_activities"`
	BackgroundActivities       string `json:"background_activities"`
	Sleep                      string `json:"sleep"`
	Nutrition                  string `json:"nutrition"`
	Weight                     string `json:"weight"`
	GeneralMeasurements        string `json:"general_measurements"`
	Diabetes                   string `json:"diabetes"`
	Records                    string `json:"records"`
	Team                       string `json:"team"`
}
