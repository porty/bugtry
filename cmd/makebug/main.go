package main

import (
	"errors"

	bugsnag "github.com/bugsnag/bugsnag-go"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		// Your Bugsnag project API key
		APIKey: "c1a224732246b15ca855cbe5388314a3",
		// The development stage of your application build, like "alpha" or
		// "production"
		ReleaseStage: "development",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/porty/bugtry/cmd/makebug/lib"},
		// more configuration options
	})

	bugsnag.Notify(errors.New("some error"))
}
