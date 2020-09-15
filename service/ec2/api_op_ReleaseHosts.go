// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// When you no longer want to use an On-Demand Dedicated Host it can be released.
// On-Demand billing is stopped and the host goes into released state. The host ID
// of Dedicated Hosts that have been released can no longer be specified in another
// request, for example, to modify the host. You must stop or terminate all
// instances on a host before it can be released. When Dedicated Hosts are
// released, it may take some time for them to stop counting toward your limit and
// you may receive capacity errors when trying to allocate new Dedicated Hosts.
// Wait a few minutes and then try again. Released hosts still appear in a
// DescribeHosts () response.
func (c *Client) ReleaseHosts(ctx context.Context, params *ReleaseHostsInput, optFns ...func(*Options)) (*ReleaseHostsOutput, error) {
	stack := middleware.NewStack("ReleaseHosts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReleaseHostsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpReleaseHostsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReleaseHosts(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ReleaseHosts",
			Err:           err,
		}
	}
	out := result.(*ReleaseHostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleaseHostsInput struct {
	// The IDs of the Dedicated Hosts to release.
	HostIds []*string
}

type ReleaseHostsOutput struct {
	// The IDs of the Dedicated Hosts that were successfully released.
	Successful []*string
	// The IDs of the Dedicated Hosts that could not be released, including an error
	// message.
	Unsuccessful []*types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReleaseHostsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReleaseHosts{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReleaseHosts{}, middleware.After)
}

func newServiceMetadataMiddleware_opReleaseHosts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReleaseHosts",
	}
}
