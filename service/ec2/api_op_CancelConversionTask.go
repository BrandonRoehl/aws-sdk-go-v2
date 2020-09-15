// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an active conversion task. The task can be the import of an instance or
// volume. The action removes all artifacts of the conversion, including a
// partially uploaded volume or instance. If the conversion is complete or is in
// the process of transferring the final disk image, the command fails and returns
// an exception. For more information, see Importing a Virtual Machine Using the
// Amazon EC2 CLI
// (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/ec2-cli-vmimport-export.html).
func (c *Client) CancelConversionTask(ctx context.Context, params *CancelConversionTaskInput, optFns ...func(*Options)) (*CancelConversionTaskOutput, error) {
	stack := middleware.NewStack("CancelConversionTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCancelConversionTaskMiddlewares(stack)
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
	addOpCancelConversionTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelConversionTask(options.Region), middleware.Before)

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
			OperationName: "CancelConversionTask",
			Err:           err,
		}
	}
	out := result.(*CancelConversionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelConversionTaskInput struct {
	// The reason for canceling the conversion task.
	ReasonMessage *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the conversion task.
	ConversionTaskId *string
}

type CancelConversionTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCancelConversionTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCancelConversionTask{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCancelConversionTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelConversionTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelConversionTask",
	}
}
