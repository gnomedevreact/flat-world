package flat

type Character struct {
	Name               string
	Personality        string
	Gender             string
	Mood               int
	Hunger             int
	Patience           int
	Energy             int
	SocialNeed         int
	Relationships      map[string]int
	Interests          []string
	CurrentAction      string
	TicksSinceLastMeal int
	LastAction         string
}

type Characters []Character
