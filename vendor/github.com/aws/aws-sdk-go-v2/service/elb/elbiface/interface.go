// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elbiface provides an interface to enable mocking the Elastic Load Balancing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elbiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elb"
)

// ELBAPI provides an interface to enable mocking the
// elb.ELB service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Elastic Load Balancing.
//    func myFunc(svc elbiface.ELBAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := elb.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockELBClient struct {
//        elbiface.ELBAPI
//    }
//    func (m *mockELBClient) AddTags(input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockELBClient{}
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
type ELBAPI interface {
	AddTagsRequest(*elb.AddTagsInput) elb.AddTagsRequest

	ApplySecurityGroupsToLoadBalancerRequest(*elb.ApplySecurityGroupsToLoadBalancerInput) elb.ApplySecurityGroupsToLoadBalancerRequest

	AttachLoadBalancerToSubnetsRequest(*elb.AttachLoadBalancerToSubnetsInput) elb.AttachLoadBalancerToSubnetsRequest

	ConfigureHealthCheckRequest(*elb.ConfigureHealthCheckInput) elb.ConfigureHealthCheckRequest

	CreateAppCookieStickinessPolicyRequest(*elb.CreateAppCookieStickinessPolicyInput) elb.CreateAppCookieStickinessPolicyRequest

	CreateLBCookieStickinessPolicyRequest(*elb.CreateLBCookieStickinessPolicyInput) elb.CreateLBCookieStickinessPolicyRequest

	CreateLoadBalancerRequest(*elb.CreateLoadBalancerInput) elb.CreateLoadBalancerRequest

	CreateLoadBalancerListenersRequest(*elb.CreateLoadBalancerListenersInput) elb.CreateLoadBalancerListenersRequest

	CreateLoadBalancerPolicyRequest(*elb.CreateLoadBalancerPolicyInput) elb.CreateLoadBalancerPolicyRequest

	DeleteLoadBalancerRequest(*elb.DeleteLoadBalancerInput) elb.DeleteLoadBalancerRequest

	DeleteLoadBalancerListenersRequest(*elb.DeleteLoadBalancerListenersInput) elb.DeleteLoadBalancerListenersRequest

	DeleteLoadBalancerPolicyRequest(*elb.DeleteLoadBalancerPolicyInput) elb.DeleteLoadBalancerPolicyRequest

	DeregisterInstancesFromLoadBalancerRequest(*elb.DeregisterInstancesFromLoadBalancerInput) elb.DeregisterInstancesFromLoadBalancerRequest

	DescribeAccountLimitsRequest(*elb.DescribeAccountLimitsInput) elb.DescribeAccountLimitsRequest

	DescribeInstanceHealthRequest(*elb.DescribeInstanceHealthInput) elb.DescribeInstanceHealthRequest

	DescribeLoadBalancerAttributesRequest(*elb.DescribeLoadBalancerAttributesInput) elb.DescribeLoadBalancerAttributesRequest

	DescribeLoadBalancerPoliciesRequest(*elb.DescribeLoadBalancerPoliciesInput) elb.DescribeLoadBalancerPoliciesRequest

	DescribeLoadBalancerPolicyTypesRequest(*elb.DescribeLoadBalancerPolicyTypesInput) elb.DescribeLoadBalancerPolicyTypesRequest

	DescribeLoadBalancersRequest(*elb.DescribeLoadBalancersInput) elb.DescribeLoadBalancersRequest

	DescribeLoadBalancersPages(*elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool) error
	DescribeLoadBalancersPagesWithContext(aws.Context, *elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool, ...aws.Option) error

	DescribeTagsRequest(*elb.DescribeTagsInput) elb.DescribeTagsRequest

	DetachLoadBalancerFromSubnetsRequest(*elb.DetachLoadBalancerFromSubnetsInput) elb.DetachLoadBalancerFromSubnetsRequest

	DisableAvailabilityZonesForLoadBalancerRequest(*elb.DisableAvailabilityZonesForLoadBalancerInput) elb.DisableAvailabilityZonesForLoadBalancerRequest

	EnableAvailabilityZonesForLoadBalancerRequest(*elb.EnableAvailabilityZonesForLoadBalancerInput) elb.EnableAvailabilityZonesForLoadBalancerRequest

	ModifyLoadBalancerAttributesRequest(*elb.ModifyLoadBalancerAttributesInput) elb.ModifyLoadBalancerAttributesRequest

	RegisterInstancesWithLoadBalancerRequest(*elb.RegisterInstancesWithLoadBalancerInput) elb.RegisterInstancesWithLoadBalancerRequest

	RemoveTagsRequest(*elb.RemoveTagsInput) elb.RemoveTagsRequest

	SetLoadBalancerListenerSSLCertificateRequest(*elb.SetLoadBalancerListenerSSLCertificateInput) elb.SetLoadBalancerListenerSSLCertificateRequest

	SetLoadBalancerPoliciesForBackendServerRequest(*elb.SetLoadBalancerPoliciesForBackendServerInput) elb.SetLoadBalancerPoliciesForBackendServerRequest

	SetLoadBalancerPoliciesOfListenerRequest(*elb.SetLoadBalancerPoliciesOfListenerInput) elb.SetLoadBalancerPoliciesOfListenerRequest

	WaitUntilAnyInstanceInService(*elb.DescribeInstanceHealthInput) error
	WaitUntilAnyInstanceInServiceWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...aws.WaiterOption) error

	WaitUntilInstanceDeregistered(*elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceDeregisteredWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...aws.WaiterOption) error

	WaitUntilInstanceInService(*elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceInServiceWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...aws.WaiterOption) error
}

var _ ELBAPI = (*elb.ELB)(nil)
