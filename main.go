package main

import (
	"flag"
	"fmt"
	"git-semver/semver"
	"log"
)

func main() {

	repoPath := flag.String("repo", ".", "Path to git repository")
	push := flag.Bool("push", false, "Push tags to remote")

	flag.Parse()

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
		fmt.Println("No release needed")
		return
	}

	var current semver.Version

	if hasTag {
		current = semver.ParseVersion(lastTag)
	} else {
		current = semver.Version{0, 0, 0}
	}
	next := semver.Bump(current, release)

	fmt.Printf("New version: v%d.%d.%d\n", next.Major, next.Minor, next.Patch)

	if err := semver.CreateTag(*repoPath, next); err != nil {
		log.Fatal(err)
	}

	if *push {
		err := semver.PushTags(*repoPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
