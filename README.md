# s3analyser

[![GoDoc](https://godoc.org/github.com/SimpleApplicationsOrg/s3analyser?status.svg)](https://godoc.org/github.com/SimpleApplicationsOrg/s3analyser)
[![Go Report Card](https://goreportcard.com/badge/github.com/SimpleApplicationsOrg/s3analyser)](https://goreportcard.com/report/github.com/SimpleApplicationsOrg/s3analyser)
[![travis-ci.org](https://travis-ci.org/SimpleApplicationsOrg/s3analyser.svg?branch=master)](http://travis-ci.org/SimpleApplicationsOrg/s3analyser?branch=master)
[![codecov.io](http://codecov.io/github/SimpleApplicationsOrg/s3analyser/coverage.svg?branch=master)](http://codecov.io/github/SimpleApplicationsOrg/s3analyser?branch=master)
```
The s3analyser is a tool that shows information about an AWS account s3 buckets.
The credentials can be set in the same way as aws-cli. More information on:
https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html
https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html

Usage:  
  -byRegion  
        Group by Region  
  -filter value  
        List of bucket names to filter  
  -profile string  
        Get credentials for profile in ~/.aws/credentials  
  -size string  
        KB, MB, GB, TB (default "KB")  
  -withStorage  
        Organize by Storage Class  
```
### Get the binaries
The latest binaries can be downloaded from https://github.com/SimpleApplicationsOrg/s3analyser/releases

### Installing with the docker image

### Using s3analyser
[![asciicast](https://asciinema.org/a/UolZxtmF7KT4hv0h8wIF8xzIs.png)](https://asciinema.org/a/UolZxtmF7KT4hv0h8wIF8xzIs)
