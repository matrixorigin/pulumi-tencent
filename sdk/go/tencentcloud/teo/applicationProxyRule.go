// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationProxyRule struct {
	pulumi.CustomResourceState

	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
	// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
	// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
	// client IP.
	ForwardClientIp pulumi.StringOutput `pulumi:"forwardClientIp"`
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort pulumi.StringOutput `pulumi:"originPort"`
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType pulumi.StringOutput `pulumi:"originType"`
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
	// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
	// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValues pulumi.StringArrayOutput `pulumi:"originValues"`
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Ports pulumi.StringArrayOutput `pulumi:"ports"`
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto pulumi.StringOutput `pulumi:"proto"`
	// Proxy ID.
	ProxyId pulumi.StringOutput `pulumi:"proxyId"`
	// Rule ID.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist pulumi.BoolOutput `pulumi:"sessionPersist"`
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
	// deactivated; `fail`: deployment failure/deactivation failure.
	Status pulumi.StringOutput `pulumi:"status"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewApplicationProxyRule registers a new resource with the given unique name, arguments, and options.
func NewApplicationProxyRule(ctx *pulumi.Context,
	name string, args *ApplicationProxyRuleArgs, opts ...pulumi.ResourceOption) (*ApplicationProxyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OriginPort == nil {
		return nil, errors.New("invalid value for required argument 'OriginPort'")
	}
	if args.OriginType == nil {
		return nil, errors.New("invalid value for required argument 'OriginType'")
	}
	if args.OriginValues == nil {
		return nil, errors.New("invalid value for required argument 'OriginValues'")
	}
	if args.Ports == nil {
		return nil, errors.New("invalid value for required argument 'Ports'")
	}
	if args.Proto == nil {
		return nil, errors.New("invalid value for required argument 'Proto'")
	}
	if args.ProxyId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationProxyRule
	err := ctx.RegisterResource("tencentcloud:Teo/applicationProxyRule:ApplicationProxyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationProxyRule gets an existing ApplicationProxyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationProxyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationProxyRuleState, opts ...pulumi.ResourceOption) (*ApplicationProxyRule, error) {
	var resource ApplicationProxyRule
	err := ctx.ReadResource("tencentcloud:Teo/applicationProxyRule:ApplicationProxyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationProxyRule resources.
type applicationProxyRuleState struct {
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
	// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
	// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
	// client IP.
	ForwardClientIp *string `pulumi:"forwardClientIp"`
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort *string `pulumi:"originPort"`
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType *string `pulumi:"originType"`
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
	// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
	// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValues []string `pulumi:"originValues"`
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Ports []string `pulumi:"ports"`
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto *string `pulumi:"proto"`
	// Proxy ID.
	ProxyId *string `pulumi:"proxyId"`
	// Rule ID.
	RuleId *string `pulumi:"ruleId"`
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist *bool `pulumi:"sessionPersist"`
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
	// deactivated; `fail`: deployment failure/deactivation failure.
	Status *string `pulumi:"status"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type ApplicationProxyRuleState struct {
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
	// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
	// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
	// client IP.
	ForwardClientIp pulumi.StringPtrInput
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort pulumi.StringPtrInput
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType pulumi.StringPtrInput
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
	// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
	// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValues pulumi.StringArrayInput
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Ports pulumi.StringArrayInput
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto pulumi.StringPtrInput
	// Proxy ID.
	ProxyId pulumi.StringPtrInput
	// Rule ID.
	RuleId pulumi.StringPtrInput
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist pulumi.BoolPtrInput
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
	// deactivated; `fail`: deployment failure/deactivation failure.
	Status pulumi.StringPtrInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (ApplicationProxyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationProxyRuleState)(nil)).Elem()
}

type applicationProxyRuleArgs struct {
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
	// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
	// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
	// client IP.
	ForwardClientIp *string `pulumi:"forwardClientIp"`
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort string `pulumi:"originPort"`
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType string `pulumi:"originType"`
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
	// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
	// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValues []string `pulumi:"originValues"`
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Ports []string `pulumi:"ports"`
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto string `pulumi:"proto"`
	// Proxy ID.
	ProxyId string `pulumi:"proxyId"`
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist *bool `pulumi:"sessionPersist"`
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
	// deactivated; `fail`: deployment failure/deactivation failure.
	Status *string `pulumi:"status"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ApplicationProxyRule resource.
type ApplicationProxyRuleArgs struct {
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
	// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
	// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
	// client IP.
	ForwardClientIp pulumi.StringPtrInput
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort pulumi.StringInput
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType pulumi.StringInput
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
	// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
	// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValues pulumi.StringArrayInput
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Ports pulumi.StringArrayInput
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto pulumi.StringInput
	// Proxy ID.
	ProxyId pulumi.StringInput
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist pulumi.BoolPtrInput
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
	// deactivated; `fail`: deployment failure/deactivation failure.
	Status pulumi.StringPtrInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (ApplicationProxyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationProxyRuleArgs)(nil)).Elem()
}

type ApplicationProxyRuleInput interface {
	pulumi.Input

	ToApplicationProxyRuleOutput() ApplicationProxyRuleOutput
	ToApplicationProxyRuleOutputWithContext(ctx context.Context) ApplicationProxyRuleOutput
}

func (*ApplicationProxyRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProxyRule)(nil)).Elem()
}

func (i *ApplicationProxyRule) ToApplicationProxyRuleOutput() ApplicationProxyRuleOutput {
	return i.ToApplicationProxyRuleOutputWithContext(context.Background())
}

func (i *ApplicationProxyRule) ToApplicationProxyRuleOutputWithContext(ctx context.Context) ApplicationProxyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProxyRuleOutput)
}

// ApplicationProxyRuleArrayInput is an input type that accepts ApplicationProxyRuleArray and ApplicationProxyRuleArrayOutput values.
// You can construct a concrete instance of `ApplicationProxyRuleArrayInput` via:
//
//	ApplicationProxyRuleArray{ ApplicationProxyRuleArgs{...} }
type ApplicationProxyRuleArrayInput interface {
	pulumi.Input

	ToApplicationProxyRuleArrayOutput() ApplicationProxyRuleArrayOutput
	ToApplicationProxyRuleArrayOutputWithContext(context.Context) ApplicationProxyRuleArrayOutput
}

type ApplicationProxyRuleArray []ApplicationProxyRuleInput

func (ApplicationProxyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationProxyRule)(nil)).Elem()
}

func (i ApplicationProxyRuleArray) ToApplicationProxyRuleArrayOutput() ApplicationProxyRuleArrayOutput {
	return i.ToApplicationProxyRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationProxyRuleArray) ToApplicationProxyRuleArrayOutputWithContext(ctx context.Context) ApplicationProxyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProxyRuleArrayOutput)
}

// ApplicationProxyRuleMapInput is an input type that accepts ApplicationProxyRuleMap and ApplicationProxyRuleMapOutput values.
// You can construct a concrete instance of `ApplicationProxyRuleMapInput` via:
//
//	ApplicationProxyRuleMap{ "key": ApplicationProxyRuleArgs{...} }
type ApplicationProxyRuleMapInput interface {
	pulumi.Input

	ToApplicationProxyRuleMapOutput() ApplicationProxyRuleMapOutput
	ToApplicationProxyRuleMapOutputWithContext(context.Context) ApplicationProxyRuleMapOutput
}

type ApplicationProxyRuleMap map[string]ApplicationProxyRuleInput

func (ApplicationProxyRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationProxyRule)(nil)).Elem()
}

func (i ApplicationProxyRuleMap) ToApplicationProxyRuleMapOutput() ApplicationProxyRuleMapOutput {
	return i.ToApplicationProxyRuleMapOutputWithContext(context.Background())
}

func (i ApplicationProxyRuleMap) ToApplicationProxyRuleMapOutputWithContext(ctx context.Context) ApplicationProxyRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProxyRuleMapOutput)
}

type ApplicationProxyRuleOutput struct{ *pulumi.OutputState }

func (ApplicationProxyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProxyRule)(nil)).Elem()
}

func (o ApplicationProxyRuleOutput) ToApplicationProxyRuleOutput() ApplicationProxyRuleOutput {
	return o
}

func (o ApplicationProxyRuleOutput) ToApplicationProxyRuleOutputWithContext(ctx context.Context) ApplicationProxyRuleOutput {
	return o
}

// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA;
// `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass
// the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the
// client IP.
func (o ApplicationProxyRuleOutput) ForwardClientIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.ForwardClientIp }).(pulumi.StringOutput)
}

// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
func (o ApplicationProxyRuleOutput) OriginPort() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.OriginPort }).(pulumi.StringOutput)
}

// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
func (o ApplicationProxyRuleOutput) OriginType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.OriginType }).(pulumi.StringOutput)
}

// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8',
// '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one
// element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
func (o ApplicationProxyRuleOutput) OriginValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringArrayOutput { return v.OriginValues }).(pulumi.StringArrayOutput)
}

// Valid values: `80` means port 80; `81-90` means port range 81-90.
func (o ApplicationProxyRuleOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringArrayOutput { return v.Ports }).(pulumi.StringArrayOutput)
}

// Protocol. Valid values: `TCP`, `UDP`.
func (o ApplicationProxyRuleOutput) Proto() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.Proto }).(pulumi.StringOutput)
}

// Proxy ID.
func (o ApplicationProxyRuleOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.ProxyId }).(pulumi.StringOutput)
}

// Rule ID.
func (o ApplicationProxyRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Specifies whether to enable session persistence. Default value is false.
func (o ApplicationProxyRuleOutput) SessionPersist() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.BoolOutput { return v.SessionPersist }).(pulumi.BoolOutput)
}

// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being
// deactivated; `fail`: deployment failure/deactivation failure.
func (o ApplicationProxyRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Site ID.
func (o ApplicationProxyRuleOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationProxyRule) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ApplicationProxyRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationProxyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationProxyRule)(nil)).Elem()
}

func (o ApplicationProxyRuleArrayOutput) ToApplicationProxyRuleArrayOutput() ApplicationProxyRuleArrayOutput {
	return o
}

func (o ApplicationProxyRuleArrayOutput) ToApplicationProxyRuleArrayOutputWithContext(ctx context.Context) ApplicationProxyRuleArrayOutput {
	return o
}

func (o ApplicationProxyRuleArrayOutput) Index(i pulumi.IntInput) ApplicationProxyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationProxyRule {
		return vs[0].([]*ApplicationProxyRule)[vs[1].(int)]
	}).(ApplicationProxyRuleOutput)
}

type ApplicationProxyRuleMapOutput struct{ *pulumi.OutputState }

func (ApplicationProxyRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationProxyRule)(nil)).Elem()
}

func (o ApplicationProxyRuleMapOutput) ToApplicationProxyRuleMapOutput() ApplicationProxyRuleMapOutput {
	return o
}

func (o ApplicationProxyRuleMapOutput) ToApplicationProxyRuleMapOutputWithContext(ctx context.Context) ApplicationProxyRuleMapOutput {
	return o
}

func (o ApplicationProxyRuleMapOutput) MapIndex(k pulumi.StringInput) ApplicationProxyRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationProxyRule {
		return vs[0].(map[string]*ApplicationProxyRule)[vs[1].(string)]
	}).(ApplicationProxyRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProxyRuleInput)(nil)).Elem(), &ApplicationProxyRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProxyRuleArrayInput)(nil)).Elem(), ApplicationProxyRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProxyRuleMapInput)(nil)).Elem(), ApplicationProxyRuleMap{})
	pulumi.RegisterOutputType(ApplicationProxyRuleOutput{})
	pulumi.RegisterOutputType(ApplicationProxyRuleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationProxyRuleMapOutput{})
}