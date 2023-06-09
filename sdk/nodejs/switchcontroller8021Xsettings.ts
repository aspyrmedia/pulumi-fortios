// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure global 802.1X settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.Switchcontroller8021Xsettings("trname", {
 *     linkDownAuth: "set-unauth",
 *     maxReauthAttempt: 3,
 *     reauthPeriod: 12,
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController 8021XSettings can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings labelname SwitchController8021XSettings
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings labelname SwitchController8021XSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Switchcontroller8021Xsettings extends pulumi.CustomResource {
    /**
     * Get an existing Switchcontroller8021Xsettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Switchcontroller8021XsettingsState, opts?: pulumi.CustomResourceOptions): Switchcontroller8021Xsettings {
        return new Switchcontroller8021Xsettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings';

    /**
     * Returns true if the given object is an instance of Switchcontroller8021Xsettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Switchcontroller8021Xsettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Switchcontroller8021Xsettings.__pulumiType;
    }

    /**
     * Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
     */
    public readonly linkDownAuth!: pulumi.Output<string>;
    /**
     * Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
     */
    public readonly mabReauth!: pulumi.Output<string>;
    /**
     * Maximum number of authentication attempts (0 - 15, default = 3).
     */
    public readonly maxReauthAttempt!: pulumi.Output<number>;
    /**
     * Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
     */
    public readonly reauthPeriod!: pulumi.Output<number>;
    /**
     * 802.1X Tx period (seconds, default=30).
     */
    public readonly txPeriod!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Switchcontroller8021Xsettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Switchcontroller8021XsettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Switchcontroller8021XsettingsArgs | Switchcontroller8021XsettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Switchcontroller8021XsettingsState | undefined;
            resourceInputs["linkDownAuth"] = state ? state.linkDownAuth : undefined;
            resourceInputs["mabReauth"] = state ? state.mabReauth : undefined;
            resourceInputs["maxReauthAttempt"] = state ? state.maxReauthAttempt : undefined;
            resourceInputs["reauthPeriod"] = state ? state.reauthPeriod : undefined;
            resourceInputs["txPeriod"] = state ? state.txPeriod : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Switchcontroller8021XsettingsArgs | undefined;
            resourceInputs["linkDownAuth"] = args ? args.linkDownAuth : undefined;
            resourceInputs["mabReauth"] = args ? args.mabReauth : undefined;
            resourceInputs["maxReauthAttempt"] = args ? args.maxReauthAttempt : undefined;
            resourceInputs["reauthPeriod"] = args ? args.reauthPeriod : undefined;
            resourceInputs["txPeriod"] = args ? args.txPeriod : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Switchcontroller8021Xsettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Switchcontroller8021Xsettings resources.
 */
export interface Switchcontroller8021XsettingsState {
    /**
     * Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
     */
    linkDownAuth?: pulumi.Input<string>;
    /**
     * Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
     */
    mabReauth?: pulumi.Input<string>;
    /**
     * Maximum number of authentication attempts (0 - 15, default = 3).
     */
    maxReauthAttempt?: pulumi.Input<number>;
    /**
     * Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
     */
    reauthPeriod?: pulumi.Input<number>;
    /**
     * 802.1X Tx period (seconds, default=30).
     */
    txPeriod?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Switchcontroller8021Xsettings resource.
 */
export interface Switchcontroller8021XsettingsArgs {
    /**
     * Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
     */
    linkDownAuth?: pulumi.Input<string>;
    /**
     * Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
     */
    mabReauth?: pulumi.Input<string>;
    /**
     * Maximum number of authentication attempts (0 - 15, default = 3).
     */
    maxReauthAttempt?: pulumi.Input<number>;
    /**
     * Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
     */
    reauthPeriod?: pulumi.Input<number>;
    /**
     * 802.1X Tx period (seconds, default=30).
     */
    txPeriod?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
