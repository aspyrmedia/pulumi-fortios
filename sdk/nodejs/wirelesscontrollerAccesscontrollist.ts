// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure WiFi bridge access control list. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * WirelessController AccessControlList can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelesscontrollerAccesscontrollist extends pulumi.CustomResource {
    /**
     * Get an existing WirelesscontrollerAccesscontrollist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelesscontrollerAccesscontrollistState, opts?: pulumi.CustomResourceOptions): WirelesscontrollerAccesscontrollist {
        return new WirelesscontrollerAccesscontrollist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelesscontrollerAccesscontrollist:WirelesscontrollerAccesscontrollist';

    /**
     * Returns true if the given object is an instance of WirelesscontrollerAccesscontrollist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelesscontrollerAccesscontrollist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelesscontrollerAccesscontrollist.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
     */
    public readonly layer3Ipv4Rules!: pulumi.Output<outputs.WirelesscontrollerAccesscontrollistLayer3Ipv4Rule[] | undefined>;
    /**
     * AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
     */
    public readonly layer3Ipv6Rules!: pulumi.Output<outputs.WirelesscontrollerAccesscontrollistLayer3Ipv6Rule[] | undefined>;
    /**
     * AP access control list name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `layer3Ipv4Rules` block supports:
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelesscontrollerAccesscontrollist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelesscontrollerAccesscontrollistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelesscontrollerAccesscontrollistArgs | WirelesscontrollerAccesscontrollistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelesscontrollerAccesscontrollistState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["layer3Ipv4Rules"] = state ? state.layer3Ipv4Rules : undefined;
            resourceInputs["layer3Ipv6Rules"] = state ? state.layer3Ipv6Rules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelesscontrollerAccesscontrollistArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["layer3Ipv4Rules"] = args ? args.layer3Ipv4Rules : undefined;
            resourceInputs["layer3Ipv6Rules"] = args ? args.layer3Ipv6Rules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelesscontrollerAccesscontrollist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelesscontrollerAccesscontrollist resources.
 */
export interface WirelesscontrollerAccesscontrollistState {
    /**
     * Description.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
     */
    layer3Ipv4Rules?: pulumi.Input<pulumi.Input<inputs.WirelesscontrollerAccesscontrollistLayer3Ipv4Rule>[]>;
    /**
     * AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
     */
    layer3Ipv6Rules?: pulumi.Input<pulumi.Input<inputs.WirelesscontrollerAccesscontrollistLayer3Ipv6Rule>[]>;
    /**
     * AP access control list name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `layer3Ipv4Rules` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelesscontrollerAccesscontrollist resource.
 */
export interface WirelesscontrollerAccesscontrollistArgs {
    /**
     * Description.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AP ACL layer3 ipv4 rule list. The structure of `layer3Ipv4Rules` block is documented below.
     */
    layer3Ipv4Rules?: pulumi.Input<pulumi.Input<inputs.WirelesscontrollerAccesscontrollistLayer3Ipv4Rule>[]>;
    /**
     * AP ACL layer3 ipv6 rule list. The structure of `layer3Ipv6Rules` block is documented below.
     */
    layer3Ipv6Rules?: pulumi.Input<pulumi.Input<inputs.WirelesscontrollerAccesscontrollistLayer3Ipv6Rule>[]>;
    /**
     * AP access control list name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `layer3Ipv4Rules` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}