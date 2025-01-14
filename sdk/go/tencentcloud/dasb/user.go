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

type User struct {
	pulumi.CustomResourceState

	// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
	AuthType pulumi.IntPtrOutput `pulumi:"authType"`
	// Department ID, such as: 1.2.3.
	DepartmentId pulumi.StringPtrOutput `pulumi:"departmentId"`
	// Email. Please provide at least one of `phone` or `email`.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// The set of user group IDs to which it belongs.
	GroupIdSets pulumi.IntArrayOutput `pulumi:"groupIdSets"`
	// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
	// provide at least one of `phone` or `email`.
	Phone pulumi.StringPtrOutput `pulumi:"phone"`
	// Real name, maximum length 20 characters, cannot contain blank characters.
	RealName pulumi.StringOutput `pulumi:"realName"`
	// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
	// '.', '_', '-'.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
	// be valid for a long time.
	ValidateFrom pulumi.StringOutput `pulumi:"validateFrom"`
	// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
	// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
	// not allowed, 1 - means access is allowed.
	ValidateTime pulumi.StringPtrOutput `pulumi:"validateTime"`
	// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
	// will be valid for a long time.
	ValidateTo pulumi.StringOutput `pulumi:"validateTo"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealName == nil {
		return nil, errors.New("invalid value for required argument 'RealName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("tencentcloud:Dasb/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("tencentcloud:Dasb/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
	AuthType *int `pulumi:"authType"`
	// Department ID, such as: 1.2.3.
	DepartmentId *string `pulumi:"departmentId"`
	// Email. Please provide at least one of `phone` or `email`.
	Email *string `pulumi:"email"`
	// The set of user group IDs to which it belongs.
	GroupIdSets []int `pulumi:"groupIdSets"`
	// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
	// provide at least one of `phone` or `email`.
	Phone *string `pulumi:"phone"`
	// Real name, maximum length 20 characters, cannot contain blank characters.
	RealName *string `pulumi:"realName"`
	// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
	// '.', '_', '-'.
	UserName *string `pulumi:"userName"`
	// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
	// be valid for a long time.
	ValidateFrom *string `pulumi:"validateFrom"`
	// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
	// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
	// not allowed, 1 - means access is allowed.
	ValidateTime *string `pulumi:"validateTime"`
	// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
	// will be valid for a long time.
	ValidateTo *string `pulumi:"validateTo"`
}

type UserState struct {
	// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
	AuthType pulumi.IntPtrInput
	// Department ID, such as: 1.2.3.
	DepartmentId pulumi.StringPtrInput
	// Email. Please provide at least one of `phone` or `email`.
	Email pulumi.StringPtrInput
	// The set of user group IDs to which it belongs.
	GroupIdSets pulumi.IntArrayInput
	// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
	// provide at least one of `phone` or `email`.
	Phone pulumi.StringPtrInput
	// Real name, maximum length 20 characters, cannot contain blank characters.
	RealName pulumi.StringPtrInput
	// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
	// '.', '_', '-'.
	UserName pulumi.StringPtrInput
	// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
	// be valid for a long time.
	ValidateFrom pulumi.StringPtrInput
	// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
	// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
	// not allowed, 1 - means access is allowed.
	ValidateTime pulumi.StringPtrInput
	// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
	// will be valid for a long time.
	ValidateTo pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
	AuthType *int `pulumi:"authType"`
	// Department ID, such as: 1.2.3.
	DepartmentId *string `pulumi:"departmentId"`
	// Email. Please provide at least one of `phone` or `email`.
	Email *string `pulumi:"email"`
	// The set of user group IDs to which it belongs.
	GroupIdSets []int `pulumi:"groupIdSets"`
	// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
	// provide at least one of `phone` or `email`.
	Phone *string `pulumi:"phone"`
	// Real name, maximum length 20 characters, cannot contain blank characters.
	RealName string `pulumi:"realName"`
	// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
	// '.', '_', '-'.
	UserName string `pulumi:"userName"`
	// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
	// be valid for a long time.
	ValidateFrom *string `pulumi:"validateFrom"`
	// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
	// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
	// not allowed, 1 - means access is allowed.
	ValidateTime *string `pulumi:"validateTime"`
	// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
	// will be valid for a long time.
	ValidateTo *string `pulumi:"validateTo"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
	AuthType pulumi.IntPtrInput
	// Department ID, such as: 1.2.3.
	DepartmentId pulumi.StringPtrInput
	// Email. Please provide at least one of `phone` or `email`.
	Email pulumi.StringPtrInput
	// The set of user group IDs to which it belongs.
	GroupIdSets pulumi.IntArrayInput
	// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
	// provide at least one of `phone` or `email`.
	Phone pulumi.StringPtrInput
	// Real name, maximum length 20 characters, cannot contain blank characters.
	RealName pulumi.StringInput
	// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
	// '.', '_', '-'.
	UserName pulumi.StringInput
	// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
	// be valid for a long time.
	ValidateFrom pulumi.StringPtrInput
	// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
	// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
	// not allowed, 1 - means access is allowed.
	ValidateTime pulumi.StringPtrInput
	// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
	// will be valid for a long time.
	ValidateTo pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
func (o UserOutput) AuthType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *User) pulumi.IntPtrOutput { return v.AuthType }).(pulumi.IntPtrOutput)
}

// Department ID, such as: 1.2.3.
func (o UserOutput) DepartmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.DepartmentId }).(pulumi.StringPtrOutput)
}

// Email. Please provide at least one of `phone` or `email`.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// The set of user group IDs to which it belongs.
func (o UserOutput) GroupIdSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *User) pulumi.IntArrayOutput { return v.GroupIdSets }).(pulumi.IntArrayOutput)
}

// Enter it in the format of country area code|mobile phone number. For example: +86|***********, +852|xxxxxxxx. Please
// provide at least one of `phone` or `email`.
func (o UserOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Phone }).(pulumi.StringPtrOutput)
}

// Real name, maximum length 20 characters, cannot contain blank characters.
func (o UserOutput) RealName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.RealName }).(pulumi.StringOutput)
}

// Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers,
// '.', '_', '-'.
func (o UserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will
// be valid for a long time.
func (o UserOutput) ValidateFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.ValidateFrom }).(pulumi.StringOutput)
}

// Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is
// allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is
// not allowed, 1 - means access is allowed.
func (o UserOutput) ValidateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.ValidateTime }).(pulumi.StringPtrOutput)
}

// User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user
// will be valid for a long time.
func (o UserOutput) ValidateTo() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.ValidateTo }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
