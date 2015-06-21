# runner

This is a very basic service written in Go that generates GET HTTP requests to
the `test-service` project  and check the response for validity.

## Usage

You need to compile the runner.go source file:  

    go build runner.go

Run the resulting binary file with one argument a URL to the service to test:  

    runner http://ipaddress

This binary is available in a docker file.  If you have docker installed you
can docker run this container:

    docker run systemzoo/runner http://ipaddress

## Development

To include the binary in the docker file the go code must be compiled with
all the dynamically linked libraries included in the resulting binary.  This
is not default settings for the `go build` command.  The build command looks
like this:

    GCO_ENABLED=0 go build -a -installsuffix runner.go

Use the included `Makefile` with the `docker` target to compile the go code,
and build the docker container locally.

    make docker


## Summary

The runner loops forever generating GET requests to the url + randomInt.

It checks the response, logs the output code and looks for the random
number in the response, logging an error if not found in the response.


## Payload for elastic search

        {
           "containerid": "0b0c75dc3a5e",
           "requestdate": "2015-06-21",
           "requesttime": "01:56:59",
           "containername": "hopeful_stallman",
           "response": {
              "request": 126,
              "duration": 2500,
              "response": 126,
              "code": 500
           }
        }
        
