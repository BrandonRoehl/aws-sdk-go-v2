// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a gateway from AWS IoT SiteWise. When you delete a gateway, some of the
// gateway's files remain in your gateway's file system. For more information, see
// Data retention
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/data-retention.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) DeleteGateway(ctx context.Context, params *DeleteGatewayInput, optFns ...func(*Options)) (*DeleteGatewayOutput, error) {
	stack := middleware.NewStack("DeleteGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteGatewayMiddlewares(stack)
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
	addOpDeleteGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGateway(options.Region), middleware.Before)

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
			OperationName: "DeleteGateway",
			Err:           err,
		}
	}
	out := result.(*DeleteGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGatewayInput struct {
	// The ID of the gateway to delete.
	GatewayId *string
}

type DeleteGatewayOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DeleteGateway",
	}
}
