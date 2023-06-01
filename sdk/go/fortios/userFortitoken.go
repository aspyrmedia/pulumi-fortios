// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiToken.
//
// ## Import
//
// # User Fortitoken can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/userFortitoken:UserFortitoken labelname {{serial_number}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/userFortitoken:UserFortitoken labelname {{serial_number}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type UserFortitoken struct {
	pulumi.CustomResourceState

	// Mobile token user activation-code.
	ActivationCode pulumi.StringOutput `pulumi:"activationCode"`
	// Mobile token user activation-code expire time.
	ActivationExpire pulumi.IntOutput `pulumi:"activationExpire"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Mobile token license.
	License pulumi.StringOutput `pulumi:"license"`
	// Device Mobile Version.
	OsVer pulumi.StringOutput `pulumi:"osVer"`
	// Device Reg ID.
	RegId pulumi.StringOutput `pulumi:"regId"`
	// Token seed.
	Seed pulumi.StringOutput `pulumi:"seed"`
	// Serial number.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Status Valid values: `active`, `lock`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserFortitoken registers a new resource with the given unique name, arguments, and options.
func NewUserFortitoken(ctx *pulumi.Context,
	name string, args *UserFortitokenArgs, opts ...pulumi.ResourceOption) (*UserFortitoken, error) {
	if args == nil {
		args = &UserFortitokenArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserFortitoken
	err := ctx.RegisterResource("fortios:index/userFortitoken:UserFortitoken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFortitoken gets an existing UserFortitoken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFortitoken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFortitokenState, opts ...pulumi.ResourceOption) (*UserFortitoken, error) {
	var resource UserFortitoken
	err := ctx.ReadResource("fortios:index/userFortitoken:UserFortitoken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFortitoken resources.
type userFortitokenState struct {
	// Mobile token user activation-code.
	ActivationCode *string `pulumi:"activationCode"`
	// Mobile token user activation-code expire time.
	ActivationExpire *int `pulumi:"activationExpire"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Mobile token license.
	License *string `pulumi:"license"`
	// Device Mobile Version.
	OsVer *string `pulumi:"osVer"`
	// Device Reg ID.
	RegId *string `pulumi:"regId"`
	// Token seed.
	Seed *string `pulumi:"seed"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Status Valid values: `active`, `lock`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserFortitokenState struct {
	// Mobile token user activation-code.
	ActivationCode pulumi.StringPtrInput
	// Mobile token user activation-code expire time.
	ActivationExpire pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Mobile token license.
	License pulumi.StringPtrInput
	// Device Mobile Version.
	OsVer pulumi.StringPtrInput
	// Device Reg ID.
	RegId pulumi.StringPtrInput
	// Token seed.
	Seed pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Status Valid values: `active`, `lock`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFortitokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFortitokenState)(nil)).Elem()
}

type userFortitokenArgs struct {
	// Mobile token user activation-code.
	ActivationCode *string `pulumi:"activationCode"`
	// Mobile token user activation-code expire time.
	ActivationExpire *int `pulumi:"activationExpire"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Mobile token license.
	License *string `pulumi:"license"`
	// Device Mobile Version.
	OsVer *string `pulumi:"osVer"`
	// Device Reg ID.
	RegId *string `pulumi:"regId"`
	// Token seed.
	Seed *string `pulumi:"seed"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Status Valid values: `active`, `lock`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserFortitoken resource.
type UserFortitokenArgs struct {
	// Mobile token user activation-code.
	ActivationCode pulumi.StringPtrInput
	// Mobile token user activation-code expire time.
	ActivationExpire pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Mobile token license.
	License pulumi.StringPtrInput
	// Device Mobile Version.
	OsVer pulumi.StringPtrInput
	// Device Reg ID.
	RegId pulumi.StringPtrInput
	// Token seed.
	Seed pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Status Valid values: `active`, `lock`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFortitokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFortitokenArgs)(nil)).Elem()
}

type UserFortitokenInput interface {
	pulumi.Input

	ToUserFortitokenOutput() UserFortitokenOutput
	ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput
}

func (*UserFortitoken) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFortitoken)(nil)).Elem()
}

func (i *UserFortitoken) ToUserFortitokenOutput() UserFortitokenOutput {
	return i.ToUserFortitokenOutputWithContext(context.Background())
}

func (i *UserFortitoken) ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenOutput)
}

// UserFortitokenArrayInput is an input type that accepts UserFortitokenArray and UserFortitokenArrayOutput values.
// You can construct a concrete instance of `UserFortitokenArrayInput` via:
//
//	UserFortitokenArray{ UserFortitokenArgs{...} }
type UserFortitokenArrayInput interface {
	pulumi.Input

	ToUserFortitokenArrayOutput() UserFortitokenArrayOutput
	ToUserFortitokenArrayOutputWithContext(context.Context) UserFortitokenArrayOutput
}

type UserFortitokenArray []UserFortitokenInput

func (UserFortitokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFortitoken)(nil)).Elem()
}

func (i UserFortitokenArray) ToUserFortitokenArrayOutput() UserFortitokenArrayOutput {
	return i.ToUserFortitokenArrayOutputWithContext(context.Background())
}

func (i UserFortitokenArray) ToUserFortitokenArrayOutputWithContext(ctx context.Context) UserFortitokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenArrayOutput)
}

// UserFortitokenMapInput is an input type that accepts UserFortitokenMap and UserFortitokenMapOutput values.
// You can construct a concrete instance of `UserFortitokenMapInput` via:
//
//	UserFortitokenMap{ "key": UserFortitokenArgs{...} }
type UserFortitokenMapInput interface {
	pulumi.Input

	ToUserFortitokenMapOutput() UserFortitokenMapOutput
	ToUserFortitokenMapOutputWithContext(context.Context) UserFortitokenMapOutput
}

type UserFortitokenMap map[string]UserFortitokenInput

func (UserFortitokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFortitoken)(nil)).Elem()
}

func (i UserFortitokenMap) ToUserFortitokenMapOutput() UserFortitokenMapOutput {
	return i.ToUserFortitokenMapOutputWithContext(context.Background())
}

func (i UserFortitokenMap) ToUserFortitokenMapOutputWithContext(ctx context.Context) UserFortitokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFortitokenMapOutput)
}

type UserFortitokenOutput struct{ *pulumi.OutputState }

func (UserFortitokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenOutput) ToUserFortitokenOutput() UserFortitokenOutput {
	return o
}

func (o UserFortitokenOutput) ToUserFortitokenOutputWithContext(ctx context.Context) UserFortitokenOutput {
	return o
}

// Mobile token user activation-code.
func (o UserFortitokenOutput) ActivationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.ActivationCode }).(pulumi.StringOutput)
}

// Mobile token user activation-code expire time.
func (o UserFortitokenOutput) ActivationExpire() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.IntOutput { return v.ActivationExpire }).(pulumi.IntOutput)
}

// Comment.
func (o UserFortitokenOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Mobile token license.
func (o UserFortitokenOutput) License() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.License }).(pulumi.StringOutput)
}

// Device Mobile Version.
func (o UserFortitokenOutput) OsVer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.OsVer }).(pulumi.StringOutput)
}

// Device Reg ID.
func (o UserFortitokenOutput) RegId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.RegId }).(pulumi.StringOutput)
}

// Token seed.
func (o UserFortitokenOutput) Seed() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.Seed }).(pulumi.StringOutput)
}

// Serial number.
func (o UserFortitokenOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// Status Valid values: `active`, `lock`.
func (o UserFortitokenOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserFortitokenOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFortitoken) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserFortitokenArrayOutput struct{ *pulumi.OutputState }

func (UserFortitokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenArrayOutput) ToUserFortitokenArrayOutput() UserFortitokenArrayOutput {
	return o
}

func (o UserFortitokenArrayOutput) ToUserFortitokenArrayOutputWithContext(ctx context.Context) UserFortitokenArrayOutput {
	return o
}

func (o UserFortitokenArrayOutput) Index(i pulumi.IntInput) UserFortitokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserFortitoken {
		return vs[0].([]*UserFortitoken)[vs[1].(int)]
	}).(UserFortitokenOutput)
}

type UserFortitokenMapOutput struct{ *pulumi.OutputState }

func (UserFortitokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFortitoken)(nil)).Elem()
}

func (o UserFortitokenMapOutput) ToUserFortitokenMapOutput() UserFortitokenMapOutput {
	return o
}

func (o UserFortitokenMapOutput) ToUserFortitokenMapOutputWithContext(ctx context.Context) UserFortitokenMapOutput {
	return o
}

func (o UserFortitokenMapOutput) MapIndex(k pulumi.StringInput) UserFortitokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserFortitoken {
		return vs[0].(map[string]*UserFortitoken)[vs[1].(string)]
	}).(UserFortitokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenInput)(nil)).Elem(), &UserFortitoken{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenArrayInput)(nil)).Elem(), UserFortitokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFortitokenMapInput)(nil)).Elem(), UserFortitokenMap{})
	pulumi.RegisterOutputType(UserFortitokenOutput{})
	pulumi.RegisterOutputType(UserFortitokenArrayOutput{})
	pulumi.RegisterOutputType(UserFortitokenMapOutput{})
}