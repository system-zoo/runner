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

### - container response payload

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

#### Docker machine metrics

    {
       "0b0c75dc3a5e9849954be2d18a32481e95874af325cc3b16cd389e5d8a0a7e9f": {
          "newtork_in": "172.17.0.12",
          "containername": "hopeful_stallman",
          "ip": "172.17.0.12",
          "cpu_user": 9.42,
          "cpu_system": 1.39,
          "status": "172.17.0.12",
          "memory": "164777984",
          "cpu_all": 10.81,
          "newtork_out": "172.17.0.12"
       },
       "58a94dafc6c505df9b1eb6b6184f46b2c8f63b71202e301bc029103f8d253c34": {
          "newtork_in": "172.17.0.18",
          "containername": "sharp_elion",
          "ip": "172.17.0.18",
          "cpu_user": 3.23,
          "cpu_system": 0.54,
          "status": "172.17.0.18",
          "memory": "154255360",
          "cpu_all": 3.77,
          "newtork_out": "172.17.0.18"
       },
       "0d35e265adbfc2db737f8e142398a0e91634f5387d42c73833b9cfbdab4367bf": {
          "newtork_in": "172.17.0.14",
          "containername": "tender_fermat",
          "ip": "172.17.0.14",
          "cpu_user": 0.36,
          "cpu_system": 0.28,
          "status": "172.17.0.14",
          "memory": "856064",
          "cpu_all": 0.64,
          "newtork_out": "172.17.0.14"
       },
       "50f9a1e1b035f8c21684ea418d89b48e99b88a523ad7b23c9cd762131758c290": {
          "newtork_in": "172.17.0.3",
          "containername": "mynginx1",
          "ip": "172.17.0.3",
          "cpu_user": 0.01,
          "cpu_system": 0.01,
          "status": "172.17.0.3",
          "memory": "1388544",
          "cpu_all": 0.02,
          "newtork_out": "172.17.0.3"
       }
    }
