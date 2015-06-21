docker:
	GCO_ENABLED=0 go build -a -installsuffix runner.go
	docker build -t systemzoo/runner .
