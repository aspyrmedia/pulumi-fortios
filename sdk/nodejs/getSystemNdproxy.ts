// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system ndproxy
 */
export function getSystemNdproxy(args?: GetSystemNdproxyArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemNdproxyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemNdproxy:getSystemNdproxy", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemNdproxy.
 */
export interface GetSystemNdproxyArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemNdproxy.
 */
export interface GetSystemNdproxyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.
     */
    readonly members: outputs.GetSystemNdproxyMember[];
    /**
     * Enable/disable neighbor discovery proxy.
     */
    readonly status: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system ndproxy
 */
export function getSystemNdproxyOutput(args?: GetSystemNdproxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemNdproxyResult> {
    return pulumi.output(args).apply((a: any) => getSystemNdproxy(a, opts))
}

/**
 * A collection of arguments for invoking getSystemNdproxy.
 */
export interface GetSystemNdproxyOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
