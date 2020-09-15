// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a new task using the specified task definition. You can allow Amazon ECS
// to place tasks for you, or you can customize how Amazon ECS places tasks using
// placement constraints and placement strategies. For more information, see
// Scheduling Tasks
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html)
// in the Amazon Elastic Container Service Developer Guide. Alternatively, you can
// use StartTask () to use your own scheduler or place tasks manually on specific
// container instances. The Amazon ECS API follows an eventual consistency model,
// due to the distributed nature of the system supporting the API. This means that
// the result of an API command you run that affects your Amazon ECS resources
// might not be immediately visible to all subsequent commands you run. Keep this
// in mind when you carry out an API command that immediately follows a previous
// API command. To manage eventual consistency, you can do the following:
//
//     *
// Confirm the state of the resource before you run a command to modify it. Run the
// DescribeTasks command using an exponential backoff algorithm to ensure that you
// allow enough time for the previous command to propagate through the system. To
// do this, run the DescribeTasks command repeatedly, starting with a couple of
// seconds of wait time and increasing gradually up to five minutes of wait time.
//
//
// * Add wait time between subsequent commands, even if the DescribeTasks command
// returns an accurate response. Apply an exponential backoff algorithm starting
// with a couple of seconds of wait time, and increase gradually up to about five
// minutes of wait time.
func (c *Client) RunTask(ctx context.Context, params *RunTaskInput, optFns ...func(*Options)) (*RunTaskOutput, error) {
	stack := middleware.NewStack("RunTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRunTaskMiddlewares(stack)
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
	addOpRunTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRunTask(options.Region), middleware.Before)

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
			OperationName: "RunTask",
			Err:           err,
		}
	}
	out := result.(*RunTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RunTaskInput struct {
	// The placement strategy objects to use for the task. You can specify a maximum of
	// five strategy rules per task.
	PlacementStrategy []*types.PlacementStrategy
	// Specifies whether to enable Amazon ECS managed tags for the task. For more
	// information, see Tagging Your Amazon ECS Resources
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide.
	EnableECSManagedTags *bool
	// The metadata that you apply to the task to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. The following basic restrictions apply to tags:
	//
	//     * Maximum number of
	// tags per resource - 50
	//
	//     * For each resource, each tag key must be unique,
	// and each tag key can have only one value.
	//
	//     * Maximum key length - 128
	// Unicode characters in UTF-8
	//
	//     * Maximum value length - 256 Unicode characters
	// in UTF-8
	//
	//     * If your tagging schema is used across multiple services and
	// resources, remember that other services may have restrictions on allowed
	// characters. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//     *
	// Tag keys and values are case-sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for either keys or values as
	// it is reserved for AWS use. You cannot edit or delete tag keys or values with
	// this prefix. Tags with this prefix do not count against your tags per resource
	// limit.
	Tags []*types.Tag
	// The capacity provider strategy to use for the task. A capacity provider strategy
	// consists of one or more capacity providers along with the base and weight to
	// assign to them. A capacity provider must be associated with the cluster to be
	// used in a capacity provider strategy. The PutClusterCapacityProviders () API is
	// used to associate a capacity provider with a cluster. Only capacity providers
	// with an ACTIVE or UPDATING status can be used. If a capacityProviderStrategy is
	// specified, the launchType parameter must be omitted. If no
	// capacityProviderStrategy or launchType is specified, the
	// defaultCapacityProviderStrategy for the cluster is used. If specifying a
	// capacity provider that uses an Auto Scaling group, the capacity provider must
	// already be created. New capacity providers can be created with the
	// CreateCapacityProvider () API operation. To use a AWS Fargate capacity provider,
	// specify either the FARGATE or FARGATE_SPOT capacity providers. The AWS Fargate
	// capacity providers are available to all accounts and only need to be associated
	// with a cluster to be used. The PutClusterCapacityProviders () API operation is
	// used to update the list of available capacity providers for a cluster after the
	// cluster is created.
	CapacityProviderStrategy []*types.CapacityProviderStrategyItem
	// The launch type on which to run your task. For more information, see Amazon ECS
	// Launch Types
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// in the Amazon Elastic Container Service Developer Guide. If a launchType is
	// specified, the capacityProviderStrategy parameter must be omitted.
	LaunchType types.LaunchType
	// Specifies whether to propagate the tags from the task definition to the task. If
	// no value is specified, the tags are not propagated. Tags can only be propagated
	// to the task during task creation. To add tags to a task after task creation, use
	// the TagResource () API action. An error will be received if you specify the
	// SERVICE option when running a task.
	PropagateTags types.PropagateTags
	// A list of container overrides in JSON format that specify the name of a
	// container in the specified task definition and the overrides it should receive.
	// You can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition or
	// Docker image) on a container or add new environment variables to it with an
	// environment override. A total of 8192 characters are allowed for overrides. This
	// limit includes the JSON formatting characters of the override structure.
	Overrides *types.TaskOverride
	// The reference ID to use for the task.
	ReferenceId *string
	// An optional tag specified when a task is started. For example, if you
	// automatically trigger a task to run a batch process job, you could apply a
	// unique identifier for that job to your task with the startedBy parameter. You
	// can then identify which tasks belong to that job by filtering the results of a
	// ListTasks () call with the startedBy value. Up to 36 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed. If a task is started
	// by an Amazon ECS service, then the startedBy parameter contains the deployment
	// ID of the service that starts it.
	StartedBy *string
	// The network configuration for the task. This parameter is required for task
	// definitions that use the awsvpc network mode to receive their own elastic
	// network interface, and it is not supported for other network modes. For more
	// information, see Task Networking
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
	// in the Amazon Elastic Container Service Developer Guide.
	NetworkConfiguration *types.NetworkConfiguration
	// The short name or full Amazon Resource Name (ARN) of the cluster on which to run
	// your task. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string
	// The name of the task group to associate with the task. The default value is the
	// family name of the task definition (for example, family:my-family-name).
	Group *string
	// An array of placement constraint objects to use for the task. You can specify up
	// to 10 constraints per task (including constraints in the task definition and
	// those specified at runtime).
	PlacementConstraints []*types.PlacementConstraint
	// The family and revision (family:revision) or full ARN of the task definition to
	// run. If a revision is not specified, the latest ACTIVE revision is used.
	TaskDefinition *string
	// The platform version the task should run. A platform version is only specified
	// for tasks using the Fargate launch type. If one is not specified, the LATEST
	// platform version is used by default. For more information, see AWS Fargate
	// Platform Versions
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string
	// The number of instantiations of the specified task to place on your cluster. You
	// can specify up to 10 tasks per call.
	Count *int32
}

type RunTaskOutput struct {
	// A full description of the tasks that were run. The tasks that were successfully
	// placed on your cluster are described here.
	Tasks []*types.Task
	// Any failures associated with the call.
	Failures []*types.Failure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRunTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRunTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRunTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opRunTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "RunTask",
	}
}
