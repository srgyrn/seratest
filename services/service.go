package service

func IsDuck(animal string) string {
	if animal == "duck" {
		return "that is indeed a duck"
	}

	return "nope! not a duck"
}
