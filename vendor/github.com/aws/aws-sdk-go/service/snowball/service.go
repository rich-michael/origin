// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package snowball

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Snowball is a petabyte-scale data transport solution that uses secure
// appliances to transfer large amounts of data between your on-premises data
// centers and Amazon Simple Storage Service (Amazon S3). The Snowball commands
// described here provide access to the same functionality that is available
// in the AWS Snowball Management Console, which enables you to create and manage
// jobs for Snowball. To transfer data locally with a Snowball appliance, you'll
// need to use the Snowball client or the Amazon S3 API adapter for Snowball.
// For more information, see the User Guide (http://docs.aws.amazon.com/AWSImportExport/latest/ug/api-reference.html).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30
type Snowball struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "snowball"  // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Snowball client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Snowball client from just a session.
//     svc := snowball.New(mySession)
//
//     // Create a Snowball client with additional configuration
//     svc := snowball.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Snowball {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Snowball {
	svc := &Snowball{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-06-30",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSIESnowballJobManagementService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Snowball operation and runs any
// custom request initialization.
func (c *Snowball) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}