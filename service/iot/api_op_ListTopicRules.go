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

// Lists the rules for the specific topic.
func (c *Client) ListTopicRules(ctx context.Context, params *ListTopicRulesInput, optFns ...func(*Options)) (*ListTopicRulesOutput, error) {
	stack := middleware.NewStack("ListTopicRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListTopicRulesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTopicRules(options.Region), middleware.Before)

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
			OperationName: "ListTopicRules",
			Err:           err,
		}
	}
	out := result.(*ListTopicRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListTopicRules operation.
type ListTopicRulesInput struct {
	// The topic.
	Topic *string
	// A token used to retrieve the next value.
	NextToken *string
	// Specifies whether the rule is disabled.
	RuleDisabled *bool
	// The maximum number of results to return.
	MaxResults *int32
}

// The output from the ListTopicRules operation.
type ListTopicRulesOutput struct {
	// A token used to retrieve the next value.
	NextToken *string
	// The rules.
	Rules []*types.TopicRuleListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListTopicRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListTopicRules{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListTopicRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTopicRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListTopicRules",
	}
}
