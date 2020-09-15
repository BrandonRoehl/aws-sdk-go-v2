// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Moves the specified phone number into the Deletion queue. A phone number must be
// disassociated from any users or Amazon Chime Voice Connectors before it can be
// deleted. Deleted phone numbers remain in the Deletion queue for 7 days before
// they are deleted permanently.
func (c *Client) DeletePhoneNumber(ctx context.Context, params *DeletePhoneNumberInput, optFns ...func(*Options)) (*DeletePhoneNumberOutput, error) {
	stack := middleware.NewStack("DeletePhoneNumber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeletePhoneNumberMiddlewares(stack)
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
	addOpDeletePhoneNumberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePhoneNumber(options.Region), middleware.Before)

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
			OperationName: "DeletePhoneNumber",
			Err:           err,
		}
	}
	out := result.(*DeletePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePhoneNumberInput struct {
	// The phone number ID.
	PhoneNumberId *string
}

type DeletePhoneNumberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeletePhoneNumberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeletePhoneNumber{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePhoneNumber{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePhoneNumber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeletePhoneNumber",
	}
}
