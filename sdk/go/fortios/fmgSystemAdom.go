// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete system adom for FortiManager.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewFmgSystemAdom(ctx, "test1", &fortios.FmgSystemAdomArgs{
//				ActionWhenConflictsOccurDuringPolicyCheck:  pulumi.String("Continue"),
//				AutoPushPolicyPackagesWhenDeviceBackOnline: pulumi.String("Enable"),
//				CentralManagementFortiap:                   pulumi.Bool(true),
//				CentralManagementSdwan:                     pulumi.Bool(false),
//				CentralManagementVpn:                       pulumi.Bool(false),
//				Mode:                                       pulumi.String("Normal"),
//				PerformPolicyCheckBeforeEveryInstall:       pulumi.Bool(true),
//				Status:                                     pulumi.Int(1),
//				Type:                                       pulumi.String("FortiCarrier"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgSystemAdom struct {
	pulumi.CustomResourceState

	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrOutput `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrOutput `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrOutput `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrOutput `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn pulumi.BoolPtrOutput `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Administrative Domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrOutput `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFmgSystemAdom registers a new resource with the given unique name, arguments, and options.
func NewFmgSystemAdom(ctx *pulumi.Context,
	name string, args *FmgSystemAdomArgs, opts ...pulumi.ResourceOption) (*FmgSystemAdom, error) {
	if args == nil {
		args = &FmgSystemAdomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FmgSystemAdom
	err := ctx.RegisterResource("fortios:index/fmgSystemAdom:FmgSystemAdom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgSystemAdom gets an existing FmgSystemAdom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgSystemAdom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgSystemAdomState, opts ...pulumi.ResourceOption) (*FmgSystemAdom, error) {
	var resource FmgSystemAdom
	err := ctx.ReadResource("fortios:index/fmgSystemAdom:FmgSystemAdom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgSystemAdom resources.
type fmgSystemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

type FmgSystemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (FmgSystemAdomState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgSystemAdomState)(nil)).Elem()
}

type fmgSystemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FmgSystemAdom resource.
type FmgSystemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (FmgSystemAdomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgSystemAdomArgs)(nil)).Elem()
}

type FmgSystemAdomInput interface {
	pulumi.Input

	ToFmgSystemAdomOutput() FmgSystemAdomOutput
	ToFmgSystemAdomOutputWithContext(ctx context.Context) FmgSystemAdomOutput
}

func (*FmgSystemAdom) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgSystemAdom)(nil)).Elem()
}

func (i *FmgSystemAdom) ToFmgSystemAdomOutput() FmgSystemAdomOutput {
	return i.ToFmgSystemAdomOutputWithContext(context.Background())
}

func (i *FmgSystemAdom) ToFmgSystemAdomOutputWithContext(ctx context.Context) FmgSystemAdomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdomOutput)
}

// FmgSystemAdomArrayInput is an input type that accepts FmgSystemAdomArray and FmgSystemAdomArrayOutput values.
// You can construct a concrete instance of `FmgSystemAdomArrayInput` via:
//
//	FmgSystemAdomArray{ FmgSystemAdomArgs{...} }
type FmgSystemAdomArrayInput interface {
	pulumi.Input

	ToFmgSystemAdomArrayOutput() FmgSystemAdomArrayOutput
	ToFmgSystemAdomArrayOutputWithContext(context.Context) FmgSystemAdomArrayOutput
}

type FmgSystemAdomArray []FmgSystemAdomInput

func (FmgSystemAdomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgSystemAdom)(nil)).Elem()
}

func (i FmgSystemAdomArray) ToFmgSystemAdomArrayOutput() FmgSystemAdomArrayOutput {
	return i.ToFmgSystemAdomArrayOutputWithContext(context.Background())
}

func (i FmgSystemAdomArray) ToFmgSystemAdomArrayOutputWithContext(ctx context.Context) FmgSystemAdomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdomArrayOutput)
}

// FmgSystemAdomMapInput is an input type that accepts FmgSystemAdomMap and FmgSystemAdomMapOutput values.
// You can construct a concrete instance of `FmgSystemAdomMapInput` via:
//
//	FmgSystemAdomMap{ "key": FmgSystemAdomArgs{...} }
type FmgSystemAdomMapInput interface {
	pulumi.Input

	ToFmgSystemAdomMapOutput() FmgSystemAdomMapOutput
	ToFmgSystemAdomMapOutputWithContext(context.Context) FmgSystemAdomMapOutput
}

type FmgSystemAdomMap map[string]FmgSystemAdomInput

func (FmgSystemAdomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgSystemAdom)(nil)).Elem()
}

func (i FmgSystemAdomMap) ToFmgSystemAdomMapOutput() FmgSystemAdomMapOutput {
	return i.ToFmgSystemAdomMapOutputWithContext(context.Background())
}

func (i FmgSystemAdomMap) ToFmgSystemAdomMapOutputWithContext(ctx context.Context) FmgSystemAdomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgSystemAdomMapOutput)
}

type FmgSystemAdomOutput struct{ *pulumi.OutputState }

func (FmgSystemAdomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgSystemAdom)(nil)).Elem()
}

func (o FmgSystemAdomOutput) ToFmgSystemAdomOutput() FmgSystemAdomOutput {
	return o
}

func (o FmgSystemAdomOutput) ToFmgSystemAdomOutputWithContext(ctx context.Context) FmgSystemAdomOutput {
	return o
}

// True or False.
func (o FmgSystemAdomOutput) ActionWhenConflictsOccurDuringPolicyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.StringPtrOutput { return v.ActionWhenConflictsOccurDuringPolicyCheck }).(pulumi.StringPtrOutput)
}

// True or False.
func (o FmgSystemAdomOutput) AutoPushPolicyPackagesWhenDeviceBackOnline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.StringPtrOutput { return v.AutoPushPolicyPackagesWhenDeviceBackOnline }).(pulumi.StringPtrOutput)
}

// True or False.
func (o FmgSystemAdomOutput) CentralManagementFortiap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementFortiap }).(pulumi.BoolPtrOutput)
}

// True or False.
func (o FmgSystemAdomOutput) CentralManagementSdwan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementSdwan }).(pulumi.BoolPtrOutput)
}

// True or False.
func (o FmgSystemAdomOutput) CentralManagementVpn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.BoolPtrOutput { return v.CentralManagementVpn }).(pulumi.BoolPtrOutput)
}

// Adom mode: Normal or Backup.
func (o FmgSystemAdomOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Administrative Domain name.
func (o FmgSystemAdomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// True or False.
func (o FmgSystemAdomOutput) PerformPolicyCheckBeforeEveryInstall() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.BoolPtrOutput { return v.PerformPolicyCheckBeforeEveryInstall }).(pulumi.BoolPtrOutput)
}

// Adom status. 0 means off and 1 means on.
func (o FmgSystemAdomOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
func (o FmgSystemAdomOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgSystemAdom) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type FmgSystemAdomArrayOutput struct{ *pulumi.OutputState }

func (FmgSystemAdomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgSystemAdom)(nil)).Elem()
}

func (o FmgSystemAdomArrayOutput) ToFmgSystemAdomArrayOutput() FmgSystemAdomArrayOutput {
	return o
}

func (o FmgSystemAdomArrayOutput) ToFmgSystemAdomArrayOutputWithContext(ctx context.Context) FmgSystemAdomArrayOutput {
	return o
}

func (o FmgSystemAdomArrayOutput) Index(i pulumi.IntInput) FmgSystemAdomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgSystemAdom {
		return vs[0].([]*FmgSystemAdom)[vs[1].(int)]
	}).(FmgSystemAdomOutput)
}

type FmgSystemAdomMapOutput struct{ *pulumi.OutputState }

func (FmgSystemAdomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgSystemAdom)(nil)).Elem()
}

func (o FmgSystemAdomMapOutput) ToFmgSystemAdomMapOutput() FmgSystemAdomMapOutput {
	return o
}

func (o FmgSystemAdomMapOutput) ToFmgSystemAdomMapOutputWithContext(ctx context.Context) FmgSystemAdomMapOutput {
	return o
}

func (o FmgSystemAdomMapOutput) MapIndex(k pulumi.StringInput) FmgSystemAdomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgSystemAdom {
		return vs[0].(map[string]*FmgSystemAdom)[vs[1].(string)]
	}).(FmgSystemAdomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdomInput)(nil)).Elem(), &FmgSystemAdom{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdomArrayInput)(nil)).Elem(), FmgSystemAdomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgSystemAdomMapInput)(nil)).Elem(), FmgSystemAdomMap{})
	pulumi.RegisterOutputType(FmgSystemAdomOutput{})
	pulumi.RegisterOutputType(FmgSystemAdomArrayOutput{})
	pulumi.RegisterOutputType(FmgSystemAdomMapOutput{})
}
