.PHONY: binary
binary:
	goreleaser build --single-target --snapshot --rm-dist
