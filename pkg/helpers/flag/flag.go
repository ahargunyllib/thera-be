package flag

import (
	"flag"
)

type Flag struct {
	SeederEntity string
}

var FlagVars = getFlags()

func getFlags() *Flag {
	seederEntity := flag.String(
		"entity",
		"",
		`Specify seeder entity to run
			Must be used with flag seeder
			Available entities can be looked up at cmd/seed/main.go file.
			Example: -entity=hospital
		`,
	)

	flag.Parse()

	flag := &Flag{
		SeederEntity: *seederEntity,
	}

	return flag
}
