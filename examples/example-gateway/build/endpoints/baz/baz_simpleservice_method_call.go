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

package bazendpoint

import (
	"context"
	"fmt"
	"runtime/debug"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/workflow"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SimpleServiceCallHandler is the handler for "/baz/call"
type SimpleServiceCallHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewSimpleServiceCallHandler creates a handler
func NewSimpleServiceCallHandler(deps *module.Dependencies) *SimpleServiceCallHandler {
	handler := &SimpleServiceCallHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpointContext(
		deps.Default.ContextExtractor, deps.Default.ContextMetrics, deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"baz", "call",
		handler.HandleRequest,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *SimpleServiceCallHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"POST", "/baz/call",
		h.endpoint,
	)
}

// HandleRequest handles "/baz/call".
func (h *SimpleServiceCallHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	defer func() {
		if r := recover(); r != nil {
			stacktrace := string(debug.Stack())
			e := errors.Errorf("enpoint panic: %v, stacktrace: %v", r, stacktrace)
			h.Dependencies.Default.ContextLogger.Error(
				ctx,
				"endpoint panic",
				zap.Error(e),
				zap.String("stacktrace", stacktrace),
				zap.String("endpoint", h.endpoint.EndpointName))

			h.endpoint.ContextMetrics.GetOrAddEndpointMetrics(ctx).Panic.Inc(1)
			res.SendError(502, "Unexpected workflow panic, recovered at endpoint.", nil)
		}
	}()

	var requestBody endpointsBazBaz.SimpleService_Call_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	if requestBody.Arg == nil {
		requestBody.Arg = &endpointsBazBaz.BazRequest{}
	}
	xTokenValue, xTokenValueExists := req.Header.Get("x-token")
	if xTokenValueExists {
		body, _ := strconv.ParseInt(xTokenValue, 10, 64)
		requestBody.I64Optional = &body
	}
	xUUIDValue, xUUIDValueExists := req.Header.Get("x-uuid")
	if xUUIDValueExists {
		body := endpointsBazBaz.UUID(xUUIDValue)
		requestBody.TestUUID = &body
	}

	// log endpoint request to downstream services
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", req.GetRawBody())))
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		h.Dependencies.Default.ContextLogger.Debug(ctx, "endpoint request to downstream", zfields...)
	}

	w := workflow.NewSimpleServiceCallWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	cliRespHeaders, err := w.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}

	res.WriteJSONBytes(204, cliRespHeaders, nil)
}
