# Specify how travis ci needs to fetch its dependencies in target
# We want to skip travis ci's default go get ./...,
# because ./... also checks the example directory in which we have main redeclared multiple times
target:
	go get -u -t .;
