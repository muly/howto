
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

validate:
	swagger validate ./swagger.yml


# The `gen` target depends on the `validate` target as
# it will only succesfully generate the code if the specification
# is valid.
# 
# Here we're specifying some flags:
# --target              the base directory for generating the files;
# --spec                path to the swagger specification;
# --exclude-main        generates only the library code and not a 
#                       sample CLI application;
# --name                the name of the application.
gen: validate
	swagger generate server \
		--target=./gen \
		--spec=./swagger.yml \
		--exclude-main \
		--name=hello



.PHONY: check_install validate gen