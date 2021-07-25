package changelog

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Tag struct {
	Previous string
	Current  string
}

type ChangelogData struct {
	Name       string
	ShortHash  string
	CommitMsg  string
	ChangeType string
	Scope      string
	Tag        Tag
}

func Render() {
	ch := ChangelogData{
		Name:       "App for Testing",
		ShortHash:  "f79ee7f",
		CommitMsg:  "This is a test commit",
		ChangeType: "feat",
		Scope:      "DE-123",
		Tag: Tag{
			Current:  "v1.0.0",
			Previous: "v0.5.9",
		},
	}
	c, e := ioutil.ReadFile("./pkg/changelog/CHANGELOG.tmpl.md")
	if e != nil {
		log.Fatal(e)
	}

	t, e := template.New("CHANGELOG").Parse(string(c))
	if e != nil {
		log.Fatal(e)
	}

	e = t.Execute(os.Stdout, ch)
	if e != nil {
		log.Fatal(e)
	}
}
