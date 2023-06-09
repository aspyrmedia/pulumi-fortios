// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios systemsnmp community
 */
export function getSystemsnmpCommunity(args: GetSystemsnmpCommunityArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemsnmpCommunityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemsnmpCommunity:getSystemsnmpCommunity", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemsnmpCommunity.
 */
export interface GetSystemsnmpCommunityArgs {
    /**
     * Specify the fosid of the desired systemsnmp community.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemsnmpCommunity.
 */
export interface GetSystemsnmpCommunityResult {
    /**
     * SNMP trap events.
     */
    readonly events: string;
    /**
     * Community ID.
     */
    readonly fosid: number;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    readonly hosts: outputs.GetSystemsnmpCommunityHost[];
    /**
     * Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
     */
    readonly hosts6s: outputs.GetSystemsnmpCommunityHosts6[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * SNMP access control MIB view.
     */
    readonly mibView: string;
    /**
     * VDOM name
     */
    readonly name: string;
    /**
     * SNMP v1 query port (default = 161).
     */
    readonly queryV1Port: number;
    /**
     * Enable/disable SNMP v1 queries.
     */
    readonly queryV1Status: string;
    /**
     * SNMP v2c query port (default = 161).
     */
    readonly queryV2cPort: number;
    /**
     * Enable/disable SNMP v2c queries.
     */
    readonly queryV2cStatus: string;
    /**
     * Enable/disable this SNMP community.
     */
    readonly status: string;
    /**
     * SNMP v1 trap local port (default = 162).
     */
    readonly trapV1Lport: number;
    /**
     * SNMP v1 trap remote port (default = 162).
     */
    readonly trapV1Rport: number;
    /**
     * Enable/disable SNMP v1 traps.
     */
    readonly trapV1Status: string;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    readonly trapV2cLport: number;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    readonly trapV2cRport: number;
    /**
     * Enable/disable SNMP v2c traps.
     */
    readonly trapV2cStatus: string;
    readonly vdomparam?: string;
    /**
     * SNMP access control VDOMs. The structure of `vdoms` block is documented below.
     */
    readonly vdoms: outputs.GetSystemsnmpCommunityVdom[];
}
/**
 * Use this data source to get information on an fortios systemsnmp community
 */
export function getSystemsnmpCommunityOutput(args: GetSystemsnmpCommunityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemsnmpCommunityResult> {
    return pulumi.output(args).apply((a: any) => getSystemsnmpCommunity(a, opts))
}

/**
 * A collection of arguments for invoking getSystemsnmpCommunity.
 */
export interface GetSystemsnmpCommunityOutputArgs {
    /**
     * Specify the fosid of the desired systemsnmp community.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
