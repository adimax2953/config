package env

import "github.com/adimax2953/config/internal/env"

func Process(prefix string, target interface{}) error {
	return env.Process(prefix, target)
}
