// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system dnsserver
 */
export function getSystemDnsserver(args: GetSystemDnsserverArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemDnsserverResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemDnsserver:getSystemDnsserver", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemDnsserver.
 */
export interface GetSystemDnsserverArgs {
    /**
     * Specify the name of the desired system dnsserver.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemDnsserver.
 */
export interface GetSystemDnsserverResult {
    /**
     * DNS filter profile.
     */
    readonly dnsfilterProfile: string;
    /**
     * DNS over HTTPS.
     */
    readonly doh: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * DNS server mode.
     */
    readonly mode: string;
    /**
     * DNS server name.
     */
    readonly name: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system dnsserver
 */
export function getSystemDnsserverOutput(args: GetSystemDnsserverOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemDnsserverResult> {
    return pulumi.output(args).apply((a: any) => getSystemDnsserver(a, opts))
}

/**
 * A collection of arguments for invoking getSystemDnsserver.
 */
export interface GetSystemDnsserverOutputArgs {
    /**
     * Specify the name of the desired system dnsserver.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
