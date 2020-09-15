// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Simulate how a set of IAM policies and optionally a resource-based policy works
// with a list of API operations and AWS resources to determine the policies'
// effective permissions. The policies are provided as strings. The simulation does
// not perform the API operations; it only checks the authorization to determine if
// the simulated policies allow or deny the operations. If you want to simulate
// existing policies that are attached to an IAM user, group, or role, use
// SimulatePrincipalPolicy () instead. Context keys are variables that are
// maintained by AWS and its services and which provide details about the context
// of an API query request. You can use the Condition element of an IAM policy to
// evaluate context keys. To get the list of context keys that the policies require
// for correct simulation, use GetContextKeysForCustomPolicy (). If the output is
// long, you can use MaxItems and Marker parameters to paginate the results.
func (c *Client) SimulateCustomPolicy(ctx context.Context, params *SimulateCustomPolicyInput, optFns ...func(*Options)) (*SimulateCustomPolicyOutput, error) {
	stack := middleware.NewStack("SimulateCustomPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSimulateCustomPolicyMiddlewares(stack)
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
	addOpSimulateCustomPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimulateCustomPolicy(options.Region), middleware.Before)

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
			OperationName: "SimulateCustomPolicy",
			Err:           err,
		}
	}
	out := result.(*SimulateCustomPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimulateCustomPolicyInput struct {
	// A list of context keys and corresponding values for the simulation to use.
	// Whenever a context key is evaluated in one of the simulated IAM permissions
	// policies, the corresponding value is supplied.
	ContextEntries []*types.ContextEntry
	// A list of names of API operations to evaluate in the simulation. Each operation
	// is evaluated against each resource. Each operation must include the service
	// identifier, such as iam:CreateUser. This operation does not support using
	// wildcards (*) in an action name.
	ActionNames []*string
	// An ARN representing the AWS account ID that specifies the owner of any simulated
	// resource that does not identify its owner in the resource ARN. Examples of
	// resource ARNs include an S3 bucket or object. If ResourceOwner is specified, it
	// is also used as the account owner of any ResourcePolicy included in the
	// simulation. If the ResourceOwner parameter is not specified, then the owner of
	// the resources and the resource policy defaults to the account of the identity
	// provided in CallerArn. This parameter is required only if you specify a
	// resource-based policy and account that owns the resource is different from the
	// account that owns the simulated calling user CallerArn. The ARN for an account
	// uses the following syntax: arn:aws:iam::AWS-account-ID:root. For example, to
	// represent the account with the 112233445566 ID, use the following ARN:
	// arn:aws:iam::112233445566-ID:root.
	ResourceOwner *string
	// Specifies the type of simulation to run. Different API operations that support
	// resource-based policies require different combinations of resources. By
	// specifying the type of simulation to run, you enable the policy simulator to
	// enforce the presence of the required resources to ensure reliable simulation
	// results. If your simulation does not match one of the following scenarios, then
	// you can omit this parameter. The following list shows each of the supported
	// scenario values and the resources that you must define to run the simulation.
	// Each of the EC2 scenarios requires that you specify instance, image, and
	// security-group resources. If your scenario includes an EBS volume, then you must
	// specify that volume as a resource. If the EC2 scenario includes VPC, then you
	// must supply the network-interface resource. If it includes an IP subnet, then
	// you must specify the subnet resource. For more information on the EC2 scenario
	// options, see Supported Platforms
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html)
	// in the Amazon EC2 User Guide.
	//
	//     * EC2-Classic-InstanceStore instance, image,
	// security-group
	//
	//     * EC2-Classic-EBS instance, image, security-group, volume
	//
	//
	// * EC2-VPC-InstanceStore instance, image, security-group, network-interface
	//
	//
	// * EC2-VPC-InstanceStore-Subnet instance, image, security-group,
	// network-interface, subnet
	//
	//     * EC2-VPC-EBS instance, image, security-group,
	// network-interface, volume
	//
	//     * EC2-VPC-EBS-Subnet instance, image,
	// security-group, network-interface, subnet, volume
	ResourceHandlingOption *string
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32
	// A resource-based policy to include in the simulation provided as a string. Each
	// resource in the simulation is treated as if it had this policy attached. You can
	// include only one resource-based policy in a simulation. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	//     * Any printable ASCII character
	// ranging from the space character (\u0020) through the end of the ASCII character
	// range
	//
	//     * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	//     * The special characters tab (\u0009), line
	// feed (\u000A), and carriage return (\u000D)
	ResourcePolicy *string
	// A list of policy documents to include in the simulation. Each document is
	// specified as a string containing the complete, valid JSON text of an IAM policy.
	// Do not include any resource-based policies in this parameter. Any resource-based
	// policy must be submitted with the ResourcePolicy parameter. The policies cannot
	// be "scope-down" policies, such as you could include in a call to
	// GetFederationToken
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetFederationToken.html)
	// or one of the AssumeRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_AssumeRole.html) API
	// operations. In other words, do not use policies designed to restrict what a user
	// can do while using the temporary credentials. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	//     * Any printable ASCII character
	// ranging from the space character (\u0020) through the end of the ASCII character
	// range
	//
	//     * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	//     * The special characters tab (\u0009), line
	// feed (\u000A), and carriage return (\u000D)
	PolicyInputList []*string
	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string
	// The ARN of the IAM user that you want to use as the simulated caller of the API
	// operations. CallerArn is required if you include a ResourcePolicy so that the
	// policy's Principal element has a value to use in evaluating the policy. You can
	// specify only the ARN of an IAM user. You cannot specify the ARN of an assumed
	// role, federated user, or a service principal.
	CallerArn *string
	// A list of ARNs of AWS resources to include in the simulation. If this parameter
	// is not provided, then the value defaults to * (all resources). Each API in the
	// ActionNames parameter is evaluated for each resource in this list. The
	// simulation determines the access result (allowed or denied) of each combination
	// and reports it in the response. The simulation does not automatically retrieve
	// policies for the specified resources. If you want to include a resource policy
	// in the simulation, then you must include the policy as a string in the
	// ResourcePolicy parameter. If you include a ResourcePolicy, then it must be
	// applicable to all of the resources included in the simulation or you receive an
	// invalid input error. For more information about ARNs, see Amazon Resource Names
	// (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	ResourceArns []*string
	// The IAM permissions boundary policy to simulate. The permissions boundary sets
	// the maximum permissions that an IAM entity can have. You can input only one
	// permissions boundary when you pass a policy to this operation. For more
	// information about permissions boundaries, see Permissions Boundaries for IAM
	// Entities
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide. The policy input is specified as a string that contains
	// the complete, valid JSON text of a permissions boundary policy. The regex
	// pattern (http://wikipedia.org/wiki/regex) used to validate this parameter is a
	// string of characters consisting of the following:
	//
	//     * Any printable ASCII
	// character ranging from the space character (\u0020) through the end of the ASCII
	// character range
	//
	//     * The printable characters in the Basic Latin and Latin-1
	// Supplement character set (through \u00FF)
	//
	//     * The special characters tab
	// (\u0009), line feed (\u000A), and carriage return (\u000D)
	PermissionsBoundaryPolicyInputList []*string
}

// Contains the response to a successful SimulatePrincipalPolicy () or
// SimulateCustomPolicy () request.
type SimulateCustomPolicyOutput struct {
	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool
	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string
	// The results of the simulation.
	EvaluationResults []*types.EvaluationResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSimulateCustomPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSimulateCustomPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSimulateCustomPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opSimulateCustomPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SimulateCustomPolicy",
	}
}
