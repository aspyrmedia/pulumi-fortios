// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure access point status (rogue | accepted | suppressed).
 *
 * ## Import
 *
 * WirelessController ApStatus can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelesscontrollerApstatus extends pulumi.CustomResource {
    /**
     * Get an existing WirelesscontrollerApstatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelesscontrollerApstatusState, opts?: pulumi.CustomResourceOptions): WirelesscontrollerApstatus {
        return new WirelesscontrollerApstatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus';

    /**
     * Returns true if the given object is an instance of WirelesscontrollerApstatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelesscontrollerApstatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelesscontrollerApstatus.__pulumiType;
    }

    /**
     * Access Point's (AP's) BSSID.
     */
    public readonly bssid!: pulumi.Output<string>;
    /**
     * AP ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Access Point's (AP's) SSID.
     */
    public readonly ssid!: pulumi.Output<string>;
    /**
     * Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelesscontrollerApstatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelesscontrollerApstatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelesscontrollerApstatusArgs | WirelesscontrollerApstatusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelesscontrollerApstatusState | undefined;
            resourceInputs["bssid"] = state ? state.bssid : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ssid"] = state ? state.ssid : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelesscontrollerApstatusArgs | undefined;
            resourceInputs["bssid"] = args ? args.bssid : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ssid"] = args ? args.ssid : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelesscontrollerApstatus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelesscontrollerApstatus resources.
 */
export interface WirelesscontrollerApstatusState {
    /**
     * Access Point's (AP's) BSSID.
     */
    bssid?: pulumi.Input<string>;
    /**
     * AP ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Access Point's (AP's) SSID.
     */
    ssid?: pulumi.Input<string>;
    /**
     * Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelesscontrollerApstatus resource.
 */
export interface WirelesscontrollerApstatusArgs {
    /**
     * Access Point's (AP's) BSSID.
     */
    bssid?: pulumi.Input<string>;
    /**
     * AP ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Access Point's (AP's) SSID.
     */
    ssid?: pulumi.Input<string>;
    /**
     * Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}