// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: app/gen-code/api/api.proto

/*
Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package api

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_GenCodeService_GenReactSourceCode_0(ctx context.Context, marshaler runtime.Marshaler, client GenCodeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenReactSourceCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GenReactSourceCode(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GenCodeService_GenReactSourceCode_0(ctx context.Context, marshaler runtime.Marshaler, server GenCodeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenReactSourceCodeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GenReactSourceCode(ctx, &protoReq)
	return msg, metadata, err

}

func request_GenCodeService_HelloWorld_0(ctx context.Context, marshaler runtime.Marshaler, client GenCodeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EmptyRequest
	var metadata runtime.ServerMetadata

	msg, err := client.HelloWorld(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GenCodeService_HelloWorld_0(ctx context.Context, marshaler runtime.Marshaler, server GenCodeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EmptyRequest
	var metadata runtime.ServerMetadata

	msg, err := server.HelloWorld(ctx, &protoReq)
	return msg, metadata, err

}

func request_GenCodeService_HealthCheck_0(ctx context.Context, marshaler runtime.Marshaler, client GenCodeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EmptyRequest
	var metadata runtime.ServerMetadata

	msg, err := client.HealthCheck(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GenCodeService_HealthCheck_0(ctx context.Context, marshaler runtime.Marshaler, server GenCodeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EmptyRequest
	var metadata runtime.ServerMetadata

	msg, err := server.HealthCheck(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterGenCodeServiceHandlerServer registers the http handlers for service GenCodeService to "mux".
// UnaryRPC     :call GenCodeServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGenCodeServiceHandlerFromEndpoint instead.
func RegisterGenCodeServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GenCodeServiceServer) error {

	mux.Handle("POST", pattern_GenCodeService_GenReactSourceCode_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/GenReactSourceCode", runtime.WithHTTPPathPattern("/gen-code-service/api/gen-react-code"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenCodeService_GenReactSourceCode_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_GenReactSourceCode_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GenCodeService_HelloWorld_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/HelloWorld", runtime.WithHTTPPathPattern("/gen-code-service/api/hello-world"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenCodeService_HelloWorld_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_HelloWorld_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GenCodeService_HealthCheck_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/HealthCheck", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenCodeService_HealthCheck_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_HealthCheck_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterGenCodeServiceHandlerFromEndpoint is same as RegisterGenCodeServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGenCodeServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGenCodeServiceHandler(ctx, mux, conn)
}

// RegisterGenCodeServiceHandler registers the http handlers for service GenCodeService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGenCodeServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGenCodeServiceHandlerClient(ctx, mux, NewGenCodeServiceClient(conn))
}

// RegisterGenCodeServiceHandlerClient registers the http handlers for service GenCodeService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GenCodeServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GenCodeServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GenCodeServiceClient" to call the correct interceptors.
func RegisterGenCodeServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GenCodeServiceClient) error {

	mux.Handle("POST", pattern_GenCodeService_GenReactSourceCode_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/GenReactSourceCode", runtime.WithHTTPPathPattern("/gen-code-service/api/gen-react-code"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenCodeService_GenReactSourceCode_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_GenReactSourceCode_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GenCodeService_HelloWorld_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/HelloWorld", runtime.WithHTTPPathPattern("/gen-code-service/api/hello-world"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenCodeService_HelloWorld_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_HelloWorld_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GenCodeService_HealthCheck_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/app.gen_code.api.GenCodeService/HealthCheck", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenCodeService_HealthCheck_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GenCodeService_HealthCheck_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_GenCodeService_GenReactSourceCode_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"gen-code-service", "api", "gen-react-code"}, ""))

	pattern_GenCodeService_HelloWorld_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"gen-code-service", "api", "hello-world"}, ""))

	pattern_GenCodeService_HealthCheck_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{""}, ""))
)

var (
	forward_GenCodeService_GenReactSourceCode_0 = runtime.ForwardResponseMessage

	forward_GenCodeService_HelloWorld_0 = runtime.ForwardResponseMessage

	forward_GenCodeService_HealthCheck_0 = runtime.ForwardResponseMessage
)
