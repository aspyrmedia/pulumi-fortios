// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall DoSpolicy
 */
export function getFirewallDoSpolicy(args: GetFirewallDoSpolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallDoSpolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallDoSpolicy:getFirewallDoSpolicy", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallDoSpolicy.
 */
export interface GetFirewallDoSpolicyArgs {
    /**
     * Specify the policyid of the desired firewall DoSpolicy.
     */
    policyid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFirewallDoSpolicy.
 */
export interface GetFirewallDoSpolicyResult {
    /**
     * Anomaly name. The structure of `anomaly` block is documented below.
     */
    readonly anomalies: outputs.GetFirewallDoSpolicyAnomaly[];
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * Destination address name from available addresses. The structure of `dstaddr` block is documented below.
     */
    readonly dstaddrs: outputs.GetFirewallDoSpolicyDstaddr[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Incoming interface name from available interfaces.
     */
    readonly interface: string;
    /**
     * Anomaly name.
     */
    readonly name: string;
    /**
     * Policy ID.
     */
    readonly policyid: number;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    readonly services: outputs.GetFirewallDoSpolicyService[];
    /**
     * Source address name from available addresses. The structure of `srcaddr` block is documented below.
     */
    readonly srcaddrs: outputs.GetFirewallDoSpolicySrcaddr[];
    /**
     * Enable/disable this anomaly.
     */
    readonly status: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewall DoSpolicy
 */
export function getFirewallDoSpolicyOutput(args: GetFirewallDoSpolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallDoSpolicyResult> {
    return pulumi.output(args).apply((a: any) => getFirewallDoSpolicy(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallDoSpolicy.
 */
export interface GetFirewallDoSpolicyOutputArgs {
    /**
     * Specify the policyid of the desired firewall DoSpolicy.
     */
    policyid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}