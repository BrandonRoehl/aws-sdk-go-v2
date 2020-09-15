// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes Predicate () objects in a Rule. Each Predicate
// object identifies a predicate, such as a ByteMatchSet () or an IPSet (), that
// specifies the web requests that you want to allow, block, or count. If you add
// more than one predicate to a Rule, a request must match all of the
// specifications to be allowed, blocked, or counted. For example, suppose that you
// add the following to a Rule:
//
//     * A ByteMatchSet that matches the value BadBot
// in the User-Agent header
//
//     * An IPSet that matches the IP address
// 192.0.2.44
//
// You then add the Rule to a WebACL and specify that you want to block
// requests that satisfy the Rule. For a request to be blocked, the User-Agent
// header in the request must contain the value BadBot and the request must
// originate from the IP address 192.0.2.44. To create and configure a Rule,
// perform the following steps:
//
//     * Create and update the predicates that you
// want to include in the Rule.
//
//     * Create the Rule. See CreateRule ().
//
//     *
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRule () request.
//
//     * Submit an UpdateRule request to
// add predicates to the Rule.
//
//     * Create and update a WebACL that contains the
// Rule. See CreateWebACL ().
//
// If you want to replace one ByteMatchSet or IPSet
// with another, you delete the existing one and add the new one. For more
// information about how to use the AWS WAF API to allow or block HTTP requests,
// see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) {
	stack := middleware.NewStack("UpdateRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRuleMiddlewares(stack)
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
	addOpUpdateRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRule(options.Region), middleware.Before)

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
			OperationName: "UpdateRule",
			Err:           err,
		}
	}
	out := result.(*UpdateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleInput struct {
	// An array of RuleUpdate objects that you want to insert into or delete from a
	// Rule (). For more information, see the applicable data types:
	//
	//     * RuleUpdate
	// (): Contains Action and Predicate
	//
	//     * Predicate (): Contains DataId, Negated,
	// and Type
	//
	//     * FieldToMatch (): Contains Data and Type
	Updates []*types.RuleUpdate
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
	// The RuleId of the Rule that you want to update. RuleId is returned by CreateRule
	// and by ListRules ().
	RuleId *string
}

type UpdateRuleOutput struct {
	// The ChangeToken that you used to submit the UpdateRule request. You can also use
	// this value to query the status of the request. For more information, see
	// GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateRule",
	}
}
