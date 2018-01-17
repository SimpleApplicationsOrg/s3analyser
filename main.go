package main

import (
	"flag"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/SimpleApplicationsOrg/s3analyser/service"
	"github.com/SimpleApplicationsOrg/s3analyser/analyser"
)

func init() {
	profile := flag.String("profile", "", "Get credentials for profile in ~/.aws/credentials")

	flag.Parse()

	if profile != nil {
		os.Setenv(external.AWSProfileEnvVar, *profile)
	}
}

func main() {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load config, " + err.Error())
	}

	s3 := service.S3Factory(cfg)

	result := analyser.Analyse(s3)

	analyser.Print(result)

}