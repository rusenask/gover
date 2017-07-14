dep:
	go get

release:dep
	@echo "++ Building gover binaries"
	gox -verbose -output="release/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-ldflags "$(LDFLAGS)" -osarch="linux/amd64 darwin/amd64 windows/amd64"