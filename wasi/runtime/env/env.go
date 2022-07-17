package env

// Boolean retrieve a boolean flag from the environment, checks each key in order
// first to parse successfully is returned.
func Boolean(fallback bool, values ...string) bool
