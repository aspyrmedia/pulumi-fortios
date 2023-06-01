// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Internet Service owner. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Firewall InternetServiceOwner can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallInternetserviceowner:FirewallInternetserviceowner labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallInternetserviceowner:FirewallInternetserviceowner labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallInternetserviceowner extends pulumi.CustomResource {
    /**
     * Get an existing FirewallInternetserviceowner resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallInternetserviceownerState, opts?: pulumi.CustomResourceOptions): FirewallInternetserviceowner {
        return new FirewallInternetserviceowner(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallInternetserviceowner:FirewallInternetserviceowner';

    /**
     * Returns true if the given object is an instance of FirewallInternetserviceowner.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallInternetserviceowner {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallInternetserviceowner.__pulumiType;
    }

    /**
     * Internet Service owner ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Internet Service owner name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallInternetserviceowner resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallInternetserviceownerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallInternetserviceownerArgs | FirewallInternetserviceownerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallInternetserviceownerState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallInternetserviceownerArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallInternetserviceowner.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallInternetserviceowner resources.
 */
export interface FirewallInternetserviceownerState {
    /**
     * Internet Service owner ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Internet Service owner name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallInternetserviceowner resource.
 */
export interface FirewallInternetserviceownerArgs {
    /**
     * Internet Service owner ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Internet Service owner name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}