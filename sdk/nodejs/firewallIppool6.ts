// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPv6 IP pools.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.FirewallIppool6("trname", {
 *     endip: "2001:3ca1:10f:1a:121b::19",
 *     startip: "2001:3ca1:10f:1a:121b::10",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Ippool6 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallIppool6:FirewallIppool6 labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallIppool6:FirewallIppool6 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallIppool6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallIppool6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallIppool6State, opts?: pulumi.CustomResourceOptions): FirewallIppool6 {
        return new FirewallIppool6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallIppool6:FirewallIppool6';

    /**
     * Returns true if the given object is an instance of FirewallIppool6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallIppool6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallIppool6.__pulumiType;
    }

    /**
     * Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
     */
    public readonly addNat46Route!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    public readonly endip!: pulumi.Output<string>;
    /**
     * IPv6 IP pool name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable NAT46. Valid values: `disable`, `enable`.
     */
    public readonly nat46!: pulumi.Output<string>;
    /**
     * First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    public readonly startip!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallIppool6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallIppool6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallIppool6Args | FirewallIppool6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallIppool6State | undefined;
            resourceInputs["addNat46Route"] = state ? state.addNat46Route : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["endip"] = state ? state.endip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nat46"] = state ? state.nat46 : undefined;
            resourceInputs["startip"] = state ? state.startip : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallIppool6Args | undefined;
            if ((!args || args.endip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endip'");
            }
            if ((!args || args.startip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startip'");
            }
            resourceInputs["addNat46Route"] = args ? args.addNat46Route : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["endip"] = args ? args.endip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nat46"] = args ? args.nat46 : undefined;
            resourceInputs["startip"] = args ? args.startip : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallIppool6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallIppool6 resources.
 */
export interface FirewallIppool6State {
    /**
     * Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
     */
    addNat46Route?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    endip?: pulumi.Input<string>;
    /**
     * IPv6 IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NAT46. Valid values: `disable`, `enable`.
     */
    nat46?: pulumi.Input<string>;
    /**
     * First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    startip?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallIppool6 resource.
 */
export interface FirewallIppool6Args {
    /**
     * Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
     */
    addNat46Route?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    endip: pulumi.Input<string>;
    /**
     * IPv6 IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NAT46. Valid values: `disable`, `enable`.
     */
    nat46?: pulumi.Input<string>;
    /**
     * First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
     */
    startip: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}