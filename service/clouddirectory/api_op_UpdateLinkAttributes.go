// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a given typed link’s attributes. Attributes to be updated must not
// contribute to the typed link’s identity, as defined by its
// IdentityAttributeOrder.
func (c *Client) UpdateLinkAttributes(ctx context.Context, params *UpdateLinkAttributesInput, optFns ...func(*Options)) (*UpdateLinkAttributesOutput, error) {
	stack := middleware.NewStack("UpdateLinkAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateLinkAttributesMiddlewares(stack)
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
	addOpUpdateLinkAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLinkAttributes(options.Region), middleware.Before)

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
			OperationName: "UpdateLinkAttributes",
			Err:           err,
		}
	}
	out := result.(*UpdateLinkAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLinkAttributesInput struct {
	// The attributes update structure.
	AttributeUpdates []*types.LinkAttributeUpdate
	// Allows a typed link specifier to be accepted as input.
	TypedLinkSpecifier *types.TypedLinkSpecifier
	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// updated typed link resides. For more information, see arns () or Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	DirectoryArn *string
}

type UpdateLinkAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateLinkAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLinkAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLinkAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateLinkAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpdateLinkAttributes",
	}
}
