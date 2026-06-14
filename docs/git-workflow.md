# Marsadyn Git Workflow

## Branching Strategy

Marsadyn follows a phase-based branching strategy with atomic commits.

## Branch Naming

```
phase-<number>-<short-name>
```

Examples:
- `phase-0-repo-inspection`
- `phase-1-project-foundation`
- `phase-2-core-data-model`
- `phase-4-kafka-pipeline`

## Workflow

### Starting a Phase

```bash
# Ensure main is up to date
git checkout main
git pull origin main

# Create phase branch
git checkout -b phase-XX-short-name
```

### During a Phase

1. Make atomic commits for each small task
2. Validate code after each commit
3. Do NOT push yet

```bash
# Example atomic commits
git add <file>
git commit -m "phase 1: initialize go backend module"

git add <file>
git commit -m "phase 1: add backend command structure"
```

### Ending a Phase

```bash
# Push phase branch
git push -u origin phase-XX-short-name

# Create PR via GitHub CLI
gh pr create --base main --head phase-XX-short-name --title "Phase XX: Description"

# Merge PR
gh pr merge --merge

# Delete branch
git checkout main
git pull origin main
git branch -d phase-XX-short-name
git push origin --delete phase-XX-short-name
```

## Commit Message Format

```
phase <number>: <small completed task>
```

### Examples

```
phase 0: document repository structure
phase 0: add architecture overview
phase 1: initialize go backend module
phase 1: add docker compose infrastructure
phase 2: add tenant schema
phase 4: add kafka producer abstraction
phase 6: add metrics query endpoint
phase 9: add frontend layout
```

### Rules

- Start with `phase <number>:`
- Use lowercase after colon
- Be specific about what was added
- One logical change per commit

## Git Identity

All commits must be authored as:

```
YASSERRMD <arafath.yasser@gmail.com>
```

Configure locally:

```bash
git config user.name "YASSERRMD"
git config user.email "arafath.yasser@gmail.com"
```

## Atomic Commit Guidelines

### Good Atomic Commits

```
phase 1: initialize go backend module
phase 1: add api command entry point
phase 1: add health check endpoint
phase 2: add tenant migration
phase 2: add application migration
phase 4: add kafka producer interface
phase 4: add metric ingestion endpoint
```

### Bad Commits (Too Large)

```
phase 1: add entire backend
phase 4: add kafka pipeline and ingestion
```

## Validation Commands

After each commit, verify:

```bash
# Backend
go fmt ./...
go vet ./...
go test ./...

# Frontend
npm run lint
npm run build

# Docker
docker compose -f deploy/docker-compose.yml config
```

## PR Description Template

```markdown
## Phase XX: Description

### Changes
- Task 1
- Task 2
- Task 3

### Validation
- [ ] go fmt ./...
- [ ] go vet ./...
- [ ] go test ./...
- [ ] npm run lint
- [ ] npm run build

### Testing
- Describe how changes were tested
```

## Branch Cleanup

After merging:

```bash
# Local cleanup
git checkout main
git pull origin main
git branch -d phase-XX-short-name

# Remote cleanup
git push origin --delete phase-XX-short-name
```

## Continuing to Next Phase

```bash
# Ensure main is current
git checkout main
git pull origin main

# Start next phase
git checkout -b phase-YY-next-name
```

## Conflict Resolution

If conflicts arise during merge:

```bash
# During PR merge
git checkout main
git pull origin main
git merge phase-XX-short-name

# Resolve conflicts
git add .
git commit -m "phase XX: resolve merge conflicts"

# Continue with merge
```

## Protected Branches

- `main` requires PR approval
- No direct commits to main
- All changes go through phase branches

## Tags

Tag releases after major phases:

```bash
git tag -a v0.1.0 -m "Phase 1 complete: Project foundation"
git push origin v0.1.0
```
