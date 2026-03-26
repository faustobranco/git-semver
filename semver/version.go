package semver

import (
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func ParseVersion(tag string) Version {
	tag = strings.TrimPrefix(tag, "v")
	parts := strings.Split(tag, ".")

	if len(parts) != 3 {
		return Version{0, 0, 0}
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	return Version{major, minor, patch}
}

func Bump(v Version, r ReleaseType) Version {
	switch r {
	case Major:
		return Version{v.Major + 1, 0, 0}
	case Minor:
		return Version{v.Major, v.Minor + 1, 0}
	case Patch:
		return Version{v.Major, v.Minor, v.Patch + 1}
	default:
		return v
	}
}
