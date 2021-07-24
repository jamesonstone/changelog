package changelog

import "log"

func Build() {
	Render()
	log.Print("create changelog entry")
}
