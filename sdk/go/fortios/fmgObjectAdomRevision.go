// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete object adom revision for FortiManager.
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
//			_, err := fortios.NewFmgObjectAdomRevision(ctx, "test1", &fortios.FmgObjectAdomRevisionArgs{
//				CreatedBy:   pulumi.String("fortinet"),
//				Description: pulumi.String("adom revision"),
//				Locked:      pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgObjectAdomRevision struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrOutput `pulumi:"locked"`
	// Adom revision name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFmgObjectAdomRevision registers a new resource with the given unique name, arguments, and options.
func NewFmgObjectAdomRevision(ctx *pulumi.Context,
	name string, args *FmgObjectAdomRevisionArgs, opts ...pulumi.ResourceOption) (*FmgObjectAdomRevision, error) {
	if args == nil {
		args = &FmgObjectAdomRevisionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FmgObjectAdomRevision
	err := ctx.RegisterResource("fortios:index/fmgObjectAdomRevision:FmgObjectAdomRevision", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgObjectAdomRevision gets an existing FmgObjectAdomRevision resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgObjectAdomRevision(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgObjectAdomRevisionState, opts ...pulumi.ResourceOption) (*FmgObjectAdomRevision, error) {
	var resource FmgObjectAdomRevision
	err := ctx.ReadResource("fortios:index/fmgObjectAdomRevision:FmgObjectAdomRevision", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgObjectAdomRevision resources.
type fmgObjectAdomRevisionState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy *string `pulumi:"createdBy"`
	// Description.
	Description *string `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked *int `pulumi:"locked"`
	// Adom revision name.
	Name *string `pulumi:"name"`
}

type FmgObjectAdomRevisionState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrInput
	// Adom revision name.
	Name pulumi.StringPtrInput
}

func (FmgObjectAdomRevisionState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgObjectAdomRevisionState)(nil)).Elem()
}

type fmgObjectAdomRevisionArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Who created this adom revision.
	CreatedBy *string `pulumi:"createdBy"`
	// Description.
	Description *string `pulumi:"description"`
	// lock. 0 means unlock and 1 means locked.
	Locked *int `pulumi:"locked"`
	// Adom revision name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FmgObjectAdomRevision resource.
type FmgObjectAdomRevisionArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Who created this adom revision.
	CreatedBy pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// lock. 0 means unlock and 1 means locked.
	Locked pulumi.IntPtrInput
	// Adom revision name.
	Name pulumi.StringPtrInput
}

func (FmgObjectAdomRevisionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgObjectAdomRevisionArgs)(nil)).Elem()
}

type FmgObjectAdomRevisionInput interface {
	pulumi.Input

	ToFmgObjectAdomRevisionOutput() FmgObjectAdomRevisionOutput
	ToFmgObjectAdomRevisionOutputWithContext(ctx context.Context) FmgObjectAdomRevisionOutput
}

func (*FmgObjectAdomRevision) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgObjectAdomRevision)(nil)).Elem()
}

func (i *FmgObjectAdomRevision) ToFmgObjectAdomRevisionOutput() FmgObjectAdomRevisionOutput {
	return i.ToFmgObjectAdomRevisionOutputWithContext(context.Background())
}

func (i *FmgObjectAdomRevision) ToFmgObjectAdomRevisionOutputWithContext(ctx context.Context) FmgObjectAdomRevisionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgObjectAdomRevisionOutput)
}

// FmgObjectAdomRevisionArrayInput is an input type that accepts FmgObjectAdomRevisionArray and FmgObjectAdomRevisionArrayOutput values.
// You can construct a concrete instance of `FmgObjectAdomRevisionArrayInput` via:
//
//	FmgObjectAdomRevisionArray{ FmgObjectAdomRevisionArgs{...} }
type FmgObjectAdomRevisionArrayInput interface {
	pulumi.Input

	ToFmgObjectAdomRevisionArrayOutput() FmgObjectAdomRevisionArrayOutput
	ToFmgObjectAdomRevisionArrayOutputWithContext(context.Context) FmgObjectAdomRevisionArrayOutput
}

type FmgObjectAdomRevisionArray []FmgObjectAdomRevisionInput

func (FmgObjectAdomRevisionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgObjectAdomRevision)(nil)).Elem()
}

func (i FmgObjectAdomRevisionArray) ToFmgObjectAdomRevisionArrayOutput() FmgObjectAdomRevisionArrayOutput {
	return i.ToFmgObjectAdomRevisionArrayOutputWithContext(context.Background())
}

func (i FmgObjectAdomRevisionArray) ToFmgObjectAdomRevisionArrayOutputWithContext(ctx context.Context) FmgObjectAdomRevisionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgObjectAdomRevisionArrayOutput)
}

// FmgObjectAdomRevisionMapInput is an input type that accepts FmgObjectAdomRevisionMap and FmgObjectAdomRevisionMapOutput values.
// You can construct a concrete instance of `FmgObjectAdomRevisionMapInput` via:
//
//	FmgObjectAdomRevisionMap{ "key": FmgObjectAdomRevisionArgs{...} }
type FmgObjectAdomRevisionMapInput interface {
	pulumi.Input

	ToFmgObjectAdomRevisionMapOutput() FmgObjectAdomRevisionMapOutput
	ToFmgObjectAdomRevisionMapOutputWithContext(context.Context) FmgObjectAdomRevisionMapOutput
}

type FmgObjectAdomRevisionMap map[string]FmgObjectAdomRevisionInput

func (FmgObjectAdomRevisionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgObjectAdomRevision)(nil)).Elem()
}

func (i FmgObjectAdomRevisionMap) ToFmgObjectAdomRevisionMapOutput() FmgObjectAdomRevisionMapOutput {
	return i.ToFmgObjectAdomRevisionMapOutputWithContext(context.Background())
}

func (i FmgObjectAdomRevisionMap) ToFmgObjectAdomRevisionMapOutputWithContext(ctx context.Context) FmgObjectAdomRevisionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgObjectAdomRevisionMapOutput)
}

type FmgObjectAdomRevisionOutput struct{ *pulumi.OutputState }

func (FmgObjectAdomRevisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgObjectAdomRevision)(nil)).Elem()
}

func (o FmgObjectAdomRevisionOutput) ToFmgObjectAdomRevisionOutput() FmgObjectAdomRevisionOutput {
	return o
}

func (o FmgObjectAdomRevisionOutput) ToFmgObjectAdomRevisionOutputWithContext(ctx context.Context) FmgObjectAdomRevisionOutput {
	return o
}

// ADOM name. default is 'root'.
func (o FmgObjectAdomRevisionOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgObjectAdomRevision) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// Who created this adom revision.
func (o FmgObjectAdomRevisionOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgObjectAdomRevision) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// Description.
func (o FmgObjectAdomRevisionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgObjectAdomRevision) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// lock. 0 means unlock and 1 means locked.
func (o FmgObjectAdomRevisionOutput) Locked() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgObjectAdomRevision) pulumi.IntPtrOutput { return v.Locked }).(pulumi.IntPtrOutput)
}

// Adom revision name.
func (o FmgObjectAdomRevisionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgObjectAdomRevision) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FmgObjectAdomRevisionArrayOutput struct{ *pulumi.OutputState }

func (FmgObjectAdomRevisionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgObjectAdomRevision)(nil)).Elem()
}

func (o FmgObjectAdomRevisionArrayOutput) ToFmgObjectAdomRevisionArrayOutput() FmgObjectAdomRevisionArrayOutput {
	return o
}

func (o FmgObjectAdomRevisionArrayOutput) ToFmgObjectAdomRevisionArrayOutputWithContext(ctx context.Context) FmgObjectAdomRevisionArrayOutput {
	return o
}

func (o FmgObjectAdomRevisionArrayOutput) Index(i pulumi.IntInput) FmgObjectAdomRevisionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgObjectAdomRevision {
		return vs[0].([]*FmgObjectAdomRevision)[vs[1].(int)]
	}).(FmgObjectAdomRevisionOutput)
}

type FmgObjectAdomRevisionMapOutput struct{ *pulumi.OutputState }

func (FmgObjectAdomRevisionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgObjectAdomRevision)(nil)).Elem()
}

func (o FmgObjectAdomRevisionMapOutput) ToFmgObjectAdomRevisionMapOutput() FmgObjectAdomRevisionMapOutput {
	return o
}

func (o FmgObjectAdomRevisionMapOutput) ToFmgObjectAdomRevisionMapOutputWithContext(ctx context.Context) FmgObjectAdomRevisionMapOutput {
	return o
}

func (o FmgObjectAdomRevisionMapOutput) MapIndex(k pulumi.StringInput) FmgObjectAdomRevisionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgObjectAdomRevision {
		return vs[0].(map[string]*FmgObjectAdomRevision)[vs[1].(string)]
	}).(FmgObjectAdomRevisionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgObjectAdomRevisionInput)(nil)).Elem(), &FmgObjectAdomRevision{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgObjectAdomRevisionArrayInput)(nil)).Elem(), FmgObjectAdomRevisionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgObjectAdomRevisionMapInput)(nil)).Elem(), FmgObjectAdomRevisionMap{})
	pulumi.RegisterOutputType(FmgObjectAdomRevisionOutput{})
	pulumi.RegisterOutputType(FmgObjectAdomRevisionArrayOutput{})
	pulumi.RegisterOutputType(FmgObjectAdomRevisionMapOutput{})
}