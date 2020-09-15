// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends you notification through CloudWatch Events when all files written to your
// file share have been uploaded to Amazon S3.  <p>AWS Storage Gateway can send a
// notification through Amazon CloudWatch Events when all files written to your
// file share up to that point in time have been uploaded to Amazon S3. These files
// include files written to the file share up to the time that you make a request
// for notification. When the upload is done, Storage Gateway sends you
// notification through an Amazon CloudWatch Event. You can configure CloudWatch
// Events to send the notification through event targets such as Amazon SNS or AWS
// Lambda function. This operation is only supported for file gateways.</p> <p>For
// more information, see <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-upload-notification">Getting
// file upload notification</a> in the <i>AWS Storage Gateway User Guide</i>.</p>
func (c *Client) NotifyWhenUploaded(ctx context.Context, params *NotifyWhenUploadedInput, optFns ...func(*Options)) (*NotifyWhenUploadedOutput, error) {
	stack := middleware.NewStack("NotifyWhenUploaded", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpNotifyWhenUploadedMiddlewares(stack)
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
	addOpNotifyWhenUploadedValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyWhenUploaded(options.Region), middleware.Before)

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
			OperationName: "NotifyWhenUploaded",
			Err:           err,
		}
	}
	out := result.(*NotifyWhenUploadedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyWhenUploadedInput struct {
	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string
}

type NotifyWhenUploadedOutput struct {
	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string
	// The randomly generated ID of the notification that was sent. This ID is in UUID
	// format.
	NotificationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpNotifyWhenUploadedMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyWhenUploaded{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyWhenUploaded{}, middleware.After)
}

func newServiceMetadataMiddleware_opNotifyWhenUploaded(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "NotifyWhenUploaded",
	}
}
