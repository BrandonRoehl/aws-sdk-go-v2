// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Request a list of field-level encryption profiles that have been created in
// CloudFront for this account.
func (c *Client) ListFieldLevelEncryptionProfiles(ctx context.Context, params *ListFieldLevelEncryptionProfilesInput, optFns ...func(*Options)) (*ListFieldLevelEncryptionProfilesOutput, error) {
	stack := middleware.NewStack("ListFieldLevelEncryptionProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListFieldLevelEncryptionProfilesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFieldLevelEncryptionProfiles(options.Region), middleware.Before)

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
			OperationName: "ListFieldLevelEncryptionProfiles",
			Err:           err,
		}
	}
	out := result.(*ListFieldLevelEncryptionProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFieldLevelEncryptionProfilesInput struct {
	// Use this when paginating results to indicate where to begin in your list of
	// profiles. The results include profiles in the list that occur after the marker.
	// To get the next page of results, set the Marker to the value of the NextMarker
	// from the current page's response (which is also the ID of the last profile on
	// that page).
	Marker *string
	// The maximum number of field-level encryption profiles you want in the response
	// body.
	MaxItems *string
}

type ListFieldLevelEncryptionProfilesOutput struct {
	// Returns a list of the field-level encryption profiles that have been created in
	// CloudFront for this account.
	FieldLevelEncryptionProfileList *types.FieldLevelEncryptionProfileList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListFieldLevelEncryptionProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListFieldLevelEncryptionProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListFieldLevelEncryptionProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFieldLevelEncryptionProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListFieldLevelEncryptionProfiles",
	}
}
