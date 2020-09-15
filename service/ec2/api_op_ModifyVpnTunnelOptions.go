// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the options for a VPN tunnel in an AWS Site-to-Site VPN connection. You
// can modify multiple options for a tunnel in a single request, but you can only
// modify one tunnel at a time. For more information, see Site-to-Site VPN Tunnel
// Options for Your Site-to-Site VPN Connection
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html) in the AWS
// Site-to-Site VPN User Guide.
func (c *Client) ModifyVpnTunnelOptions(ctx context.Context, params *ModifyVpnTunnelOptionsInput, optFns ...func(*Options)) (*ModifyVpnTunnelOptionsOutput, error) {
	stack := middleware.NewStack("ModifyVpnTunnelOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyVpnTunnelOptionsMiddlewares(stack)
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
	addOpModifyVpnTunnelOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnTunnelOptions(options.Region), middleware.Before)

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
			OperationName: "ModifyVpnTunnelOptions",
			Err:           err,
		}
	}
	out := result.(*ModifyVpnTunnelOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnTunnelOptionsInput struct {
	// The ID of the AWS Site-to-Site VPN connection.
	VpnConnectionId *string
	// The tunnel options to modify.
	TunnelOptions *types.ModifyVpnTunnelOptionsSpecification
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The external IP address of the VPN tunnel.
	VpnTunnelOutsideIpAddress *string
}

type ModifyVpnTunnelOptionsOutput struct {
	// Describes a VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyVpnTunnelOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnTunnelOptions{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnTunnelOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyVpnTunnelOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpnTunnelOptions",
	}
}
