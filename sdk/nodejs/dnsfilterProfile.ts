// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure DNS domain filter profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.DnsfilterProfile("trname", {
 *     blockAction: "redirect",
 *     blockBotnet: "disable",
 *     domainFilter: {
 *         domainFilterTable: 0,
 *     },
 *     ftgdDns: {
 *         filters: [
 *             {
 *                 action: "block",
 *                 category: 26,
 *                 id: 1,
 *                 log: "enable",
 *             },
 *             {
 *                 action: "block",
 *                 category: 61,
 *                 id: 2,
 *                 log: "enable",
 *             },
 *             {
 *                 action: "block",
 *                 category: 86,
 *                 id: 3,
 *                 log: "enable",
 *             },
 *             {
 *                 action: "block",
 *                 category: 88,
 *                 id: 4,
 *                 log: "enable",
 *             },
 *         ],
 *     },
 *     logAllDomain: "disable",
 *     redirectPortal: "0.0.0.0",
 *     safeSearch: "disable",
 *     sdnsDomainLog: "enable",
 *     sdnsFtgdErrLog: "enable",
 *     youtubeRestrict: "strict",
 * });
 * ```
 *
 * ## Import
 *
 * Dnsfilter Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/dnsfilterProfile:DnsfilterProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/dnsfilterProfile:DnsfilterProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DnsfilterProfile extends pulumi.CustomResource {
    /**
     * Get an existing DnsfilterProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsfilterProfileState, opts?: pulumi.CustomResourceOptions): DnsfilterProfile {
        return new DnsfilterProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dnsfilterProfile:DnsfilterProfile';

    /**
     * Returns true if the given object is an instance of DnsfilterProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsfilterProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsfilterProfile.__pulumiType;
    }

    /**
     * Action to take for blocked domains.
     */
    public readonly blockAction!: pulumi.Output<string>;
    /**
     * Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
     */
    public readonly blockBotnet!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * DNS translation settings. The structure of `dnsTranslation` block is documented below.
     */
    public readonly dnsTranslations!: pulumi.Output<outputs.DnsfilterProfileDnsTranslation[] | undefined>;
    /**
     * Domain filter settings. The structure of `domainFilter` block is documented below.
     */
    public readonly domainFilter!: pulumi.Output<outputs.DnsfilterProfileDomainFilter>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
     */
    public readonly externalIpBlocklists!: pulumi.Output<outputs.DnsfilterProfileExternalIpBlocklist[] | undefined>;
    /**
     * FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
     */
    public readonly ftgdDns!: pulumi.Output<outputs.DnsfilterProfileFtgdDns>;
    /**
     * Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
     */
    public readonly logAllDomain!: pulumi.Output<string>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IP address of the SDNS redirect portal.
     */
    public readonly redirectPortal!: pulumi.Output<string>;
    /**
     * IPv6 address of the SDNS redirect portal.
     */
    public readonly redirectPortal6!: pulumi.Output<string>;
    /**
     * Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
     */
    public readonly safeSearch!: pulumi.Output<string>;
    /**
     * Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
     */
    public readonly sdnsDomainLog!: pulumi.Output<string>;
    /**
     * Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
     */
    public readonly sdnsFtgdErrLog!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
     */
    public readonly youtubeRestrict!: pulumi.Output<string>;

    /**
     * Create a DnsfilterProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DnsfilterProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsfilterProfileArgs | DnsfilterProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsfilterProfileState | undefined;
            resourceInputs["blockAction"] = state ? state.blockAction : undefined;
            resourceInputs["blockBotnet"] = state ? state.blockBotnet : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dnsTranslations"] = state ? state.dnsTranslations : undefined;
            resourceInputs["domainFilter"] = state ? state.domainFilter : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["externalIpBlocklists"] = state ? state.externalIpBlocklists : undefined;
            resourceInputs["ftgdDns"] = state ? state.ftgdDns : undefined;
            resourceInputs["logAllDomain"] = state ? state.logAllDomain : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redirectPortal"] = state ? state.redirectPortal : undefined;
            resourceInputs["redirectPortal6"] = state ? state.redirectPortal6 : undefined;
            resourceInputs["safeSearch"] = state ? state.safeSearch : undefined;
            resourceInputs["sdnsDomainLog"] = state ? state.sdnsDomainLog : undefined;
            resourceInputs["sdnsFtgdErrLog"] = state ? state.sdnsFtgdErrLog : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["youtubeRestrict"] = state ? state.youtubeRestrict : undefined;
        } else {
            const args = argsOrState as DnsfilterProfileArgs | undefined;
            resourceInputs["blockAction"] = args ? args.blockAction : undefined;
            resourceInputs["blockBotnet"] = args ? args.blockBotnet : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dnsTranslations"] = args ? args.dnsTranslations : undefined;
            resourceInputs["domainFilter"] = args ? args.domainFilter : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["externalIpBlocklists"] = args ? args.externalIpBlocklists : undefined;
            resourceInputs["ftgdDns"] = args ? args.ftgdDns : undefined;
            resourceInputs["logAllDomain"] = args ? args.logAllDomain : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redirectPortal"] = args ? args.redirectPortal : undefined;
            resourceInputs["redirectPortal6"] = args ? args.redirectPortal6 : undefined;
            resourceInputs["safeSearch"] = args ? args.safeSearch : undefined;
            resourceInputs["sdnsDomainLog"] = args ? args.sdnsDomainLog : undefined;
            resourceInputs["sdnsFtgdErrLog"] = args ? args.sdnsFtgdErrLog : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["youtubeRestrict"] = args ? args.youtubeRestrict : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsfilterProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsfilterProfile resources.
 */
export interface DnsfilterProfileState {
    /**
     * Action to take for blocked domains.
     */
    blockAction?: pulumi.Input<string>;
    /**
     * Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
     */
    blockBotnet?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * DNS translation settings. The structure of `dnsTranslation` block is documented below.
     */
    dnsTranslations?: pulumi.Input<pulumi.Input<inputs.DnsfilterProfileDnsTranslation>[]>;
    /**
     * Domain filter settings. The structure of `domainFilter` block is documented below.
     */
    domainFilter?: pulumi.Input<inputs.DnsfilterProfileDomainFilter>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
     */
    externalIpBlocklists?: pulumi.Input<pulumi.Input<inputs.DnsfilterProfileExternalIpBlocklist>[]>;
    /**
     * FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
     */
    ftgdDns?: pulumi.Input<inputs.DnsfilterProfileFtgdDns>;
    /**
     * Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
     */
    logAllDomain?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * IP address of the SDNS redirect portal.
     */
    redirectPortal?: pulumi.Input<string>;
    /**
     * IPv6 address of the SDNS redirect portal.
     */
    redirectPortal6?: pulumi.Input<string>;
    /**
     * Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
     */
    safeSearch?: pulumi.Input<string>;
    /**
     * Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
     */
    sdnsDomainLog?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
     */
    sdnsFtgdErrLog?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
     */
    youtubeRestrict?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsfilterProfile resource.
 */
export interface DnsfilterProfileArgs {
    /**
     * Action to take for blocked domains.
     */
    blockAction?: pulumi.Input<string>;
    /**
     * Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
     */
    blockBotnet?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * DNS translation settings. The structure of `dnsTranslation` block is documented below.
     */
    dnsTranslations?: pulumi.Input<pulumi.Input<inputs.DnsfilterProfileDnsTranslation>[]>;
    /**
     * Domain filter settings. The structure of `domainFilter` block is documented below.
     */
    domainFilter?: pulumi.Input<inputs.DnsfilterProfileDomainFilter>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
     */
    externalIpBlocklists?: pulumi.Input<pulumi.Input<inputs.DnsfilterProfileExternalIpBlocklist>[]>;
    /**
     * FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
     */
    ftgdDns?: pulumi.Input<inputs.DnsfilterProfileFtgdDns>;
    /**
     * Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
     */
    logAllDomain?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * IP address of the SDNS redirect portal.
     */
    redirectPortal?: pulumi.Input<string>;
    /**
     * IPv6 address of the SDNS redirect portal.
     */
    redirectPortal6?: pulumi.Input<string>;
    /**
     * Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
     */
    safeSearch?: pulumi.Input<string>;
    /**
     * Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
     */
    sdnsDomainLog?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
     */
    sdnsFtgdErrLog?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
     */
    youtubeRestrict?: pulumi.Input<string>;
}
