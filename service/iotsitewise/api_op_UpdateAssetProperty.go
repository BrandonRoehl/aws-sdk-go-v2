// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an asset property's alias and notification state. This operation
// overwrites the property's existing alias and notification state. To keep your
// existing property's alias or notification state, you must include the existing
// values in the UpdateAssetProperty request. For more information, see
// DescribeAssetProperty
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetProperty.html).
func (c *Client) UpdateAssetProperty(ctx context.Context, params *UpdateAssetPropertyInput, optFns ...func(*Options)) (*UpdateAssetPropertyOutput, error) {
	stack := middleware.NewStack("UpdateAssetProperty", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateAssetPropertyMiddlewares(stack)
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
	addIdempotencyToken_opUpdateAssetPropertyMiddleware(stack, options)
	addOpUpdateAssetPropertyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssetProperty(options.Region), middleware.Before)

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
			OperationName: "UpdateAssetProperty",
			Err:           err,
		}
	}
	out := result.(*UpdateAssetPropertyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssetPropertyInput struct {
	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping Industrial Data Streams to Asset Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide. If you omit this parameter, the alias is
	// removed from the property.
	PropertyAlias *string
	// The ID of the asset property to be updated.
	PropertyId *string
	// The ID of the asset to be updated.
	AssetId *string
	// The MQTT notification state (enabled or disabled) for this asset property. When
	// the notification state is enabled, AWS IoT SiteWise publishes property value
	// updates to a unique MQTT topic. For more information, see Interacting with Other
	// Services
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/interact-with-other-services.html)
	// in the AWS IoT SiteWise User Guide. If you omit this parameter, the notification
	// state is set to DISABLED.
	PropertyNotificationState types.PropertyNotificationState
}

type UpdateAssetPropertyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateAssetPropertyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssetProperty{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssetProperty{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateAssetProperty struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAssetProperty) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAssetProperty) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAssetPropertyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAssetPropertyInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateAssetPropertyMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAssetProperty{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAssetProperty(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "UpdateAssetProperty",
	}
}
