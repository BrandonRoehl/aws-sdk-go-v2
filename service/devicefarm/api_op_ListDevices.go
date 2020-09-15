// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about unique device types.
func (c *Client) ListDevices(ctx context.Context, params *ListDevicesInput, optFns ...func(*Options)) (*ListDevicesOutput, error) {
	stack := middleware.NewStack("ListDevices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDevicesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDevices(options.Region), middleware.Before)

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
			OperationName: "ListDevices",
			Err:           err,
		}
	}
	out := result.(*ListDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the result of a list devices request.
type ListDevicesInput struct {
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The Amazon Resource Name (ARN) of the project.
	Arn *string
	// Used to select a set of devices. A filter is made up of an attribute, an
	// operator, and one or more values.
	//
	//     * Attribute: The aspect of a device such
	// as platform or model used as the selection criteria in a device filter. Allowed
	// values include:
	//
	//         * ARN: The Amazon Resource Name (ARN) of the device
	// (for example, arn:aws:devicefarm:us-west-2::device:12345Example).
	//
	//         *
	// PLATFORM: The device platform. Valid values are ANDROID or IOS.
	//
	//         *
	// OS_VERSION: The operating system version (for example, 10.3.2).
	//
	//         *
	// MODEL: The device model (for example, iPad 5th Gen).
	//
	//         * AVAILABILITY:
	// The current availability of the device. Valid values are AVAILABLE,
	// HIGHLY_AVAILABLE, BUSY, or TEMPORARY_NOT_AVAILABLE.
	//
	//         * FORM_FACTOR: The
	// device form factor. Valid values are PHONE or TABLET.
	//
	//         * MANUFACTURER:
	// The device manufacturer (for example, Apple).
	//
	//         * REMOTE_ACCESS_ENABLED:
	// Whether the device is enabled for remote access. Valid values are TRUE or
	// FALSE.
	//
	//         * REMOTE_DEBUG_ENABLED: Whether the device is enabled for remote
	// debugging. Valid values are TRUE or FALSE. Because remote debugging is no longer
	// supported
	// (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html),
	// this attribute is ignored.
	//
	//         * INSTANCE_ARN: The Amazon Resource Name
	// (ARN) of the device instance.
	//
	//         * INSTANCE_LABELS: The label of the
	// device instance.
	//
	//         * FLEET_TYPE: The fleet type. Valid values are PUBLIC
	// or PRIVATE.
	//
	//     * Operator: The filter operator.
	//
	//         * The EQUALS operator
	// is available for every attribute except INSTANCE_LABELS.
	//
	//         * The CONTAINS
	// operator is available for the INSTANCE_LABELS and MODEL attributes.
	//
	//         *
	// The IN and NOT_IN operators are available for the ARN, OS_VERSION, MODEL,
	// MANUFACTURER, and INSTANCE_ARN attributes.
	//
	//         * The LESS_THAN,
	// GREATER_THAN, LESS_THAN_OR_EQUALS, and GREATER_THAN_OR_EQUALS operators are also
	// available for the OS_VERSION attribute.
	//
	//     * Values: An array of one or more
	// filter values.
	//
	//         * The IN and NOT_IN operators take a values array that
	// has one or more elements.
	//
	//         * The other operators require an array with a
	// single element.
	//
	//         * In a request, the AVAILABILITY attribute takes the
	// following values: AVAILABLE, HIGHLY_AVAILABLE, BUSY, or TEMPORARY_NOT_AVAILABLE.
	Filters []*types.DeviceFilter
}

// Represents the result of a list devices operation.
type ListDevicesOutput struct {
	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string
	// Information about the devices.
	Devices []*types.Device

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDevicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDevices{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDevices{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDevices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListDevices",
	}
}
