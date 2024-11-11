# Joule Reader

> Joule (J) is the unit of energy in the SI that is equal to the amount of work done when a force of one newton displaces a mass through a distance of one meter in the direction of that force.

It happens to be the `J` unit, which is also a java testing framework and what we're interested in.

Joule reader allows one to easily read JUnit XML reports

## Usage

```bash
# If you want to read the formatted version of the XML files
joule-reader build/junit-test-files/*

# If you want to only print the files that have a failure or an error, add the `-x` flag.
# If no tests are red, then it returns 0, otherwise it will have the number of failed/errored test suites
# in the exit code
joule-reader -x build/junit-test-files/*
```
