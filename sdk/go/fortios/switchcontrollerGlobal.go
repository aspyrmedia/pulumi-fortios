// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch global settings.
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
//			_, err := fortios.NewSwitchcontrollerGlobal(ctx, "trname", &fortios.SwitchcontrollerGlobalArgs{
//				AllowMultipleInterfaces: pulumi.String("disable"),
//				HttpsImagePush:          pulumi.String("disable"),
//				LogMacLimitViolations:   pulumi.String("disable"),
//				MacAgingInterval:        pulumi.Int(332),
//				MacRetentionPeriod:      pulumi.Int(24),
//				MacViolationTimer:       pulumi.Int(0),
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
// # SwitchController Global can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal labelname SwitchControllerGlobal
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal labelname SwitchControllerGlobal
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerGlobal struct {
	pulumi.CustomResourceState

	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringOutput `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringOutput `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands SwitchcontrollerGlobalCustomCommandArrayOutput `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringOutput `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringOutput `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries SwitchcontrollerGlobalDisableDiscoveryArrayOutput `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringOutput `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringOutput `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringOutput `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringOutput `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntOutput `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringOutput `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntOutput `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntOutput `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringOutput `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringOutput `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringOutput `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringOutput `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringOutput `pulumi:"vlanOptimization"`
}

// NewSwitchcontrollerGlobal registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerGlobal(ctx *pulumi.Context,
	name string, args *SwitchcontrollerGlobalArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerGlobal, error) {
	if args == nil {
		args = &SwitchcontrollerGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerGlobal
	err := ctx.RegisterResource("fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerGlobal gets an existing SwitchcontrollerGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerGlobalState, opts ...pulumi.ResourceOption) (*SwitchcontrollerGlobal, error) {
	var resource SwitchcontrollerGlobal
	err := ctx.ReadResource("fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerGlobal resources.
type switchcontrollerGlobalState struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces *string `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink *string `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands []SwitchcontrollerGlobalCustomCommand `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan *string `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList *string `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries []SwitchcontrollerGlobalDisableDiscovery `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce *string `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush *string `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations *string `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval *int `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging *string `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod *int `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer *int `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode *string `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution *string `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice *string `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode *string `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization *string `pulumi:"vlanOptimization"`
}

type SwitchcontrollerGlobalState struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringPtrInput
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands SwitchcontrollerGlobalCustomCommandArrayInput
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringPtrInput
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringPtrInput
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries SwitchcontrollerGlobalDisableDiscoveryArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringPtrInput
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringPtrInput
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntPtrInput
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringPtrInput
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntPtrInput
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntPtrInput
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringPtrInput
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringPtrInput
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringPtrInput
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringPtrInput
}

func (SwitchcontrollerGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerGlobalState)(nil)).Elem()
}

type switchcontrollerGlobalArgs struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces *string `pulumi:"allowMultipleInterfaces"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink *string `pulumi:"bounceQuarantinedLink"`
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands []SwitchcontrollerGlobalCustomCommand `pulumi:"customCommands"`
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan *string `pulumi:"defaultVirtualSwitchVlan"`
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList *string `pulumi:"dhcpServerAccessList"`
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries []SwitchcontrollerGlobalDisableDiscovery `pulumi:"disableDiscoveries"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce *string `pulumi:"fipsEnforce"`
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization *string `pulumi:"firmwareProvisionOnAuthorization"`
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush *string `pulumi:"httpsImagePush"`
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations *string `pulumi:"logMacLimitViolations"`
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval *int `pulumi:"macAgingInterval"`
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging *string `pulumi:"macEventLogging"`
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod *int `pulumi:"macRetentionPeriod"`
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer *int `pulumi:"macViolationTimer"`
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode *string `pulumi:"quarantineMode"`
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution *string `pulumi:"snDnsResolution"`
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice *string `pulumi:"updateUserDevice"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode *string `pulumi:"vlanAllMode"`
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization *string `pulumi:"vlanOptimization"`
}

// The set of arguments for constructing a SwitchcontrollerGlobal resource.
type SwitchcontrollerGlobalArgs struct {
	// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
	AllowMultipleInterfaces pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceQuarantinedLink pulumi.StringPtrInput
	// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
	CustomCommands SwitchcontrollerGlobalCustomCommandArrayInput
	// Default VLAN for ports when added to the virtual-switch.
	DefaultVirtualSwitchVlan pulumi.StringPtrInput
	// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
	DhcpServerAccessList pulumi.StringPtrInput
	// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
	DisableDiscoveries SwitchcontrollerGlobalDisableDiscoveryArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
	FipsEnforce pulumi.StringPtrInput
	// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
	FirmwareProvisionOnAuthorization pulumi.StringPtrInput
	// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
	HttpsImagePush pulumi.StringPtrInput
	// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
	LogMacLimitViolations pulumi.StringPtrInput
	// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
	MacAgingInterval pulumi.IntPtrInput
	// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
	MacEventLogging pulumi.StringPtrInput
	// Time in hours after which an inactive MAC is removed from client DB.
	MacRetentionPeriod pulumi.IntPtrInput
	// Set timeout for Learning Limit Violations (0 = disabled).
	MacViolationTimer pulumi.IntPtrInput
	// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
	QuarantineMode pulumi.StringPtrInput
	// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
	SnDnsResolution pulumi.StringPtrInput
	// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
	UpdateUserDevice pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
	VlanAllMode pulumi.StringPtrInput
	// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
	VlanOptimization pulumi.StringPtrInput
}

func (SwitchcontrollerGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerGlobalArgs)(nil)).Elem()
}

type SwitchcontrollerGlobalInput interface {
	pulumi.Input

	ToSwitchcontrollerGlobalOutput() SwitchcontrollerGlobalOutput
	ToSwitchcontrollerGlobalOutputWithContext(ctx context.Context) SwitchcontrollerGlobalOutput
}

func (*SwitchcontrollerGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerGlobal)(nil)).Elem()
}

func (i *SwitchcontrollerGlobal) ToSwitchcontrollerGlobalOutput() SwitchcontrollerGlobalOutput {
	return i.ToSwitchcontrollerGlobalOutputWithContext(context.Background())
}

func (i *SwitchcontrollerGlobal) ToSwitchcontrollerGlobalOutputWithContext(ctx context.Context) SwitchcontrollerGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerGlobalOutput)
}

// SwitchcontrollerGlobalArrayInput is an input type that accepts SwitchcontrollerGlobalArray and SwitchcontrollerGlobalArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerGlobalArrayInput` via:
//
//	SwitchcontrollerGlobalArray{ SwitchcontrollerGlobalArgs{...} }
type SwitchcontrollerGlobalArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerGlobalArrayOutput() SwitchcontrollerGlobalArrayOutput
	ToSwitchcontrollerGlobalArrayOutputWithContext(context.Context) SwitchcontrollerGlobalArrayOutput
}

type SwitchcontrollerGlobalArray []SwitchcontrollerGlobalInput

func (SwitchcontrollerGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerGlobal)(nil)).Elem()
}

func (i SwitchcontrollerGlobalArray) ToSwitchcontrollerGlobalArrayOutput() SwitchcontrollerGlobalArrayOutput {
	return i.ToSwitchcontrollerGlobalArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerGlobalArray) ToSwitchcontrollerGlobalArrayOutputWithContext(ctx context.Context) SwitchcontrollerGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerGlobalArrayOutput)
}

// SwitchcontrollerGlobalMapInput is an input type that accepts SwitchcontrollerGlobalMap and SwitchcontrollerGlobalMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerGlobalMapInput` via:
//
//	SwitchcontrollerGlobalMap{ "key": SwitchcontrollerGlobalArgs{...} }
type SwitchcontrollerGlobalMapInput interface {
	pulumi.Input

	ToSwitchcontrollerGlobalMapOutput() SwitchcontrollerGlobalMapOutput
	ToSwitchcontrollerGlobalMapOutputWithContext(context.Context) SwitchcontrollerGlobalMapOutput
}

type SwitchcontrollerGlobalMap map[string]SwitchcontrollerGlobalInput

func (SwitchcontrollerGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerGlobal)(nil)).Elem()
}

func (i SwitchcontrollerGlobalMap) ToSwitchcontrollerGlobalMapOutput() SwitchcontrollerGlobalMapOutput {
	return i.ToSwitchcontrollerGlobalMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerGlobalMap) ToSwitchcontrollerGlobalMapOutputWithContext(ctx context.Context) SwitchcontrollerGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerGlobalMapOutput)
}

type SwitchcontrollerGlobalOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerGlobal)(nil)).Elem()
}

func (o SwitchcontrollerGlobalOutput) ToSwitchcontrollerGlobalOutput() SwitchcontrollerGlobalOutput {
	return o
}

func (o SwitchcontrollerGlobalOutput) ToSwitchcontrollerGlobalOutputWithContext(ctx context.Context) SwitchcontrollerGlobalOutput {
	return o
}

// Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) AllowMultipleInterfaces() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.AllowMultipleInterfaces }).(pulumi.StringOutput)
}

// Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
func (o SwitchcontrollerGlobalOutput) BounceQuarantinedLink() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.BounceQuarantinedLink }).(pulumi.StringOutput)
}

// List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
func (o SwitchcontrollerGlobalOutput) CustomCommands() SwitchcontrollerGlobalCustomCommandArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) SwitchcontrollerGlobalCustomCommandArrayOutput {
		return v.CustomCommands
	}).(SwitchcontrollerGlobalCustomCommandArrayOutput)
}

// Default VLAN for ports when added to the virtual-switch.
func (o SwitchcontrollerGlobalOutput) DefaultVirtualSwitchVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.DefaultVirtualSwitchVlan }).(pulumi.StringOutput)
}

// Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) DhcpServerAccessList() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.DhcpServerAccessList }).(pulumi.StringOutput)
}

// Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
func (o SwitchcontrollerGlobalOutput) DisableDiscoveries() SwitchcontrollerGlobalDisableDiscoveryArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) SwitchcontrollerGlobalDisableDiscoveryArrayOutput {
		return v.DisableDiscoveries
	}).(SwitchcontrollerGlobalDisableDiscoveryArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerGlobalOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
func (o SwitchcontrollerGlobalOutput) FipsEnforce() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.FipsEnforce }).(pulumi.StringOutput)
}

// Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) FirmwareProvisionOnAuthorization() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.FirmwareProvisionOnAuthorization }).(pulumi.StringOutput)
}

// Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) HttpsImagePush() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.HttpsImagePush }).(pulumi.StringOutput)
}

// Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) LogMacLimitViolations() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.LogMacLimitViolations }).(pulumi.StringOutput)
}

// Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
func (o SwitchcontrollerGlobalOutput) MacAgingInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.IntOutput { return v.MacAgingInterval }).(pulumi.IntOutput)
}

// Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) MacEventLogging() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.MacEventLogging }).(pulumi.StringOutput)
}

// Time in hours after which an inactive MAC is removed from client DB.
func (o SwitchcontrollerGlobalOutput) MacRetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.IntOutput { return v.MacRetentionPeriod }).(pulumi.IntOutput)
}

// Set timeout for Learning Limit Violations (0 = disabled).
func (o SwitchcontrollerGlobalOutput) MacViolationTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.IntOutput { return v.MacViolationTimer }).(pulumi.IntOutput)
}

// Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
func (o SwitchcontrollerGlobalOutput) QuarantineMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.QuarantineMode }).(pulumi.StringOutput)
}

// Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) SnDnsResolution() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.SnDnsResolution }).(pulumi.StringOutput)
}

// Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
func (o SwitchcontrollerGlobalOutput) UpdateUserDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.UpdateUserDevice }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerGlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
func (o SwitchcontrollerGlobalOutput) VlanAllMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.VlanAllMode }).(pulumi.StringOutput)
}

// FortiLink VLAN optimization. Valid values: `enable`, `disable`.
func (o SwitchcontrollerGlobalOutput) VlanOptimization() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerGlobal) pulumi.StringOutput { return v.VlanOptimization }).(pulumi.StringOutput)
}

type SwitchcontrollerGlobalArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerGlobal)(nil)).Elem()
}

func (o SwitchcontrollerGlobalArrayOutput) ToSwitchcontrollerGlobalArrayOutput() SwitchcontrollerGlobalArrayOutput {
	return o
}

func (o SwitchcontrollerGlobalArrayOutput) ToSwitchcontrollerGlobalArrayOutputWithContext(ctx context.Context) SwitchcontrollerGlobalArrayOutput {
	return o
}

func (o SwitchcontrollerGlobalArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerGlobal {
		return vs[0].([]*SwitchcontrollerGlobal)[vs[1].(int)]
	}).(SwitchcontrollerGlobalOutput)
}

type SwitchcontrollerGlobalMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerGlobal)(nil)).Elem()
}

func (o SwitchcontrollerGlobalMapOutput) ToSwitchcontrollerGlobalMapOutput() SwitchcontrollerGlobalMapOutput {
	return o
}

func (o SwitchcontrollerGlobalMapOutput) ToSwitchcontrollerGlobalMapOutputWithContext(ctx context.Context) SwitchcontrollerGlobalMapOutput {
	return o
}

func (o SwitchcontrollerGlobalMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerGlobal {
		return vs[0].(map[string]*SwitchcontrollerGlobal)[vs[1].(string)]
	}).(SwitchcontrollerGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerGlobalInput)(nil)).Elem(), &SwitchcontrollerGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerGlobalArrayInput)(nil)).Elem(), SwitchcontrollerGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerGlobalMapInput)(nil)).Elem(), SwitchcontrollerGlobalMap{})
	pulumi.RegisterOutputType(SwitchcontrollerGlobalOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerGlobalArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerGlobalMapOutput{})
}
