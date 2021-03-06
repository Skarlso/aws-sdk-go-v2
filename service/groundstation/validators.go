// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelContact struct {
}

func (*validateOpCancelContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateConfig struct {
}

func (*validateOpCreateConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataflowEndpointGroup struct {
}

func (*validateOpCreateDataflowEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataflowEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataflowEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataflowEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMissionProfile struct {
}

func (*validateOpCreateMissionProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMissionProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMissionProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMissionProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteConfig struct {
}

func (*validateOpDeleteConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataflowEndpointGroup struct {
}

func (*validateOpDeleteDataflowEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataflowEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataflowEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataflowEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMissionProfile struct {
}

func (*validateOpDeleteMissionProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMissionProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMissionProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMissionProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeContact struct {
}

func (*validateOpDescribeContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetConfig struct {
}

func (*validateOpGetConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDataflowEndpointGroup struct {
}

func (*validateOpGetDataflowEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDataflowEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDataflowEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDataflowEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMinuteUsage struct {
}

func (*validateOpGetMinuteUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMinuteUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMinuteUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMinuteUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMissionProfile struct {
}

func (*validateOpGetMissionProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMissionProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMissionProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMissionProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSatellite struct {
}

func (*validateOpGetSatellite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSatellite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSatelliteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSatelliteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListContacts struct {
}

func (*validateOpListContacts) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListContacts) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListContactsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListContactsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReserveContact struct {
}

func (*validateOpReserveContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReserveContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReserveContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReserveContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateConfig struct {
}

func (*validateOpUpdateConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMissionProfile struct {
}

func (*validateOpUpdateMissionProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMissionProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMissionProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMissionProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelContact{}, middleware.After)
}

func addOpCreateConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateConfig{}, middleware.After)
}

func addOpCreateDataflowEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataflowEndpointGroup{}, middleware.After)
}

func addOpCreateMissionProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMissionProfile{}, middleware.After)
}

func addOpDeleteConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteConfig{}, middleware.After)
}

func addOpDeleteDataflowEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataflowEndpointGroup{}, middleware.After)
}

func addOpDeleteMissionProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMissionProfile{}, middleware.After)
}

func addOpDescribeContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeContact{}, middleware.After)
}

func addOpGetConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetConfig{}, middleware.After)
}

func addOpGetDataflowEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDataflowEndpointGroup{}, middleware.After)
}

func addOpGetMinuteUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMinuteUsage{}, middleware.After)
}

func addOpGetMissionProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMissionProfile{}, middleware.After)
}

func addOpGetSatelliteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSatellite{}, middleware.After)
}

func addOpListContactsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListContacts{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpReserveContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReserveContact{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateConfig{}, middleware.After)
}

func addOpUpdateMissionProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMissionProfile{}, middleware.After)
}

func validateDataflowEndpoint(v *types.DataflowEndpoint) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataflowEndpoint"}
	if v.Address != nil {
		if err := validateSocketAddress(v.Address); err != nil {
			invalidParams.AddNested("Address", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEndpointDetails(v *types.EndpointDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointDetails"}
	if v.Endpoint != nil {
		if err := validateDataflowEndpoint(v.Endpoint); err != nil {
			invalidParams.AddNested("Endpoint", err.(smithy.InvalidParamsError))
		}
	}
	if v.SecurityDetails != nil {
		if err := validateSecurityDetails(v.SecurityDetails); err != nil {
			invalidParams.AddNested("SecurityDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEndpointDetailsList(v []types.EndpointDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointDetailsList"}
	for i := range v {
		if err := validateEndpointDetails(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSecurityDetails(v *types.SecurityDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SecurityDetails"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSocketAddress(v *types.SocketAddress) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SocketAddress"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Port == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Port"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelContactInput(v *CancelContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateConfigInput(v *CreateConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateConfigInput"}
	if v.ConfigData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigData"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataflowEndpointGroupInput(v *CreateDataflowEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataflowEndpointGroupInput"}
	if v.EndpointDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointDetails"))
	} else if v.EndpointDetails != nil {
		if err := validateEndpointDetailsList(v.EndpointDetails); err != nil {
			invalidParams.AddNested("EndpointDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMissionProfileInput(v *CreateMissionProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMissionProfileInput"}
	if v.DataflowEdges == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataflowEdges"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.TrackingConfigArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrackingConfigArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteConfigInput(v *DeleteConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteConfigInput"}
	if v.ConfigId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigId"))
	}
	if len(v.ConfigType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataflowEndpointGroupInput(v *DeleteDataflowEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataflowEndpointGroupInput"}
	if v.DataflowEndpointGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataflowEndpointGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMissionProfileInput(v *DeleteMissionProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMissionProfileInput"}
	if v.MissionProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MissionProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeContactInput(v *DescribeContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetConfigInput(v *GetConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetConfigInput"}
	if v.ConfigId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigId"))
	}
	if len(v.ConfigType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDataflowEndpointGroupInput(v *GetDataflowEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDataflowEndpointGroupInput"}
	if v.DataflowEndpointGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataflowEndpointGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMinuteUsageInput(v *GetMinuteUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMinuteUsageInput"}
	if v.Month == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Month"))
	}
	if v.Year == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Year"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMissionProfileInput(v *GetMissionProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMissionProfileInput"}
	if v.MissionProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MissionProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSatelliteInput(v *GetSatelliteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSatelliteInput"}
	if v.SatelliteId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SatelliteId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListContactsInput(v *ListContactsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListContactsInput"}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.StatusList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StatusList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReserveContactInput(v *ReserveContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReserveContactInput"}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if v.GroundStation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroundStation"))
	}
	if v.MissionProfileArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MissionProfileArn"))
	}
	if v.SatelliteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SatelliteArn"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateConfigInput(v *UpdateConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateConfigInput"}
	if v.ConfigData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigData"))
	}
	if v.ConfigId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigId"))
	}
	if len(v.ConfigType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigType"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMissionProfileInput(v *UpdateMissionProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMissionProfileInput"}
	if v.MissionProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MissionProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
