package main

import (
	"flag"
	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/analyser"
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"github.com/SimpleApplicationsOrg/s3analyser/service"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"os"
)

func main() {

	profile := flag.String("profile", "", "Get credentials for profile in ~/.aws/credentials")
	size := flag.String("size", "KB", "KB, MB, GB, TB")
	withStorage := flag.Bool("withStorage", false, "Organize by Storage Class")
	byRegion := flag.Bool("byRegion", false, "Group by Region")

	var filter model.FilterMap
	flag.Var(&filter, "filter", "List of bucket names to filter")

	flag.Parse()

	if *profile != "" {
		os.Setenv(external.AWSProfileEnvVar, *profile)
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config, %v\n", err)
		os.Exit(1)
	}

	s3 := service.S3Factory(cfg)

	s3Analyser := analyser.Factory(*byRegion, *withStorage, filter, *size)

	result, err := s3Analyser.Analyse(s3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to analyse s3, %v\n", err)
		os.Exit(1)
	}

	s3Analyser.Print(result)

}
