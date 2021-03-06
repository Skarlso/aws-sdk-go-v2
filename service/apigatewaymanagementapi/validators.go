// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewaymanagementapi

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteConnection struct {
}

func (*validateOpDeleteConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetConnection struct {
}

func (*validateOpGetConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPostToConnection struct {
}

func (*validateOpPostToConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPostToConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PostToConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPostToConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteConnection{}, middleware.After)
}

func addOpGetConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetConnection{}, middleware.After)
}

func addOpPostToConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPostToConnection{}, middleware.After)
}

func validateOpDeleteConnectionInput(v *DeleteConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteConnectionInput"}
	if v.ConnectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetConnectionInput(v *GetConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetConnectionInput"}
	if v.ConnectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPostToConnectionInput(v *PostToConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PostToConnectionInput"}
	if v.ConnectionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionId"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
