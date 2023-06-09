// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure VideoFilter profile. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Videofilter Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VideofilterProfile extends pulumi.CustomResource {
    /**
     * Get an existing VideofilterProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VideofilterProfileState, opts?: pulumi.CustomResourceOptions): VideofilterProfile {
        return new VideofilterProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/videofilterProfile:VideofilterProfile';

    /**
     * Returns true if the given object is an instance of VideofilterProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VideofilterProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VideofilterProfile.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
     */
    public readonly dailymotion!: pulumi.Output<string>;
    /**
     * Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
     */
    public readonly fortiguardCategory!: pulumi.Output<outputs.VideofilterProfileFortiguardCategory>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
     */
    public readonly vimeo!: pulumi.Output<string>;
    /**
     * Enable/disable YouTube video source. Valid values: `enable`, `disable`.
     */
    public readonly youtube!: pulumi.Output<string>;
    /**
     * Set YouTube channel filter.
     */
    public readonly youtubeChannelFilter!: pulumi.Output<number>;

    /**
     * Create a VideofilterProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VideofilterProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VideofilterProfileArgs | VideofilterProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VideofilterProfileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dailymotion"] = state ? state.dailymotion : undefined;
            resourceInputs["fortiguardCategory"] = state ? state.fortiguardCategory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vimeo"] = state ? state.vimeo : undefined;
            resourceInputs["youtube"] = state ? state.youtube : undefined;
            resourceInputs["youtubeChannelFilter"] = state ? state.youtubeChannelFilter : undefined;
        } else {
            const args = argsOrState as VideofilterProfileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dailymotion"] = args ? args.dailymotion : undefined;
            resourceInputs["fortiguardCategory"] = args ? args.fortiguardCategory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vimeo"] = args ? args.vimeo : undefined;
            resourceInputs["youtube"] = args ? args.youtube : undefined;
            resourceInputs["youtubeChannelFilter"] = args ? args.youtubeChannelFilter : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VideofilterProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VideofilterProfile resources.
 */
export interface VideofilterProfileState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
     */
    dailymotion?: pulumi.Input<string>;
    /**
     * Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
     */
    fortiguardCategory?: pulumi.Input<inputs.VideofilterProfileFortiguardCategory>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
     */
    vimeo?: pulumi.Input<string>;
    /**
     * Enable/disable YouTube video source. Valid values: `enable`, `disable`.
     */
    youtube?: pulumi.Input<string>;
    /**
     * Set YouTube channel filter.
     */
    youtubeChannelFilter?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VideofilterProfile resource.
 */
export interface VideofilterProfileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
     */
    dailymotion?: pulumi.Input<string>;
    /**
     * Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
     */
    fortiguardCategory?: pulumi.Input<inputs.VideofilterProfileFortiguardCategory>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
     */
    vimeo?: pulumi.Input<string>;
    /**
     * Enable/disable YouTube video source. Valid values: `enable`, `disable`.
     */
    youtube?: pulumi.Input<string>;
    /**
     * Set YouTube channel filter.
     */
    youtubeChannelFilter?: pulumi.Input<number>;
}
