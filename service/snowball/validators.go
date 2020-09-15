// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCancelCluster struct {
}

func (*validateOpCancelCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelJob struct {
}

func (*validateOpCancelJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAddress struct {
}

func (*validateOpCreateAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateJob struct {
}

func (*validateOpCreateJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAddress struct {
}

func (*validateOpDescribeAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCluster struct {
}

func (*validateOpDescribeCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJob struct {
}

func (*validateOpDescribeJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetJobManifest struct {
}

func (*validateOpGetJobManifest) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetJobManifest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetJobManifestInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetJobManifestInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetJobUnlockCode struct {
}

func (*validateOpGetJobUnlockCode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetJobUnlockCode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetJobUnlockCodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetJobUnlockCodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSoftwareUpdates struct {
}

func (*validateOpGetSoftwareUpdates) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSoftwareUpdates) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSoftwareUpdatesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSoftwareUpdatesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListClusterJobs struct {
}

func (*validateOpListClusterJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListClusterJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListClusterJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListClusterJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCluster struct {
}

func (*validateOpUpdateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateJob struct {
}

func (*validateOpUpdateJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelCluster{}, middleware.After)
}

func addOpCancelJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelJob{}, middleware.After)
}

func addOpCreateAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAddress{}, middleware.After)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateJob{}, middleware.After)
}

func addOpDescribeAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAddress{}, middleware.After)
}

func addOpDescribeClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCluster{}, middleware.After)
}

func addOpDescribeJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJob{}, middleware.After)
}

func addOpGetJobManifestValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetJobManifest{}, middleware.After)
}

func addOpGetJobUnlockCodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetJobUnlockCode{}, middleware.After)
}

func addOpGetSoftwareUpdatesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSoftwareUpdates{}, middleware.After)
}

func addOpListClusterJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListClusterJobs{}, middleware.After)
}

func addOpUpdateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCluster{}, middleware.After)
}

func addOpUpdateJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateJob{}, middleware.After)
}

func validateEc2AmiResource(v *types.Ec2AmiResource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Ec2AmiResource"}
	if v.AmiId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AmiId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEc2AmiResourceList(v []*types.Ec2AmiResource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Ec2AmiResourceList"}
	for i := range v {
		if err := validateEc2AmiResource(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobResource(v *types.JobResource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobResource"}
	if v.Ec2AmiResources != nil {
		if err := validateEc2AmiResourceList(v.Ec2AmiResources); err != nil {
			invalidParams.AddNested("Ec2AmiResources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelClusterInput(v *CancelClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelClusterInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelJobInput(v *CancelJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAddressInput(v *CreateAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAddressInput"}
	if v.Address == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Address"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.AddressId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddressId"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if len(v.ShippingOption) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ShippingOption"))
	}
	if len(v.JobType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("JobType"))
	}
	if v.Resources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Resources"))
	} else if v.Resources != nil {
		if err := validateJobResource(v.Resources); err != nil {
			invalidParams.AddNested("Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateJobInput(v *CreateJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateJobInput"}
	if v.Resources != nil {
		if err := validateJobResource(v.Resources); err != nil {
			invalidParams.AddNested("Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAddressInput(v *DescribeAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAddressInput"}
	if v.AddressId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddressId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeClusterInput(v *DescribeClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeClusterInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeJobInput(v *DescribeJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetJobManifestInput(v *GetJobManifestInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetJobManifestInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetJobUnlockCodeInput(v *GetJobUnlockCodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetJobUnlockCodeInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSoftwareUpdatesInput(v *GetSoftwareUpdatesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSoftwareUpdatesInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListClusterJobsInput(v *ListClusterJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListClusterJobsInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateClusterInput(v *UpdateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateClusterInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if v.Resources != nil {
		if err := validateJobResource(v.Resources); err != nil {
			invalidParams.AddNested("Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateJobInput(v *UpdateJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.Resources != nil {
		if err := validateJobResource(v.Resources); err != nil {
			invalidParams.AddNested("Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
