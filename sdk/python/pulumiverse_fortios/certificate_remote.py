# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CertificateRemoteArgs', 'CertificateRemote']

@pulumi.input_type
class CertificateRemoteArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CertificateRemote resource.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] remote: Remote certificate.
        :param pulumi.Input[str] source: Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if remote is not None:
            pulumi.set(__self__, "remote", remote)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter
    def remote(self) -> Optional[pulumi.Input[str]]:
        """
        Remote certificate.
        """
        return pulumi.get(self, "remote")

    @remote.setter
    def remote(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _CertificateRemoteState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CertificateRemote resources.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] remote: Remote certificate.
        :param pulumi.Input[str] source: Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if remote is not None:
            pulumi.set(__self__, "remote", remote)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter
    def remote(self) -> Optional[pulumi.Input[str]]:
        """
        Remote certificate.
        """
        return pulumi.get(self, "remote")

    @remote.setter
    def remote(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class CertificateRemote(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Remote certificate as a PEM file. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Certificate Remote can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateRemote:CertificateRemote labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateRemote:CertificateRemote labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] remote: Remote certificate.
        :param pulumi.Input[str] source: Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CertificateRemoteArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Remote certificate as a PEM file. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Certificate Remote can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateRemote:CertificateRemote labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateRemote:CertificateRemote labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CertificateRemoteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateRemoteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 remote: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateRemoteArgs.__new__(CertificateRemoteArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["range"] = range
            __props__.__dict__["remote"] = remote
            __props__.__dict__["source"] = source
            __props__.__dict__["vdomparam"] = vdomparam
        super(CertificateRemote, __self__).__init__(
            'fortios:index/certificateRemote:CertificateRemote',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            range: Optional[pulumi.Input[str]] = None,
            remote: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'CertificateRemote':
        """
        Get an existing CertificateRemote resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] remote: Remote certificate.
        :param pulumi.Input[str] source: Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateRemoteState.__new__(_CertificateRemoteState)

        __props__.__dict__["name"] = name
        __props__.__dict__["range"] = range
        __props__.__dict__["remote"] = remote
        __props__.__dict__["source"] = source
        __props__.__dict__["vdomparam"] = vdomparam
        return CertificateRemote(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def range(self) -> pulumi.Output[str]:
        """
        Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @property
    @pulumi.getter
    def remote(self) -> pulumi.Output[str]:
        """
        Remote certificate.
        """
        return pulumi.get(self, "remote")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
