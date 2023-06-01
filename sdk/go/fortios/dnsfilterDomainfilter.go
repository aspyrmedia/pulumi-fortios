// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DNS domain filters.
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
//			_, err := fortios.NewDnsfilterDomainfilter(ctx, "trname", &fortios.DnsfilterDomainfilterArgs{
//				Entries: fortios.DnsfilterDomainfilterEntryArray{
//					&fortios.DnsfilterDomainfilterEntryArgs{
//						Action: pulumi.String("block"),
//						Domain: pulumi.String("bac.com"),
//						Id:     pulumi.Int(1),
//						Status: pulumi.String("enable"),
//						Type:   pulumi.String("simple"),
//					},
//				},
//				Fosid: pulumi.Int(1),
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
// # Dnsfilter DomainFilter can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/dnsfilterDomainfilter:DnsfilterDomainfilter labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/dnsfilterDomainfilter:DnsfilterDomainfilter labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type DnsfilterDomainfilter struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// DNS domain filter entries. The structure of `entries` block is documented below.
	Entries DnsfilterDomainfilterEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDnsfilterDomainfilter registers a new resource with the given unique name, arguments, and options.
func NewDnsfilterDomainfilter(ctx *pulumi.Context,
	name string, args *DnsfilterDomainfilterArgs, opts ...pulumi.ResourceOption) (*DnsfilterDomainfilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DnsfilterDomainfilter
	err := ctx.RegisterResource("fortios:index/dnsfilterDomainfilter:DnsfilterDomainfilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsfilterDomainfilter gets an existing DnsfilterDomainfilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsfilterDomainfilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsfilterDomainfilterState, opts ...pulumi.ResourceOption) (*DnsfilterDomainfilter, error) {
	var resource DnsfilterDomainfilter
	err := ctx.ReadResource("fortios:index/dnsfilterDomainfilter:DnsfilterDomainfilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsfilterDomainfilter resources.
type dnsfilterDomainfilterState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DNS domain filter entries. The structure of `entries` block is documented below.
	Entries []DnsfilterDomainfilterEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DnsfilterDomainfilterState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DNS domain filter entries. The structure of `entries` block is documented below.
	Entries DnsfilterDomainfilterEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnsfilterDomainfilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsfilterDomainfilterState)(nil)).Elem()
}

type dnsfilterDomainfilterArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// DNS domain filter entries. The structure of `entries` block is documented below.
	Entries []DnsfilterDomainfilterEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DnsfilterDomainfilter resource.
type DnsfilterDomainfilterArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// DNS domain filter entries. The structure of `entries` block is documented below.
	Entries DnsfilterDomainfilterEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DnsfilterDomainfilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsfilterDomainfilterArgs)(nil)).Elem()
}

type DnsfilterDomainfilterInput interface {
	pulumi.Input

	ToDnsfilterDomainfilterOutput() DnsfilterDomainfilterOutput
	ToDnsfilterDomainfilterOutputWithContext(ctx context.Context) DnsfilterDomainfilterOutput
}

func (*DnsfilterDomainfilter) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsfilterDomainfilter)(nil)).Elem()
}

func (i *DnsfilterDomainfilter) ToDnsfilterDomainfilterOutput() DnsfilterDomainfilterOutput {
	return i.ToDnsfilterDomainfilterOutputWithContext(context.Background())
}

func (i *DnsfilterDomainfilter) ToDnsfilterDomainfilterOutputWithContext(ctx context.Context) DnsfilterDomainfilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterDomainfilterOutput)
}

// DnsfilterDomainfilterArrayInput is an input type that accepts DnsfilterDomainfilterArray and DnsfilterDomainfilterArrayOutput values.
// You can construct a concrete instance of `DnsfilterDomainfilterArrayInput` via:
//
//	DnsfilterDomainfilterArray{ DnsfilterDomainfilterArgs{...} }
type DnsfilterDomainfilterArrayInput interface {
	pulumi.Input

	ToDnsfilterDomainfilterArrayOutput() DnsfilterDomainfilterArrayOutput
	ToDnsfilterDomainfilterArrayOutputWithContext(context.Context) DnsfilterDomainfilterArrayOutput
}

type DnsfilterDomainfilterArray []DnsfilterDomainfilterInput

func (DnsfilterDomainfilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsfilterDomainfilter)(nil)).Elem()
}

func (i DnsfilterDomainfilterArray) ToDnsfilterDomainfilterArrayOutput() DnsfilterDomainfilterArrayOutput {
	return i.ToDnsfilterDomainfilterArrayOutputWithContext(context.Background())
}

func (i DnsfilterDomainfilterArray) ToDnsfilterDomainfilterArrayOutputWithContext(ctx context.Context) DnsfilterDomainfilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterDomainfilterArrayOutput)
}

// DnsfilterDomainfilterMapInput is an input type that accepts DnsfilterDomainfilterMap and DnsfilterDomainfilterMapOutput values.
// You can construct a concrete instance of `DnsfilterDomainfilterMapInput` via:
//
//	DnsfilterDomainfilterMap{ "key": DnsfilterDomainfilterArgs{...} }
type DnsfilterDomainfilterMapInput interface {
	pulumi.Input

	ToDnsfilterDomainfilterMapOutput() DnsfilterDomainfilterMapOutput
	ToDnsfilterDomainfilterMapOutputWithContext(context.Context) DnsfilterDomainfilterMapOutput
}

type DnsfilterDomainfilterMap map[string]DnsfilterDomainfilterInput

func (DnsfilterDomainfilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsfilterDomainfilter)(nil)).Elem()
}

func (i DnsfilterDomainfilterMap) ToDnsfilterDomainfilterMapOutput() DnsfilterDomainfilterMapOutput {
	return i.ToDnsfilterDomainfilterMapOutputWithContext(context.Background())
}

func (i DnsfilterDomainfilterMap) ToDnsfilterDomainfilterMapOutputWithContext(ctx context.Context) DnsfilterDomainfilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterDomainfilterMapOutput)
}

type DnsfilterDomainfilterOutput struct{ *pulumi.OutputState }

func (DnsfilterDomainfilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsfilterDomainfilter)(nil)).Elem()
}

func (o DnsfilterDomainfilterOutput) ToDnsfilterDomainfilterOutput() DnsfilterDomainfilterOutput {
	return o
}

func (o DnsfilterDomainfilterOutput) ToDnsfilterDomainfilterOutputWithContext(ctx context.Context) DnsfilterDomainfilterOutput {
	return o
}

// Optional comments.
func (o DnsfilterDomainfilterOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DnsfilterDomainfilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// DNS domain filter entries. The structure of `entries` block is documented below.
func (o DnsfilterDomainfilterOutput) Entries() DnsfilterDomainfilterEntryArrayOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) DnsfilterDomainfilterEntryArrayOutput { return v.Entries }).(DnsfilterDomainfilterEntryArrayOutput)
}

// ID.
func (o DnsfilterDomainfilterOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o DnsfilterDomainfilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DnsfilterDomainfilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsfilterDomainfilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DnsfilterDomainfilterArrayOutput struct{ *pulumi.OutputState }

func (DnsfilterDomainfilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsfilterDomainfilter)(nil)).Elem()
}

func (o DnsfilterDomainfilterArrayOutput) ToDnsfilterDomainfilterArrayOutput() DnsfilterDomainfilterArrayOutput {
	return o
}

func (o DnsfilterDomainfilterArrayOutput) ToDnsfilterDomainfilterArrayOutputWithContext(ctx context.Context) DnsfilterDomainfilterArrayOutput {
	return o
}

func (o DnsfilterDomainfilterArrayOutput) Index(i pulumi.IntInput) DnsfilterDomainfilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsfilterDomainfilter {
		return vs[0].([]*DnsfilterDomainfilter)[vs[1].(int)]
	}).(DnsfilterDomainfilterOutput)
}

type DnsfilterDomainfilterMapOutput struct{ *pulumi.OutputState }

func (DnsfilterDomainfilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsfilterDomainfilter)(nil)).Elem()
}

func (o DnsfilterDomainfilterMapOutput) ToDnsfilterDomainfilterMapOutput() DnsfilterDomainfilterMapOutput {
	return o
}

func (o DnsfilterDomainfilterMapOutput) ToDnsfilterDomainfilterMapOutputWithContext(ctx context.Context) DnsfilterDomainfilterMapOutput {
	return o
}

func (o DnsfilterDomainfilterMapOutput) MapIndex(k pulumi.StringInput) DnsfilterDomainfilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsfilterDomainfilter {
		return vs[0].(map[string]*DnsfilterDomainfilter)[vs[1].(string)]
	}).(DnsfilterDomainfilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterDomainfilterInput)(nil)).Elem(), &DnsfilterDomainfilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterDomainfilterArrayInput)(nil)).Elem(), DnsfilterDomainfilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterDomainfilterMapInput)(nil)).Elem(), DnsfilterDomainfilterMap{})
	pulumi.RegisterOutputType(DnsfilterDomainfilterOutput{})
	pulumi.RegisterOutputType(DnsfilterDomainfilterArrayOutput{})
	pulumi.RegisterOutputType(DnsfilterDomainfilterMapOutput{})
}