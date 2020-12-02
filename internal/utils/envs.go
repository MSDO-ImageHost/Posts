package utils

import (
	"fmt"
	"os"
)

func CheckEnvs(envs []string) (err error) {

	for _, env := range envs {
		if os.Getenv(env) == "" {
			return fmt.Errorf("%s is not configured!", env)
		}
	}
	return nil
}
