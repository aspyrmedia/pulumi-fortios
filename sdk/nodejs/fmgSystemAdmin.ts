// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports modifying system admin setting for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.FmgSystemAdmin("test1", {
 *     httpPort: 80,
 *     httpsPort: 443,
 *     idleTimeout: 20,
 * });
 * ```
 */
export class FmgSystemAdmin extends pulumi.CustomResource {
    /**
     * Get an existing FmgSystemAdmin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FmgSystemAdminState, opts?: pulumi.CustomResourceOptions): FmgSystemAdmin {
        return new FmgSystemAdmin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fmgSystemAdmin:FmgSystemAdmin';

    /**
     * Returns true if the given object is an instance of FmgSystemAdmin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FmgSystemAdmin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FmgSystemAdmin.__pulumiType;
    }

    /**
     * Http port.
     */
    public readonly httpPort!: pulumi.Output<number | undefined>;
    /**
     * Https port.
     */
    public readonly httpsPort!: pulumi.Output<number | undefined>;
    /**
     * Idle Timeout(1-480 minute).
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;

    /**
     * Create a FmgSystemAdmin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FmgSystemAdminArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FmgSystemAdminArgs | FmgSystemAdminState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FmgSystemAdminState | undefined;
            resourceInputs["httpPort"] = state ? state.httpPort : undefined;
            resourceInputs["httpsPort"] = state ? state.httpsPort : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
        } else {
            const args = argsOrState as FmgSystemAdminArgs | undefined;
            resourceInputs["httpPort"] = args ? args.httpPort : undefined;
            resourceInputs["httpsPort"] = args ? args.httpsPort : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FmgSystemAdmin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FmgSystemAdmin resources.
 */
export interface FmgSystemAdminState {
    /**
     * Http port.
     */
    httpPort?: pulumi.Input<number>;
    /**
     * Https port.
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * Idle Timeout(1-480 minute).
     */
    idleTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a FmgSystemAdmin resource.
 */
export interface FmgSystemAdminArgs {
    /**
     * Http port.
     */
    httpPort?: pulumi.Input<number>;
    /**
     * Https port.
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * Idle Timeout(1-480 minute).
     */
    idleTimeout?: pulumi.Input<number>;
}
