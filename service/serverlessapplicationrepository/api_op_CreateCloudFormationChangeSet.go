// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS CloudFormation change set for the given application.
func (c *Client) CreateCloudFormationChangeSet(ctx context.Context, params *CreateCloudFormationChangeSetInput, optFns ...func(*Options)) (*CreateCloudFormationChangeSetOutput, error) {
	stack := middleware.NewStack("CreateCloudFormationChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateCloudFormationChangeSetMiddlewares(stack)
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
	addOpCreateCloudFormationChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFormationChangeSet(options.Region), middleware.Before)

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
			OperationName: "CreateCloudFormationChangeSet",
			Err:           err,
		}
	}
	out := result.(*CreateCloudFormationChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCloudFormationChangeSetInput struct {
	// The UUID returned by CreateCloudFormationTemplate.Pattern:
	// [0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}
	TemplateId *string
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	ChangeSetName *string
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	StackName *string
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	Tags []*types.Tag
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	RollbackConfiguration *types.RollbackConfiguration
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	ResourceTypes []*string
	// A list of values that you must specify before you can deploy certain
	// applications. Some applications might include resources that can affect
	// permissions in your AWS account, for example, by creating new AWS Identity and
	// Access Management (IAM) users. For those applications, you must explicitly
	// acknowledge their capabilities by specifying this parameter.The only valid
	// values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, and
	// CAPABILITY_AUTO_EXPAND.The following resources require you to specify
	// CAPABILITY_IAM or CAPABILITY_NAMED_IAM: AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html),
	// AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html),
	// AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// and AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html).
	// If the application contains IAM resources, you can specify either CAPABILITY_IAM
	// or CAPABILITY_NAMED_IAM. If the application contains IAM resources with custom
	// names, you must specify CAPABILITY_NAMED_IAM.The following resources require you
	// to specify CAPABILITY_RESOURCE_POLICY: AWS::Lambda::Permission
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html),
	// AWS::IAM:Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// AWS::ApplicationAutoScaling::ScalingPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html),
	// AWS::S3::BucketPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html),
	// AWS::SQS::QueuePolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html),
	// and AWS::SNS:TopicPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html).Applications
	// that contain one or more nested applications require you to specify
	// CAPABILITY_AUTO_EXPAND.If your application template contains any of the above
	// resources, we recommend that you review all permissions associated with the
	// application before deploying. If you don't specify this parameter for an
	// application that requires capabilities, the call will fail.
	Capabilities []*string
	// A list of parameter values for the parameters of the application.
	ParameterOverrides []*types.ParameterValue
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	Description *string
	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	ClientToken *string
	// The Amazon Resource Name (ARN) of the application.
	ApplicationId *string
	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation CreateChangeSet
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)
	// API.
	NotificationArns []*string
}

type CreateCloudFormationChangeSetOutput struct {
	// The Amazon Resource Name (ARN) of the change set.Length constraints: Minimum
	// length of 1.Pattern: ARN:[-a-zA-Z0-9:/]*
	ChangeSetId *string
	// The application Amazon Resource Name (ARN).
	ApplicationId *string
	// The unique ID of the stack.
	StackId *string
	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateCloudFormationChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateCloudFormationChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCloudFormationChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCloudFormationChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "CreateCloudFormationChangeSet",
	}
}
