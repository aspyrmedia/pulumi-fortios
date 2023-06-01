// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Access Proxy virtual hosts. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Firewall AccessProxyVirtualHost can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessproxyvirtualhost:FirewallAccessproxyvirtualhost labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessproxyvirtualhost:FirewallAccessproxyvirtualhost labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAccessproxyvirtualhost extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAccessproxyvirtualhost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAccessproxyvirtualhostState, opts?: pulumi.CustomResourceOptions): FirewallAccessproxyvirtualhost {
        return new FirewallAccessproxyvirtualhost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAccessproxyvirtualhost:FirewallAccessproxyvirtualhost';

    /**
     * Returns true if the given object is an instance of FirewallAccessproxyvirtualhost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAccessproxyvirtualhost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAccessproxyvirtualhost.__pulumiType;
    }

    /**
     * The host name.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Type of host pattern. Valid values: `sub-string`, `wildcard`.
     */
    public readonly hostType!: pulumi.Output<string>;
    /**
     * Virtual host name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Access-proxy-virtual-host replacement message override group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * SSL certificate for this host.
     */
    public readonly sslCertificate!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallAccessproxyvirtualhost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAccessproxyvirtualhostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAccessproxyvirtualhostArgs | FirewallAccessproxyvirtualhostState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAccessproxyvirtualhostState | undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["hostType"] = state ? state.hostType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallAccessproxyvirtualhostArgs | undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["hostType"] = args ? args.hostType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAccessproxyvirtualhost.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAccessproxyvirtualhost resources.
 */
export interface FirewallAccessproxyvirtualhostState {
    /**
     * The host name.
     */
    host?: pulumi.Input<string>;
    /**
     * Type of host pattern. Valid values: `sub-string`, `wildcard`.
     */
    hostType?: pulumi.Input<string>;
    /**
     * Virtual host name.
     */
    name?: pulumi.Input<string>;
    /**
     * Access-proxy-virtual-host replacement message override group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * SSL certificate for this host.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAccessproxyvirtualhost resource.
 */
export interface FirewallAccessproxyvirtualhostArgs {
    /**
     * The host name.
     */
    host?: pulumi.Input<string>;
    /**
     * Type of host pattern. Valid values: `sub-string`, `wildcard`.
     */
    hostType?: pulumi.Input<string>;
    /**
     * Virtual host name.
     */
    name?: pulumi.Input<string>;
    /**
     * Access-proxy-virtual-host replacement message override group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * SSL certificate for this host.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}