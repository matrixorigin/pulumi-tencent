// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterDatabases struct {
	pulumi.CustomResourceState

	// Character Set Type.
	CharacterSet pulumi.StringOutput `pulumi:"characterSet"`
	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Sort Rules.
	CollateRule pulumi.StringOutput `pulumi:"collateRule"`
	// Database name.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Remarks.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Authorize user host permissions.
	UserHostPrivileges ClusterDatabasesUserHostPrivilegeArrayOutput `pulumi:"userHostPrivileges"`
}

// NewClusterDatabases registers a new resource with the given unique name, arguments, and options.
func NewClusterDatabases(ctx *pulumi.Context,
	name string, args *ClusterDatabasesArgs, opts ...pulumi.ResourceOption) (*ClusterDatabases, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CharacterSet == nil {
		return nil, errors.New("invalid value for required argument 'CharacterSet'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.CollateRule == nil {
		return nil, errors.New("invalid value for required argument 'CollateRule'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterDatabases
	err := ctx.RegisterResource("tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterDatabases gets an existing ClusterDatabases resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterDatabases(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterDatabasesState, opts ...pulumi.ResourceOption) (*ClusterDatabases, error) {
	var resource ClusterDatabases
	err := ctx.ReadResource("tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterDatabases resources.
type clusterDatabasesState struct {
	// Character Set Type.
	CharacterSet *string `pulumi:"characterSet"`
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Sort Rules.
	CollateRule *string `pulumi:"collateRule"`
	// Database name.
	DbName *string `pulumi:"dbName"`
	// Remarks.
	Description *string `pulumi:"description"`
	// Authorize user host permissions.
	UserHostPrivileges []ClusterDatabasesUserHostPrivilege `pulumi:"userHostPrivileges"`
}

type ClusterDatabasesState struct {
	// Character Set Type.
	CharacterSet pulumi.StringPtrInput
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Sort Rules.
	CollateRule pulumi.StringPtrInput
	// Database name.
	DbName pulumi.StringPtrInput
	// Remarks.
	Description pulumi.StringPtrInput
	// Authorize user host permissions.
	UserHostPrivileges ClusterDatabasesUserHostPrivilegeArrayInput
}

func (ClusterDatabasesState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabasesState)(nil)).Elem()
}

type clusterDatabasesArgs struct {
	// Character Set Type.
	CharacterSet string `pulumi:"characterSet"`
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Sort Rules.
	CollateRule string `pulumi:"collateRule"`
	// Database name.
	DbName string `pulumi:"dbName"`
	// Remarks.
	Description *string `pulumi:"description"`
	// Authorize user host permissions.
	UserHostPrivileges []ClusterDatabasesUserHostPrivilege `pulumi:"userHostPrivileges"`
}

// The set of arguments for constructing a ClusterDatabases resource.
type ClusterDatabasesArgs struct {
	// Character Set Type.
	CharacterSet pulumi.StringInput
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Sort Rules.
	CollateRule pulumi.StringInput
	// Database name.
	DbName pulumi.StringInput
	// Remarks.
	Description pulumi.StringPtrInput
	// Authorize user host permissions.
	UserHostPrivileges ClusterDatabasesUserHostPrivilegeArrayInput
}

func (ClusterDatabasesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabasesArgs)(nil)).Elem()
}

type ClusterDatabasesInput interface {
	pulumi.Input

	ToClusterDatabasesOutput() ClusterDatabasesOutput
	ToClusterDatabasesOutputWithContext(ctx context.Context) ClusterDatabasesOutput
}

func (*ClusterDatabases) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDatabases)(nil)).Elem()
}

func (i *ClusterDatabases) ToClusterDatabasesOutput() ClusterDatabasesOutput {
	return i.ToClusterDatabasesOutputWithContext(context.Background())
}

func (i *ClusterDatabases) ToClusterDatabasesOutputWithContext(ctx context.Context) ClusterDatabasesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDatabasesOutput)
}

// ClusterDatabasesArrayInput is an input type that accepts ClusterDatabasesArray and ClusterDatabasesArrayOutput values.
// You can construct a concrete instance of `ClusterDatabasesArrayInput` via:
//
//	ClusterDatabasesArray{ ClusterDatabasesArgs{...} }
type ClusterDatabasesArrayInput interface {
	pulumi.Input

	ToClusterDatabasesArrayOutput() ClusterDatabasesArrayOutput
	ToClusterDatabasesArrayOutputWithContext(context.Context) ClusterDatabasesArrayOutput
}

type ClusterDatabasesArray []ClusterDatabasesInput

func (ClusterDatabasesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterDatabases)(nil)).Elem()
}

func (i ClusterDatabasesArray) ToClusterDatabasesArrayOutput() ClusterDatabasesArrayOutput {
	return i.ToClusterDatabasesArrayOutputWithContext(context.Background())
}

func (i ClusterDatabasesArray) ToClusterDatabasesArrayOutputWithContext(ctx context.Context) ClusterDatabasesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDatabasesArrayOutput)
}

// ClusterDatabasesMapInput is an input type that accepts ClusterDatabasesMap and ClusterDatabasesMapOutput values.
// You can construct a concrete instance of `ClusterDatabasesMapInput` via:
//
//	ClusterDatabasesMap{ "key": ClusterDatabasesArgs{...} }
type ClusterDatabasesMapInput interface {
	pulumi.Input

	ToClusterDatabasesMapOutput() ClusterDatabasesMapOutput
	ToClusterDatabasesMapOutputWithContext(context.Context) ClusterDatabasesMapOutput
}

type ClusterDatabasesMap map[string]ClusterDatabasesInput

func (ClusterDatabasesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterDatabases)(nil)).Elem()
}

func (i ClusterDatabasesMap) ToClusterDatabasesMapOutput() ClusterDatabasesMapOutput {
	return i.ToClusterDatabasesMapOutputWithContext(context.Background())
}

func (i ClusterDatabasesMap) ToClusterDatabasesMapOutputWithContext(ctx context.Context) ClusterDatabasesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDatabasesMapOutput)
}

type ClusterDatabasesOutput struct{ *pulumi.OutputState }

func (ClusterDatabasesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDatabases)(nil)).Elem()
}

func (o ClusterDatabasesOutput) ToClusterDatabasesOutput() ClusterDatabasesOutput {
	return o
}

func (o ClusterDatabasesOutput) ToClusterDatabasesOutputWithContext(ctx context.Context) ClusterDatabasesOutput {
	return o
}

// Character Set Type.
func (o ClusterDatabasesOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterDatabases) pulumi.StringOutput { return v.CharacterSet }).(pulumi.StringOutput)
}

// Cluster ID.
func (o ClusterDatabasesOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterDatabases) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Sort Rules.
func (o ClusterDatabasesOutput) CollateRule() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterDatabases) pulumi.StringOutput { return v.CollateRule }).(pulumi.StringOutput)
}

// Database name.
func (o ClusterDatabasesOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterDatabases) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// Remarks.
func (o ClusterDatabasesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDatabases) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Authorize user host permissions.
func (o ClusterDatabasesOutput) UserHostPrivileges() ClusterDatabasesUserHostPrivilegeArrayOutput {
	return o.ApplyT(func(v *ClusterDatabases) ClusterDatabasesUserHostPrivilegeArrayOutput { return v.UserHostPrivileges }).(ClusterDatabasesUserHostPrivilegeArrayOutput)
}

type ClusterDatabasesArrayOutput struct{ *pulumi.OutputState }

func (ClusterDatabasesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterDatabases)(nil)).Elem()
}

func (o ClusterDatabasesArrayOutput) ToClusterDatabasesArrayOutput() ClusterDatabasesArrayOutput {
	return o
}

func (o ClusterDatabasesArrayOutput) ToClusterDatabasesArrayOutputWithContext(ctx context.Context) ClusterDatabasesArrayOutput {
	return o
}

func (o ClusterDatabasesArrayOutput) Index(i pulumi.IntInput) ClusterDatabasesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterDatabases {
		return vs[0].([]*ClusterDatabases)[vs[1].(int)]
	}).(ClusterDatabasesOutput)
}

type ClusterDatabasesMapOutput struct{ *pulumi.OutputState }

func (ClusterDatabasesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterDatabases)(nil)).Elem()
}

func (o ClusterDatabasesMapOutput) ToClusterDatabasesMapOutput() ClusterDatabasesMapOutput {
	return o
}

func (o ClusterDatabasesMapOutput) ToClusterDatabasesMapOutputWithContext(ctx context.Context) ClusterDatabasesMapOutput {
	return o
}

func (o ClusterDatabasesMapOutput) MapIndex(k pulumi.StringInput) ClusterDatabasesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterDatabases {
		return vs[0].(map[string]*ClusterDatabases)[vs[1].(string)]
	}).(ClusterDatabasesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterDatabasesInput)(nil)).Elem(), &ClusterDatabases{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterDatabasesArrayInput)(nil)).Elem(), ClusterDatabasesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterDatabasesMapInput)(nil)).Elem(), ClusterDatabasesMap{})
	pulumi.RegisterOutputType(ClusterDatabasesOutput{})
	pulumi.RegisterOutputType(ClusterDatabasesArrayOutput{})
	pulumi.RegisterOutputType(ClusterDatabasesMapOutput{})
}