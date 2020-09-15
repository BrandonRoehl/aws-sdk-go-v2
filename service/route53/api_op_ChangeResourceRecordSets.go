// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates, changes, or deletes a resource record set, which contains authoritative
// DNS information for a specified domain name or subdomain name. For example, you
// can use ChangeResourceRecordSets to create a resource record set that routes
// traffic for test.example.com to a web server that has an IP address of
// 192.0.2.44.  <p> <b>Deleting Resource Record Sets</b> </p> <p>To delete a
// resource record set, you must specify all the same values that you specified
// when you created it.</p> <p> <b>Change Batches and Transactional Changes</b>
// </p> <p>The request body must include a document with a
// <code>ChangeResourceRecordSetsRequest</code> element. The request body contains
// a list of change items, known as a change batch. Change batches are considered
// transactional changes. Route 53 validates the changes in the request and then
// either makes all or none of the changes in the change batch request. This
// ensures that DNS routing isn't adversely affected by partial changes to the
// resource record sets in a hosted zone. </p> <p>For example, suppose a change
// batch request contains two changes: it deletes the <code>CNAME</code> resource
// record set for www.example.com and creates an alias resource record set for
// www.example.com. If validation for both records succeeds, Route 53 deletes the
// first resource record set and creates the second resource record set in a single
// operation. If validation for either the <code>DELETE</code> or the
// <code>CREATE</code> action fails, then the request is canceled, and the original
// <code>CNAME</code> record continues to exist.</p> <note> <p>If you try to delete
// the same resource record set more than once in a single change batch, Route 53
// returns an <code>InvalidChangeBatch</code> error.</p> </note> <p> <b>Traffic
// Flow</b> </p> <p>To create resource record sets for complex routing
// configurations, use either the traffic flow visual editor in the Route 53
// console or the API actions for traffic policies and traffic policy instances.
// Save the configuration as a traffic policy, then associate the traffic policy
// with one or more domain names (such as example.com) or subdomain names (such as
// www.example.com), in the same hosted zone or in multiple hosted zones. You can
// roll back the updates if the new configuration isn't performing as expected. For
// more information, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/traffic-flow.html">Using
// Traffic Flow to Route DNS Traffic</a> in the <i>Amazon Route 53 Developer
// Guide</i>.</p> <p> <b>Create, Delete, and Upsert</b> </p> <p>Use
// <code>ChangeResourceRecordsSetsRequest</code> to perform the following
// actions:</p> <ul> <li> <p> <code>CREATE</code>: Creates a resource record set
// that has the specified values.</p> </li> <li> <p> <code>DELETE</code>: Deletes
// an existing resource record set that has the specified values.</p> </li> <li>
// <p> <code>UPSERT</code>: If a resource record set does not already exist, AWS
// creates it. If a resource set does exist, Route 53 updates it with the values in
// the request. </p> </li> </ul> <p> <b>Syntaxes for Creating, Updating, and
// Deleting Resource Record Sets</b> </p> <p>The syntax for a request depends on
// the type of resource record set that you want to create, delete, or update, such
// as weighted, alias, or failover. The XML elements in your request must appear in
// the order listed in the syntax. </p> <p>For an example for each type of resource
// record set, see "Examples."</p> <p>Don't refer to the syntax in the "Parameter
// Syntax" section, which includes all of the elements for every kind of resource
// record set that you can create, delete, or update by using
// <code>ChangeResourceRecordSets</code>. </p> <p> <b>Change Propagation to Route
// 53 DNS Servers</b> </p> <p>When you submit a
// <code>ChangeResourceRecordSets</code> request, Route 53 propagates your changes
// to all of the Route 53 authoritative DNS servers. While your changes are
// propagating, <code>GetChange</code> returns a status of <code>PENDING</code>.
// When propagation is complete, <code>GetChange</code> returns a status of
// <code>INSYNC</code>. Changes generally propagate to all Route 53 name servers
// within 60 seconds. For more information, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html">GetChange</a>.</p>
// <p> <b>Limits on ChangeResourceRecordSets Requests</b> </p> <p>For information
// about the limits on a <code>ChangeResourceRecordSets</code> request, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html">Limits</a>
// in the <i>Amazon Route 53 Developer Guide</i>.</p>
func (c *Client) ChangeResourceRecordSets(ctx context.Context, params *ChangeResourceRecordSetsInput, optFns ...func(*Options)) (*ChangeResourceRecordSetsOutput, error) {
	stack := middleware.NewStack("ChangeResourceRecordSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpChangeResourceRecordSetsMiddlewares(stack)
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
	addOpChangeResourceRecordSetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opChangeResourceRecordSets(options.Region), middleware.Before)

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
			OperationName: "ChangeResourceRecordSets",
			Err:           err,
		}
	}
	out := result.(*ChangeResourceRecordSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains change information for the resource record set.
type ChangeResourceRecordSetsInput struct {
	// A complex type that contains an optional comment and the Changes element.
	ChangeBatch *types.ChangeBatch
	// The ID of the hosted zone that contains the resource record sets that you want
	// to change.
	HostedZoneId *string
}

// A complex type containing the response for the request.
type ChangeResourceRecordSetsOutput struct {
	// A complex type that contains information about changes made to your hosted zone.
	// This element contains an ID that you use when performing a GetChange
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html)
	// action to get detailed information about the change.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpChangeResourceRecordSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpChangeResourceRecordSets{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpChangeResourceRecordSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opChangeResourceRecordSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ChangeResourceRecordSets",
	}
}
