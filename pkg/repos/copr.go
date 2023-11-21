package repos

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

type Copr struct {
	Author        string
	Reponame      string
	ReleaseServer string
}

const (
	HUB string = "copr.fedorainfracloud.org"
)

// Initialize new Copr if Argument matches the {Author}/{Reponame} pattern
func NewCopr(args []string) Copr {
	repoP := "[A-Za-z]+/[A-Za-z]+"

	if match, err := regexp.MatchString(repoP, args[0]); err != nil {
		log.Fatal(err)
	} else if !match {
		log.Fatalf(
			"Bad Copr Project Format Error: use format `copr_username/copr_projectname` to reference copr project",
		)
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
		"https://" + HUB + "/coprs/{{.Author}}/{{.Reponame}}/repo/fedora-{{.ReleaseServer}}/{{.Author}}-{{.Reponame}}-fedora-.repo",
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

// red the .repo file in /etc/yum.repos.d
func (c Copr) getRepoFilePath() string {
	tmpl := template.Must(template.New("path").Parse(
		"/etc/yum.repos.d/_copr:" + HUB + ":{{.Author}}:{{.Reponame}}.repo",
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

// either writes enabled=1 order enabled=0 into the repo file
func writeEnabled(configPath string, enable int) error {
	if !fileExists(configPath) {
		return errors.New("File does not exist!")
	}
	read, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		configPath,
		[]byte(
			strings.Replace(string(read),
				"enabled="+fmt.Sprint(1-enable), "enabled="+fmt.Sprint(enable), -1)), 0); err != nil {
		return err
	} else {
		return nil
	}
}

// if repo is installed, it enables ist otherwise it installes it
func (c Copr) Enable() {
	configPath := c.getRepoFilePath()

	if fileExists(configPath) {
		if err := writeEnabled(configPath, 1); err != nil {
			log.Fatal("Error Disabling COPR repo", err)
		} else {
			log.Println("Enabled COPR Repo " + c.Author + "/" + c.Reponame)
			os.Exit(0)
		}
	}

	config := c.getRepoConfig()

	err := os.WriteFile(configPath, []byte(config), 0644)
	if err != nil {
		log.Fatal("Error writing .repo file: ", err)
	}

	log.Println("Enabled COPR Repo " + c.Author + "/" + c.Reponame)
	os.Exit(0)
}

// if repo is enabled it disables it
func (c Copr) Disable() {
	configPath := c.getRepoFilePath()

	if !fileExists(configPath) {
		log.Println("No COPR Repo " + c.Author + "/" + c.Reponame + " is enabled")
		os.Exit(0)
	}

	if err := writeEnabled(configPath, 0); err != nil {
		log.Fatal("Error Disabling COPR repo", err)
	} else {
		log.Println("Disabled COPR Repo " + c.Author + "/" + c.Reponame)
		os.Exit(0)
	}
}

// remove the repo file
func (c Copr) Remove() {
	configPath := c.getRepoFilePath()

	if fileExists(configPath) {
		err := os.Remove(configPath)
		if err != nil {
			log.Fatal("Error removing .repo file:", err)
		}

		log.Println("Removed COPR Repo " + c.Author + "/" + c.Reponame)
		os.Exit(0)
	}

	log.Println("No COPR Repo " + c.Author + "/" + c.Reponame + " is enabled")
	os.Exit(0)
}

// list installed coprs
func ListCoprs(flags *pflag.FlagSet) {
	repoConfigs, err := os.ReadDir("/etc/yum.repos.d/")

	if err != nil {
		log.Fatal("Error listing coprs:", err)
	}

	for _, file := range repoConfigs {
		fname := file.Name()

		if en, _ := flags.GetBool("enabled"); en {
			c, err := os.ReadFile("/etc/yum.repos.d/" + fname)
			if err != nil {
				log.Fatal("Error listing coprs:", err)
			}
			if !strings.Contains(string(c), "enabled=1") {
				break
			}
		}

		if en, _ := flags.GetBool("disabled"); en {
			c, err := os.ReadFile("/etc/yum.repos.d/" + fname)
			if err != nil {
				log.Fatal("Error listing coprs:", err)
			}
			if !strings.Contains(string(c), "enabled=0") {
				break
			}
		}

		if strings.HasPrefix(fname, "_copr:") {
			fname = strings.TrimPrefix(file.Name(), "_copr:")
			fname = strings.TrimSuffix(fname, ".repo")
			fname = strings.ReplaceAll(fname, ":", "/")
			fmt.Println(fname)
		}
	}
}
