// +build integration

//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	_ "github.com/aws/aws-sdk-go-v2/internal/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@swf", func() {
		gucumber.World["client"] = swf.New(integration.Config())
	})
}
