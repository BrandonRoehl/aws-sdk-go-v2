// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the instance profiles that have the specified path prefix. If there are
// none, the operation returns an empty list. For more information about instance
// profiles, go to About Instance Profiles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
// You can paginate the results using the MaxItems and Marker parameters.
func (c *Client) ListInstanceProfiles(ctx context.Context, params *ListInstanceProfilesInput, optFns ...func(*Options)) (*ListInstanceProfilesOutput, error) {
	stack := middleware.NewStack("ListInstanceProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListInstanceProfilesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceProfiles(options.Region), middleware.Before)

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
			OperationName: "ListInstanceProfiles",
			Err:           err,
		}
	}
	out := result.(*ListInstanceProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceProfilesInput struct {
	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string
	// The path prefix for filtering the results. For example, the prefix
	// /application_abc/component_xyz/ gets all instance profiles whose path starts
	// with /application_abc/component_xyz/. This parameter is optional. If it is not
	// included, it defaults to a slash (/), listing all instance profiles. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of either a forward slash (/) by itself or a
	// string that must begin and end with forward slashes. In addition, it can contain
	// any ASCII character from the ! (\u0021) through the DEL character (\u007F),
	// including most punctuation characters, digits, and upper and lowercased letters.
	PathPrefix *string
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32
}

// Contains the response to a successful ListInstanceProfiles () request.
type ListInstanceProfilesOutput struct {
	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string
	// A list of instance profiles.
	InstanceProfiles []*types.InstanceProfile
	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListInstanceProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListInstanceProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListInstanceProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opListInstanceProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListInstanceProfiles",
	}
}
