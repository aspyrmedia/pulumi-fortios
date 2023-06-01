// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemVdomexception`.
 */
export function getSystemVdomexceptionlist(args?: GetSystemVdomexceptionlistArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemVdomexceptionlistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemVdomexceptionlist:getSystemVdomexceptionlist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemVdomexceptionlist.
 */
export interface GetSystemVdomexceptionlistArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemVdomexceptionlist.
 */
export interface GetSystemVdomexceptionlistResult {
    readonly filter?: string;
    /**
     * A list of the `fortios.SystemVdomexception`.
     */
    readonly fosidlists: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.SystemVdomexception`.
 */
export function getSystemVdomexceptionlistOutput(args?: GetSystemVdomexceptionlistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemVdomexceptionlistResult> {
    return pulumi.output(args).apply((a: any) => getSystemVdomexceptionlist(a, opts))
}

/**
 * A collection of arguments for invoking getSystemVdomexceptionlist.
 */
export interface GetSystemVdomexceptionlistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}