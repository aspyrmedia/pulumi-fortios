// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 to IPv4 virtual IPs. Applies to FortiOS Version `<= 7.0.0`.
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
//			_, err := fortios.NewFirewallVip64(ctx, "trname", &fortios.FirewallVip64Args{
//				ArpReply:    pulumi.String("enable"),
//				Color:       pulumi.Int(0),
//				Extip:       pulumi.String("2001:db8:99:203::22"),
//				Extport:     pulumi.String("0-65535"),
//				Fosid:       pulumi.Int(0),
//				LdbMethod:   pulumi.String("static"),
//				Mappedip:    pulumi.String("1.1.1.1"),
//				Mappedport:  pulumi.String("0-65535"),
//				Portforward: pulumi.String("disable"),
//				Protocol:    pulumi.String("tcp"),
//				Type:        pulumi.String("static-nat"),
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
// # Firewall Vip64 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallVip64:FirewallVip64 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallVip64:FirewallVip64 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallVip64 struct {
	pulumi.CustomResourceState

	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringOutput `pulumi:"extip"`
	// External service port.
	Extport pulumi.StringOutput `pulumi:"extport"`
	// Custom defined id.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringOutput `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport pulumi.StringOutput `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors FirewallVip64MonitorArrayOutput `pulumi:"monitors"`
	// VIP64 name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringOutput `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip64RealserverArrayOutput `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip64SrcFilterArrayOutput `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallVip64 registers a new resource with the given unique name, arguments, and options.
func NewFirewallVip64(ctx *pulumi.Context,
	name string, args *FirewallVip64Args, opts ...pulumi.ResourceOption) (*FirewallVip64, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedip == nil {
		return nil, errors.New("invalid value for required argument 'Mappedip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallVip64
	err := ctx.RegisterResource("fortios:index/firewallVip64:FirewallVip64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVip64 gets an existing FirewallVip64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVip64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVip64State, opts ...pulumi.ResourceOption) (*FirewallVip64, error) {
	var resource FirewallVip64
	err := ctx.ReadResource("fortios:index/firewallVip64:FirewallVip64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVip64 resources.
type firewallVip64State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip *string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip *string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors []FirewallVip64Monitor `pulumi:"monitors"`
	// VIP64 name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []FirewallVip64Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []FirewallVip64SrcFilter `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallVip64State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringPtrInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringPtrInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors FirewallVip64MonitorArrayInput
	// VIP64 name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip64RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip64SrcFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVip64State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip64State)(nil)).Elem()
}

type firewallVip64Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors []FirewallVip64Monitor `pulumi:"monitors"`
	// VIP64 name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []FirewallVip64Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []FirewallVip64SrcFilter `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVip64 resource.
type FirewallVip64Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors FirewallVip64MonitorArrayInput
	// VIP64 name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip64RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip64SrcFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVip64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip64Args)(nil)).Elem()
}

type FirewallVip64Input interface {
	pulumi.Input

	ToFirewallVip64Output() FirewallVip64Output
	ToFirewallVip64OutputWithContext(ctx context.Context) FirewallVip64Output
}

func (*FirewallVip64) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip64)(nil)).Elem()
}

func (i *FirewallVip64) ToFirewallVip64Output() FirewallVip64Output {
	return i.ToFirewallVip64OutputWithContext(context.Background())
}

func (i *FirewallVip64) ToFirewallVip64OutputWithContext(ctx context.Context) FirewallVip64Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip64Output)
}

// FirewallVip64ArrayInput is an input type that accepts FirewallVip64Array and FirewallVip64ArrayOutput values.
// You can construct a concrete instance of `FirewallVip64ArrayInput` via:
//
//	FirewallVip64Array{ FirewallVip64Args{...} }
type FirewallVip64ArrayInput interface {
	pulumi.Input

	ToFirewallVip64ArrayOutput() FirewallVip64ArrayOutput
	ToFirewallVip64ArrayOutputWithContext(context.Context) FirewallVip64ArrayOutput
}

type FirewallVip64Array []FirewallVip64Input

func (FirewallVip64Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVip64)(nil)).Elem()
}

func (i FirewallVip64Array) ToFirewallVip64ArrayOutput() FirewallVip64ArrayOutput {
	return i.ToFirewallVip64ArrayOutputWithContext(context.Background())
}

func (i FirewallVip64Array) ToFirewallVip64ArrayOutputWithContext(ctx context.Context) FirewallVip64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip64ArrayOutput)
}

// FirewallVip64MapInput is an input type that accepts FirewallVip64Map and FirewallVip64MapOutput values.
// You can construct a concrete instance of `FirewallVip64MapInput` via:
//
//	FirewallVip64Map{ "key": FirewallVip64Args{...} }
type FirewallVip64MapInput interface {
	pulumi.Input

	ToFirewallVip64MapOutput() FirewallVip64MapOutput
	ToFirewallVip64MapOutputWithContext(context.Context) FirewallVip64MapOutput
}

type FirewallVip64Map map[string]FirewallVip64Input

func (FirewallVip64Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVip64)(nil)).Elem()
}

func (i FirewallVip64Map) ToFirewallVip64MapOutput() FirewallVip64MapOutput {
	return i.ToFirewallVip64MapOutputWithContext(context.Background())
}

func (i FirewallVip64Map) ToFirewallVip64MapOutputWithContext(ctx context.Context) FirewallVip64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip64MapOutput)
}

type FirewallVip64Output struct{ *pulumi.OutputState }

func (FirewallVip64Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip64)(nil)).Elem()
}

func (o FirewallVip64Output) ToFirewallVip64Output() FirewallVip64Output {
	return o
}

func (o FirewallVip64Output) ToFirewallVip64OutputWithContext(ctx context.Context) FirewallVip64Output {
	return o
}

// Enable ARP reply. Valid values: `disable`, `enable`.
func (o FirewallVip64Output) ArpReply() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.ArpReply }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o FirewallVip64Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o FirewallVip64Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallVip64Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Start-external-IP [-end-external-IP].
func (o FirewallVip64Output) Extip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Extip }).(pulumi.StringOutput)
}

// External service port.
func (o FirewallVip64Output) Extport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Extport }).(pulumi.StringOutput)
}

// Custom defined id.
func (o FirewallVip64Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
func (o FirewallVip64Output) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Start-mapped-IP [-end-mapped-IP].
func (o FirewallVip64Output) Mappedip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Mappedip }).(pulumi.StringOutput)
}

// Mapped service port.
func (o FirewallVip64Output) Mappedport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Mappedport }).(pulumi.StringOutput)
}

// Health monitors. The structure of `monitor` block is documented below.
func (o FirewallVip64Output) Monitors() FirewallVip64MonitorArrayOutput {
	return o.ApplyT(func(v *FirewallVip64) FirewallVip64MonitorArrayOutput { return v.Monitors }).(FirewallVip64MonitorArrayOutput)
}

// VIP64 name.
func (o FirewallVip64Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable port forwarding. Valid values: `disable`, `enable`.
func (o FirewallVip64Output) Portforward() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Portforward }).(pulumi.StringOutput)
}

// Mapped port protocol. Valid values: `tcp`, `udp`.
func (o FirewallVip64Output) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Real servers. The structure of `realservers` block is documented below.
func (o FirewallVip64Output) Realservers() FirewallVip64RealserverArrayOutput {
	return o.ApplyT(func(v *FirewallVip64) FirewallVip64RealserverArrayOutput { return v.Realservers }).(FirewallVip64RealserverArrayOutput)
}

// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
func (o FirewallVip64Output) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
func (o FirewallVip64Output) SrcFilters() FirewallVip64SrcFilterArrayOutput {
	return o.ApplyT(func(v *FirewallVip64) FirewallVip64SrcFilterArrayOutput { return v.SrcFilters }).(FirewallVip64SrcFilterArrayOutput)
}

// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
func (o FirewallVip64Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallVip64Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallVip64Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip64) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallVip64ArrayOutput struct{ *pulumi.OutputState }

func (FirewallVip64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVip64)(nil)).Elem()
}

func (o FirewallVip64ArrayOutput) ToFirewallVip64ArrayOutput() FirewallVip64ArrayOutput {
	return o
}

func (o FirewallVip64ArrayOutput) ToFirewallVip64ArrayOutputWithContext(ctx context.Context) FirewallVip64ArrayOutput {
	return o
}

func (o FirewallVip64ArrayOutput) Index(i pulumi.IntInput) FirewallVip64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVip64 {
		return vs[0].([]*FirewallVip64)[vs[1].(int)]
	}).(FirewallVip64Output)
}

type FirewallVip64MapOutput struct{ *pulumi.OutputState }

func (FirewallVip64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVip64)(nil)).Elem()
}

func (o FirewallVip64MapOutput) ToFirewallVip64MapOutput() FirewallVip64MapOutput {
	return o
}

func (o FirewallVip64MapOutput) ToFirewallVip64MapOutputWithContext(ctx context.Context) FirewallVip64MapOutput {
	return o
}

func (o FirewallVip64MapOutput) MapIndex(k pulumi.StringInput) FirewallVip64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVip64 {
		return vs[0].(map[string]*FirewallVip64)[vs[1].(string)]
	}).(FirewallVip64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip64Input)(nil)).Elem(), &FirewallVip64{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip64ArrayInput)(nil)).Elem(), FirewallVip64Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip64MapInput)(nil)).Elem(), FirewallVip64Map{})
	pulumi.RegisterOutputType(FirewallVip64Output{})
	pulumi.RegisterOutputType(FirewallVip64ArrayOutput{})
	pulumi.RegisterOutputType(FirewallVip64MapOutput{})
}
