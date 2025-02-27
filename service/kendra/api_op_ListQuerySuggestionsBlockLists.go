// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the block lists used for query suggestions for an index. For information
// on the current quota limits for block lists, see Quotas for Amazon Kendra
// (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
func (c *Client) ListQuerySuggestionsBlockLists(ctx context.Context, params *ListQuerySuggestionsBlockListsInput, optFns ...func(*Options)) (*ListQuerySuggestionsBlockListsOutput, error) {
	if params == nil {
		params = &ListQuerySuggestionsBlockListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQuerySuggestionsBlockLists", params, optFns, addOperationListQuerySuggestionsBlockListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQuerySuggestionsBlockListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQuerySuggestionsBlockListsInput struct {

	// The identifier of the index for a list of all block lists that exist for that
	// index. For information on the current quota limits for block lists, see Quotas
	// for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
	//
	// This member is required.
	IndexId *string

	// The maximum number of block lists to return.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of block lists
	// (BlockListSummaryItems).
	NextToken *string
}

type ListQuerySuggestionsBlockListsOutput struct {

	// Summary items for a block list. This includes summary items on the block list
	// ID, block list name, when the block list was created, when the block list was
	// last updated, and the count of block words/phrases in the block list. For
	// information on the current quota limits for block lists, see Quotas for Amazon
	// Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
	BlockListSummaryItems []types.QuerySuggestionsBlockListSummary

	// If the response is truncated, Amazon Kendra returns this token that you can use
	// in the subsequent request to retrieve the next set of block lists.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQuerySuggestionsBlockListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQuerySuggestionsBlockLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQuerySuggestionsBlockLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListQuerySuggestionsBlockListsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQuerySuggestionsBlockLists(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListQuerySuggestionsBlockLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListQuerySuggestionsBlockLists",
	}
}
