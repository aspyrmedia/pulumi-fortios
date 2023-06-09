// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Bluetooth Low Energy profile.
 *
 * ## Import
 *
 * WirelessController BleProfile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerBleprofile:WirelesscontrollerBleprofile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerBleprofile:WirelesscontrollerBleprofile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelesscontrollerBleprofile extends pulumi.CustomResource {
    /**
     * Get an existing WirelesscontrollerBleprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelesscontrollerBleprofileState, opts?: pulumi.CustomResourceOptions): WirelesscontrollerBleprofile {
        return new WirelesscontrollerBleprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelesscontrollerBleprofile:WirelesscontrollerBleprofile';

    /**
     * Returns true if the given object is an instance of WirelesscontrollerBleprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelesscontrollerBleprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelesscontrollerBleprofile.__pulumiType;
    }

    /**
     * Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
     */
    public readonly advertising!: pulumi.Output<string>;
    /**
     * Beacon interval (default = 100 msec).
     */
    public readonly beaconInterval!: pulumi.Output<number>;
    /**
     * Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
     */
    public readonly bleScanning!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Eddystone instance ID.
     */
    public readonly eddystoneInstance!: pulumi.Output<string>;
    /**
     * Eddystone namespace ID.
     */
    public readonly eddystoneNamespace!: pulumi.Output<string>;
    /**
     * Eddystone URL.
     */
    public readonly eddystoneUrl!: pulumi.Output<string>;
    /**
     * Eddystone encoded URL hexadecimal string
     */
    public readonly eddystoneUrlEncodeHex!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly ibeaconUuid!: pulumi.Output<string>;
    /**
     * Major ID.
     */
    public readonly majorId!: pulumi.Output<number>;
    /**
     * Minor ID.
     */
    public readonly minorId!: pulumi.Output<number>;
    /**
     * Bluetooth Low Energy profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
     */
    public readonly txpower!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelesscontrollerBleprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelesscontrollerBleprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelesscontrollerBleprofileArgs | WirelesscontrollerBleprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelesscontrollerBleprofileState | undefined;
            resourceInputs["advertising"] = state ? state.advertising : undefined;
            resourceInputs["beaconInterval"] = state ? state.beaconInterval : undefined;
            resourceInputs["bleScanning"] = state ? state.bleScanning : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["eddystoneInstance"] = state ? state.eddystoneInstance : undefined;
            resourceInputs["eddystoneNamespace"] = state ? state.eddystoneNamespace : undefined;
            resourceInputs["eddystoneUrl"] = state ? state.eddystoneUrl : undefined;
            resourceInputs["eddystoneUrlEncodeHex"] = state ? state.eddystoneUrlEncodeHex : undefined;
            resourceInputs["ibeaconUuid"] = state ? state.ibeaconUuid : undefined;
            resourceInputs["majorId"] = state ? state.majorId : undefined;
            resourceInputs["minorId"] = state ? state.minorId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["txpower"] = state ? state.txpower : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelesscontrollerBleprofileArgs | undefined;
            resourceInputs["advertising"] = args ? args.advertising : undefined;
            resourceInputs["beaconInterval"] = args ? args.beaconInterval : undefined;
            resourceInputs["bleScanning"] = args ? args.bleScanning : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["eddystoneInstance"] = args ? args.eddystoneInstance : undefined;
            resourceInputs["eddystoneNamespace"] = args ? args.eddystoneNamespace : undefined;
            resourceInputs["eddystoneUrl"] = args ? args.eddystoneUrl : undefined;
            resourceInputs["eddystoneUrlEncodeHex"] = args ? args.eddystoneUrlEncodeHex : undefined;
            resourceInputs["ibeaconUuid"] = args ? args.ibeaconUuid : undefined;
            resourceInputs["majorId"] = args ? args.majorId : undefined;
            resourceInputs["minorId"] = args ? args.minorId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["txpower"] = args ? args.txpower : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelesscontrollerBleprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelesscontrollerBleprofile resources.
 */
export interface WirelesscontrollerBleprofileState {
    /**
     * Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
     */
    advertising?: pulumi.Input<string>;
    /**
     * Beacon interval (default = 100 msec).
     */
    beaconInterval?: pulumi.Input<number>;
    /**
     * Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
     */
    bleScanning?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Eddystone instance ID.
     */
    eddystoneInstance?: pulumi.Input<string>;
    /**
     * Eddystone namespace ID.
     */
    eddystoneNamespace?: pulumi.Input<string>;
    /**
     * Eddystone URL.
     */
    eddystoneUrl?: pulumi.Input<string>;
    /**
     * Eddystone encoded URL hexadecimal string
     */
    eddystoneUrlEncodeHex?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    ibeaconUuid?: pulumi.Input<string>;
    /**
     * Major ID.
     */
    majorId?: pulumi.Input<number>;
    /**
     * Minor ID.
     */
    minorId?: pulumi.Input<number>;
    /**
     * Bluetooth Low Energy profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
     */
    txpower?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelesscontrollerBleprofile resource.
 */
export interface WirelesscontrollerBleprofileArgs {
    /**
     * Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
     */
    advertising?: pulumi.Input<string>;
    /**
     * Beacon interval (default = 100 msec).
     */
    beaconInterval?: pulumi.Input<number>;
    /**
     * Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
     */
    bleScanning?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Eddystone instance ID.
     */
    eddystoneInstance?: pulumi.Input<string>;
    /**
     * Eddystone namespace ID.
     */
    eddystoneNamespace?: pulumi.Input<string>;
    /**
     * Eddystone URL.
     */
    eddystoneUrl?: pulumi.Input<string>;
    /**
     * Eddystone encoded URL hexadecimal string
     */
    eddystoneUrlEncodeHex?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    ibeaconUuid?: pulumi.Input<string>;
    /**
     * Major ID.
     */
    majorId?: pulumi.Input<number>;
    /**
     * Minor ID.
     */
    minorId?: pulumi.Input<number>;
    /**
     * Bluetooth Low Energy profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
     */
    txpower?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
