// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP to MAC address pairs in the IP/MAC binding table.
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
//			_, err := fortios.NewFirewallipmacbindingTable(ctx, "trname", &fortios.FirewallipmacbindingTableArgs{
//				Ip:     pulumi.String("1.1.1.1"),
//				Mac:    pulumi.String("00:01:6c:06:a6:29"),
//				SeqNum: pulumi.Int(1),
//				Status: pulumi.String("disable"),
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
// # FirewallIpmacbinding Table can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallipmacbindingTable:FirewallipmacbindingTable labelname {{seq_num}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallipmacbindingTable:FirewallipmacbindingTable labelname {{seq_num}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallipmacbindingTable struct {
	pulumi.CustomResourceState

	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringOutput `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name pulumi.StringOutput `pulumi:"name"`
	// Entry number.
	SeqNum pulumi.IntOutput `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallipmacbindingTable registers a new resource with the given unique name, arguments, and options.
func NewFirewallipmacbindingTable(ctx *pulumi.Context,
	name string, args *FirewallipmacbindingTableArgs, opts ...pulumi.ResourceOption) (*FirewallipmacbindingTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallipmacbindingTable
	err := ctx.RegisterResource("fortios:index/firewallipmacbindingTable:FirewallipmacbindingTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallipmacbindingTable gets an existing FirewallipmacbindingTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallipmacbindingTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallipmacbindingTableState, opts ...pulumi.ResourceOption) (*FirewallipmacbindingTable, error) {
	var resource FirewallipmacbindingTable
	err := ctx.ReadResource("fortios:index/firewallipmacbindingTable:FirewallipmacbindingTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallipmacbindingTable resources.
type firewallipmacbindingTableState struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip *string `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac *string `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name *string `pulumi:"name"`
	// Entry number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallipmacbindingTableState struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringPtrInput
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringPtrInput
	// Name of the pair (optional, default = no name).
	Name pulumi.StringPtrInput
	// Entry number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallipmacbindingTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallipmacbindingTableState)(nil)).Elem()
}

type firewallipmacbindingTableArgs struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip string `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac *string `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name *string `pulumi:"name"`
	// Entry number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallipmacbindingTable resource.
type FirewallipmacbindingTableArgs struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringInput
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringPtrInput
	// Name of the pair (optional, default = no name).
	Name pulumi.StringPtrInput
	// Entry number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallipmacbindingTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallipmacbindingTableArgs)(nil)).Elem()
}

type FirewallipmacbindingTableInput interface {
	pulumi.Input

	ToFirewallipmacbindingTableOutput() FirewallipmacbindingTableOutput
	ToFirewallipmacbindingTableOutputWithContext(ctx context.Context) FirewallipmacbindingTableOutput
}

func (*FirewallipmacbindingTable) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallipmacbindingTable)(nil)).Elem()
}

func (i *FirewallipmacbindingTable) ToFirewallipmacbindingTableOutput() FirewallipmacbindingTableOutput {
	return i.ToFirewallipmacbindingTableOutputWithContext(context.Background())
}

func (i *FirewallipmacbindingTable) ToFirewallipmacbindingTableOutputWithContext(ctx context.Context) FirewallipmacbindingTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingTableOutput)
}

// FirewallipmacbindingTableArrayInput is an input type that accepts FirewallipmacbindingTableArray and FirewallipmacbindingTableArrayOutput values.
// You can construct a concrete instance of `FirewallipmacbindingTableArrayInput` via:
//
//	FirewallipmacbindingTableArray{ FirewallipmacbindingTableArgs{...} }
type FirewallipmacbindingTableArrayInput interface {
	pulumi.Input

	ToFirewallipmacbindingTableArrayOutput() FirewallipmacbindingTableArrayOutput
	ToFirewallipmacbindingTableArrayOutputWithContext(context.Context) FirewallipmacbindingTableArrayOutput
}

type FirewallipmacbindingTableArray []FirewallipmacbindingTableInput

func (FirewallipmacbindingTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallipmacbindingTable)(nil)).Elem()
}

func (i FirewallipmacbindingTableArray) ToFirewallipmacbindingTableArrayOutput() FirewallipmacbindingTableArrayOutput {
	return i.ToFirewallipmacbindingTableArrayOutputWithContext(context.Background())
}

func (i FirewallipmacbindingTableArray) ToFirewallipmacbindingTableArrayOutputWithContext(ctx context.Context) FirewallipmacbindingTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingTableArrayOutput)
}

// FirewallipmacbindingTableMapInput is an input type that accepts FirewallipmacbindingTableMap and FirewallipmacbindingTableMapOutput values.
// You can construct a concrete instance of `FirewallipmacbindingTableMapInput` via:
//
//	FirewallipmacbindingTableMap{ "key": FirewallipmacbindingTableArgs{...} }
type FirewallipmacbindingTableMapInput interface {
	pulumi.Input

	ToFirewallipmacbindingTableMapOutput() FirewallipmacbindingTableMapOutput
	ToFirewallipmacbindingTableMapOutputWithContext(context.Context) FirewallipmacbindingTableMapOutput
}

type FirewallipmacbindingTableMap map[string]FirewallipmacbindingTableInput

func (FirewallipmacbindingTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallipmacbindingTable)(nil)).Elem()
}

func (i FirewallipmacbindingTableMap) ToFirewallipmacbindingTableMapOutput() FirewallipmacbindingTableMapOutput {
	return i.ToFirewallipmacbindingTableMapOutputWithContext(context.Background())
}

func (i FirewallipmacbindingTableMap) ToFirewallipmacbindingTableMapOutputWithContext(ctx context.Context) FirewallipmacbindingTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallipmacbindingTableMapOutput)
}

type FirewallipmacbindingTableOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallipmacbindingTable)(nil)).Elem()
}

func (o FirewallipmacbindingTableOutput) ToFirewallipmacbindingTableOutput() FirewallipmacbindingTableOutput {
	return o
}

func (o FirewallipmacbindingTableOutput) ToFirewallipmacbindingTableOutputWithContext(ctx context.Context) FirewallipmacbindingTableOutput {
	return o
}

// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
func (o FirewallipmacbindingTableOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
func (o FirewallipmacbindingTableOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Name of the pair (optional, default = no name).
func (o FirewallipmacbindingTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Entry number.
func (o FirewallipmacbindingTableOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.IntOutput { return v.SeqNum }).(pulumi.IntOutput)
}

// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
func (o FirewallipmacbindingTableOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallipmacbindingTableOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallipmacbindingTable) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallipmacbindingTableArrayOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallipmacbindingTable)(nil)).Elem()
}

func (o FirewallipmacbindingTableArrayOutput) ToFirewallipmacbindingTableArrayOutput() FirewallipmacbindingTableArrayOutput {
	return o
}

func (o FirewallipmacbindingTableArrayOutput) ToFirewallipmacbindingTableArrayOutputWithContext(ctx context.Context) FirewallipmacbindingTableArrayOutput {
	return o
}

func (o FirewallipmacbindingTableArrayOutput) Index(i pulumi.IntInput) FirewallipmacbindingTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallipmacbindingTable {
		return vs[0].([]*FirewallipmacbindingTable)[vs[1].(int)]
	}).(FirewallipmacbindingTableOutput)
}

type FirewallipmacbindingTableMapOutput struct{ *pulumi.OutputState }

func (FirewallipmacbindingTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallipmacbindingTable)(nil)).Elem()
}

func (o FirewallipmacbindingTableMapOutput) ToFirewallipmacbindingTableMapOutput() FirewallipmacbindingTableMapOutput {
	return o
}

func (o FirewallipmacbindingTableMapOutput) ToFirewallipmacbindingTableMapOutputWithContext(ctx context.Context) FirewallipmacbindingTableMapOutput {
	return o
}

func (o FirewallipmacbindingTableMapOutput) MapIndex(k pulumi.StringInput) FirewallipmacbindingTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallipmacbindingTable {
		return vs[0].(map[string]*FirewallipmacbindingTable)[vs[1].(string)]
	}).(FirewallipmacbindingTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingTableInput)(nil)).Elem(), &FirewallipmacbindingTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingTableArrayInput)(nil)).Elem(), FirewallipmacbindingTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallipmacbindingTableMapInput)(nil)).Elem(), FirewallipmacbindingTableMap{})
	pulumi.RegisterOutputType(FirewallipmacbindingTableOutput{})
	pulumi.RegisterOutputType(FirewallipmacbindingTableArrayOutput{})
	pulumi.RegisterOutputType(FirewallipmacbindingTableMapOutput{})
}