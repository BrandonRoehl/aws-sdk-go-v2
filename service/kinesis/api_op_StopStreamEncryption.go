// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables server-side encryption for a specified stream. Stopping encryption is
// an asynchronous operation. Upon receiving the request, Kinesis Data Streams
// returns immediately and sets the status of the stream to UPDATING. After the
// update is complete, Kinesis Data Streams sets the status of the stream back to
// ACTIVE. Stopping encryption normally takes a few seconds to complete, but it can
// take minutes. You can continue to read and write data to your stream while its
// status is UPDATING. Once the status of the stream is ACTIVE, records written to
// the stream are no longer encrypted by Kinesis Data Streams. API Limits: You can
// successfully disable server-side encryption 25 times in a rolling 24-hour
// period. Note: It can take up to 5 seconds after the stream is in an ACTIVE
// status before all records written to the stream are no longer subject to
// encryption. After you disabled encryption, you can verify that encryption is not
// applied by inspecting the API response from PutRecord or PutRecords.
func (c *Client) StopStreamEncryption(ctx context.Context, params *StopStreamEncryptionInput, optFns ...func(*Options)) (*StopStreamEncryptionOutput, error) {
	stack := middleware.NewStack("StopStreamEncryption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopStreamEncryptionMiddlewares(stack)
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
	addOpStopStreamEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopStreamEncryption(options.Region), middleware.Before)

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
			OperationName: "StopStreamEncryption",
			Err:           err,
		}
	}
	out := result.(*StopStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopStreamEncryptionInput struct {
	// The GUID for the customer-managed AWS KMS key to use for encryption. This value
	// can be a globally unique identifier, a fully specified Amazon Resource Name
	// (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can
	// also use a master key owned by Kinesis Data Streams by specifying the alias
	// aws/kinesis.
	//
	//     * Key ARN example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//
	// * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//     *
	// Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//     *
	// Alias name example: alias/MyAliasName
	//
	//     * Master key owned by Kinesis Data
	// Streams: alias/aws/kinesis
	KeyId *string
	// The name of the stream on which to stop encrypting records.
	StreamName *string
	// The encryption type. The only valid value is KMS.
	EncryptionType types.EncryptionType
}

type StopStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopStreamEncryptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopStreamEncryption{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopStreamEncryption{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopStreamEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "StopStreamEncryption",
	}
}
