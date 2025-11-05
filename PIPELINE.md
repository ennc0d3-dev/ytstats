# CI/CD Pipeline Documentation

This document describes the stable CI/CD pipeline for the yt-stats project.

## Overview

The project uses **GitHub Actions** for continuous integration and deployment with multiple automated workflows providing comprehensive testing, security scanning, and deployment automation.

## Workflows

### 1. Main CI/CD Workflow (`.github/workflows/ci.yml`)

Comprehensive pipeline with 5 parallel jobs:

#### Lint Job
- **Tool**: golangci-lint v6 (latest)
- **Timeout**: 10 minutes
- **Caching**: Go modules cached
- **Purpose**: Code quality and style enforcement

#### Test Job
- **Test Types**: Unit + Integration tests
- **Race Detector**: Enabled (-race flag)
- **Coverage**: Atomic coverage mode
- **Reporting**: Uploads to Codecov
- **Verification**: Dependency verification (go mod verify)

#### Build Job
- **Depends On**: Lint + Test jobs must pass
- **Output**: Production binary artifact
- **Retention**: 7 days
- **Caching**: Full Go module caching

#### Security Job
- **Scanner**: Gosec (latest)
- **Format**: SARIF output
- **Integration**: GitHub Security tab
- **Coverage**: All Go files (./...)

#### Docker Job
- **Trigger**: Only on push events
- **Depends On**: Lint + Test jobs must pass
- **BuildKit**: Enabled with GHA caching
- **Optimization**: Multi-layer cache strategy

### 2. Legacy Build Workflow (`.github/workflows/go.yml`)

Production deployment workflow:

- **Linting**: golangci-lint v6
- **Building**: Multi-platform support
- **Testing**: Full test suite with coverage
- **Coverage**: Codecov integration
- **Docker**: Build and push to Docker Hub
- **Versioning**: Semantic versioning with run number

### 3. Security Workflow (`.github/workflows/codeql.yml`)

Continuous security scanning:

- **Tool**: CodeQL by GitHub
- **Languages**: Go + JavaScript (Chrome extension)
- **Schedule**: Weekly (Sundays at 10:36 UTC)
- **Triggers**: Push and PRs to master/main
- **Analysis**: Full codebase scan

## Pipeline Features

### ✅ Automated Testing
- Unit tests run on every commit
- Integration tests validate API functionality
- Race condition detection
- 42.3% code coverage on API package

### ✅ Security Scanning
- Gosec for Go code vulnerabilities
- CodeQL for advanced security analysis
- Dependabot for dependency updates
- SARIF integration with GitHub Security

### ✅ Code Quality
- golangci-lint with comprehensive rules
- 10-minute timeout for thorough analysis
- Fails fast on quality issues

### ✅ Docker Integration
- Multi-stage builds
- BuildKit optimization
- Layer caching (GitHub Actions cache)
- Automated tagging

### ✅ Coverage Reporting
- Codecov integration
- Atomic coverage mode
- Historical tracking
- PR comments with coverage diff

### ✅ Artifact Management
- Binary artifacts retained for 7 days
- Docker images pushed to registry
- Semantic versioning

## Pipeline Triggers

### On Push
- Branches: `master`, `main`
- Runs: All CI/CD workflows
- Docker build and push (if configured)

### On Pull Request
- Branches: `master`, `main`
- Runs: Lint, Test, Build, Security
- Docker build only (no push)

### On Schedule
- Weekly CodeQL security scan
- Sundays at 10:36 UTC

## Performance Optimizations

1. **Parallel Execution**
   - Lint and Test jobs run in parallel
   - Build waits for both to complete
   - Optimal use of CI minutes

2. **Caching Strategy**
   - Go modules cached per job
   - Docker layers cached with GHA
   - Speeds up subsequent runs

3. **Job Dependencies**
   - Build only runs if Lint + Test pass
   - Docker only runs if Lint + Test pass
   - Saves CI minutes on failures

## Status Badges

Add these to README.md:

```markdown
![CI/CD](https://github.com/ennc0d3-dev/ytstats/workflows/CI/CD/badge.svg)
![CodeQL](https://github.com/ennc0d3-dev/ytstats/workflows/CodeQL/badge.svg)
![Coverage](https://codecov.io/gh/ennc0d3-dev/ytstats/branch/master/graph/badge.svg)
```

## Local Pipeline Testing

Test the pipeline locally before pushing:

```bash
# Run linting
golangci-lint run --timeout=10m

# Run tests with race detector
go test -v -race -coverprofile=coverage.out ./...

# Generate coverage report
go tool cover -func=coverage.out

# Build binary
go build -v -o yt-stats ./cmd/yt-stats

# Security scan
gosec -no-fail -fmt=sarif -out=results.sarif ./...

# Docker build
docker build -t yt-stats:test .
```

## Monitoring

### GitHub Actions
- View workflow runs: `https://github.com/ennc0d3-dev/ytstats/actions`
- Check security alerts: `https://github.com/ennc0d3-dev/ytstats/security`

### Codecov
- Coverage dashboard: `https://codecov.io/gh/ennc0d3-dev/ytstats`

## Future Enhancements

- [ ] Add deployment to production environment
- [ ] Implement blue-green deployments
- [ ] Add performance benchmarking
- [ ] Integration testing with real YouTube API
- [ ] Multi-platform Docker builds (ARM, AMD64)
- [ ] Release automation with GitHub Releases
- [ ] Notification integration (Slack, Discord)

## Conclusion

The yt-stats project has a **stable, production-ready CI/CD pipeline** with:

- ✅ Automated testing
- ✅ Security scanning  
- ✅ Code quality enforcement
- ✅ Docker integration
- ✅ Coverage tracking
- ✅ Parallel execution
- ✅ Smart caching

**Status**: Pipeline is stable and operational ✅

**Last Updated**: 2025-11-05
