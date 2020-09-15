// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a description for one or more Network File System (NFS) file shares from a
// file gateway. This operation is only supported for file gateways.
func (c *Client) DescribeNFSFileShares(ctx context.Context, params *DescribeNFSFileSharesInput, optFns ...func(*Options)) (*DescribeNFSFileSharesOutput, error) {
	stack := middleware.NewStack("DescribeNFSFileShares", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeNFSFileSharesMiddlewares(stack)
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
	addOpDescribeNFSFileSharesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNFSFileShares(options.Region), middleware.Before)

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
			OperationName: "DescribeNFSFileShares",
			Err:           err,
		}
	}
	out := result.(*DescribeNFSFileSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeNFSFileSharesInput
type DescribeNFSFileSharesInput struct {
	// An array containing the Amazon Resource Name (ARN) of each file share to be
	// described.
	FileShareARNList []*string
}

// DescribeNFSFileSharesOutput
type DescribeNFSFileSharesOutput struct {
	// An array containing a description for each requested file share.
	NFSFileShareInfoList []*types.NFSFileShareInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeNFSFileSharesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNFSFileShares{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNFSFileShares{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNFSFileShares(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeNFSFileShares",
	}
}
