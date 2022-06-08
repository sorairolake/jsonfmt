#
# SPDX-License-Identifier: Apache-2.0 OR MIT
#
# Copyright (C) 2022 Shun Sakai
#

alias all := default

# Run default recipe
default: build

# Build a package
@build:
    mkdir -p ./build
    go build -o ./build/jsonfmt ./cmd/jsonfmt

# Remove generated artifacts
@clean:
    go clean
    rm -rf ./build

# Run tests
@test:
    go test ./...

# Run `golangci-lint run`
@golangci-lint:
    golangci-lint run -E gofmt,goimports

# Run the formatter (`go fmt` and `goimports`)
fmt: gofmt goimports

# Run `go fmt`
@gofmt:
    go fmt ./...

# Run `goimports`
@goimports:
    find . -type f -name "*.go" -execdir goimports -w {} +

# Run the linter (`go vet` and `staticcheck`)
lint: vet staticcheck

# Run `go vet`
@vet:
    go vet ./...

# Run `staticcheck`
@staticcheck:
    staticcheck ./...

# Build a man page
@build-man-page:
    go run ./tools/build.go

# Run the linter for GitHub Actions workflow files
@lint-github-actions:
    actionlint

# Run the formatter for the README
@fmt-readme:
    npx prettier -w README.md
