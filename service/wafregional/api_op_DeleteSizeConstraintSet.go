// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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
// global use. Permanently deletes a SizeConstraintSet (). You can't delete a
// SizeConstraintSet if it's still used in any Rules or if it still includes any
// SizeConstraint () objects (any filters). If you just want to remove a
// SizeConstraintSet from a Rule, use UpdateRule (). To permanently delete a
// SizeConstraintSet, perform the following steps:
//
//     * Update the
// SizeConstraintSet to remove filters, if any. For more information, see
// UpdateSizeConstraintSet ().
//
//     * Use GetChangeToken () to get the change token
// that you provide in the ChangeToken parameter of a DeleteSizeConstraintSet
// request.
//
//     * Submit a DeleteSizeConstraintSet request.
func (c *Client) DeleteSizeConstraintSet(ctx context.Context, params *DeleteSizeConstraintSetInput, optFns ...func(*Options)) (*DeleteSizeConstraintSetOutput, error) {
	stack := middleware.NewStack("DeleteSizeConstraintSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteSizeConstraintSetMiddlewares(stack)
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
	addOpDeleteSizeConstraintSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSizeConstraintSet(options.Region), middleware.Before)

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
			OperationName: "DeleteSizeConstraintSet",
			Err:           err,
		}
	}
	out := result.(*DeleteSizeConstraintSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSizeConstraintSetInput struct {
	// The SizeConstraintSetId of the SizeConstraintSet () that you want to delete.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet () and by
	// ListSizeConstraintSets ().
	SizeConstraintSetId *string
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
}

type DeleteSizeConstraintSetOutput struct {
	// The ChangeToken that you used to submit the DeleteSizeConstraintSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteSizeConstraintSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSizeConstraintSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSizeConstraintSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSizeConstraintSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "DeleteSizeConstraintSet",
	}
}
