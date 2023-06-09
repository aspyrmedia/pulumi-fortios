// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SystemGeoipoverride("trname", {description: "TEST for country"});
 * ```
 *
 * ## Import
 *
 * System GeoipOverride can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemGeoipoverride:SystemGeoipoverride labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemGeoipoverride:SystemGeoipoverride labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemGeoipoverride extends pulumi.CustomResource {
    /**
     * Get an existing SystemGeoipoverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemGeoipoverrideState, opts?: pulumi.CustomResourceOptions): SystemGeoipoverride {
        return new SystemGeoipoverride(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemGeoipoverride:SystemGeoipoverride';

    /**
     * Returns true if the given object is an instance of SystemGeoipoverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemGeoipoverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemGeoipoverride.__pulumiType;
    }

    /**
     * Two character Country ID code.
     */
    public readonly countryId!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
     */
    public readonly ip6Ranges!: pulumi.Output<outputs.SystemGeoipoverrideIp6Range[] | undefined>;
    /**
     * Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
     */
    public readonly ipRanges!: pulumi.Output<outputs.SystemGeoipoverrideIpRange[] | undefined>;
    /**
     * Location name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemGeoipoverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemGeoipoverrideArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemGeoipoverrideArgs | SystemGeoipoverrideState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemGeoipoverrideState | undefined;
            resourceInputs["countryId"] = state ? state.countryId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["ip6Ranges"] = state ? state.ip6Ranges : undefined;
            resourceInputs["ipRanges"] = state ? state.ipRanges : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemGeoipoverrideArgs | undefined;
            resourceInputs["countryId"] = args ? args.countryId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["ip6Ranges"] = args ? args.ip6Ranges : undefined;
            resourceInputs["ipRanges"] = args ? args.ipRanges : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemGeoipoverride.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemGeoipoverride resources.
 */
export interface SystemGeoipoverrideState {
    /**
     * Two character Country ID code.
     */
    countryId?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
     */
    ip6Ranges?: pulumi.Input<pulumi.Input<inputs.SystemGeoipoverrideIp6Range>[]>;
    /**
     * Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.SystemGeoipoverrideIpRange>[]>;
    /**
     * Location name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemGeoipoverride resource.
 */
export interface SystemGeoipoverrideArgs {
    /**
     * Two character Country ID code.
     */
    countryId?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Table of IPv6 ranges assigned to country. The structure of `ip6Range` block is documented below.
     */
    ip6Ranges?: pulumi.Input<pulumi.Input<inputs.SystemGeoipoverrideIp6Range>[]>;
    /**
     * Table of IP ranges assigned to country. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.SystemGeoipoverrideIpRange>[]>;
    /**
     * Location name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
