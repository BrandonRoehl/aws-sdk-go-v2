// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details for an events configuration that allows a bot to receive outgoing
// events, such as an HTTPS endpoint or Lambda function ARN.
func (c *Client) GetEventsConfiguration(ctx context.Context, params *GetEventsConfigurationInput, optFns ...func(*Options)) (*GetEventsConfigurationOutput, error) {
	stack := middleware.NewStack("GetEventsConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetEventsConfigurationMiddlewares(stack)
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
	addOpGetEventsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventsConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetEventsConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetEventsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventsConfigurationInput struct {
	// The Amazon Chime account ID.
	AccountId *string
	// The bot ID.
	BotId *string
}

type GetEventsConfigurationOutput struct {
	// The events configuration details.
	EventsConfiguration *types.EventsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetEventsConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetEventsConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventsConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEventsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetEventsConfiguration",
	}
}
