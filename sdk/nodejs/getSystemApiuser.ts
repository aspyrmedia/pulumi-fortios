// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system apiuser
 */
export function getSystemApiuser(args: GetSystemApiuserArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemApiuserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemApiuser:getSystemApiuser", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemApiuser.
 */
export interface GetSystemApiuserArgs {
    /**
     * Specify the name of the desired system apiuser.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemApiuser.
 */
export interface GetSystemApiuserResult {
    /**
     * Admin user access profile.
     */
    readonly accprofile: string;
    /**
     * Admin user password.
     */
    readonly apiKey: string;
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
     */
    readonly corsAllowOrigin: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Virtual domain name.
     */
    readonly name: string;
    /**
     * Enable/disable peer authentication.
     */
    readonly peerAuth: string;
    /**
     * Peer group name.
     */
    readonly peerGroup: string;
    /**
     * Schedule name.
     */
    readonly schedule: string;
    /**
     * Trusthost. The structure of `trusthost` block is documented below.
     */
    readonly trusthosts: outputs.GetSystemApiuserTrusthost[];
    readonly vdomparam?: string;
    /**
     * Virtual domains. The structure of `vdom` block is documented below.
     */
    readonly vdoms: outputs.GetSystemApiuserVdom[];
}
/**
 * Use this data source to get information on an fortios system apiuser
 */
export function getSystemApiuserOutput(args: GetSystemApiuserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemApiuserResult> {
    return pulumi.output(args).apply((a: any) => getSystemApiuser(a, opts))
}

/**
 * A collection of arguments for invoking getSystemApiuser.
 */
export interface GetSystemApiuserOutputArgs {
    /**
     * Specify the name of the desired system apiuser.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}