// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewallshaper trafficshaper
 */
export function getFirewallshaperTrafficshaper(args: GetFirewallshaperTrafficshaperArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallshaperTrafficshaperResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallshaperTrafficshaper:getFirewallshaperTrafficshaper", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallshaperTrafficshaper.
 */
export interface GetFirewallshaperTrafficshaperArgs {
    /**
     * Specify the name of the desired firewallshaper trafficshaper.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFirewallshaperTrafficshaper.
 */
export interface GetFirewallshaperTrafficshaperResult {
    /**
     * Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
     */
    readonly bandwidthUnit: string;
    /**
     * Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
     */
    readonly diffserv: string;
    /**
     * DiffServ setting to be applied to traffic accepted by this shaper.
     */
    readonly diffservcode: string;
    /**
     * Select DSCP marking method.
     */
    readonly dscpMarkingMethod: string;
    /**
     * Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
     */
    readonly exceedBandwidth: number;
    /**
     * Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
     */
    readonly exceedClassId: number;
    /**
     * DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    readonly exceedDscp: string;
    /**
     * Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
     */
    readonly guaranteedBandwidth: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
     */
    readonly maximumBandwidth: number;
    /**
     * DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    readonly maximumDscp: string;
    /**
     * Traffic shaper name.
     */
    readonly name: string;
    /**
     * Per-packet size overhead used in rate computations.
     */
    readonly overhead: number;
    /**
     * Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
     */
    readonly perPolicy: string;
    /**
     * Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
     */
    readonly priority: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewallshaper trafficshaper
 */
export function getFirewallshaperTrafficshaperOutput(args: GetFirewallshaperTrafficshaperOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallshaperTrafficshaperResult> {
    return pulumi.output(args).apply((a: any) => getFirewallshaperTrafficshaper(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallshaperTrafficshaper.
 */
export interface GetFirewallshaperTrafficshaperOutputArgs {
    /**
     * Specify the name of the desired firewallshaper trafficshaper.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
