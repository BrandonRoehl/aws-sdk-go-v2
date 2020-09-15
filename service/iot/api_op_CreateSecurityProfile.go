// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Device Defender security profile.
func (c *Client) CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) {
	stack := middleware.NewStack("CreateSecurityProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateSecurityProfileMiddlewares(stack)
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
	addOpCreateSecurityProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityProfile(options.Region), middleware.Before)

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
			OperationName: "CreateSecurityProfile",
			Err:           err,
		}
	}
	out := result.(*CreateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityProfileInput struct {
	// Specifies the destinations to which alerts are sent. (Alerts are always sent to
	// the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets map[string]*types.AlertTarget
	// Metadata that can be used to manage the security profile.
	Tags []*types.Tag
	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here. Note: This API field is deprecated. Please use
	// CreateSecurityProfileRequest$additionalMetricsToRetainV2 () instead.
	AdditionalMetricsToRetain []*string
	// The name you are giving to the security profile.
	SecurityProfileName *string
	// A description of the security profile.
	SecurityProfileDescription *string
	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetainV2 []*types.MetricToRetain
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []*types.Behavior
}

type CreateSecurityProfileOutput struct {
	// The ARN of the security profile.
	SecurityProfileArn *string
	// The name you gave to the security profile.
	SecurityProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateSecurityProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateSecurityProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSecurityProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSecurityProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateSecurityProfile",
	}
}
