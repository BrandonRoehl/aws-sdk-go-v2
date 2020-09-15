// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// To deregister a consumer, provide its ARN. Alternatively, you can provide the
// ARN of the data stream and the name you gave the consumer when you registered
// it. You may also provide all three parameters, as long as they don't conflict
// with each other. If you don't know the name or ARN of the consumer that you want
// to deregister, you can use the ListStreamConsumers () operation to get a list of
// the descriptions of all the consumers that are currently registered with a given
// data stream. The description of a consumer contains its name and ARN. This
// operation has a limit of five transactions per second per account.
func (c *Client) DeregisterStreamConsumer(ctx context.Context, params *DeregisterStreamConsumerInput, optFns ...func(*Options)) (*DeregisterStreamConsumerOutput, error) {
	stack := middleware.NewStack("DeregisterStreamConsumer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterStreamConsumerMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterStreamConsumer(options.Region), middleware.Before)

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
			OperationName: "DeregisterStreamConsumer",
			Err:           err,
		}
	}
	out := result.(*DeregisterStreamConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterStreamConsumerInput struct {
	// The ARN of the Kinesis data stream that the consumer is registered with. For
	// more information, see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	StreamARN *string
	// The ARN returned by Kinesis Data Streams when you registered the consumer. If
	// you don't know the ARN of the consumer that you want to deregister, you can use
	// the ListStreamConsumers operation to get a list of the descriptions of all the
	// consumers that are currently registered with a given data stream. The
	// description of a consumer contains its ARN.
	ConsumerARN *string
	// The name that you gave to the consumer.
	ConsumerName *string
}

type DeregisterStreamConsumerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterStreamConsumerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterStreamConsumer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterStreamConsumer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterStreamConsumer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DeregisterStreamConsumer",
	}
}
