package semver

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Functional (affect version)
// feat → new feature → minor
// fix → correction → patch
// BREAKING CHANGE or ! → major

// Non-functional (ignored for version)
// docs → documentation
// chore → internal tasks
// refactor → code without changing behavior
// test → tests
// style → formatting

func GetLastTag(repoPath string) (string, bool, error) {
	cmd := exec.Command("git", "-C", repoPath, "describe", "--tags", "--abbrev=0")

	out, err := cmd.Output()
	if err != nil {
		return "", false, nil // no tags
	}

	return strings.TrimSpace(string(out)), true, nil
}

func GetCommitsSince(repoPath, tag string, hasTag bool) ([]string, error) {
	var cmd *exec.Cmd

	if !hasTag {
		cmd = exec.Command("git", "-C", repoPath, "log", "--pretty=format:%s")
	} else {
		rangeRef := fmt.Sprintf("%s..HEAD", tag)
		cmd = exec.Command("git", "-C", repoPath, "log", rangeRef, "--pretty=format:%s")
	}

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	var commits []string
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		if strings.HasPrefix(l, "Merge") {
			continue
		}
		commits = append(commits, l)
	}

	return commits, nil
}

func CreateTag(repoPath string, v Version) error {
	tag := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)

	cmd := exec.Command("git", "-C", repoPath, "tag", tag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func PushTags(repoPath string) error {
	cmd := exec.Command("git", "-C", repoPath, "push", "origin", "--tags")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
