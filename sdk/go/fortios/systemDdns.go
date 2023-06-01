// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DDNS.
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
//			_, err := fortios.NewSystemDdns(ctx, "trname", &fortios.SystemDdnsArgs{
//				BoundIp:      pulumi.String("0.0.0.0"),
//				ClearText:    pulumi.String("disable"),
//				DdnsAuth:     pulumi.String("disable"),
//				DdnsDomain:   pulumi.String("www.s.com"),
//				DdnsPassword: pulumi.String("ewewcd"),
//				DdnsServer:   pulumi.String("tzo.com"),
//				DdnsServerIp: pulumi.String("0.0.0.0"),
//				DdnsTtl:      pulumi.Int(300),
//				DdnsUsername: pulumi.String("sie2ae"),
//				Ddnsid:       pulumi.Int(1),
//				MonitorInterfaces: fortios.SystemDdnsMonitorInterfaceArray{
//					&fortios.SystemDdnsMonitorInterfaceArgs{
//						InterfaceName: pulumi.String("port2"),
//					},
//				},
//				SslCertificate: pulumi.String("Fortinet_Factory"),
//				UpdateInterval: pulumi.Int(300),
//				UsePublicIp:    pulumi.String("disable"),
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
// # System Ddns can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemDdns:SystemDdns labelname {{ddnsid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemDdns:SystemDdns labelname {{ddnsid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemDdns struct {
	pulumi.CustomResourceState

	// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
	AddrType pulumi.StringOutput `pulumi:"addrType"`
	// Bound IP address.
	BoundIp pulumi.StringOutput `pulumi:"boundIp"`
	// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
	ClearText pulumi.StringOutput `pulumi:"clearText"`
	// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
	DdnsAuth pulumi.StringOutput `pulumi:"ddnsAuth"`
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain pulumi.StringOutput `pulumi:"ddnsDomain"`
	// DDNS update key (base 64 encoding).
	DdnsKey pulumi.StringOutput `pulumi:"ddnsKey"`
	// DDNS update key name.
	DdnsKeyname pulumi.StringOutput `pulumi:"ddnsKeyname"`
	// DDNS password.
	DdnsPassword pulumi.StringPtrOutput `pulumi:"ddnsPassword"`
	// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
	DdnsServer pulumi.StringOutput `pulumi:"ddnsServer"`
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs SystemDdnsDdnsServerAddrArrayOutput `pulumi:"ddnsServerAddrs"`
	// Generic DDNS server IP.
	DdnsServerIp pulumi.StringOutput `pulumi:"ddnsServerIp"`
	// DDNS Serial Number.
	DdnsSn pulumi.StringOutput `pulumi:"ddnsSn"`
	// Time-to-live for DDNS packets.
	DdnsTtl pulumi.IntOutput `pulumi:"ddnsTtl"`
	// DDNS user name.
	DdnsUsername pulumi.StringOutput `pulumi:"ddnsUsername"`
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone pulumi.StringOutput `pulumi:"ddnsZone"`
	// DDNS ID.
	Ddnsid pulumi.IntOutput `pulumi:"ddnsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces SystemDdnsMonitorInterfaceArrayOutput `pulumi:"monitorInterfaces"`
	// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Name of local certificate for SSL connections.
	SslCertificate pulumi.StringOutput `pulumi:"sslCertificate"`
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval pulumi.IntOutput `pulumi:"updateInterval"`
	// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
	UsePublicIp pulumi.StringOutput `pulumi:"usePublicIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDdns registers a new resource with the given unique name, arguments, and options.
func NewSystemDdns(ctx *pulumi.Context,
	name string, args *SystemDdnsArgs, opts ...pulumi.ResourceOption) (*SystemDdns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DdnsServer == nil {
		return nil, errors.New("invalid value for required argument 'DdnsServer'")
	}
	if args.MonitorInterfaces == nil {
		return nil, errors.New("invalid value for required argument 'MonitorInterfaces'")
	}
	if args.DdnsKey != nil {
		args.DdnsKey = pulumi.ToSecret(args.DdnsKey).(pulumi.StringPtrInput)
	}
	if args.DdnsPassword != nil {
		args.DdnsPassword = pulumi.ToSecret(args.DdnsPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ddnsKey",
		"ddnsPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemDdns
	err := ctx.RegisterResource("fortios:index/systemDdns:SystemDdns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDdns gets an existing SystemDdns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDdns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDdnsState, opts ...pulumi.ResourceOption) (*SystemDdns, error) {
	var resource SystemDdns
	err := ctx.ReadResource("fortios:index/systemDdns:SystemDdns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDdns resources.
type systemDdnsState struct {
	// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
	AddrType *string `pulumi:"addrType"`
	// Bound IP address.
	BoundIp *string `pulumi:"boundIp"`
	// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
	ClearText *string `pulumi:"clearText"`
	// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
	DdnsAuth *string `pulumi:"ddnsAuth"`
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain *string `pulumi:"ddnsDomain"`
	// DDNS update key (base 64 encoding).
	DdnsKey *string `pulumi:"ddnsKey"`
	// DDNS update key name.
	DdnsKeyname *string `pulumi:"ddnsKeyname"`
	// DDNS password.
	DdnsPassword *string `pulumi:"ddnsPassword"`
	// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
	DdnsServer *string `pulumi:"ddnsServer"`
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs []SystemDdnsDdnsServerAddr `pulumi:"ddnsServerAddrs"`
	// Generic DDNS server IP.
	DdnsServerIp *string `pulumi:"ddnsServerIp"`
	// DDNS Serial Number.
	DdnsSn *string `pulumi:"ddnsSn"`
	// Time-to-live for DDNS packets.
	DdnsTtl *int `pulumi:"ddnsTtl"`
	// DDNS user name.
	DdnsUsername *string `pulumi:"ddnsUsername"`
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone *string `pulumi:"ddnsZone"`
	// DDNS ID.
	Ddnsid *int `pulumi:"ddnsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces []SystemDdnsMonitorInterface `pulumi:"monitorInterfaces"`
	// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
	ServerType *string `pulumi:"serverType"`
	// Name of local certificate for SSL connections.
	SslCertificate *string `pulumi:"sslCertificate"`
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval *int `pulumi:"updateInterval"`
	// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
	UsePublicIp *string `pulumi:"usePublicIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemDdnsState struct {
	// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
	AddrType pulumi.StringPtrInput
	// Bound IP address.
	BoundIp pulumi.StringPtrInput
	// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
	ClearText pulumi.StringPtrInput
	// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
	DdnsAuth pulumi.StringPtrInput
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain pulumi.StringPtrInput
	// DDNS update key (base 64 encoding).
	DdnsKey pulumi.StringPtrInput
	// DDNS update key name.
	DdnsKeyname pulumi.StringPtrInput
	// DDNS password.
	DdnsPassword pulumi.StringPtrInput
	// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
	DdnsServer pulumi.StringPtrInput
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs SystemDdnsDdnsServerAddrArrayInput
	// Generic DDNS server IP.
	DdnsServerIp pulumi.StringPtrInput
	// DDNS Serial Number.
	DdnsSn pulumi.StringPtrInput
	// Time-to-live for DDNS packets.
	DdnsTtl pulumi.IntPtrInput
	// DDNS user name.
	DdnsUsername pulumi.StringPtrInput
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone pulumi.StringPtrInput
	// DDNS ID.
	Ddnsid pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces SystemDdnsMonitorInterfaceArrayInput
	// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
	ServerType pulumi.StringPtrInput
	// Name of local certificate for SSL connections.
	SslCertificate pulumi.StringPtrInput
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval pulumi.IntPtrInput
	// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
	UsePublicIp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDdnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDdnsState)(nil)).Elem()
}

type systemDdnsArgs struct {
	// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
	AddrType *string `pulumi:"addrType"`
	// Bound IP address.
	BoundIp *string `pulumi:"boundIp"`
	// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
	ClearText *string `pulumi:"clearText"`
	// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
	DdnsAuth *string `pulumi:"ddnsAuth"`
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain *string `pulumi:"ddnsDomain"`
	// DDNS update key (base 64 encoding).
	DdnsKey *string `pulumi:"ddnsKey"`
	// DDNS update key name.
	DdnsKeyname *string `pulumi:"ddnsKeyname"`
	// DDNS password.
	DdnsPassword *string `pulumi:"ddnsPassword"`
	// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
	DdnsServer string `pulumi:"ddnsServer"`
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs []SystemDdnsDdnsServerAddr `pulumi:"ddnsServerAddrs"`
	// Generic DDNS server IP.
	DdnsServerIp *string `pulumi:"ddnsServerIp"`
	// DDNS Serial Number.
	DdnsSn *string `pulumi:"ddnsSn"`
	// Time-to-live for DDNS packets.
	DdnsTtl *int `pulumi:"ddnsTtl"`
	// DDNS user name.
	DdnsUsername *string `pulumi:"ddnsUsername"`
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone *string `pulumi:"ddnsZone"`
	// DDNS ID.
	Ddnsid *int `pulumi:"ddnsid"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces []SystemDdnsMonitorInterface `pulumi:"monitorInterfaces"`
	// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
	ServerType *string `pulumi:"serverType"`
	// Name of local certificate for SSL connections.
	SslCertificate *string `pulumi:"sslCertificate"`
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval *int `pulumi:"updateInterval"`
	// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
	UsePublicIp *string `pulumi:"usePublicIp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDdns resource.
type SystemDdnsArgs struct {
	// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
	AddrType pulumi.StringPtrInput
	// Bound IP address.
	BoundIp pulumi.StringPtrInput
	// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
	ClearText pulumi.StringPtrInput
	// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
	DdnsAuth pulumi.StringPtrInput
	// Your fully qualified domain name (for example, yourname.DDNS.com).
	DdnsDomain pulumi.StringPtrInput
	// DDNS update key (base 64 encoding).
	DdnsKey pulumi.StringPtrInput
	// DDNS update key name.
	DdnsKeyname pulumi.StringPtrInput
	// DDNS password.
	DdnsPassword pulumi.StringPtrInput
	// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
	DdnsServer pulumi.StringInput
	// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
	DdnsServerAddrs SystemDdnsDdnsServerAddrArrayInput
	// Generic DDNS server IP.
	DdnsServerIp pulumi.StringPtrInput
	// DDNS Serial Number.
	DdnsSn pulumi.StringPtrInput
	// Time-to-live for DDNS packets.
	DdnsTtl pulumi.IntPtrInput
	// DDNS user name.
	DdnsUsername pulumi.StringPtrInput
	// Zone of your domain name (for example, DDNS.com).
	DdnsZone pulumi.StringPtrInput
	// DDNS ID.
	Ddnsid pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Monitored interface. The structure of `monitorInterface` block is documented below.
	MonitorInterfaces SystemDdnsMonitorInterfaceArrayInput
	// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
	ServerType pulumi.StringPtrInput
	// Name of local certificate for SSL connections.
	SslCertificate pulumi.StringPtrInput
	// DDNS update interval (60 - 2592000 sec, default = 300).
	UpdateInterval pulumi.IntPtrInput
	// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
	UsePublicIp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDdnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDdnsArgs)(nil)).Elem()
}

type SystemDdnsInput interface {
	pulumi.Input

	ToSystemDdnsOutput() SystemDdnsOutput
	ToSystemDdnsOutputWithContext(ctx context.Context) SystemDdnsOutput
}

func (*SystemDdns) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDdns)(nil)).Elem()
}

func (i *SystemDdns) ToSystemDdnsOutput() SystemDdnsOutput {
	return i.ToSystemDdnsOutputWithContext(context.Background())
}

func (i *SystemDdns) ToSystemDdnsOutputWithContext(ctx context.Context) SystemDdnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDdnsOutput)
}

// SystemDdnsArrayInput is an input type that accepts SystemDdnsArray and SystemDdnsArrayOutput values.
// You can construct a concrete instance of `SystemDdnsArrayInput` via:
//
//	SystemDdnsArray{ SystemDdnsArgs{...} }
type SystemDdnsArrayInput interface {
	pulumi.Input

	ToSystemDdnsArrayOutput() SystemDdnsArrayOutput
	ToSystemDdnsArrayOutputWithContext(context.Context) SystemDdnsArrayOutput
}

type SystemDdnsArray []SystemDdnsInput

func (SystemDdnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDdns)(nil)).Elem()
}

func (i SystemDdnsArray) ToSystemDdnsArrayOutput() SystemDdnsArrayOutput {
	return i.ToSystemDdnsArrayOutputWithContext(context.Background())
}

func (i SystemDdnsArray) ToSystemDdnsArrayOutputWithContext(ctx context.Context) SystemDdnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDdnsArrayOutput)
}

// SystemDdnsMapInput is an input type that accepts SystemDdnsMap and SystemDdnsMapOutput values.
// You can construct a concrete instance of `SystemDdnsMapInput` via:
//
//	SystemDdnsMap{ "key": SystemDdnsArgs{...} }
type SystemDdnsMapInput interface {
	pulumi.Input

	ToSystemDdnsMapOutput() SystemDdnsMapOutput
	ToSystemDdnsMapOutputWithContext(context.Context) SystemDdnsMapOutput
}

type SystemDdnsMap map[string]SystemDdnsInput

func (SystemDdnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDdns)(nil)).Elem()
}

func (i SystemDdnsMap) ToSystemDdnsMapOutput() SystemDdnsMapOutput {
	return i.ToSystemDdnsMapOutputWithContext(context.Background())
}

func (i SystemDdnsMap) ToSystemDdnsMapOutputWithContext(ctx context.Context) SystemDdnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDdnsMapOutput)
}

type SystemDdnsOutput struct{ *pulumi.OutputState }

func (SystemDdnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDdns)(nil)).Elem()
}

func (o SystemDdnsOutput) ToSystemDdnsOutput() SystemDdnsOutput {
	return o
}

func (o SystemDdnsOutput) ToSystemDdnsOutputWithContext(ctx context.Context) SystemDdnsOutput {
	return o
}

// Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
func (o SystemDdnsOutput) AddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.AddrType }).(pulumi.StringOutput)
}

// Bound IP address.
func (o SystemDdnsOutput) BoundIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.BoundIp }).(pulumi.StringOutput)
}

// Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
func (o SystemDdnsOutput) ClearText() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.ClearText }).(pulumi.StringOutput)
}

// Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
func (o SystemDdnsOutput) DdnsAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsAuth }).(pulumi.StringOutput)
}

// Your fully qualified domain name (for example, yourname.DDNS.com).
func (o SystemDdnsOutput) DdnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsDomain }).(pulumi.StringOutput)
}

// DDNS update key (base 64 encoding).
func (o SystemDdnsOutput) DdnsKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsKey }).(pulumi.StringOutput)
}

// DDNS update key name.
func (o SystemDdnsOutput) DdnsKeyname() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsKeyname }).(pulumi.StringOutput)
}

// DDNS password.
func (o SystemDdnsOutput) DdnsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringPtrOutput { return v.DdnsPassword }).(pulumi.StringPtrOutput)
}

// Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
func (o SystemDdnsOutput) DdnsServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsServer }).(pulumi.StringOutput)
}

// Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
func (o SystemDdnsOutput) DdnsServerAddrs() SystemDdnsDdnsServerAddrArrayOutput {
	return o.ApplyT(func(v *SystemDdns) SystemDdnsDdnsServerAddrArrayOutput { return v.DdnsServerAddrs }).(SystemDdnsDdnsServerAddrArrayOutput)
}

// Generic DDNS server IP.
func (o SystemDdnsOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsServerIp }).(pulumi.StringOutput)
}

// DDNS Serial Number.
func (o SystemDdnsOutput) DdnsSn() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsSn }).(pulumi.StringOutput)
}

// Time-to-live for DDNS packets.
func (o SystemDdnsOutput) DdnsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.IntOutput { return v.DdnsTtl }).(pulumi.IntOutput)
}

// DDNS user name.
func (o SystemDdnsOutput) DdnsUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsUsername }).(pulumi.StringOutput)
}

// Zone of your domain name (for example, DDNS.com).
func (o SystemDdnsOutput) DdnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.DdnsZone }).(pulumi.StringOutput)
}

// DDNS ID.
func (o SystemDdnsOutput) Ddnsid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.IntOutput { return v.Ddnsid }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemDdnsOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Monitored interface. The structure of `monitorInterface` block is documented below.
func (o SystemDdnsOutput) MonitorInterfaces() SystemDdnsMonitorInterfaceArrayOutput {
	return o.ApplyT(func(v *SystemDdns) SystemDdnsMonitorInterfaceArrayOutput { return v.MonitorInterfaces }).(SystemDdnsMonitorInterfaceArrayOutput)
}

// Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
func (o SystemDdnsOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// Name of local certificate for SSL connections.
func (o SystemDdnsOutput) SslCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.SslCertificate }).(pulumi.StringOutput)
}

// DDNS update interval (60 - 2592000 sec, default = 300).
func (o SystemDdnsOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.IntOutput { return v.UpdateInterval }).(pulumi.IntOutput)
}

// Enable/disable use of public IP address. Valid values: `disable`, `enable`.
func (o SystemDdnsOutput) UsePublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringOutput { return v.UsePublicIp }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemDdnsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDdns) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemDdnsArrayOutput struct{ *pulumi.OutputState }

func (SystemDdnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDdns)(nil)).Elem()
}

func (o SystemDdnsArrayOutput) ToSystemDdnsArrayOutput() SystemDdnsArrayOutput {
	return o
}

func (o SystemDdnsArrayOutput) ToSystemDdnsArrayOutputWithContext(ctx context.Context) SystemDdnsArrayOutput {
	return o
}

func (o SystemDdnsArrayOutput) Index(i pulumi.IntInput) SystemDdnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDdns {
		return vs[0].([]*SystemDdns)[vs[1].(int)]
	}).(SystemDdnsOutput)
}

type SystemDdnsMapOutput struct{ *pulumi.OutputState }

func (SystemDdnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDdns)(nil)).Elem()
}

func (o SystemDdnsMapOutput) ToSystemDdnsMapOutput() SystemDdnsMapOutput {
	return o
}

func (o SystemDdnsMapOutput) ToSystemDdnsMapOutputWithContext(ctx context.Context) SystemDdnsMapOutput {
	return o
}

func (o SystemDdnsMapOutput) MapIndex(k pulumi.StringInput) SystemDdnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDdns {
		return vs[0].(map[string]*SystemDdns)[vs[1].(string)]
	}).(SystemDdnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDdnsInput)(nil)).Elem(), &SystemDdns{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDdnsArrayInput)(nil)).Elem(), SystemDdnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDdnsMapInput)(nil)).Elem(), SystemDdnsMap{})
	pulumi.RegisterOutputType(SystemDdnsOutput{})
	pulumi.RegisterOutputType(SystemDdnsArrayOutput{})
	pulumi.RegisterOutputType(SystemDdnsMapOutput{})
}