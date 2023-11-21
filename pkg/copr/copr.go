package copr

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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
	getRepoFilePath() string
	Enable()
	Disable()
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

	return Copr{args[0], args[1], string(out)[:2]}
}

// Get .repo config from https://copr.fedorainfracloud.org/
func (c Copr) getRepoConfig() string {
	tmpl := template.Must(template.New("url").Parse(
		"https://copr.fedorainfracloud.org/coprs/{{.Author}}/{{.Reponame}}/repo/fedora-{{.ReleaseServer}}/{{.Author}}-{{.Reponame}}-fedora-.repo",
	))

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, c); err != nil {
		log.Fatal("Error executing template: ", err)
	}

	url := buf.String()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal("Error requesting repo config: ", err)
	} else if res.StatusCode != 200 {
		log.Fatalf("Error requesting repo config: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading request body: ", err)
	}

	return string(body)
}

func (c Copr) getRepoFilePath() string {
	tmpl := template.Must(template.New("path").Parse(
		"/etc/yum.repos.d/{{.Author}}-{{.Reponame}}.repo",
	))

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, c); err != nil {
		log.Fatal("Error getting .repo file path: ", err)
	}

	return buf.String()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (c Copr) Enable() {
	configPath := c.getRepoFilePath()

	if fileExists(configPath) {
		log.Println("COPR Repo " + c.Author + "/" + c.Reponame + " is allready enabled")
		os.Exit(0)
	}

	config := c.getRepoConfig()

	err := os.WriteFile(configPath, []byte(config), 0644)
	if err != nil {
		log.Fatal("Error writing .repo file: ", err)
	}

	log.Println("Enabled COPR Repo " + c.Author + "/" + c.Reponame)
	os.Exit(0)
}

func (c Copr) Disable() {
	configPath := c.getRepoFilePath()

	if fileExists(configPath) {
		err := os.Remove(configPath)
		if err != nil {
			log.Fatal("Error removing .repo file:", err)
		}

		log.Println("Disabled COPR Repo " + c.Author + "/" + c.Reponame)
		os.Exit(0)
	}

	log.Println("No COPR Repo " + c.Author + "/" + c.Reponame + " is enabled")
	os.Exit(0)
}
