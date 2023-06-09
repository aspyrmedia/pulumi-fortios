// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallAddrgrp6`.
 */
export function getFirewallAddrgrp6list(args?: GetFirewallAddrgrp6listArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallAddrgrp6listResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallAddrgrp6list:getFirewallAddrgrp6list", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallAddrgrp6list.
 */
export interface GetFirewallAddrgrp6listArgs {
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
 * A collection of values returned by getFirewallAddrgrp6list.
 */
export interface GetFirewallAddrgrp6listResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallAddrgrp6`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.FirewallAddrgrp6`.
 */
export function getFirewallAddrgrp6listOutput(args?: GetFirewallAddrgrp6listOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallAddrgrp6listResult> {
    return pulumi.output(args).apply((a: any) => getFirewallAddrgrp6list(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallAddrgrp6list.
 */
export interface GetFirewallAddrgrp6listOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
