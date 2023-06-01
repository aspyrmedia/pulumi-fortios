// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete devicemanager script for FortiManager.
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
//			_, err := fortios.NewFmgDevicemanagerScript(ctx, "test1", &fortios.FmgDevicemanagerScriptArgs{
//				Content:     pulumi.String("config system interface \n edit port3 \n	 set vdom \"root\"\n	 set ip 10.7.0.200 255.255.0.0 \n	 set allowaccess ping http https\n	 next \n end\n"),
//				Description: pulumi.String("description"),
//				Target:      pulumi.String("remote_device"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgDevicemanagerScript struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Script content, only cli script is supported now
	Content pulumi.StringOutput `pulumi:"content"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Script name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
	Target pulumi.StringPtrOutput `pulumi:"target"`
}

// NewFmgDevicemanagerScript registers a new resource with the given unique name, arguments, and options.
func NewFmgDevicemanagerScript(ctx *pulumi.Context,
	name string, args *FmgDevicemanagerScriptArgs, opts ...pulumi.ResourceOption) (*FmgDevicemanagerScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FmgDevicemanagerScript
	err := ctx.RegisterResource("fortios:index/fmgDevicemanagerScript:FmgDevicemanagerScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgDevicemanagerScript gets an existing FmgDevicemanagerScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgDevicemanagerScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgDevicemanagerScriptState, opts ...pulumi.ResourceOption) (*FmgDevicemanagerScript, error) {
	var resource FmgDevicemanagerScript
	err := ctx.ReadResource("fortios:index/fmgDevicemanagerScript:FmgDevicemanagerScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgDevicemanagerScript resources.
type fmgDevicemanagerScriptState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Script content, only cli script is supported now
	Content *string `pulumi:"content"`
	// Description.
	Description *string `pulumi:"description"`
	// Script name.
	Name *string `pulumi:"name"`
	// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
	Target *string `pulumi:"target"`
}

type FmgDevicemanagerScriptState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Script content, only cli script is supported now
	Content pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Script name.
	Name pulumi.StringPtrInput
	// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
	Target pulumi.StringPtrInput
}

func (FmgDevicemanagerScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgDevicemanagerScriptState)(nil)).Elem()
}

type fmgDevicemanagerScriptArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Script content, only cli script is supported now
	Content string `pulumi:"content"`
	// Description.
	Description *string `pulumi:"description"`
	// Script name.
	Name *string `pulumi:"name"`
	// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
	Target *string `pulumi:"target"`
}

// The set of arguments for constructing a FmgDevicemanagerScript resource.
type FmgDevicemanagerScriptArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Script content, only cli script is supported now
	Content pulumi.StringInput
	// Description.
	Description pulumi.StringPtrInput
	// Script name.
	Name pulumi.StringPtrInput
	// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
	Target pulumi.StringPtrInput
}

func (FmgDevicemanagerScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgDevicemanagerScriptArgs)(nil)).Elem()
}

type FmgDevicemanagerScriptInput interface {
	pulumi.Input

	ToFmgDevicemanagerScriptOutput() FmgDevicemanagerScriptOutput
	ToFmgDevicemanagerScriptOutputWithContext(ctx context.Context) FmgDevicemanagerScriptOutput
}

func (*FmgDevicemanagerScript) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgDevicemanagerScript)(nil)).Elem()
}

func (i *FmgDevicemanagerScript) ToFmgDevicemanagerScriptOutput() FmgDevicemanagerScriptOutput {
	return i.ToFmgDevicemanagerScriptOutputWithContext(context.Background())
}

func (i *FmgDevicemanagerScript) ToFmgDevicemanagerScriptOutputWithContext(ctx context.Context) FmgDevicemanagerScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerScriptOutput)
}

// FmgDevicemanagerScriptArrayInput is an input type that accepts FmgDevicemanagerScriptArray and FmgDevicemanagerScriptArrayOutput values.
// You can construct a concrete instance of `FmgDevicemanagerScriptArrayInput` via:
//
//	FmgDevicemanagerScriptArray{ FmgDevicemanagerScriptArgs{...} }
type FmgDevicemanagerScriptArrayInput interface {
	pulumi.Input

	ToFmgDevicemanagerScriptArrayOutput() FmgDevicemanagerScriptArrayOutput
	ToFmgDevicemanagerScriptArrayOutputWithContext(context.Context) FmgDevicemanagerScriptArrayOutput
}

type FmgDevicemanagerScriptArray []FmgDevicemanagerScriptInput

func (FmgDevicemanagerScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgDevicemanagerScript)(nil)).Elem()
}

func (i FmgDevicemanagerScriptArray) ToFmgDevicemanagerScriptArrayOutput() FmgDevicemanagerScriptArrayOutput {
	return i.ToFmgDevicemanagerScriptArrayOutputWithContext(context.Background())
}

func (i FmgDevicemanagerScriptArray) ToFmgDevicemanagerScriptArrayOutputWithContext(ctx context.Context) FmgDevicemanagerScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerScriptArrayOutput)
}

// FmgDevicemanagerScriptMapInput is an input type that accepts FmgDevicemanagerScriptMap and FmgDevicemanagerScriptMapOutput values.
// You can construct a concrete instance of `FmgDevicemanagerScriptMapInput` via:
//
//	FmgDevicemanagerScriptMap{ "key": FmgDevicemanagerScriptArgs{...} }
type FmgDevicemanagerScriptMapInput interface {
	pulumi.Input

	ToFmgDevicemanagerScriptMapOutput() FmgDevicemanagerScriptMapOutput
	ToFmgDevicemanagerScriptMapOutputWithContext(context.Context) FmgDevicemanagerScriptMapOutput
}

type FmgDevicemanagerScriptMap map[string]FmgDevicemanagerScriptInput

func (FmgDevicemanagerScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgDevicemanagerScript)(nil)).Elem()
}

func (i FmgDevicemanagerScriptMap) ToFmgDevicemanagerScriptMapOutput() FmgDevicemanagerScriptMapOutput {
	return i.ToFmgDevicemanagerScriptMapOutputWithContext(context.Background())
}

func (i FmgDevicemanagerScriptMap) ToFmgDevicemanagerScriptMapOutputWithContext(ctx context.Context) FmgDevicemanagerScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerScriptMapOutput)
}

type FmgDevicemanagerScriptOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgDevicemanagerScript)(nil)).Elem()
}

func (o FmgDevicemanagerScriptOutput) ToFmgDevicemanagerScriptOutput() FmgDevicemanagerScriptOutput {
	return o
}

func (o FmgDevicemanagerScriptOutput) ToFmgDevicemanagerScriptOutputWithContext(ctx context.Context) FmgDevicemanagerScriptOutput {
	return o
}

// ADOM name. default is 'root'.
func (o FmgDevicemanagerScriptOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerScript) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// Script content, only cli script is supported now
func (o FmgDevicemanagerScriptOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgDevicemanagerScript) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Description.
func (o FmgDevicemanagerScriptOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerScript) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Script name.
func (o FmgDevicemanagerScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgDevicemanagerScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
func (o FmgDevicemanagerScriptOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerScript) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

type FmgDevicemanagerScriptArrayOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgDevicemanagerScript)(nil)).Elem()
}

func (o FmgDevicemanagerScriptArrayOutput) ToFmgDevicemanagerScriptArrayOutput() FmgDevicemanagerScriptArrayOutput {
	return o
}

func (o FmgDevicemanagerScriptArrayOutput) ToFmgDevicemanagerScriptArrayOutputWithContext(ctx context.Context) FmgDevicemanagerScriptArrayOutput {
	return o
}

func (o FmgDevicemanagerScriptArrayOutput) Index(i pulumi.IntInput) FmgDevicemanagerScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgDevicemanagerScript {
		return vs[0].([]*FmgDevicemanagerScript)[vs[1].(int)]
	}).(FmgDevicemanagerScriptOutput)
}

type FmgDevicemanagerScriptMapOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgDevicemanagerScript)(nil)).Elem()
}

func (o FmgDevicemanagerScriptMapOutput) ToFmgDevicemanagerScriptMapOutput() FmgDevicemanagerScriptMapOutput {
	return o
}

func (o FmgDevicemanagerScriptMapOutput) ToFmgDevicemanagerScriptMapOutputWithContext(ctx context.Context) FmgDevicemanagerScriptMapOutput {
	return o
}

func (o FmgDevicemanagerScriptMapOutput) MapIndex(k pulumi.StringInput) FmgDevicemanagerScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgDevicemanagerScript {
		return vs[0].(map[string]*FmgDevicemanagerScript)[vs[1].(string)]
	}).(FmgDevicemanagerScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerScriptInput)(nil)).Elem(), &FmgDevicemanagerScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerScriptArrayInput)(nil)).Elem(), FmgDevicemanagerScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerScriptMapInput)(nil)).Elem(), FmgDevicemanagerScriptMap{})
	pulumi.RegisterOutputType(FmgDevicemanagerScriptOutput{})
	pulumi.RegisterOutputType(FmgDevicemanagerScriptArrayOutput{})
	pulumi.RegisterOutputType(FmgDevicemanagerScriptMapOutput{})
}