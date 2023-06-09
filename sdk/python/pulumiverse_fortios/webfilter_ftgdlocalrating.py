# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterFtgdlocalratingArgs', 'WebfilterFtgdlocalrating']

@pulumi.input_type
class WebfilterFtgdlocalratingArgs:
    def __init__(__self__, *,
                 rating: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterFtgdlocalrating resource.
        :param pulumi.Input[str] rating: Local rating.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] status: Enable/disable local rating. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url: URL to rate locally.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "rating", rating)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def rating(self) -> pulumi.Input[str]:
        """
        Local rating.
        """
        return pulumi.get(self, "rating")

    @rating.setter
    def rating(self, value: pulumi.Input[str]):
        pulumi.set(self, "rating", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable local rating. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL to rate locally.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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
class _WebfilterFtgdlocalratingState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 rating: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterFtgdlocalrating resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] rating: Local rating.
        :param pulumi.Input[str] status: Enable/disable local rating. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url: URL to rate locally.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if rating is not None:
            pulumi.set(__self__, "rating", rating)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def rating(self) -> Optional[pulumi.Input[str]]:
        """
        Local rating.
        """
        return pulumi.get(self, "rating")

    @rating.setter
    def rating(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rating", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable local rating. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL to rate locally.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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


class WebfilterFtgdlocalrating(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 rating: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure local FortiGuard Web Filter local ratings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebfilterFtgdlocalrating("trname",
            rating="1",
            status="enable",
            url="sgala.com")
        ```

        ## Import

        Webfilter FtgdLocalRating can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterFtgdlocalrating:WebfilterFtgdlocalrating labelname {{url}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterFtgdlocalrating:WebfilterFtgdlocalrating labelname {{url}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] rating: Local rating.
        :param pulumi.Input[str] status: Enable/disable local rating. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url: URL to rate locally.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebfilterFtgdlocalratingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure local FortiGuard Web Filter local ratings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebfilterFtgdlocalrating("trname",
            rating="1",
            status="enable",
            url="sgala.com")
        ```

        ## Import

        Webfilter FtgdLocalRating can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterFtgdlocalrating:WebfilterFtgdlocalrating labelname {{url}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterFtgdlocalrating:WebfilterFtgdlocalrating labelname {{url}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebfilterFtgdlocalratingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterFtgdlocalratingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 rating: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebfilterFtgdlocalratingArgs.__new__(WebfilterFtgdlocalratingArgs)

            __props__.__dict__["comment"] = comment
            if rating is None and not opts.urn:
                raise TypeError("Missing required property 'rating'")
            __props__.__dict__["rating"] = rating
            __props__.__dict__["status"] = status
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebfilterFtgdlocalrating, __self__).__init__(
            'fortios:index/webfilterFtgdlocalrating:WebfilterFtgdlocalrating',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            rating: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebfilterFtgdlocalrating':
        """
        Get an existing WebfilterFtgdlocalrating resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] rating: Local rating.
        :param pulumi.Input[str] status: Enable/disable local rating. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url: URL to rate locally.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterFtgdlocalratingState.__new__(_WebfilterFtgdlocalratingState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["rating"] = rating
        __props__.__dict__["status"] = status
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return WebfilterFtgdlocalrating(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def rating(self) -> pulumi.Output[str]:
        """
        Local rating.
        """
        return pulumi.get(self, "rating")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable local rating. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL to rate locally.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

