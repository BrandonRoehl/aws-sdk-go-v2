// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified directory as a publisher to the specified SNS topic.
func (c *Client) DeregisterEventTopic(ctx context.Context, params *DeregisterEventTopicInput, optFns ...func(*Options)) (*DeregisterEventTopicOutput, error) {
	stack := middleware.NewStack("DeregisterEventTopic", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterEventTopicMiddlewares(stack)
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
	addOpDeregisterEventTopicValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterEventTopic(options.Region), middleware.Before)

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
			OperationName: "DeregisterEventTopic",
			Err:           err,
		}
	}
	out := result.(*DeregisterEventTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Removes the specified directory as a publisher to the specified SNS topic.
type DeregisterEventTopicInput struct {
	// The name of the SNS topic from which to remove the directory as a publisher.
	TopicName *string
	// The Directory ID to remove as a publisher. This directory will no longer send
	// messages to the specified SNS topic.
	DirectoryId *string
}

// The result of a DeregisterEventTopic request.
type DeregisterEventTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterEventTopicMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterEventTopic{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterEventTopic{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterEventTopic(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeregisterEventTopic",
	}
}
