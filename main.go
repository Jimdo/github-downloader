package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"io"

	"github.com/Luzifer/gh-private-dl/privatehub"
)

var (
	ghToken string
)

func init() {
	flag.StringVar(&ghToken, "github-token", os.Getenv("GITHUB_TOKEN"), "GitHub token to use for authentication")
}

func main() {
	flag.Parse()

	if flag.NArg() != 3 {
		abort("You need to pass three arguments: <repository> <version> <file>")
	}

	var (
		repo     = flag.Arg(0)
		version  = flag.Arg(1)
		filename = flag.Arg(2)
	)

	log.Printf("Downloading %s in version %s from %s", filename, version, repo)
	downloadURL, err := privatehub.GetDownloadURL(repo, version, filename, ghToken)
	if err != nil {
		abort("Error generating download URL: %s", err)
	}

	tmpFile, err := os.Create(fmt.Sprintf("%s.tmp", filename))
	if err != nil {
		abort("Error creating tmp file: %s", err)
	}
	defer tmpFile.Close()

	resp, err := http.Get(downloadURL)
	if err != nil {
		abort("Error downloading file from GitHub: %s", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		abort("Could not write to tmp file: %s", err)
	}

	if err := os.Chmod(tmpFile.Name(), 0775); err != nil {
		abort("Could not chmod tmp file: %s", err)
	}

	if err := os.Rename(tmpFile.Name(), filename); err != nil {
		abort("Could not move file to final location: %s", err)
	}

	log.Printf("Done, file was written to %s", filename)
}

func abort(format string, params ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf("%s\n", fmt.Sprintf(format, params...)))
	os.Exit(1)
}
