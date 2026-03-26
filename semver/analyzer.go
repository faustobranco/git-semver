package semver

import (
	"fmt"
	"regexp"
	"strings"
)

type ReleaseType int

const (
	None ReleaseType = iota
	Patch
	Minor
	Major
)

var (
	majorPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)BREAKING CHANGE`),
		regexp.MustCompile(`!:`), // feat!: ...
	}

	minorPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)^feat(\(.+\))?:`),
		regexp.MustCompile(`(?i)^feature(\(.+\))?:`),
		regexp.MustCompile(`(?i)^\[feat\]`),
	}

	patchPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)^fix(\(.+\))?:`),
		regexp.MustCompile(`(?i)^bugfix(\(.+\))?:`),
		regexp.MustCompile(`(?i)^hotfix(\(.+\))?:`),
	}
)

func matchesAny(message string, patterns []*regexp.Regexp) bool {
	for _, p := range patterns {
		if p.MatchString(message) {
			return true
		}
	}
	return false
}

func prefixPattern(prefix string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`(?i)^%s(\(.+\))?:`, prefix))
}

func AnalyzeCommit(message string) ReleaseType {
	msg := strings.TrimSpace(message)

	// Major (breaking changes)
	if matchesAny(msg, majorPatterns) {
		return Major
	}

	// Minor (features)
	if matchesAny(msg, minorPatterns) {
		return Minor
	}

	// Patch (fixes)
	if matchesAny(msg, patchPatterns) {
		return Patch
	}

	return None
}
