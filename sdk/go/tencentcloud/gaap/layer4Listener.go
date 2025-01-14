// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Layer4Listener struct {
	pulumi.CustomResourceState

	// UDP origin station health check probe port.
	CheckPort pulumi.IntOutput `pulumi:"checkPort"`
	// UDP origin server health type. PORT means check port, and PING means PING.
	CheckType pulumi.StringOutput `pulumi:"checkType"`
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
	// listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrOutput `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
	// `interval`.
	ConnectTimeout pulumi.IntPtrOutput `pulumi:"connectTimeout"`
	// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
	// PORT.
	ContextType pulumi.StringOutput `pulumi:"contextType"`
	// Creation time of the layer4 listener.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether health check is enable, default value is `false`.
	HealthCheck pulumi.BoolPtrOutput `pulumi:"healthCheck"`
	// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
	// be healthy. Range from 1 to 10. Default value is 1.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port of the layer4 listener.
	Port pulumi.IntOutput `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId pulumi.StringOutput `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayOutput `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
	// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringOutput `pulumi:"realserverType"`
	// UDP source server health check port detects received messages. Only used when the health check type is PORT.
	RecvContext pulumi.StringOutput `pulumi:"recvContext"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrOutput `pulumi:"scheduler"`
	// UDP source server health check port detection sends messages. Only used when health check type is PORT.
	SendContext pulumi.StringOutput `pulumi:"sendContext"`
	// Status of the layer4 listener.
	Status pulumi.IntOutput `pulumi:"status"`
	// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
	// Range from 1 to 10. Default value is 1.
	UnhealthyThreshold pulumi.IntPtrOutput `pulumi:"unhealthyThreshold"`
}

// NewLayer4Listener registers a new resource with the given unique name, arguments, and options.
func NewLayer4Listener(ctx *pulumi.Context,
	name string, args *Layer4ListenerArgs, opts ...pulumi.ResourceOption) (*Layer4Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProxyId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyId'")
	}
	if args.RealserverType == nil {
		return nil, errors.New("invalid value for required argument 'RealserverType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Layer4Listener
	err := ctx.RegisterResource("tencentcloud:Gaap/layer4Listener:Layer4Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayer4Listener gets an existing Layer4Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayer4Listener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Layer4ListenerState, opts ...pulumi.ResourceOption) (*Layer4Listener, error) {
	var resource Layer4Listener
	err := ctx.ReadResource("tencentcloud:Gaap/layer4Listener:Layer4Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Layer4Listener resources.
type layer4ListenerState struct {
	// UDP origin station health check probe port.
	CheckPort *int `pulumi:"checkPort"`
	// UDP origin server health type. PORT means check port, and PING means PING.
	CheckType *string `pulumi:"checkType"`
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
	// listeners of `TCP` protocol.
	ClientIpMethod *int `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
	// `interval`.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
	// PORT.
	ContextType *string `pulumi:"contextType"`
	// Creation time of the layer4 listener.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether health check is enable, default value is `false`.
	HealthCheck *bool `pulumi:"healthCheck"`
	// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
	// be healthy. Range from 1 to 10. Default value is 1.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// Interval of the health check, default value is 5s.
	Interval *int `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Port of the layer4 listener.
	Port *int `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol *string `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId *string `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets []Layer4ListenerRealserverBindSet `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
	// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType *string `pulumi:"realserverType"`
	// UDP source server health check port detects received messages. Only used when the health check type is PORT.
	RecvContext *string `pulumi:"recvContext"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
	// UDP source server health check port detection sends messages. Only used when health check type is PORT.
	SendContext *string `pulumi:"sendContext"`
	// Status of the layer4 listener.
	Status *int `pulumi:"status"`
	// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
	// Range from 1 to 10. Default value is 1.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

type Layer4ListenerState struct {
	// UDP origin station health check probe port.
	CheckPort pulumi.IntPtrInput
	// UDP origin server health type. PORT means check port, and PING means PING.
	CheckType pulumi.StringPtrInput
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
	// listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrInput
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
	// `interval`.
	ConnectTimeout pulumi.IntPtrInput
	// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
	// PORT.
	ContextType pulumi.StringPtrInput
	// Creation time of the layer4 listener.
	CreateTime pulumi.StringPtrInput
	// Indicates whether health check is enable, default value is `false`.
	HealthCheck pulumi.BoolPtrInput
	// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
	// be healthy. Range from 1 to 10. Default value is 1.
	HealthyThreshold pulumi.IntPtrInput
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrInput
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Port of the layer4 listener.
	Port pulumi.IntPtrInput
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringPtrInput
	// ID of the GAAP proxy.
	ProxyId pulumi.StringPtrInput
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
	// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringPtrInput
	// UDP source server health check port detects received messages. Only used when the health check type is PORT.
	RecvContext pulumi.StringPtrInput
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
	// UDP source server health check port detection sends messages. Only used when health check type is PORT.
	SendContext pulumi.StringPtrInput
	// Status of the layer4 listener.
	Status pulumi.IntPtrInput
	// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
	// Range from 1 to 10. Default value is 1.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (Layer4ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*layer4ListenerState)(nil)).Elem()
}

type layer4ListenerArgs struct {
	// UDP origin station health check probe port.
	CheckPort *int `pulumi:"checkPort"`
	// UDP origin server health type. PORT means check port, and PING means PING.
	CheckType *string `pulumi:"checkType"`
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
	// listeners of `TCP` protocol.
	ClientIpMethod *int `pulumi:"clientIpMethod"`
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
	// `interval`.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
	// PORT.
	ContextType *string `pulumi:"contextType"`
	// Indicates whether health check is enable, default value is `false`.
	HealthCheck *bool `pulumi:"healthCheck"`
	// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
	// be healthy. Range from 1 to 10. Default value is 1.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// Interval of the health check, default value is 5s.
	Interval *int `pulumi:"interval"`
	// Name of the layer4 listener, the maximum length is 30.
	Name *string `pulumi:"name"`
	// Port of the layer4 listener.
	Port int `pulumi:"port"`
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol string `pulumi:"protocol"`
	// ID of the GAAP proxy.
	ProxyId string `pulumi:"proxyId"`
	// An information list of GAAP realserver.
	RealserverBindSets []Layer4ListenerRealserverBindSet `pulumi:"realserverBindSets"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
	// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType string `pulumi:"realserverType"`
	// UDP source server health check port detects received messages. Only used when the health check type is PORT.
	RecvContext *string `pulumi:"recvContext"`
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
	// UDP source server health check port detection sends messages. Only used when health check type is PORT.
	SendContext *string `pulumi:"sendContext"`
	// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
	// Range from 1 to 10. Default value is 1.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a Layer4Listener resource.
type Layer4ListenerArgs struct {
	// UDP origin station health check probe port.
	CheckPort pulumi.IntPtrInput
	// UDP origin server health type. PORT means check port, and PING means PING.
	CheckType pulumi.StringPtrInput
	// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
	// listeners of `TCP` protocol.
	ClientIpMethod pulumi.IntPtrInput
	// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
	// `interval`.
	ConnectTimeout pulumi.IntPtrInput
	// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
	// PORT.
	ContextType pulumi.StringPtrInput
	// Indicates whether health check is enable, default value is `false`.
	HealthCheck pulumi.BoolPtrInput
	// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
	// be healthy. Range from 1 to 10. Default value is 1.
	HealthyThreshold pulumi.IntPtrInput
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrInput
	// Name of the layer4 listener, the maximum length is 30.
	Name pulumi.StringPtrInput
	// Port of the layer4 listener.
	Port pulumi.IntInput
	// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
	Protocol pulumi.StringInput
	// ID of the GAAP proxy.
	ProxyId pulumi.StringInput
	// An information list of GAAP realserver.
	RealserverBindSets Layer4ListenerRealserverBindSetArrayInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
	// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
	RealserverType pulumi.StringInput
	// UDP source server health check port detects received messages. Only used when the health check type is PORT.
	RecvContext pulumi.StringPtrInput
	// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
	// UDP source server health check port detection sends messages. Only used when health check type is PORT.
	SendContext pulumi.StringPtrInput
	// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
	// Range from 1 to 10. Default value is 1.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (Layer4ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layer4ListenerArgs)(nil)).Elem()
}

type Layer4ListenerInput interface {
	pulumi.Input

	ToLayer4ListenerOutput() Layer4ListenerOutput
	ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput
}

func (*Layer4Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Layer4Listener)(nil)).Elem()
}

func (i *Layer4Listener) ToLayer4ListenerOutput() Layer4ListenerOutput {
	return i.ToLayer4ListenerOutputWithContext(context.Background())
}

func (i *Layer4Listener) ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerOutput)
}

// Layer4ListenerArrayInput is an input type that accepts Layer4ListenerArray and Layer4ListenerArrayOutput values.
// You can construct a concrete instance of `Layer4ListenerArrayInput` via:
//
//	Layer4ListenerArray{ Layer4ListenerArgs{...} }
type Layer4ListenerArrayInput interface {
	pulumi.Input

	ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput
	ToLayer4ListenerArrayOutputWithContext(context.Context) Layer4ListenerArrayOutput
}

type Layer4ListenerArray []Layer4ListenerInput

func (Layer4ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Layer4Listener)(nil)).Elem()
}

func (i Layer4ListenerArray) ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput {
	return i.ToLayer4ListenerArrayOutputWithContext(context.Background())
}

func (i Layer4ListenerArray) ToLayer4ListenerArrayOutputWithContext(ctx context.Context) Layer4ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerArrayOutput)
}

// Layer4ListenerMapInput is an input type that accepts Layer4ListenerMap and Layer4ListenerMapOutput values.
// You can construct a concrete instance of `Layer4ListenerMapInput` via:
//
//	Layer4ListenerMap{ "key": Layer4ListenerArgs{...} }
type Layer4ListenerMapInput interface {
	pulumi.Input

	ToLayer4ListenerMapOutput() Layer4ListenerMapOutput
	ToLayer4ListenerMapOutputWithContext(context.Context) Layer4ListenerMapOutput
}

type Layer4ListenerMap map[string]Layer4ListenerInput

func (Layer4ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Layer4Listener)(nil)).Elem()
}

func (i Layer4ListenerMap) ToLayer4ListenerMapOutput() Layer4ListenerMapOutput {
	return i.ToLayer4ListenerMapOutputWithContext(context.Background())
}

func (i Layer4ListenerMap) ToLayer4ListenerMapOutputWithContext(ctx context.Context) Layer4ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4ListenerMapOutput)
}

type Layer4ListenerOutput struct{ *pulumi.OutputState }

func (Layer4ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerOutput) ToLayer4ListenerOutput() Layer4ListenerOutput {
	return o
}

func (o Layer4ListenerOutput) ToLayer4ListenerOutputWithContext(ctx context.Context) Layer4ListenerOutput {
	return o
}

// UDP origin station health check probe port.
func (o Layer4ListenerOutput) CheckPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntOutput { return v.CheckPort }).(pulumi.IntOutput)
}

// UDP origin server health type. PORT means check port, and PING means PING.
func (o Layer4ListenerOutput) CheckType() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.CheckType }).(pulumi.StringOutput)
}

// The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports
// listeners of `TCP` protocol.
func (o Layer4ListenerOutput) ClientIpMethod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.ClientIpMethod }).(pulumi.IntPtrOutput)
}

// Timeout of the health check response, should less than interval, default value is 2s. NOTES: Require less than
// `interval`.
func (o Layer4ListenerOutput) ConnectTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.ConnectTimeout }).(pulumi.IntPtrOutput)
}

// UDP source station health check port probe message type: TEXT represents text. Only used when the health check type is
// PORT.
func (o Layer4ListenerOutput) ContextType() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.ContextType }).(pulumi.StringOutput)
}

// Creation time of the layer4 listener.
func (o Layer4ListenerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether health check is enable, default value is `false`.
func (o Layer4ListenerOutput) HealthCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.BoolPtrOutput { return v.HealthCheck }).(pulumi.BoolPtrOutput)
}

// Health threshold, which indicates how many consecutive inspections are successful, the source station is determined to
// be healthy. Range from 1 to 10. Default value is 1.
func (o Layer4ListenerOutput) HealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.HealthyThreshold }).(pulumi.IntPtrOutput)
}

// Interval of the health check, default value is 5s.
func (o Layer4ListenerOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Name of the layer4 listener, the maximum length is 30.
func (o Layer4ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port of the layer4 listener.
func (o Layer4ListenerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Protocol of the layer4 listener. Valid value: `TCP` and `UDP`.
func (o Layer4ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// ID of the GAAP proxy.
func (o Layer4ListenerOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.ProxyId }).(pulumi.StringOutput)
}

// An information list of GAAP realserver.
func (o Layer4ListenerOutput) RealserverBindSets() Layer4ListenerRealserverBindSetArrayOutput {
	return o.ApplyT(func(v *Layer4Listener) Layer4ListenerRealserverBindSetArrayOutput { return v.RealserverBindSets }).(Layer4ListenerRealserverBindSetArrayOutput)
}

// Type of the realserver. Valid value: `IP` and `DOMAIN`. NOTES: when the `protocol` is specified as `TCP` and the
// `scheduler` is specified as `wrr`, the item can only be set to `IP`.
func (o Layer4ListenerOutput) RealserverType() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.RealserverType }).(pulumi.StringOutput)
}

// UDP source server health check port detects received messages. Only used when the health check type is PORT.
func (o Layer4ListenerOutput) RecvContext() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.RecvContext }).(pulumi.StringOutput)
}

// Scheduling policy of the layer4 listener, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
func (o Layer4ListenerOutput) Scheduler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringPtrOutput { return v.Scheduler }).(pulumi.StringPtrOutput)
}

// UDP source server health check port detection sends messages. Only used when health check type is PORT.
func (o Layer4ListenerOutput) SendContext() pulumi.StringOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.StringOutput { return v.SendContext }).(pulumi.StringOutput)
}

// Status of the layer4 listener.
func (o Layer4ListenerOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Unhealthy threshold, which indicates how many consecutive check failures the source station is considered unhealthy.
// Range from 1 to 10. Default value is 1.
func (o Layer4ListenerOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Layer4Listener) pulumi.IntPtrOutput { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type Layer4ListenerArrayOutput struct{ *pulumi.OutputState }

func (Layer4ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerArrayOutput) ToLayer4ListenerArrayOutput() Layer4ListenerArrayOutput {
	return o
}

func (o Layer4ListenerArrayOutput) ToLayer4ListenerArrayOutputWithContext(ctx context.Context) Layer4ListenerArrayOutput {
	return o
}

func (o Layer4ListenerArrayOutput) Index(i pulumi.IntInput) Layer4ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Layer4Listener {
		return vs[0].([]*Layer4Listener)[vs[1].(int)]
	}).(Layer4ListenerOutput)
}

type Layer4ListenerMapOutput struct{ *pulumi.OutputState }

func (Layer4ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Layer4Listener)(nil)).Elem()
}

func (o Layer4ListenerMapOutput) ToLayer4ListenerMapOutput() Layer4ListenerMapOutput {
	return o
}

func (o Layer4ListenerMapOutput) ToLayer4ListenerMapOutputWithContext(ctx context.Context) Layer4ListenerMapOutput {
	return o
}

func (o Layer4ListenerMapOutput) MapIndex(k pulumi.StringInput) Layer4ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Layer4Listener {
		return vs[0].(map[string]*Layer4Listener)[vs[1].(string)]
	}).(Layer4ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerInput)(nil)).Elem(), &Layer4Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerArrayInput)(nil)).Elem(), Layer4ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Layer4ListenerMapInput)(nil)).Elem(), Layer4ListenerMap{})
	pulumi.RegisterOutputType(Layer4ListenerOutput{})
	pulumi.RegisterOutputType(Layer4ListenerArrayOutput{})
	pulumi.RegisterOutputType(Layer4ListenerMapOutput{})
}
