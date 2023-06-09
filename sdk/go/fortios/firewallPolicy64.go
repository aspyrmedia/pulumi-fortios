// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 to IPv4 policies. Applies to FortiOS Version `<= 7.0.0`.
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
//			_, err := fortios.NewFirewallPolicy64(ctx, "trname", &fortios.FirewallPolicy64Args{
//				Action: pulumi.String("accept"),
//				Dstaddrs: fortios.FirewallPolicy64DstaddrArray{
//					&fortios.FirewallPolicy64DstaddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				Dstintf:       pulumi.String("port4"),
//				Fixedport:     pulumi.String("disable"),
//				Ippool:        pulumi.String("disable"),
//				Logtraffic:    pulumi.String("disable"),
//				PermitAnyHost: pulumi.String("disable"),
//				Policyid:      pulumi.Int(1),
//				Schedule:      pulumi.String("always"),
//				Services: fortios.FirewallPolicy64ServiceArray{
//					&fortios.FirewallPolicy64ServiceArgs{
//						Name: pulumi.String("ALL"),
//					},
//				},
//				Srcaddrs: fortios.FirewallPolicy64SrcaddrArray{
//					&fortios.FirewallPolicy64SrcaddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				Srcintf:        pulumi.String("port3"),
//				Status:         pulumi.String("enable"),
//				TcpMssReceiver: pulumi.Int(0),
//				TcpMssSender:   pulumi.Int(0),
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
// # Firewall Policy64 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallPolicy64:FirewallPolicy64 labelname {{policyid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallPolicy64:FirewallPolicy64 labelname {{policyid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallPolicy64 struct {
	pulumi.CustomResourceState

	// Policy action. Valid values: `accept`, `deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy64DstaddrArrayOutput `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf pulumi.StringOutput `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringOutput `pulumi:"fixedport"`
	// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
	Ippool pulumi.StringOutput `pulumi:"ippool"`
	// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringOutput `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringOutput `pulumi:"logtrafficStart"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Per-IP traffic shaper.
	PerIpShaper pulumi.StringOutput `pulumi:"perIpShaper"`
	// Enable/disable permit any host in. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringOutput `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy64PoolnameArrayOutput `pulumi:"poolnames"`
	// Schedule name.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy64ServiceArrayOutput `pulumi:"services"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy64SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf pulumi.StringOutput `pulumi:"srcintf"`
	// Enable/disable policy status. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// TCP MSS value of receiver.
	TcpMssReceiver pulumi.IntOutput `pulumi:"tcpMssReceiver"`
	// TCP MSS value of sender.
	TcpMssSender pulumi.IntOutput `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper pulumi.StringOutput `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringOutput `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallPolicy64 registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy64(ctx *pulumi.Context,
	name string, args *FirewallPolicy64Args, opts ...pulumi.ResourceOption) (*FirewallPolicy64, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Dstintf == nil {
		return nil, errors.New("invalid value for required argument 'Dstintf'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	if args.Srcintf == nil {
		return nil, errors.New("invalid value for required argument 'Srcintf'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallPolicy64
	err := ctx.RegisterResource("fortios:index/firewallPolicy64:FirewallPolicy64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy64 gets an existing FirewallPolicy64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicy64State, opts ...pulumi.ResourceOption) (*FirewallPolicy64, error) {
	var resource FirewallPolicy64
	err := ctx.ReadResource("fortios:index/firewallPolicy64:FirewallPolicy64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy64 resources.
type firewallPolicy64State struct {
	// Policy action. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallPolicy64Dstaddr `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf *string `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
	Fixedport *string `pulumi:"fixedport"`
	// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
	Ippool *string `pulumi:"ippool"`
	// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart *string `pulumi:"logtrafficStart"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Per-IP traffic shaper.
	PerIpShaper *string `pulumi:"perIpShaper"`
	// Enable/disable permit any host in. Valid values: `enable`, `disable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames []FirewallPolicy64Poolname `pulumi:"poolnames"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services []FirewallPolicy64Service `pulumi:"services"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallPolicy64Srcaddr `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf *string `pulumi:"srcintf"`
	// Enable/disable policy status. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// TCP MSS value of receiver.
	TcpMssReceiver *int `pulumi:"tcpMssReceiver"`
	// TCP MSS value of sender.
	TcpMssSender *int `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper *string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse *string `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallPolicy64State struct {
	// Policy action. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy64DstaddrArrayInput
	// Destination interface name.
	Dstintf pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringPtrInput
	// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
	Ippool pulumi.StringPtrInput
	// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Per-IP traffic shaper.
	PerIpShaper pulumi.StringPtrInput
	// Enable/disable permit any host in. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy64PoolnameArrayInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy64ServiceArrayInput
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy64SrcaddrArrayInput
	// Source interface name.
	Srcintf pulumi.StringPtrInput
	// Enable/disable policy status. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// TCP MSS value of receiver.
	TcpMssReceiver pulumi.IntPtrInput
	// TCP MSS value of sender.
	TcpMssSender pulumi.IntPtrInput
	// Traffic shaper.
	TrafficShaper pulumi.StringPtrInput
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallPolicy64State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicy64State)(nil)).Elem()
}

type firewallPolicy64Args struct {
	// Policy action. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallPolicy64Dstaddr `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf string `pulumi:"dstintf"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
	Fixedport *string `pulumi:"fixedport"`
	// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
	Ippool *string `pulumi:"ippool"`
	// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart *string `pulumi:"logtrafficStart"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Per-IP traffic shaper.
	PerIpShaper *string `pulumi:"perIpShaper"`
	// Enable/disable permit any host in. Valid values: `enable`, `disable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames []FirewallPolicy64Poolname `pulumi:"poolnames"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services []FirewallPolicy64Service `pulumi:"services"`
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallPolicy64Srcaddr `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf string `pulumi:"srcintf"`
	// Enable/disable policy status. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// TCP MSS value of receiver.
	TcpMssReceiver *int `pulumi:"tcpMssReceiver"`
	// TCP MSS value of sender.
	TcpMssSender *int `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper *string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse *string `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallPolicy64 resource.
type FirewallPolicy64Args struct {
	// Policy action. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy64DstaddrArrayInput
	// Destination interface name.
	Dstintf pulumi.StringInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringPtrInput
	// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
	Ippool pulumi.StringPtrInput
	// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Per-IP traffic shaper.
	PerIpShaper pulumi.StringPtrInput
	// Enable/disable permit any host in. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Policy IP pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy64PoolnameArrayInput
	// Schedule name.
	Schedule pulumi.StringInput
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy64ServiceArrayInput
	// Source address name. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy64SrcaddrArrayInput
	// Source interface name.
	Srcintf pulumi.StringInput
	// Enable/disable policy status. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// TCP MSS value of receiver.
	TcpMssReceiver pulumi.IntPtrInput
	// TCP MSS value of sender.
	TcpMssSender pulumi.IntPtrInput
	// Traffic shaper.
	TrafficShaper pulumi.StringPtrInput
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallPolicy64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicy64Args)(nil)).Elem()
}

type FirewallPolicy64Input interface {
	pulumi.Input

	ToFirewallPolicy64Output() FirewallPolicy64Output
	ToFirewallPolicy64OutputWithContext(ctx context.Context) FirewallPolicy64Output
}

func (*FirewallPolicy64) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy64)(nil)).Elem()
}

func (i *FirewallPolicy64) ToFirewallPolicy64Output() FirewallPolicy64Output {
	return i.ToFirewallPolicy64OutputWithContext(context.Background())
}

func (i *FirewallPolicy64) ToFirewallPolicy64OutputWithContext(ctx context.Context) FirewallPolicy64Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy64Output)
}

// FirewallPolicy64ArrayInput is an input type that accepts FirewallPolicy64Array and FirewallPolicy64ArrayOutput values.
// You can construct a concrete instance of `FirewallPolicy64ArrayInput` via:
//
//	FirewallPolicy64Array{ FirewallPolicy64Args{...} }
type FirewallPolicy64ArrayInput interface {
	pulumi.Input

	ToFirewallPolicy64ArrayOutput() FirewallPolicy64ArrayOutput
	ToFirewallPolicy64ArrayOutputWithContext(context.Context) FirewallPolicy64ArrayOutput
}

type FirewallPolicy64Array []FirewallPolicy64Input

func (FirewallPolicy64Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy64)(nil)).Elem()
}

func (i FirewallPolicy64Array) ToFirewallPolicy64ArrayOutput() FirewallPolicy64ArrayOutput {
	return i.ToFirewallPolicy64ArrayOutputWithContext(context.Background())
}

func (i FirewallPolicy64Array) ToFirewallPolicy64ArrayOutputWithContext(ctx context.Context) FirewallPolicy64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy64ArrayOutput)
}

// FirewallPolicy64MapInput is an input type that accepts FirewallPolicy64Map and FirewallPolicy64MapOutput values.
// You can construct a concrete instance of `FirewallPolicy64MapInput` via:
//
//	FirewallPolicy64Map{ "key": FirewallPolicy64Args{...} }
type FirewallPolicy64MapInput interface {
	pulumi.Input

	ToFirewallPolicy64MapOutput() FirewallPolicy64MapOutput
	ToFirewallPolicy64MapOutputWithContext(context.Context) FirewallPolicy64MapOutput
}

type FirewallPolicy64Map map[string]FirewallPolicy64Input

func (FirewallPolicy64Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy64)(nil)).Elem()
}

func (i FirewallPolicy64Map) ToFirewallPolicy64MapOutput() FirewallPolicy64MapOutput {
	return i.ToFirewallPolicy64MapOutputWithContext(context.Background())
}

func (i FirewallPolicy64Map) ToFirewallPolicy64MapOutputWithContext(ctx context.Context) FirewallPolicy64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy64MapOutput)
}

type FirewallPolicy64Output struct{ *pulumi.OutputState }

func (FirewallPolicy64Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy64)(nil)).Elem()
}

func (o FirewallPolicy64Output) ToFirewallPolicy64Output() FirewallPolicy64Output {
	return o
}

func (o FirewallPolicy64Output) ToFirewallPolicy64OutputWithContext(ctx context.Context) FirewallPolicy64Output {
	return o
}

// Policy action. Valid values: `accept`, `deny`.
func (o FirewallPolicy64Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Comment.
func (o FirewallPolicy64Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Destination address name. The structure of `dstaddr` block is documented below.
func (o FirewallPolicy64Output) Dstaddrs() FirewallPolicy64DstaddrArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy64) FirewallPolicy64DstaddrArrayOutput { return v.Dstaddrs }).(FirewallPolicy64DstaddrArrayOutput)
}

// Destination interface name.
func (o FirewallPolicy64Output) Dstintf() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Dstintf }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallPolicy64Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable policy fixed port. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) Fixedport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Fixedport }).(pulumi.StringOutput)
}

// Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) Ippool() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Ippool }).(pulumi.StringOutput)
}

// Enable/disable policy log traffic. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Logtraffic }).(pulumi.StringOutput)
}

// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.LogtrafficStart }).(pulumi.StringOutput)
}

// Policy name.
func (o FirewallPolicy64Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Per-IP traffic shaper.
func (o FirewallPolicy64Output) PerIpShaper() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.PerIpShaper }).(pulumi.StringOutput)
}

// Enable/disable permit any host in. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) PermitAnyHost() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.PermitAnyHost }).(pulumi.StringOutput)
}

// Policy ID.
func (o FirewallPolicy64Output) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.IntOutput { return v.Policyid }).(pulumi.IntOutput)
}

// Policy IP pool names. The structure of `poolname` block is documented below.
func (o FirewallPolicy64Output) Poolnames() FirewallPolicy64PoolnameArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy64) FirewallPolicy64PoolnameArrayOutput { return v.Poolnames }).(FirewallPolicy64PoolnameArrayOutput)
}

// Schedule name.
func (o FirewallPolicy64Output) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// Service name. The structure of `service` block is documented below.
func (o FirewallPolicy64Output) Services() FirewallPolicy64ServiceArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy64) FirewallPolicy64ServiceArrayOutput { return v.Services }).(FirewallPolicy64ServiceArrayOutput)
}

// Source address name. The structure of `srcaddr` block is documented below.
func (o FirewallPolicy64Output) Srcaddrs() FirewallPolicy64SrcaddrArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy64) FirewallPolicy64SrcaddrArrayOutput { return v.Srcaddrs }).(FirewallPolicy64SrcaddrArrayOutput)
}

// Source interface name.
func (o FirewallPolicy64Output) Srcintf() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Srcintf }).(pulumi.StringOutput)
}

// Enable/disable policy status. Valid values: `enable`, `disable`.
func (o FirewallPolicy64Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// TCP MSS value of receiver.
func (o FirewallPolicy64Output) TcpMssReceiver() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.IntOutput { return v.TcpMssReceiver }).(pulumi.IntOutput)
}

// TCP MSS value of sender.
func (o FirewallPolicy64Output) TcpMssSender() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.IntOutput { return v.TcpMssSender }).(pulumi.IntOutput)
}

// Traffic shaper.
func (o FirewallPolicy64Output) TrafficShaper() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.TrafficShaper }).(pulumi.StringOutput)
}

// Reverse traffic shaper.
func (o FirewallPolicy64Output) TrafficShaperReverse() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.TrafficShaperReverse }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallPolicy64Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallPolicy64Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy64) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallPolicy64ArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicy64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy64)(nil)).Elem()
}

func (o FirewallPolicy64ArrayOutput) ToFirewallPolicy64ArrayOutput() FirewallPolicy64ArrayOutput {
	return o
}

func (o FirewallPolicy64ArrayOutput) ToFirewallPolicy64ArrayOutputWithContext(ctx context.Context) FirewallPolicy64ArrayOutput {
	return o
}

func (o FirewallPolicy64ArrayOutput) Index(i pulumi.IntInput) FirewallPolicy64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallPolicy64 {
		return vs[0].([]*FirewallPolicy64)[vs[1].(int)]
	}).(FirewallPolicy64Output)
}

type FirewallPolicy64MapOutput struct{ *pulumi.OutputState }

func (FirewallPolicy64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy64)(nil)).Elem()
}

func (o FirewallPolicy64MapOutput) ToFirewallPolicy64MapOutput() FirewallPolicy64MapOutput {
	return o
}

func (o FirewallPolicy64MapOutput) ToFirewallPolicy64MapOutputWithContext(ctx context.Context) FirewallPolicy64MapOutput {
	return o
}

func (o FirewallPolicy64MapOutput) MapIndex(k pulumi.StringInput) FirewallPolicy64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallPolicy64 {
		return vs[0].(map[string]*FirewallPolicy64)[vs[1].(string)]
	}).(FirewallPolicy64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy64Input)(nil)).Elem(), &FirewallPolicy64{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy64ArrayInput)(nil)).Elem(), FirewallPolicy64Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy64MapInput)(nil)).Elem(), FirewallPolicy64Map{})
	pulumi.RegisterOutputType(FirewallPolicy64Output{})
	pulumi.RegisterOutputType(FirewallPolicy64ArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicy64MapOutput{})
}
