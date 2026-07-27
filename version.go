package main

// Injected at build time by goreleaser via -ldflags -X
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)
