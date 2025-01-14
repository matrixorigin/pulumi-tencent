// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dasb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Acl struct {
	pulumi.CustomResourceState

	// Associate high-risk DB template IDs.
	AcTemplateIdSets pulumi.StringArrayOutput `pulumi:"acTemplateIdSets"`
	// Associated accounts.
	AccountSets pulumi.StringArrayOutput `pulumi:"accountSets"`
	// Allow access credential,default allow.
	AllowAccessCredential pulumi.BoolPtrOutput `pulumi:"allowAccessCredential"`
	// Allow any account.
	AllowAnyAccount pulumi.BoolOutput `pulumi:"allowAnyAccount"`
	// Allow clip file down.
	AllowClipFileDown pulumi.BoolPtrOutput `pulumi:"allowClipFileDown"`
	// Allow clip file up.
	AllowClipFileUp pulumi.BoolPtrOutput `pulumi:"allowClipFileUp"`
	// Allow clip text down.
	AllowClipTextDown pulumi.BoolPtrOutput `pulumi:"allowClipTextDown"`
	// Allow clip text up.
	AllowClipTextUp pulumi.BoolPtrOutput `pulumi:"allowClipTextUp"`
	// Allow disk file download.
	AllowDiskFileDown pulumi.BoolPtrOutput `pulumi:"allowDiskFileDown"`
	// Allow disk file upload.
	AllowDiskFileUp pulumi.BoolPtrOutput `pulumi:"allowDiskFileUp"`
	// Allow disk redirect.
	AllowDiskRedirect pulumi.BoolOutput `pulumi:"allowDiskRedirect"`
	// Allow sftp file delete.
	AllowFileDel pulumi.BoolPtrOutput `pulumi:"allowFileDel"`
	// Allow sftp file download.
	AllowFileDown pulumi.BoolPtrOutput `pulumi:"allowFileDown"`
	// Allow sftp up file.
	AllowFileUp pulumi.BoolPtrOutput `pulumi:"allowFileUp"`
	// Allow shell file download.
	AllowShellFileDown pulumi.BoolPtrOutput `pulumi:"allowShellFileDown"`
	// Allow shell file upload.
	AllowShellFileUp pulumi.BoolPtrOutput `pulumi:"allowShellFileUp"`
	// Associated high-risk command template ID.
	CmdTemplateIdSets pulumi.IntArrayOutput `pulumi:"cmdTemplateIdSets"`
	// Department id.
	DepartmentId pulumi.StringPtrOutput `pulumi:"departmentId"`
	// Associated device group ID.
	DeviceGroupIdSets pulumi.IntArrayOutput `pulumi:"deviceGroupIdSets"`
	// Associated collection of device IDs.
	DeviceIdSets pulumi.IntArrayOutput `pulumi:"deviceIdSets"`
	// File transfer download size limit (reserved parameter, currently unused).
	MaxFileDownSize pulumi.IntPtrOutput `pulumi:"maxFileDownSize"`
	// File upload transfer size limit (artifact parameter, currently unused).
	MaxFileUpSize pulumi.IntPtrOutput `pulumi:"maxFileUpSize"`
	// Acl name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Associated user group ID.
	UserGroupIdSets pulumi.IntArrayOutput `pulumi:"userGroupIdSets"`
	// Associated set of user IDs.
	UserIdSets pulumi.IntArrayOutput `pulumi:"userIdSets"`
	// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateFrom pulumi.StringOutput `pulumi:"validateFrom"`
	// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateTo pulumi.StringOutput `pulumi:"validateTo"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowAnyAccount == nil {
		return nil, errors.New("invalid value for required argument 'AllowAnyAccount'")
	}
	if args.AllowDiskRedirect == nil {
		return nil, errors.New("invalid value for required argument 'AllowDiskRedirect'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Acl
	err := ctx.RegisterResource("tencentcloud:Dasb/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("tencentcloud:Dasb/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// Associate high-risk DB template IDs.
	AcTemplateIdSets []string `pulumi:"acTemplateIdSets"`
	// Associated accounts.
	AccountSets []string `pulumi:"accountSets"`
	// Allow access credential,default allow.
	AllowAccessCredential *bool `pulumi:"allowAccessCredential"`
	// Allow any account.
	AllowAnyAccount *bool `pulumi:"allowAnyAccount"`
	// Allow clip file down.
	AllowClipFileDown *bool `pulumi:"allowClipFileDown"`
	// Allow clip file up.
	AllowClipFileUp *bool `pulumi:"allowClipFileUp"`
	// Allow clip text down.
	AllowClipTextDown *bool `pulumi:"allowClipTextDown"`
	// Allow clip text up.
	AllowClipTextUp *bool `pulumi:"allowClipTextUp"`
	// Allow disk file download.
	AllowDiskFileDown *bool `pulumi:"allowDiskFileDown"`
	// Allow disk file upload.
	AllowDiskFileUp *bool `pulumi:"allowDiskFileUp"`
	// Allow disk redirect.
	AllowDiskRedirect *bool `pulumi:"allowDiskRedirect"`
	// Allow sftp file delete.
	AllowFileDel *bool `pulumi:"allowFileDel"`
	// Allow sftp file download.
	AllowFileDown *bool `pulumi:"allowFileDown"`
	// Allow sftp up file.
	AllowFileUp *bool `pulumi:"allowFileUp"`
	// Allow shell file download.
	AllowShellFileDown *bool `pulumi:"allowShellFileDown"`
	// Allow shell file upload.
	AllowShellFileUp *bool `pulumi:"allowShellFileUp"`
	// Associated high-risk command template ID.
	CmdTemplateIdSets []int `pulumi:"cmdTemplateIdSets"`
	// Department id.
	DepartmentId *string `pulumi:"departmentId"`
	// Associated device group ID.
	DeviceGroupIdSets []int `pulumi:"deviceGroupIdSets"`
	// Associated collection of device IDs.
	DeviceIdSets []int `pulumi:"deviceIdSets"`
	// File transfer download size limit (reserved parameter, currently unused).
	MaxFileDownSize *int `pulumi:"maxFileDownSize"`
	// File upload transfer size limit (artifact parameter, currently unused).
	MaxFileUpSize *int `pulumi:"maxFileUpSize"`
	// Acl name.
	Name *string `pulumi:"name"`
	// Associated user group ID.
	UserGroupIdSets []int `pulumi:"userGroupIdSets"`
	// Associated set of user IDs.
	UserIdSets []int `pulumi:"userIdSets"`
	// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateFrom *string `pulumi:"validateFrom"`
	// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateTo *string `pulumi:"validateTo"`
}

type AclState struct {
	// Associate high-risk DB template IDs.
	AcTemplateIdSets pulumi.StringArrayInput
	// Associated accounts.
	AccountSets pulumi.StringArrayInput
	// Allow access credential,default allow.
	AllowAccessCredential pulumi.BoolPtrInput
	// Allow any account.
	AllowAnyAccount pulumi.BoolPtrInput
	// Allow clip file down.
	AllowClipFileDown pulumi.BoolPtrInput
	// Allow clip file up.
	AllowClipFileUp pulumi.BoolPtrInput
	// Allow clip text down.
	AllowClipTextDown pulumi.BoolPtrInput
	// Allow clip text up.
	AllowClipTextUp pulumi.BoolPtrInput
	// Allow disk file download.
	AllowDiskFileDown pulumi.BoolPtrInput
	// Allow disk file upload.
	AllowDiskFileUp pulumi.BoolPtrInput
	// Allow disk redirect.
	AllowDiskRedirect pulumi.BoolPtrInput
	// Allow sftp file delete.
	AllowFileDel pulumi.BoolPtrInput
	// Allow sftp file download.
	AllowFileDown pulumi.BoolPtrInput
	// Allow sftp up file.
	AllowFileUp pulumi.BoolPtrInput
	// Allow shell file download.
	AllowShellFileDown pulumi.BoolPtrInput
	// Allow shell file upload.
	AllowShellFileUp pulumi.BoolPtrInput
	// Associated high-risk command template ID.
	CmdTemplateIdSets pulumi.IntArrayInput
	// Department id.
	DepartmentId pulumi.StringPtrInput
	// Associated device group ID.
	DeviceGroupIdSets pulumi.IntArrayInput
	// Associated collection of device IDs.
	DeviceIdSets pulumi.IntArrayInput
	// File transfer download size limit (reserved parameter, currently unused).
	MaxFileDownSize pulumi.IntPtrInput
	// File upload transfer size limit (artifact parameter, currently unused).
	MaxFileUpSize pulumi.IntPtrInput
	// Acl name.
	Name pulumi.StringPtrInput
	// Associated user group ID.
	UserGroupIdSets pulumi.IntArrayInput
	// Associated set of user IDs.
	UserIdSets pulumi.IntArrayInput
	// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateFrom pulumi.StringPtrInput
	// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateTo pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// Associate high-risk DB template IDs.
	AcTemplateIdSets []string `pulumi:"acTemplateIdSets"`
	// Associated accounts.
	AccountSets []string `pulumi:"accountSets"`
	// Allow access credential,default allow.
	AllowAccessCredential *bool `pulumi:"allowAccessCredential"`
	// Allow any account.
	AllowAnyAccount bool `pulumi:"allowAnyAccount"`
	// Allow clip file down.
	AllowClipFileDown *bool `pulumi:"allowClipFileDown"`
	// Allow clip file up.
	AllowClipFileUp *bool `pulumi:"allowClipFileUp"`
	// Allow clip text down.
	AllowClipTextDown *bool `pulumi:"allowClipTextDown"`
	// Allow clip text up.
	AllowClipTextUp *bool `pulumi:"allowClipTextUp"`
	// Allow disk file download.
	AllowDiskFileDown *bool `pulumi:"allowDiskFileDown"`
	// Allow disk file upload.
	AllowDiskFileUp *bool `pulumi:"allowDiskFileUp"`
	// Allow disk redirect.
	AllowDiskRedirect bool `pulumi:"allowDiskRedirect"`
	// Allow sftp file delete.
	AllowFileDel *bool `pulumi:"allowFileDel"`
	// Allow sftp file download.
	AllowFileDown *bool `pulumi:"allowFileDown"`
	// Allow sftp up file.
	AllowFileUp *bool `pulumi:"allowFileUp"`
	// Allow shell file download.
	AllowShellFileDown *bool `pulumi:"allowShellFileDown"`
	// Allow shell file upload.
	AllowShellFileUp *bool `pulumi:"allowShellFileUp"`
	// Associated high-risk command template ID.
	CmdTemplateIdSets []int `pulumi:"cmdTemplateIdSets"`
	// Department id.
	DepartmentId *string `pulumi:"departmentId"`
	// Associated device group ID.
	DeviceGroupIdSets []int `pulumi:"deviceGroupIdSets"`
	// Associated collection of device IDs.
	DeviceIdSets []int `pulumi:"deviceIdSets"`
	// File transfer download size limit (reserved parameter, currently unused).
	MaxFileDownSize *int `pulumi:"maxFileDownSize"`
	// File upload transfer size limit (artifact parameter, currently unused).
	MaxFileUpSize *int `pulumi:"maxFileUpSize"`
	// Acl name.
	Name *string `pulumi:"name"`
	// Associated user group ID.
	UserGroupIdSets []int `pulumi:"userGroupIdSets"`
	// Associated set of user IDs.
	UserIdSets []int `pulumi:"userIdSets"`
	// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateFrom *string `pulumi:"validateFrom"`
	// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateTo *string `pulumi:"validateTo"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// Associate high-risk DB template IDs.
	AcTemplateIdSets pulumi.StringArrayInput
	// Associated accounts.
	AccountSets pulumi.StringArrayInput
	// Allow access credential,default allow.
	AllowAccessCredential pulumi.BoolPtrInput
	// Allow any account.
	AllowAnyAccount pulumi.BoolInput
	// Allow clip file down.
	AllowClipFileDown pulumi.BoolPtrInput
	// Allow clip file up.
	AllowClipFileUp pulumi.BoolPtrInput
	// Allow clip text down.
	AllowClipTextDown pulumi.BoolPtrInput
	// Allow clip text up.
	AllowClipTextUp pulumi.BoolPtrInput
	// Allow disk file download.
	AllowDiskFileDown pulumi.BoolPtrInput
	// Allow disk file upload.
	AllowDiskFileUp pulumi.BoolPtrInput
	// Allow disk redirect.
	AllowDiskRedirect pulumi.BoolInput
	// Allow sftp file delete.
	AllowFileDel pulumi.BoolPtrInput
	// Allow sftp file download.
	AllowFileDown pulumi.BoolPtrInput
	// Allow sftp up file.
	AllowFileUp pulumi.BoolPtrInput
	// Allow shell file download.
	AllowShellFileDown pulumi.BoolPtrInput
	// Allow shell file upload.
	AllowShellFileUp pulumi.BoolPtrInput
	// Associated high-risk command template ID.
	CmdTemplateIdSets pulumi.IntArrayInput
	// Department id.
	DepartmentId pulumi.StringPtrInput
	// Associated device group ID.
	DeviceGroupIdSets pulumi.IntArrayInput
	// Associated collection of device IDs.
	DeviceIdSets pulumi.IntArrayInput
	// File transfer download size limit (reserved parameter, currently unused).
	MaxFileDownSize pulumi.IntPtrInput
	// File upload transfer size limit (artifact parameter, currently unused).
	MaxFileUpSize pulumi.IntPtrInput
	// Acl name.
	Name pulumi.StringPtrInput
	// Associated user group ID.
	UserGroupIdSets pulumi.IntArrayInput
	// Associated set of user IDs.
	UserIdSets pulumi.IntArrayInput
	// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateFrom pulumi.StringPtrInput
	// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
	// the access rights will be valid for a long time.
	ValidateTo pulumi.StringPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//	AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//	AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

// Associate high-risk DB template IDs.
func (o AclOutput) AcTemplateIdSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringArrayOutput { return v.AcTemplateIdSets }).(pulumi.StringArrayOutput)
}

// Associated accounts.
func (o AclOutput) AccountSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringArrayOutput { return v.AccountSets }).(pulumi.StringArrayOutput)
}

// Allow access credential,default allow.
func (o AclOutput) AllowAccessCredential() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowAccessCredential }).(pulumi.BoolPtrOutput)
}

// Allow any account.
func (o AclOutput) AllowAnyAccount() pulumi.BoolOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolOutput { return v.AllowAnyAccount }).(pulumi.BoolOutput)
}

// Allow clip file down.
func (o AclOutput) AllowClipFileDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowClipFileDown }).(pulumi.BoolPtrOutput)
}

// Allow clip file up.
func (o AclOutput) AllowClipFileUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowClipFileUp }).(pulumi.BoolPtrOutput)
}

// Allow clip text down.
func (o AclOutput) AllowClipTextDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowClipTextDown }).(pulumi.BoolPtrOutput)
}

// Allow clip text up.
func (o AclOutput) AllowClipTextUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowClipTextUp }).(pulumi.BoolPtrOutput)
}

// Allow disk file download.
func (o AclOutput) AllowDiskFileDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowDiskFileDown }).(pulumi.BoolPtrOutput)
}

// Allow disk file upload.
func (o AclOutput) AllowDiskFileUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowDiskFileUp }).(pulumi.BoolPtrOutput)
}

// Allow disk redirect.
func (o AclOutput) AllowDiskRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolOutput { return v.AllowDiskRedirect }).(pulumi.BoolOutput)
}

// Allow sftp file delete.
func (o AclOutput) AllowFileDel() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowFileDel }).(pulumi.BoolPtrOutput)
}

// Allow sftp file download.
func (o AclOutput) AllowFileDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowFileDown }).(pulumi.BoolPtrOutput)
}

// Allow sftp up file.
func (o AclOutput) AllowFileUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowFileUp }).(pulumi.BoolPtrOutput)
}

// Allow shell file download.
func (o AclOutput) AllowShellFileDown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowShellFileDown }).(pulumi.BoolPtrOutput)
}

// Allow shell file upload.
func (o AclOutput) AllowShellFileUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.AllowShellFileUp }).(pulumi.BoolPtrOutput)
}

// Associated high-risk command template ID.
func (o AclOutput) CmdTemplateIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntArrayOutput { return v.CmdTemplateIdSets }).(pulumi.IntArrayOutput)
}

// Department id.
func (o AclOutput) DepartmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.DepartmentId }).(pulumi.StringPtrOutput)
}

// Associated device group ID.
func (o AclOutput) DeviceGroupIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntArrayOutput { return v.DeviceGroupIdSets }).(pulumi.IntArrayOutput)
}

// Associated collection of device IDs.
func (o AclOutput) DeviceIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntArrayOutput { return v.DeviceIdSets }).(pulumi.IntArrayOutput)
}

// File transfer download size limit (reserved parameter, currently unused).
func (o AclOutput) MaxFileDownSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntPtrOutput { return v.MaxFileDownSize }).(pulumi.IntPtrOutput)
}

// File upload transfer size limit (artifact parameter, currently unused).
func (o AclOutput) MaxFileUpSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntPtrOutput { return v.MaxFileUpSize }).(pulumi.IntPtrOutput)
}

// Acl name.
func (o AclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Associated user group ID.
func (o AclOutput) UserGroupIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntArrayOutput { return v.UserGroupIdSets }).(pulumi.IntArrayOutput)
}

// Associated set of user IDs.
func (o AclOutput) UserIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Acl) pulumi.IntArrayOutput { return v.UserIdSets }).(pulumi.IntArrayOutput)
}

// Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in,
// the access rights will be valid for a long time.
func (o AclOutput) ValidateFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ValidateFrom }).(pulumi.StringOutput)
}

// Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in,
// the access rights will be valid for a long time.
func (o AclOutput) ValidateTo() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ValidateTo }).(pulumi.StringOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].([]*Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].(map[string]*Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclArrayInput)(nil)).Elem(), AclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclMapInput)(nil)).Elem(), AclMap{})
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
