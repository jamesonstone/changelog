package changelog

import "log"

func Create() {
	Render()
	log.Print("create changelog entry")
}
