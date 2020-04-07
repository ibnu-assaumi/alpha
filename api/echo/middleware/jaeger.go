package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Jaeger middleware to wrap from http inbound (request from client)
func Jaeger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tracer := opentracing.GlobalTracer()
		operationName := fmt.Sprintf("%s %s%s", req.Method, req.Host, req.URL.Path)

		var span opentracing.Span
		var ctx context.Context
		if spanCtx, err := opentracing.Tracer.Extract(tracer, opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
			span, ctx = opentracing.StartSpanFromContext(req.Context(), operationName)
			ext.SpanKindRPCServer.Set(span)
		} else {
			span = tracer.StartSpan(operationName, ext.RPCServerOption((spanCtx)))
			ctx = opentracing.ContextWithSpan(req.Context(), span)
			ext.SpanKindRPCClient.Set(span)
		}

		body, _ := ioutil.ReadAll(req.Body)
		bodyString := string(body)
		span.SetTag("request.body", bodyString)
		span.SetTag("request.cookies", req.Cookies())
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // reuse body

		header, _ := json.Marshal(req.Header)
		span.SetTag("request.headers", string(header))
		ext.HTTPUrl.Set(span, req.Host+req.RequestURI)
		ext.HTTPMethod.Set(span, req.Method)

		span.LogEvent("start_handling_request")

		defer func() {
			span.LogEvent("complete_handling_request")
		}()

		h.ServeHTTP(w, req.WithContext(ctx))
	})
}
