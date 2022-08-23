// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto

package envoy_config_filter_network_thrift_proxy_v2alpha1

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/v2/pkg/api/envoy/api/v2/core"
	route "github.com/datawire/ambassador/v2/pkg/api/envoy/api/v2/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetMatch() *RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Route) GetRoute() *RouteAction {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Inverts whatever matching is done in the :ref:`method_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.method_name>` or
	// :ref:`service_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.service_name>` fields.
	// Cannot be combined with wildcard matching as that would result in routes never being matched.
	//
	// .. note::
	//
	//   This does not invert matching done as part of the :ref:`headers field
	//   <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.headers>` field. To
	//   invert header matching, see :ref:`invert_match
	//   <envoy_api_field_route.HeaderMatcher.invert_match>`.
	Invert bool `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config). Note that this only applies for Thrift transports and/or
	// protocols that support headers.
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{2}
}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (x *RouteMatch) GetMethodName() string {
	if x, ok := x.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (x *RouteMatch) GetServiceName() string {
	if x, ok := x.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (x *RouteMatch) GetInvert() bool {
	if x != nil {
		return x.Invert
	}
	return false
}

func (x *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	// If specified, the route must exactly match the request method name. As a special case, an
	// empty string matches any request method name.
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	// If specified, the route must have the service name as the request method name prefix. As a
	// special case, an empty string matches any service name. Only relevant when service
	// multiplexing.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

// [#next-free-field: 7]
type RouteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field will be considered.
	// Note that this will be merged with what's provided in :ref:`WeightedCluster.metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.metadata_match>`,
	// with values there taking precedence. Keys and values should be provided under the "envoy.lb"
	// metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// Specifies a set of rate limit configurations that could be applied to the route.
	// N.B. Thrift service or method name matching can be achieved by specifying a RequestHeaders
	// action with the header name ":method-name".
	RateLimits []*route.RateLimit `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Strip the service prefix from the method name, if there's a prefix. For
	// example, the method call Service:method would end up being just method.
	StripServiceName bool `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
}

func (x *RouteAction) Reset() {
	*x = RouteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAction) ProtoMessage() {}

func (x *RouteAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAction.ProtoReflect.Descriptor instead.
func (*RouteAction) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{3}
}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (x *RouteAction) GetCluster() string {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (x *RouteAction) GetClusterHeader() string {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (x *RouteAction) GetMetadataMatch() *core.Metadata {
	if x != nil {
		return x.MetadataMatch
	}
	return nil
}

func (x *RouteAction) GetRateLimits() []*route.RateLimit {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

func (x *RouteAction) GetStripServiceName() bool {
	if x != nil {
		return x.StripServiceName
	}
	return false
}

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	// Indicates a single upstream cluster to which the request should be routed
	// to.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	// Multiple upstream clusters can be specified for a given route. The
	// request is routed to one of the upstream clusters based on weights
	// assigned to each cluster.
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	// Envoy will determine the cluster to route to by reading the value of the
	// Thrift header named by cluster_header from the request headers. If the
	// header is not found or the referenced cluster does not exist Envoy will
	// respond with an unknown method exception or an internal error exception,
	// respectively.
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

// Allows for specification of multiple upstream clusters along with weights that indicate the
// percentage of traffic to be forwarded to each cluster. The router selects an upstream cluster
// based on these weights.
type WeightedCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies one or more upstream clusters associated with the route.
	Clusters []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *WeightedCluster) Reset() {
	*x = WeightedCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedCluster) ProtoMessage() {}

func (x *WeightedCluster) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightedCluster.ProtoReflect.Descriptor instead.
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{4}
}

func (x *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is determined by its
	// weight. The sum of weights across all entries in the clusters array determines the total
	// weight.
	Weight *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field, combined with what's
	// provided in :ref:`RouteAction's metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteAction.metadata_match>`,
	// will be considered. Values here will take precedence. Keys and values should be provided
	// under the "envoy.lb" metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
}

func (x *WeightedCluster_ClusterWeight) Reset() {
	*x = WeightedCluster_ClusterWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedCluster_ClusterWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedCluster_ClusterWeight) ProtoMessage() {}

func (x *WeightedCluster_ClusterWeight) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightedCluster_ClusterWeight.ProtoReflect.Descriptor instead.
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP(), []int{4, 0}
}

func (x *WeightedCluster_ClusterWeight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if x != nil {
		return x.MetadataMatch
	}
	return nil
}

var File_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72,
	0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22,
	0xc6, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x5d, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x5e, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x21, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x42, 0x16, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xa3, 0x03, 0x0a,
	0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x71, 0x0a, 0x11, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x10, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3e, 0x0a, 0x0b, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0a,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x18, 0x0a, 0x11, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x22, 0xbb, 0x02, 0x0a, 0x0f, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x76, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xaf,
	0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x42, 0x0a, 0x0e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x42, 0x8f, 0x01, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74,
	0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x32, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescData = file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDesc
)

func file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescData)
	})
	return file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDescData
}

var file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil),            // 0: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration
	(*Route)(nil),                         // 1: envoy.config.filter.network.thrift_proxy.v2alpha1.Route
	(*RouteMatch)(nil),                    // 2: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch
	(*RouteAction)(nil),                   // 3: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction
	(*WeightedCluster)(nil),               // 4: envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster
	(*WeightedCluster_ClusterWeight)(nil), // 5: envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight
	(*route.HeaderMatcher)(nil),           // 6: envoy.api.v2.route.HeaderMatcher
	(*core.Metadata)(nil),                 // 7: envoy.api.v2.core.Metadata
	(*route.RateLimit)(nil),               // 8: envoy.api.v2.route.RateLimit
	(*wrappers.UInt32Value)(nil),          // 9: google.protobuf.UInt32Value
}
var file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_depIdxs = []int32{
	1,  // 0: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration.routes:type_name -> envoy.config.filter.network.thrift_proxy.v2alpha1.Route
	2,  // 1: envoy.config.filter.network.thrift_proxy.v2alpha1.Route.match:type_name -> envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch
	3,  // 2: envoy.config.filter.network.thrift_proxy.v2alpha1.Route.route:type_name -> envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction
	6,  // 3: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch.headers:type_name -> envoy.api.v2.route.HeaderMatcher
	4,  // 4: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction.weighted_clusters:type_name -> envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster
	7,  // 5: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction.metadata_match:type_name -> envoy.api.v2.core.Metadata
	8,  // 6: envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction.rate_limits:type_name -> envoy.api.v2.route.RateLimit
	5,  // 7: envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.clusters:type_name -> envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight
	9,  // 8: envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.weight:type_name -> google.protobuf.UInt32Value
	7,  // 9: envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.metadata_match:type_name -> envoy.api.v2.core.Metadata
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_init() }
func file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_init() {
	if File_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedCluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedCluster_ClusterWeight); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto = out.File
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_rawDesc = nil
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_goTypes = nil
	file_envoy_config_filter_network_thrift_proxy_v2alpha1_route_proto_depIdxs = nil
}
