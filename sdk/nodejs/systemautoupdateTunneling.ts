// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure web proxy tunnelling for the FDN.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SystemautoupdateTunneling("trname", {
 *     port: 0,
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * SystemAutoupdate Tunneling can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemautoupdateTunneling:SystemautoupdateTunneling labelname SystemAutoupdateTunneling
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemautoupdateTunneling:SystemautoupdateTunneling labelname SystemAutoupdateTunneling
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemautoupdateTunneling extends pulumi.CustomResource {
    /**
     * Get an existing SystemautoupdateTunneling resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemautoupdateTunnelingState, opts?: pulumi.CustomResourceOptions): SystemautoupdateTunneling {
        return new SystemautoupdateTunneling(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemautoupdateTunneling:SystemautoupdateTunneling';

    /**
     * Returns true if the given object is an instance of SystemautoupdateTunneling.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemautoupdateTunneling {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemautoupdateTunneling.__pulumiType;
    }

    /**
     * Web proxy IP address or FQDN.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Web proxy password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Web proxy port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Web proxy username.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemautoupdateTunneling resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemautoupdateTunnelingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemautoupdateTunnelingArgs | SystemautoupdateTunnelingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemautoupdateTunnelingState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemautoupdateTunnelingArgs | undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemautoupdateTunneling.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemautoupdateTunneling resources.
 */
export interface SystemautoupdateTunnelingState {
    /**
     * Web proxy IP address or FQDN.
     */
    address?: pulumi.Input<string>;
    /**
     * Web proxy password.
     */
    password?: pulumi.Input<string>;
    /**
     * Web proxy port.
     */
    port?: pulumi.Input<number>;
    /**
     * Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Web proxy username.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemautoupdateTunneling resource.
 */
export interface SystemautoupdateTunnelingArgs {
    /**
     * Web proxy IP address or FQDN.
     */
    address?: pulumi.Input<string>;
    /**
     * Web proxy password.
     */
    password?: pulumi.Input<string>;
    /**
     * Web proxy port.
     */
    port?: pulumi.Input<number>;
    /**
     * Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Web proxy username.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
