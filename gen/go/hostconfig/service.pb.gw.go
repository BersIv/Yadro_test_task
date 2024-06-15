// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/service.proto

/*
Package  hostconfig is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package hostconfig

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

func request_HostConfig_ChangeHostname_0(ctx context.Context, marshaler runtime.Marshaler, client HostConfigClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChangeHostnameRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ChangeHostname(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_HostConfig_ChangeHostname_0(ctx context.Context, marshaler runtime.Marshaler, server HostConfigServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChangeHostnameRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ChangeHostname(ctx, &protoReq)
	return msg, metadata, err

}

func request_HostConfig_ListDNSServers_0(ctx context.Context, marshaler runtime.Marshaler, client HostConfigClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListDNSServersRequest
	var metadata runtime.ServerMetadata

	msg, err := client.ListDNSServers(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_HostConfig_ListDNSServers_0(ctx context.Context, marshaler runtime.Marshaler, server HostConfigServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListDNSServersRequest
	var metadata runtime.ServerMetadata

	msg, err := server.ListDNSServers(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterHostConfigHandlerServer registers the http handlers for service HostConfig to "mux".
// UnaryRPC     :call HostConfigServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterHostConfigHandlerFromEndpoint instead.
func RegisterHostConfigHandlerServer(ctx context.Context, mux *runtime.ServeMux, server HostConfigServer) error {

	mux.Handle("POST", pattern_HostConfig_ChangeHostname_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/hostconfig.HostConfig/ChangeHostname", runtime.WithHTTPPathPattern("/hostname"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_HostConfig_ChangeHostname_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_HostConfig_ChangeHostname_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_HostConfig_ListDNSServers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/hostconfig.HostConfig/ListDNSServers", runtime.WithHTTPPathPattern("/dns"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_HostConfig_ListDNSServers_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_HostConfig_ListDNSServers_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterHostConfigHandlerFromEndpoint is same as RegisterHostConfigHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterHostConfigHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterHostConfigHandler(ctx, mux, conn)
}

// RegisterHostConfigHandler registers the http handlers for service HostConfig to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterHostConfigHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterHostConfigHandlerClient(ctx, mux, NewHostConfigClient(conn))
}

// RegisterHostConfigHandlerClient registers the http handlers for service HostConfig
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "HostConfigClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "HostConfigClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "HostConfigClient" to call the correct interceptors.
func RegisterHostConfigHandlerClient(ctx context.Context, mux *runtime.ServeMux, client HostConfigClient) error {

	mux.Handle("POST", pattern_HostConfig_ChangeHostname_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/hostconfig.HostConfig/ChangeHostname", runtime.WithHTTPPathPattern("/hostname"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_HostConfig_ChangeHostname_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_HostConfig_ChangeHostname_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_HostConfig_ListDNSServers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/hostconfig.HostConfig/ListDNSServers", runtime.WithHTTPPathPattern("/dns"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_HostConfig_ListDNSServers_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_HostConfig_ListDNSServers_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_HostConfig_ChangeHostname_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"hostname"}, ""))

	pattern_HostConfig_ListDNSServers_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"dns"}, ""))
)

var (
	forward_HostConfig_ChangeHostname_0 = runtime.ForwardResponseMessage

	forward_HostConfig_ListDNSServers_0 = runtime.ForwardResponseMessage
)
