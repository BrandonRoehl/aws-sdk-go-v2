// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an infrastructure configuration.
func (c *Client) GetInfrastructureConfiguration(ctx context.Context, params *GetInfrastructureConfigurationInput, optFns ...func(*Options)) (*GetInfrastructureConfigurationOutput, error) {
	stack := middleware.NewStack("GetInfrastructureConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetInfrastructureConfigurationMiddlewares(stack)
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
	addOpGetInfrastructureConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInfrastructureConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetInfrastructureConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetInfrastructureConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetInfrastructureConfiguration request object.
type GetInfrastructureConfigurationInput struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration that you want
	// to retrieve.
	InfrastructureConfigurationArn *string
}

// GetInfrastructureConfiguration response object.
type GetInfrastructureConfigurationOutput struct {
	// The infrastructure configuration object.
	InfrastructureConfiguration *types.InfrastructureConfiguration
	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetInfrastructureConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetInfrastructureConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInfrastructureConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInfrastructureConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "GetInfrastructureConfiguration",
	}
}
