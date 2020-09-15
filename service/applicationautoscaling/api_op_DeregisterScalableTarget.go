// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters an Application Auto Scaling scalable target when you have finished
// using it. To see which resources have been registered, use
// DescribeScalableTargets
// (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html).
// Deregistering a scalable target deletes the scaling policies and the scheduled
// actions that are associated with it.
func (c *Client) DeregisterScalableTarget(ctx context.Context, params *DeregisterScalableTargetInput, optFns ...func(*Options)) (*DeregisterScalableTargetOutput, error) {
	stack := middleware.NewStack("DeregisterScalableTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterScalableTargetMiddlewares(stack)
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
	addOpDeregisterScalableTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterScalableTarget(options.Region), middleware.Before)

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
			OperationName: "DeregisterScalableTarget",
			Err:           err,
		}
	}
	out := result.(*DeregisterScalableTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterScalableTargetInput struct {
	// The namespace of the AWS service that provides the resource. For a resource
	// provided by your own application or service, use custom-resource instead.
	ServiceNamespace types.ServiceNamespace
	// The scalable dimension associated with the scalable target. This string consists
	// of the service namespace, resource type, and scaling property.
	//
	//     *
	// ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//     *
	// ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet
	// request.
	//
	//     * elasticmapreduce:instancegroup:InstanceCount - The instance
	// count of an EMR Instance Group.
	//
	//     * appstream:fleet:DesiredCapacity - The
	// desired capacity of an AppStream 2.0 fleet.
	//
	//     *
	// dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB
	// table.
	//
	//     * dynamodb:table:WriteCapacityUnits - The provisioned write capacity
	// for a DynamoDB table.
	//
	//     * dynamodb:index:ReadCapacityUnits - The provisioned
	// read capacity for a DynamoDB global secondary index.
	//
	//     *
	// dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a
	// DynamoDB global secondary index.
	//
	//     * rds:cluster:ReadReplicaCount - The count
	// of Aurora Replicas in an Aurora DB cluster. Available for Aurora
	// MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	//
	//     *
	// sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an
	// Amazon SageMaker model endpoint variant.
	//
	//     *
	// custom-resource:ResourceType:Property - The scalable dimension for a custom
	// resource provided by your own application or service.
	//
	//     *
	// comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend document classification endpoint.
	//
	//
	// * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a
	// Lambda function.
	//
	//     * cassandra:table:ReadCapacityUnits - The provisioned read
	// capacity for an Amazon Keyspaces table.
	//
	//     *
	// cassandra:table:WriteCapacityUnits - The provisioned write capacity for an
	// Amazon Keyspaces table.
	ScalableDimension types.ScalableDimension
	// The identifier of the resource associated with the scalable target. This string
	// consists of the resource type and unique identifier.
	//
	//     * ECS service - The
	// resource type is service and the unique identifier is the cluster name and
	// service name. Example: service/default/sample-webapp.
	//
	//     * Spot Fleet request
	// - The resource type is spot-fleet-request and the unique identifier is the Spot
	// Fleet request ID. Example:
	// spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//     * EMR cluster
	// - The resource type is instancegroup and the unique identifier is the cluster ID
	// and instance group ID. Example:
	// instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//     * AppStream 2.0 fleet - The
	// resource type is fleet and the unique identifier is the fleet name. Example:
	// fleet/sample-fleet.
	//
	//     * DynamoDB table - The resource type is table and the
	// unique identifier is the table name. Example: table/my-table.
	//
	//     * DynamoDB
	// global secondary index - The resource type is index and the unique identifier is
	// the index name. Example: table/my-table/index/my-table-index.
	//
	//     * Aurora DB
	// cluster - The resource type is cluster and the unique identifier is the cluster
	// name. Example: cluster:my-db-cluster.
	//
	//     * Amazon SageMaker endpoint variant -
	// The resource type is variant and the unique identifier is the resource ID.
	// Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//     * Custom resources
	// are not supported with a resource type. This parameter must specify the
	// OutputValue from the CloudFormation template stack used to access the resources.
	// The unique identifier is defined by the service provider. More information is
	// available in our GitHub repository
	// (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	//     * Amazon
	// Comprehend document classification endpoint - The resource type and unique
	// identifier are specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	//
	// * Lambda provisioned concurrency - The resource type is function and the unique
	// identifier is the function name with a function version or alias name suffix
	// that is not $LATEST. Example: function:my-function:prod or
	// function:my-function:1.
	//
	//     * Amazon Keyspaces table - The resource type is
	// table and the unique identifier is the table name. Example:
	// keyspace/mykeyspace/table/mytable.
	ResourceId *string
}

type DeregisterScalableTargetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterScalableTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterScalableTarget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterScalableTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterScalableTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-autoscaling",
		OperationName: "DeregisterScalableTarget",
	}
}
