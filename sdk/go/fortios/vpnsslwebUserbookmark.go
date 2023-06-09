// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure SSL VPN user bookmark.
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
//			_, err := fortios.NewVpnsslwebUserbookmark(ctx, "trname", &fortios.VpnsslwebUserbookmarkArgs{
//				CustomLang: pulumi.String("big5"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # VpnSslWeb UserBookmark can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/vpnsslwebUserbookmark:VpnsslwebUserbookmark labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/vpnsslwebUserbookmark:VpnsslwebUserbookmark labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type VpnsslwebUserbookmark struct {
	pulumi.CustomResourceState

	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks VpnsslwebUserbookmarkBookmarkArrayOutput `pulumi:"bookmarks"`
	// Personal language.
	CustomLang pulumi.StringOutput `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// User and group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnsslwebUserbookmark registers a new resource with the given unique name, arguments, and options.
func NewVpnsslwebUserbookmark(ctx *pulumi.Context,
	name string, args *VpnsslwebUserbookmarkArgs, opts ...pulumi.ResourceOption) (*VpnsslwebUserbookmark, error) {
	if args == nil {
		args = &VpnsslwebUserbookmarkArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnsslwebUserbookmark
	err := ctx.RegisterResource("fortios:index/vpnsslwebUserbookmark:VpnsslwebUserbookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnsslwebUserbookmark gets an existing VpnsslwebUserbookmark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnsslwebUserbookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnsslwebUserbookmarkState, opts ...pulumi.ResourceOption) (*VpnsslwebUserbookmark, error) {
	var resource VpnsslwebUserbookmark
	err := ctx.ReadResource("fortios:index/vpnsslwebUserbookmark:VpnsslwebUserbookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnsslwebUserbookmark resources.
type vpnsslwebUserbookmarkState struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks []VpnsslwebUserbookmarkBookmark `pulumi:"bookmarks"`
	// Personal language.
	CustomLang *string `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// User and group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpnsslwebUserbookmarkState struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks VpnsslwebUserbookmarkBookmarkArrayInput
	// Personal language.
	CustomLang pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// User and group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnsslwebUserbookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnsslwebUserbookmarkState)(nil)).Elem()
}

type vpnsslwebUserbookmarkArgs struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks []VpnsslwebUserbookmarkBookmark `pulumi:"bookmarks"`
	// Personal language.
	CustomLang *string `pulumi:"customLang"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// User and group name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnsslwebUserbookmark resource.
type VpnsslwebUserbookmarkArgs struct {
	// Bookmark table. The structure of `bookmarks` block is documented below.
	Bookmarks VpnsslwebUserbookmarkBookmarkArrayInput
	// Personal language.
	CustomLang pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// User and group name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnsslwebUserbookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnsslwebUserbookmarkArgs)(nil)).Elem()
}

type VpnsslwebUserbookmarkInput interface {
	pulumi.Input

	ToVpnsslwebUserbookmarkOutput() VpnsslwebUserbookmarkOutput
	ToVpnsslwebUserbookmarkOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkOutput
}

func (*VpnsslwebUserbookmark) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnsslwebUserbookmark)(nil)).Elem()
}

func (i *VpnsslwebUserbookmark) ToVpnsslwebUserbookmarkOutput() VpnsslwebUserbookmarkOutput {
	return i.ToVpnsslwebUserbookmarkOutputWithContext(context.Background())
}

func (i *VpnsslwebUserbookmark) ToVpnsslwebUserbookmarkOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnsslwebUserbookmarkOutput)
}

// VpnsslwebUserbookmarkArrayInput is an input type that accepts VpnsslwebUserbookmarkArray and VpnsslwebUserbookmarkArrayOutput values.
// You can construct a concrete instance of `VpnsslwebUserbookmarkArrayInput` via:
//
//	VpnsslwebUserbookmarkArray{ VpnsslwebUserbookmarkArgs{...} }
type VpnsslwebUserbookmarkArrayInput interface {
	pulumi.Input

	ToVpnsslwebUserbookmarkArrayOutput() VpnsslwebUserbookmarkArrayOutput
	ToVpnsslwebUserbookmarkArrayOutputWithContext(context.Context) VpnsslwebUserbookmarkArrayOutput
}

type VpnsslwebUserbookmarkArray []VpnsslwebUserbookmarkInput

func (VpnsslwebUserbookmarkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnsslwebUserbookmark)(nil)).Elem()
}

func (i VpnsslwebUserbookmarkArray) ToVpnsslwebUserbookmarkArrayOutput() VpnsslwebUserbookmarkArrayOutput {
	return i.ToVpnsslwebUserbookmarkArrayOutputWithContext(context.Background())
}

func (i VpnsslwebUserbookmarkArray) ToVpnsslwebUserbookmarkArrayOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnsslwebUserbookmarkArrayOutput)
}

// VpnsslwebUserbookmarkMapInput is an input type that accepts VpnsslwebUserbookmarkMap and VpnsslwebUserbookmarkMapOutput values.
// You can construct a concrete instance of `VpnsslwebUserbookmarkMapInput` via:
//
//	VpnsslwebUserbookmarkMap{ "key": VpnsslwebUserbookmarkArgs{...} }
type VpnsslwebUserbookmarkMapInput interface {
	pulumi.Input

	ToVpnsslwebUserbookmarkMapOutput() VpnsslwebUserbookmarkMapOutput
	ToVpnsslwebUserbookmarkMapOutputWithContext(context.Context) VpnsslwebUserbookmarkMapOutput
}

type VpnsslwebUserbookmarkMap map[string]VpnsslwebUserbookmarkInput

func (VpnsslwebUserbookmarkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnsslwebUserbookmark)(nil)).Elem()
}

func (i VpnsslwebUserbookmarkMap) ToVpnsslwebUserbookmarkMapOutput() VpnsslwebUserbookmarkMapOutput {
	return i.ToVpnsslwebUserbookmarkMapOutputWithContext(context.Background())
}

func (i VpnsslwebUserbookmarkMap) ToVpnsslwebUserbookmarkMapOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnsslwebUserbookmarkMapOutput)
}

type VpnsslwebUserbookmarkOutput struct{ *pulumi.OutputState }

func (VpnsslwebUserbookmarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnsslwebUserbookmark)(nil)).Elem()
}

func (o VpnsslwebUserbookmarkOutput) ToVpnsslwebUserbookmarkOutput() VpnsslwebUserbookmarkOutput {
	return o
}

func (o VpnsslwebUserbookmarkOutput) ToVpnsslwebUserbookmarkOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkOutput {
	return o
}

// Bookmark table. The structure of `bookmarks` block is documented below.
func (o VpnsslwebUserbookmarkOutput) Bookmarks() VpnsslwebUserbookmarkBookmarkArrayOutput {
	return o.ApplyT(func(v *VpnsslwebUserbookmark) VpnsslwebUserbookmarkBookmarkArrayOutput { return v.Bookmarks }).(VpnsslwebUserbookmarkBookmarkArrayOutput)
}

// Personal language.
func (o VpnsslwebUserbookmarkOutput) CustomLang() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnsslwebUserbookmark) pulumi.StringOutput { return v.CustomLang }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VpnsslwebUserbookmarkOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnsslwebUserbookmark) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// User and group name.
func (o VpnsslwebUserbookmarkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnsslwebUserbookmark) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VpnsslwebUserbookmarkOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnsslwebUserbookmark) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnsslwebUserbookmarkArrayOutput struct{ *pulumi.OutputState }

func (VpnsslwebUserbookmarkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnsslwebUserbookmark)(nil)).Elem()
}

func (o VpnsslwebUserbookmarkArrayOutput) ToVpnsslwebUserbookmarkArrayOutput() VpnsslwebUserbookmarkArrayOutput {
	return o
}

func (o VpnsslwebUserbookmarkArrayOutput) ToVpnsslwebUserbookmarkArrayOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkArrayOutput {
	return o
}

func (o VpnsslwebUserbookmarkArrayOutput) Index(i pulumi.IntInput) VpnsslwebUserbookmarkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnsslwebUserbookmark {
		return vs[0].([]*VpnsslwebUserbookmark)[vs[1].(int)]
	}).(VpnsslwebUserbookmarkOutput)
}

type VpnsslwebUserbookmarkMapOutput struct{ *pulumi.OutputState }

func (VpnsslwebUserbookmarkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnsslwebUserbookmark)(nil)).Elem()
}

func (o VpnsslwebUserbookmarkMapOutput) ToVpnsslwebUserbookmarkMapOutput() VpnsslwebUserbookmarkMapOutput {
	return o
}

func (o VpnsslwebUserbookmarkMapOutput) ToVpnsslwebUserbookmarkMapOutputWithContext(ctx context.Context) VpnsslwebUserbookmarkMapOutput {
	return o
}

func (o VpnsslwebUserbookmarkMapOutput) MapIndex(k pulumi.StringInput) VpnsslwebUserbookmarkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnsslwebUserbookmark {
		return vs[0].(map[string]*VpnsslwebUserbookmark)[vs[1].(string)]
	}).(VpnsslwebUserbookmarkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnsslwebUserbookmarkInput)(nil)).Elem(), &VpnsslwebUserbookmark{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnsslwebUserbookmarkArrayInput)(nil)).Elem(), VpnsslwebUserbookmarkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnsslwebUserbookmarkMapInput)(nil)).Elem(), VpnsslwebUserbookmarkMap{})
	pulumi.RegisterOutputType(VpnsslwebUserbookmarkOutput{})
	pulumi.RegisterOutputType(VpnsslwebUserbookmarkArrayOutput{})
	pulumi.RegisterOutputType(VpnsslwebUserbookmarkMapOutput{})
}
