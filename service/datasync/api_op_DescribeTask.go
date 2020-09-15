// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata about a task.
func (c *Client) DescribeTask(ctx context.Context, params *DescribeTaskInput, optFns ...func(*Options)) (*DescribeTaskOutput, error) {
	stack := middleware.NewStack("DescribeTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTaskMiddlewares(stack)
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
	addOpDescribeTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTask(options.Region), middleware.Before)

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
			OperationName: "DescribeTask",
			Err:           err,
		}
	}
	out := result.(*DescribeTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskRequest
type DescribeTaskInput struct {
	// The Amazon Resource Name (ARN) of the task to describe.
	TaskArn *string
}

// DescribeTaskResponse
type DescribeTaskOutput struct {
	// The Amazon Resource Name (ARN) of the destination ENIs (Elastic Network
	// Interface) that was created for your subnet.
	DestinationNetworkInterfaceArns []*string
	// Errors that AWS DataSync encountered during execution of the task. You can use
	// this error code to help troubleshoot issues.
	ErrorCode *string
	// The Amazon Resource Name (ARN) of the source file system's location.
	SourceLocationArn *string
	// The time that the task was created.
	CreationTime *time.Time
	// Detailed description of an error that was encountered during the task execution.
	// You can use this information to help troubleshoot issues.
	ErrorDetail *string
	// The status of the task that was described.  <p>For detailed information about
	// task execution statuses, see Understanding Task Statuses in the <i>AWS DataSync
	// User Guide.</i> </p>
	Status types.TaskStatus
	// The set of configuration options that control the behavior of a single execution
	// of the task that occurs when you call StartTaskExecution. You can configure
	// these options to preserve metadata such as user ID (UID) and group (GID), file
	// permissions, data integrity verification, and so on. For each individual task
	// execution, you can override these options by specifying the overriding
	// OverrideOptions value to operation.
	Options *types.Options
	// The Amazon Resource Name (ARN) of the AWS storage resource's location.
	DestinationLocationArn *string
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that was used
	// to monitor and log events in the task.  <p>For more information on these groups,
	// see Working with Log Groups and Log Streams in the <i>Amazon CloudWatch User
	// Guide</i>.</p>
	CloudWatchLogGroupArn *string
	// The Amazon Resource Name (ARN) of the source ENIs (Elastic Network Interface)
	// that was created for your subnet.
	SourceNetworkInterfaceArns []*string
	// The schedule used to periodically transfer files from a source to a destination
	// location.
	Schedule *types.TaskSchedule
	// The name of the task that was described.
	Name *string
	// The Amazon Resource Name (ARN) of the task execution that is syncing files.
	CurrentTaskExecutionArn *string
	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []*types.FilterRule
	// The Amazon Resource Name (ARN) of the task that was described.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTask",
	}
}
