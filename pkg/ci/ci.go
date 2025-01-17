package ci

import "github.com/debricked/cli/pkg/ci/env"

type ICi interface {
	Identify() bool
	Map() (env.Env, error)
}
