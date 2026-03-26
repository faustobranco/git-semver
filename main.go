package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"git-semver/semver"
	"log"
)

var version = "dev"

type Result struct {
	CurrentVersion string `json:"current_version"`
	NextVersion    string `json:"next_version"`
	ReleaseType    string `json:"release_type"`
	HasRelease     bool   `json:"has_release"`
	Pushed         bool   `json:"pushed"`
}

func releaseTypeToString(r semver.ReleaseType) string {
	switch r {
	case semver.Major:
		return "major"
	case semver.Minor:
		return "minor"
	case semver.Patch:
		return "patch"
	default:
		return "none"
	}
}

func main() {

	repoPath := flag.String("repo", ".", "Path to git repository")
	push := flag.Bool("push", false, "Push tags to remote")
	showVersion := flag.Bool("version", false, "Show version")
	jsonOutput := flag.Bool("json", false, "Output result in JSON")

	flag.Parse()

	if *showVersion {
		fmt.Printf("git-semver version %s\n", version)
		return
	}

	lastTag, hasTag, err := semver.GetLastTag(*repoPath)
	if err != nil {
		log.Fatal(err)
	}

	commits, err := semver.GetCommitsSince(*repoPath, lastTag, hasTag)
	if err != nil {
		log.Fatal(err)
	}
	release := semver.None

	for _, c := range commits {
		r := semver.AnalyzeCommit(c)
		if r > release {
			release = r
		}
	}

	if release == semver.None {
		result := Result{
			CurrentVersion: lastTag,
			ReleaseType:    "none",
			HasRelease:     false,
			Pushed:         false,
		}

		if *jsonOutput {
			data, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(data))
		} else {
			fmt.Println("No release needed")
		}

		return
	}

	current := semver.ParseVersion(lastTag)
	next := semver.Bump(current, release)

	result := Result{
		CurrentVersion: lastTag,
		NextVersion:    fmt.Sprintf("v%d.%d.%d", next.Major, next.Minor, next.Patch),
		ReleaseType:    releaseTypeToString(release),
		HasRelease:     true,
		Pushed:         false,
	}

	if *push {
		if err := semver.CreateTag(*repoPath, next); err != nil {
			log.Fatal(err)
		}

		if err := semver.PushTags(*repoPath); err != nil {
			log.Fatal(err)
		}

		result.Pushed = true
	}

	if *jsonOutput {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
		return
	}

	fmt.Printf("New version: %s\n", result.NextVersion)

	if !*push {
		fmt.Println("Dry-run mode (use --push to create tag)")
	}
}
