# runner

This is a very basic service that generates GET HTTP requests to the 
`test-service` and check the response for validity.

## Usage

runner url

## Summary

The runner loops forever generating GET requests to the url + randomInt.

It checks the response, logs the output code and looks for the random
number in the response, logging an error if not found in the response.

