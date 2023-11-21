package copr

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type Copr struct {
	Author        string
	Reponame      string
	ReleaseServer string
}

type Repo interface {
	getRepoConfig() string
	Enable()
}

// Initialize new Copr if Argument matches the {Author}/{Reponame} pattern
func NewCopr(args []string) Copr {
	repoP := "[A-Za-z]+/[A-Za-z]+"

	if match, err := regexp.MatchString(repoP, args[0]); err != nil {
		log.Fatal(err)
	} else if !match {
		log.Fatalf("Repository Identifier must match {Author}/{Reponame} pattern.")
	}
	args = strings.Split(args[0], "/")

	// get fedora releaseserver
	out, err := exec.Command("rpm", "-E", "%fedora").Output()

	if err != nil {
		log.Fatal("Error executing rpm command: ", err)
	}

	return Copr{args[0], args[1], string(out)}
}
