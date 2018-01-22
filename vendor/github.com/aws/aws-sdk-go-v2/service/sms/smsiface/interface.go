// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package smsiface provides an interface to enable mocking the AWS Server Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package smsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sms"
)

// SMSAPI provides an interface to enable mocking the
// sms.SMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Server Migration Service.
//    func myFunc(svc smsiface.SMSAPI) bool {
//        // Make svc.CreateReplicationJob request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sms.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSMSClient struct {
//        smsiface.SMSAPI
//    }
//    func (m *mockSMSClient) CreateReplicationJob(input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSMSClient{}
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
type SMSAPI interface {
	CreateReplicationJobRequest(*sms.CreateReplicationJobInput) sms.CreateReplicationJobRequest

	DeleteReplicationJobRequest(*sms.DeleteReplicationJobInput) sms.DeleteReplicationJobRequest

	DeleteServerCatalogRequest(*sms.DeleteServerCatalogInput) sms.DeleteServerCatalogRequest

	DisassociateConnectorRequest(*sms.DisassociateConnectorInput) sms.DisassociateConnectorRequest

	GetConnectorsRequest(*sms.GetConnectorsInput) sms.GetConnectorsRequest

	GetConnectorsPages(*sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool) error
	GetConnectorsPagesWithContext(aws.Context, *sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool, ...aws.Option) error

	GetReplicationJobsRequest(*sms.GetReplicationJobsInput) sms.GetReplicationJobsRequest

	GetReplicationJobsPages(*sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool) error
	GetReplicationJobsPagesWithContext(aws.Context, *sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool, ...aws.Option) error

	GetReplicationRunsRequest(*sms.GetReplicationRunsInput) sms.GetReplicationRunsRequest

	GetReplicationRunsPages(*sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool) error
	GetReplicationRunsPagesWithContext(aws.Context, *sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool, ...aws.Option) error

	GetServersRequest(*sms.GetServersInput) sms.GetServersRequest

	GetServersPages(*sms.GetServersInput, func(*sms.GetServersOutput, bool) bool) error
	GetServersPagesWithContext(aws.Context, *sms.GetServersInput, func(*sms.GetServersOutput, bool) bool, ...aws.Option) error

	ImportServerCatalogRequest(*sms.ImportServerCatalogInput) sms.ImportServerCatalogRequest

	StartOnDemandReplicationRunRequest(*sms.StartOnDemandReplicationRunInput) sms.StartOnDemandReplicationRunRequest

	UpdateReplicationJobRequest(*sms.UpdateReplicationJobInput) sms.UpdateReplicationJobRequest
}

var _ SMSAPI = (*sms.SMS)(nil)
