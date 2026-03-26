
# git-semver

A lightweight Go tool to automatically determine the next semantic version of a Git repository and create a tag based on commit messages.

----------

## What it does

This tool:

1.  Reads commits since the last tag
    
2.  Determines the next version (major / minor / patch)
    
3.  Creates a Git tag
    
4.  Optionally pushes the tag to the remote
    

It is inspired by tools like semantic-release, but intentionally simpler and more flexible for real-world teams.

----------

## Versioning Logic

Version bump is determined from commit messages using a flexible interpretation of Conventional Commits.

### Supported formats

The tool supports multiple commit styles:

#### Standard (Conventional Commits)

```
feat: add login
fix(api): resolve error
feat!: breaking change

```

#### Custom formats (common in teams)

```
[FEAT] - add new feature
[FIX] - bug fix
[BREAKING] - change API

```

----------

## Version Rules

Commit Type

Version Bump

`feat`, `[FEAT]`

Minor

`fix`, `[FIX]`

Patch

`BREAKING`, `!`

Major

others

Ignored

----------

## First Release

If the repository has **no tags**, the tool will:

-   Analyze all commits
    
-   Start from `v0.0.0`
    
-   Apply normal bump rules
    

Example:

```
[FEAT] initial commit → v0.1.0

```

----------

## Installation

```bash
go build -ldflags "-X main.version=v0.2.0" -o git-semver

```

Or run directly:

```bash
go run .

```

----------

## Usage

### Run in current repository

```bash
git-semver

```

----------

### Run against another repository

```bash
git-semver --repo ../other-project

```

----------

### Create and push tag

```bash
git-semver --repo ../other-project --push

```

----------

## Flags

Flag

Description

Default

`--repo`

Path to Git repository

`.`

`--push`

Push tags to remote

`false`

`--json`

json format output

`{
  "current_version": "v0.2.1",
  "next_version": "",
  "release_type": "none",
  "has_release": false,
  "pushed": false
}`

`--version`

git-semver Version

`v0.0.0`
----------

## Requirements

-   Git installed and available in PATH
    
-   Repository must:
    
    -   be cloned locally
        
    -   contain commit history
        
    -   have correct remote configured (for push)
        

----------

## Important Notes

### Shallow clones

This tool requires full history and tags.

If using CI (Jenkins, GitHub Actions, etc.), ensure:

```
fetch-depth: 0

```

----------

### No relevant commits

If no commits match version rules:

```
No release needed

```

----------

## 💡 Design Philosophy

-   Simple and predictable
    
-   Works with real-world commit habits
    
-   No strict enforcement of commit standards
    
-   Easy to extend and customize
    

----------

## 📄 License

MIT
