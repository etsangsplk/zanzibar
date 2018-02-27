// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bazclient

import (
	"context"
	"errors"

	"github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/wire"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

// SecondServiceEchoBinaryFunc is the handler function for "echoBinary" method of thrift service "SecondService".
type SecondServiceEchoBinaryFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBinary_Args,
) ([]byte, map[string]string, error)

// NewSecondServiceEchoBinaryHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoBinaryHandler(f SecondServiceEchoBinaryFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoBinaryHandler{f}
}

// SecondServiceEchoBinaryHandler handles the "echoBinary" method call of thrift service "SecondService".
type SecondServiceEchoBinaryHandler struct {
	echobinary SecondServiceEchoBinaryFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoBinaryHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoBinary_Args
	var res clientsBazBaz.SecondService_EchoBinary_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echobinary(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoBoolFunc is the handler function for "echoBool" method of thrift service "SecondService".
type SecondServiceEchoBoolFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBool_Args,
) (bool, map[string]string, error)

// NewSecondServiceEchoBoolHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoBoolHandler(f SecondServiceEchoBoolFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoBoolHandler{f}
}

// SecondServiceEchoBoolHandler handles the "echoBool" method call of thrift service "SecondService".
type SecondServiceEchoBoolHandler struct {
	echobool SecondServiceEchoBoolFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoBoolHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoBool_Args
	var res clientsBazBaz.SecondService_EchoBool_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echobool(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoDoubleFunc is the handler function for "echoDouble" method of thrift service "SecondService".
type SecondServiceEchoDoubleFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoDouble_Args,
) (float64, map[string]string, error)

// NewSecondServiceEchoDoubleHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoDoubleHandler(f SecondServiceEchoDoubleFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoDoubleHandler{f}
}

// SecondServiceEchoDoubleHandler handles the "echoDouble" method call of thrift service "SecondService".
type SecondServiceEchoDoubleHandler struct {
	echodouble SecondServiceEchoDoubleFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoDoubleHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoDouble_Args
	var res clientsBazBaz.SecondService_EchoDouble_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echodouble(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoEnumFunc is the handler function for "echoEnum" method of thrift service "SecondService".
type SecondServiceEchoEnumFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoEnum_Args,
) (clientsBazBaz.Fruit, map[string]string, error)

// NewSecondServiceEchoEnumHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoEnumHandler(f SecondServiceEchoEnumFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoEnumHandler{f}
}

// SecondServiceEchoEnumHandler handles the "echoEnum" method call of thrift service "SecondService".
type SecondServiceEchoEnumHandler struct {
	echoenum SecondServiceEchoEnumFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoEnumHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoEnum_Args
	var res clientsBazBaz.SecondService_EchoEnum_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echoenum(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoI16Func is the handler function for "echoI16" method of thrift service "SecondService".
type SecondServiceEchoI16Func func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI16_Args,
) (int16, map[string]string, error)

// NewSecondServiceEchoI16Handler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoI16Handler(f SecondServiceEchoI16Func) zanzibar.TChannelHandler {
	return &SecondServiceEchoI16Handler{f}
}

// SecondServiceEchoI16Handler handles the "echoI16" method call of thrift service "SecondService".
type SecondServiceEchoI16Handler struct {
	echoi16 SecondServiceEchoI16Func
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoI16Handler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoI16_Args
	var res clientsBazBaz.SecondService_EchoI16_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echoi16(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoI32Func is the handler function for "echoI32" method of thrift service "SecondService".
type SecondServiceEchoI32Func func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI32_Args,
) (int32, map[string]string, error)

// NewSecondServiceEchoI32Handler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoI32Handler(f SecondServiceEchoI32Func) zanzibar.TChannelHandler {
	return &SecondServiceEchoI32Handler{f}
}

// SecondServiceEchoI32Handler handles the "echoI32" method call of thrift service "SecondService".
type SecondServiceEchoI32Handler struct {
	echoi32 SecondServiceEchoI32Func
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoI32Handler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoI32_Args
	var res clientsBazBaz.SecondService_EchoI32_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echoi32(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoI64Func is the handler function for "echoI64" method of thrift service "SecondService".
type SecondServiceEchoI64Func func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI64_Args,
) (int64, map[string]string, error)

// NewSecondServiceEchoI64Handler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoI64Handler(f SecondServiceEchoI64Func) zanzibar.TChannelHandler {
	return &SecondServiceEchoI64Handler{f}
}

// SecondServiceEchoI64Handler handles the "echoI64" method call of thrift service "SecondService".
type SecondServiceEchoI64Handler struct {
	echoi64 SecondServiceEchoI64Func
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoI64Handler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoI64_Args
	var res clientsBazBaz.SecondService_EchoI64_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echoi64(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoI8Func is the handler function for "echoI8" method of thrift service "SecondService".
type SecondServiceEchoI8Func func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI8_Args,
) (int8, map[string]string, error)

// NewSecondServiceEchoI8Handler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoI8Handler(f SecondServiceEchoI8Func) zanzibar.TChannelHandler {
	return &SecondServiceEchoI8Handler{f}
}

// SecondServiceEchoI8Handler handles the "echoI8" method call of thrift service "SecondService".
type SecondServiceEchoI8Handler struct {
	echoi8 SecondServiceEchoI8Func
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoI8Handler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoI8_Args
	var res clientsBazBaz.SecondService_EchoI8_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echoi8(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStringFunc is the handler function for "echoString" method of thrift service "SecondService".
type SecondServiceEchoStringFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoString_Args,
) (string, map[string]string, error)

// NewSecondServiceEchoStringHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStringHandler(f SecondServiceEchoStringFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStringHandler{f}
}

// SecondServiceEchoStringHandler handles the "echoString" method call of thrift service "SecondService".
type SecondServiceEchoStringHandler struct {
	echostring SecondServiceEchoStringFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStringHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoString_Args
	var res clientsBazBaz.SecondService_EchoString_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostring(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStringListFunc is the handler function for "echoStringList" method of thrift service "SecondService".
type SecondServiceEchoStringListFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringList_Args,
) ([]string, map[string]string, error)

// NewSecondServiceEchoStringListHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStringListHandler(f SecondServiceEchoStringListFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStringListHandler{f}
}

// SecondServiceEchoStringListHandler handles the "echoStringList" method call of thrift service "SecondService".
type SecondServiceEchoStringListHandler struct {
	echostringlist SecondServiceEchoStringListFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStringListHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoStringList_Args
	var res clientsBazBaz.SecondService_EchoStringList_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostringlist(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStringMapFunc is the handler function for "echoStringMap" method of thrift service "SecondService".
type SecondServiceEchoStringMapFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringMap_Args,
) (map[string]*clientsBazBase.BazResponse, map[string]string, error)

// NewSecondServiceEchoStringMapHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStringMapHandler(f SecondServiceEchoStringMapFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStringMapHandler{f}
}

// SecondServiceEchoStringMapHandler handles the "echoStringMap" method call of thrift service "SecondService".
type SecondServiceEchoStringMapHandler struct {
	echostringmap SecondServiceEchoStringMapFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStringMapHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoStringMap_Args
	var res clientsBazBaz.SecondService_EchoStringMap_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostringmap(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStringSetFunc is the handler function for "echoStringSet" method of thrift service "SecondService".
type SecondServiceEchoStringSetFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringSet_Args,
) (map[string]struct{}, map[string]string, error)

// NewSecondServiceEchoStringSetHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStringSetHandler(f SecondServiceEchoStringSetFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStringSetHandler{f}
}

// SecondServiceEchoStringSetHandler handles the "echoStringSet" method call of thrift service "SecondService".
type SecondServiceEchoStringSetHandler struct {
	echostringset SecondServiceEchoStringSetFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStringSetHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoStringSet_Args
	var res clientsBazBaz.SecondService_EchoStringSet_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostringset(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStructListFunc is the handler function for "echoStructList" method of thrift service "SecondService".
type SecondServiceEchoStructListFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStructList_Args,
) ([]*clientsBazBase.BazResponse, map[string]string, error)

// NewSecondServiceEchoStructListHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStructListHandler(f SecondServiceEchoStructListFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStructListHandler{f}
}

// SecondServiceEchoStructListHandler handles the "echoStructList" method call of thrift service "SecondService".
type SecondServiceEchoStructListHandler struct {
	echostructlist SecondServiceEchoStructListFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStructListHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoStructList_Args
	var res clientsBazBaz.SecondService_EchoStructList_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostructlist(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoStructSetFunc is the handler function for "echoStructSet" method of thrift service "SecondService".
type SecondServiceEchoStructSetFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStructSet_Args,
) ([]*clientsBazBase.BazResponse, map[string]string, error)

// NewSecondServiceEchoStructSetHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoStructSetHandler(f SecondServiceEchoStructSetFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoStructSetHandler{f}
}

// SecondServiceEchoStructSetHandler handles the "echoStructSet" method call of thrift service "SecondService".
type SecondServiceEchoStructSetHandler struct {
	echostructset SecondServiceEchoStructSetFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoStructSetHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoStructSet_Args
	var res clientsBazBaz.SecondService_EchoStructSet_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostructset(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SecondServiceEchoTypedefFunc is the handler function for "echoTypedef" method of thrift service "SecondService".
type SecondServiceEchoTypedefFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoTypedef_Args,
) (clientsBazBase.UUID, map[string]string, error)

// NewSecondServiceEchoTypedefHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoTypedefHandler(f SecondServiceEchoTypedefFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoTypedefHandler{f}
}

// SecondServiceEchoTypedefHandler handles the "echoTypedef" method call of thrift service "SecondService".
type SecondServiceEchoTypedefHandler struct {
	echotypedef SecondServiceEchoTypedefFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoTypedefHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_EchoTypedef_Args
	var res clientsBazBaz.SecondService_EchoTypedef_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echotypedef(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceCallFunc is the handler function for "call" method of thrift service "SimpleService".
type SimpleServiceCallFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Call_Args,
) (map[string]string, error)

// NewSimpleServiceCallHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceCallHandler(f SimpleServiceCallFunc) zanzibar.TChannelHandler {
	return &SimpleServiceCallHandler{f}
}

// SimpleServiceCallHandler handles the "call" method call of thrift service "SimpleService".
type SimpleServiceCallHandler struct {
	call SimpleServiceCallFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceCallHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Call_Args
	var res clientsBazBaz.SimpleService_Call_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.call(ctx, reqHeaders, &req)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for call returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		default:
			return false, nil, nil, err
		}
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceCompareFunc is the handler function for "compare" method of thrift service "SimpleService".
type SimpleServiceCompareFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Compare_Args,
) (*clientsBazBase.BazResponse, map[string]string, error)

// NewSimpleServiceCompareHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceCompareHandler(f SimpleServiceCompareFunc) zanzibar.TChannelHandler {
	return &SimpleServiceCompareHandler{f}
}

// SimpleServiceCompareHandler handles the "compare" method call of thrift service "SimpleService".
type SimpleServiceCompareHandler struct {
	compare SimpleServiceCompareFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceCompareHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Compare_Args
	var res clientsBazBaz.SimpleService_Compare_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.compare(ctx, reqHeaders, &req)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for compare returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		case *clientsBazBaz.OtherAuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for compare returned non-nil error type *OtherAuthErr but nil value",
				)
			}
			res.OtherAuthErr = v
		default:
			return false, nil, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServicePingFunc is the handler function for "ping" method of thrift service "SimpleService".
type SimpleServicePingFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (*clientsBazBase.BazResponse, map[string]string, error)

// NewSimpleServicePingHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServicePingHandler(f SimpleServicePingFunc) zanzibar.TChannelHandler {
	return &SimpleServicePingHandler{f}
}

// SimpleServicePingHandler handles the "ping" method call of thrift service "SimpleService".
type SimpleServicePingHandler struct {
	ping SimpleServicePingFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServicePingHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Ping_Args
	var res clientsBazBaz.SimpleService_Ping_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.ping(ctx, reqHeaders)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceSillyNoopFunc is the handler function for "sillyNoop" method of thrift service "SimpleService".
type SimpleServiceSillyNoopFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error)

// NewSimpleServiceSillyNoopHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceSillyNoopHandler(f SimpleServiceSillyNoopFunc) zanzibar.TChannelHandler {
	return &SimpleServiceSillyNoopHandler{f}
}

// SimpleServiceSillyNoopHandler handles the "sillyNoop" method call of thrift service "SimpleService".
type SimpleServiceSillyNoopHandler struct {
	sillynoop SimpleServiceSillyNoopFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceSillyNoopHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_SillyNoop_Args
	var res clientsBazBaz.SimpleService_SillyNoop_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.sillynoop(ctx, reqHeaders)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for sillyNoop returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		case *clientsBazBase.ServerErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for sillyNoop returned non-nil error type *ServerErr but nil value",
				)
			}
			res.ServerErr = v
		default:
			return false, nil, nil, err
		}
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceTestUUIDFunc is the handler function for "testUuid" method of thrift service "SimpleService".
type SimpleServiceTestUUIDFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error)

// NewSimpleServiceTestUUIDHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceTestUUIDHandler(f SimpleServiceTestUUIDFunc) zanzibar.TChannelHandler {
	return &SimpleServiceTestUUIDHandler{f}
}

// SimpleServiceTestUUIDHandler handles the "testUuid" method call of thrift service "SimpleService".
type SimpleServiceTestUUIDHandler struct {
	testuuid SimpleServiceTestUUIDFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceTestUUIDHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_TestUuid_Args
	var res clientsBazBaz.SimpleService_TestUuid_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.testuuid(ctx, reqHeaders)

	if err != nil {
		return false, nil, nil, err
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceTransFunc is the handler function for "trans" method of thrift service "SimpleService".
type SimpleServiceTransFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Trans_Args,
) (*clientsBazBase.TransStruct, map[string]string, error)

// NewSimpleServiceTransHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceTransHandler(f SimpleServiceTransFunc) zanzibar.TChannelHandler {
	return &SimpleServiceTransHandler{f}
}

// SimpleServiceTransHandler handles the "trans" method call of thrift service "SimpleService".
type SimpleServiceTransHandler struct {
	trans SimpleServiceTransFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceTransHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Trans_Args
	var res clientsBazBaz.SimpleService_Trans_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.trans(ctx, reqHeaders, &req)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for trans returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		case *clientsBazBaz.OtherAuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for trans returned non-nil error type *OtherAuthErr but nil value",
				)
			}
			res.OtherAuthErr = v
		default:
			return false, nil, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceURLTestFunc is the handler function for "urlTest" method of thrift service "SimpleService".
type SimpleServiceURLTestFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error)

// NewSimpleServiceURLTestHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceURLTestHandler(f SimpleServiceURLTestFunc) zanzibar.TChannelHandler {
	return &SimpleServiceURLTestHandler{f}
}

// SimpleServiceURLTestHandler handles the "urlTest" method call of thrift service "SimpleService".
type SimpleServiceURLTestHandler struct {
	urltest SimpleServiceURLTestFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceURLTestHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_UrlTest_Args
	var res clientsBazBaz.SimpleService_UrlTest_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.urltest(ctx, reqHeaders)

	if err != nil {
		return false, nil, nil, err
	}

	return err == nil, &res, respHeaders, nil
}
