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

// Gets information about the entries for a specified managed prefix list.
func (c *Client) GetManagedPrefixListEntries(ctx context.Context, params *GetManagedPrefixListEntriesInput, optFns ...func(*Options)) (*GetManagedPrefixListEntriesOutput, error) {
	stack := middleware.NewStack("GetManagedPrefixListEntries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetManagedPrefixListEntriesMiddlewares(stack)
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
	addOpGetManagedPrefixListEntriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedPrefixListEntries(options.Region), middleware.Before)

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
			OperationName: "GetManagedPrefixListEntries",
			Err:           err,
		}
	}
	out := result.(*GetManagedPrefixListEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedPrefixListEntriesInput struct {
	// The version of the prefix list for which to return the entries. The default is
	// the current version.
	TargetVersion *int64
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// The token for the next page of results.
	NextToken *string
	// The ID of the prefix list.
	PrefixListId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type GetManagedPrefixListEntriesOutput struct {
	// Information about the prefix list entries.
	Entries []*types.PrefixListEntry
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetManagedPrefixListEntriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetManagedPrefixListEntries{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetManagedPrefixListEntries{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetManagedPrefixListEntries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetManagedPrefixListEntries",
	}
}
