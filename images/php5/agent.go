package main

import (
	"regexp"
	"strings"

	"github.com/h2so5/sango/src"
)

func run(files []string, in sango.Input, out *sango.Output) (string, []string) {
	return "php", files
}

var r = regexp.MustCompile("\\(.+\\)")

func version() string {
	v, _ := sango.System(".", "", "php", "-v")
	v = strings.Split(v, "\n")[0]
	v = string(r.ReplaceAll([]byte(v), []byte("")))
	return v
}

func test() ([]string, string, string) {
	return []string{"test/hello.php"}, "", "Hello World"
}

func main() {
	sango.Run(sango.AgentOption{
		RunCmd: run,
		VerCmd: version,
		Test:   test,
	})
}