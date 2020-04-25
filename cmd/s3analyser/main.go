package main

import (
	"flag"
	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/analyser"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/service"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"os"
	"time"
)

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "The s3analyser is a tool that shows information about an AWS account s3 buckets.\n"+
			"The credentials can be set in the same way as aws-cli. More information on:\n"+
			"https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html\n"+
			"https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html\n\n"+
			"Usage:\n")
		flag.PrintDefaults()
	}
}

func main() {
	initial := time.Now()

	profile := flag.String("profile", "", "Get credentials for profile in ~/.aws/credentials")
	size := flag.String("size", "KB", "KB, MB, GB, TB")
	withStorage := flag.Bool("withStorage", false, "Organize by Storage Class")
	byRegion := flag.Bool("byRegion", false, "Group by Region")

	var filter model.FilterMap
	flag.Var(&filter, "filter", "List of bucket names to filter")
	flag.Parse()

	if *profile != "" {
		os.Setenv("AWS_PROFILE", *profile)
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config, %v\n", err)
		os.Exit(1)
	}

	s := service.New(cfg)
	objects, err := s.Objects(filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get objects from s3 %v", err)
		os.Exit(1)
	}

	s3Analyser := analyser.New(*byRegion, *withStorage, *size)

	result, err := s3Analyser.Analyse(objects)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to analyse s3. "+
			"Please make sure your credentials are correct set and you have access to AWS.\n"+
			"%v\n", err)
		os.Exit(1)
	}

	s3Analyser.Print(os.Stdout, result)

	fmt.Println("time:", time.Now().Sub(initial))
}
