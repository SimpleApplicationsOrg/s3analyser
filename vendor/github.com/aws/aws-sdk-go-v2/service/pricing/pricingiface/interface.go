// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pricingiface provides an interface to enable mocking the AWS Price List Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pricingiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
)

// PricingAPI provides an interface to enable mocking the
// pricing.Pricing service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Price List Service.
//    func myFunc(svc pricingiface.PricingAPI) bool {
//        // Make svc.DescribeServices request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := pricing.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPricingClient struct {
//        pricingiface.PricingAPI
//    }
//    func (m *mockPricingClient) DescribeServices(input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPricingClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PricingAPI interface {
	DescribeServicesRequest(*pricing.DescribeServicesInput) pricing.DescribeServicesRequest

	DescribeServicesPages(*pricing.DescribeServicesInput, func(*pricing.DescribeServicesOutput, bool) bool) error
	DescribeServicesPagesWithContext(aws.Context, *pricing.DescribeServicesInput, func(*pricing.DescribeServicesOutput, bool) bool, ...aws.Option) error

	GetAttributeValuesRequest(*pricing.GetAttributeValuesInput) pricing.GetAttributeValuesRequest

	GetAttributeValuesPages(*pricing.GetAttributeValuesInput, func(*pricing.GetAttributeValuesOutput, bool) bool) error
	GetAttributeValuesPagesWithContext(aws.Context, *pricing.GetAttributeValuesInput, func(*pricing.GetAttributeValuesOutput, bool) bool, ...aws.Option) error

	GetProductsRequest(*pricing.GetProductsInput) pricing.GetProductsRequest

	GetProductsPages(*pricing.GetProductsInput, func(*pricing.GetProductsOutput, bool) bool) error
	GetProductsPagesWithContext(aws.Context, *pricing.GetProductsInput, func(*pricing.GetProductsOutput, bool) bool, ...aws.Option) error
}

var _ PricingAPI = (*pricing.Pricing)(nil)
