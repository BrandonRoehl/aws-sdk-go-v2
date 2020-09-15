// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provisions a thing in the device registry. RegisterThing calls other AWS IoT
// control plane APIs. These calls might exceed your account level  AWS IoT
// Throttling Limits
// (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_iot)
// and cause throttle errors. Please contact AWS Customer Support
// (https://console.aws.amazon.com/support/home) to raise your throttling limits if
// necessary.
func (c *Client) RegisterThing(ctx context.Context, params *RegisterThingInput, optFns ...func(*Options)) (*RegisterThingOutput, error) {
	stack := middleware.NewStack("RegisterThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegisterThingMiddlewares(stack)
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
	addOpRegisterThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterThing(options.Region), middleware.Before)

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
			OperationName: "RegisterThing",
			Err:           err,
		}
	}
	out := result.(*RegisterThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterThingInput struct {
	// The parameters for provisioning a thing. See Provisioning Templates
	// (https://docs.aws.amazon.com/iot/latest/developerguide/provision-template.html)
	// for more information.
	Parameters map[string]*string
	// The provisioning template. See Provisioning Devices That Have Device
	// Certificates
	// (https://docs.aws.amazon.com/iot/latest/developerguide/provision-w-cert.html)
	// for more information.
	TemplateBody *string
}

type RegisterThingOutput struct {
	// ARNs for the generated resources.
	ResourceArns map[string]*string
	// The certificate data, in PEM format.
	CertificatePem *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegisterThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegisterThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RegisterThing",
	}
}
