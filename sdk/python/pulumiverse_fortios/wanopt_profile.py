# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['WanoptProfileArgs', 'WanoptProfile']

@pulumi.input_type
class WanoptProfileArgs:
    def __init__(__self__, *,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input['WanoptProfileCifsArgs']] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input['WanoptProfileFtpArgs']] = None,
                 http: Optional[pulumi.Input['WanoptProfileHttpArgs']] = None,
                 mapi: Optional[pulumi.Input['WanoptProfileMapiArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input['WanoptProfileTcpArgs']] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WanoptProfile resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input['WanoptProfileCifsArgs'] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input['WanoptProfileFtpArgs'] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input['WanoptProfileHttpArgs'] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input['WanoptProfileMapiArgs'] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input['WanoptProfileTcpArgs'] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_group is not None:
            pulumi.set(__self__, "auth_group", auth_group)
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftp is not None:
            pulumi.set(__self__, "ftp", ftp)
        if http is not None:
            pulumi.set(__self__, "http", http)
        if mapi is not None:
            pulumi.set(__self__, "mapi", mapi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tcp is not None:
            pulumi.set(__self__, "tcp", tcp)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @auth_group.setter
    def auth_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_group", value)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input['WanoptProfileCifsArgs']]:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input['WanoptProfileCifsArgs']]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftp(self) -> Optional[pulumi.Input['WanoptProfileFtpArgs']]:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @ftp.setter
    def ftp(self, value: Optional[pulumi.Input['WanoptProfileFtpArgs']]):
        pulumi.set(self, "ftp", value)

    @property
    @pulumi.getter
    def http(self) -> Optional[pulumi.Input['WanoptProfileHttpArgs']]:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: Optional[pulumi.Input['WanoptProfileHttpArgs']]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def mapi(self) -> Optional[pulumi.Input['WanoptProfileMapiArgs']]:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @mapi.setter
    def mapi(self, value: Optional[pulumi.Input['WanoptProfileMapiArgs']]):
        pulumi.set(self, "mapi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tcp(self) -> Optional[pulumi.Input['WanoptProfileTcpArgs']]:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @tcp.setter
    def tcp(self, value: Optional[pulumi.Input['WanoptProfileTcpArgs']]):
        pulumi.set(self, "tcp", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

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
class _WanoptProfileState:
    def __init__(__self__, *,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input['WanoptProfileCifsArgs']] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input['WanoptProfileFtpArgs']] = None,
                 http: Optional[pulumi.Input['WanoptProfileHttpArgs']] = None,
                 mapi: Optional[pulumi.Input['WanoptProfileMapiArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input['WanoptProfileTcpArgs']] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WanoptProfile resources.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input['WanoptProfileCifsArgs'] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input['WanoptProfileFtpArgs'] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input['WanoptProfileHttpArgs'] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input['WanoptProfileMapiArgs'] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input['WanoptProfileTcpArgs'] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_group is not None:
            pulumi.set(__self__, "auth_group", auth_group)
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftp is not None:
            pulumi.set(__self__, "ftp", ftp)
        if http is not None:
            pulumi.set(__self__, "http", http)
        if mapi is not None:
            pulumi.set(__self__, "mapi", mapi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tcp is not None:
            pulumi.set(__self__, "tcp", tcp)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @auth_group.setter
    def auth_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_group", value)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input['WanoptProfileCifsArgs']]:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input['WanoptProfileCifsArgs']]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftp(self) -> Optional[pulumi.Input['WanoptProfileFtpArgs']]:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @ftp.setter
    def ftp(self, value: Optional[pulumi.Input['WanoptProfileFtpArgs']]):
        pulumi.set(self, "ftp", value)

    @property
    @pulumi.getter
    def http(self) -> Optional[pulumi.Input['WanoptProfileHttpArgs']]:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: Optional[pulumi.Input['WanoptProfileHttpArgs']]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def mapi(self) -> Optional[pulumi.Input['WanoptProfileMapiArgs']]:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @mapi.setter
    def mapi(self, value: Optional[pulumi.Input['WanoptProfileMapiArgs']]):
        pulumi.set(self, "mapi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tcp(self) -> Optional[pulumi.Input['WanoptProfileTcpArgs']]:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @tcp.setter
    def tcp(self, value: Optional[pulumi.Input['WanoptProfileTcpArgs']]):
        pulumi.set(self, "tcp", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

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


class WanoptProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input[pulumi.InputType['WanoptProfileCifsArgs']]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileFtpArgs']]] = None,
                 http: Optional[pulumi.Input[pulumi.InputType['WanoptProfileHttpArgs']]] = None,
                 mapi: Optional[pulumi.Input[pulumi.InputType['WanoptProfileMapiArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileTcpArgs']]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WAN optimization profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WanoptProfile("trname",
            cifs=fortios.WanoptProfileCifsArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=445,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            comments="test",
            ftp=fortios.WanoptProfileFtpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=21,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            http=fortios.WanoptProfileHttpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=80,
                prefer_chunking="fix",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_non_http="disable",
                tunnel_sharing="private",
                unknown_http_version="tunnel",
            ),
            mapi=fortios.WanoptProfileMapiArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=135,
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            tcp=fortios.WanoptProfileTcpArgs(
                byte_caching="disable",
                byte_caching_opt="mem-only",
                log_traffic="enable",
                port="1-65535",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_sharing="private",
            ),
            transparent="enable")
        ```

        ## Import

        Wanopt Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input[pulumi.InputType['WanoptProfileCifsArgs']] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[pulumi.InputType['WanoptProfileFtpArgs']] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[pulumi.InputType['WanoptProfileHttpArgs']] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input[pulumi.InputType['WanoptProfileMapiArgs']] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[pulumi.InputType['WanoptProfileTcpArgs']] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WanoptProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WAN optimization profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WanoptProfile("trname",
            cifs=fortios.WanoptProfileCifsArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=445,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            comments="test",
            ftp=fortios.WanoptProfileFtpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=21,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            http=fortios.WanoptProfileHttpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=80,
                prefer_chunking="fix",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_non_http="disable",
                tunnel_sharing="private",
                unknown_http_version="tunnel",
            ),
            mapi=fortios.WanoptProfileMapiArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=135,
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            tcp=fortios.WanoptProfileTcpArgs(
                byte_caching="disable",
                byte_caching_opt="mem-only",
                log_traffic="enable",
                port="1-65535",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_sharing="private",
            ),
            transparent="enable")
        ```

        ## Import

        Wanopt Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WanoptProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WanoptProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input[pulumi.InputType['WanoptProfileCifsArgs']]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileFtpArgs']]] = None,
                 http: Optional[pulumi.Input[pulumi.InputType['WanoptProfileHttpArgs']]] = None,
                 mapi: Optional[pulumi.Input[pulumi.InputType['WanoptProfileMapiArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileTcpArgs']]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WanoptProfileArgs.__new__(WanoptProfileArgs)

            __props__.__dict__["auth_group"] = auth_group
            __props__.__dict__["cifs"] = cifs
            __props__.__dict__["comments"] = comments
            __props__.__dict__["ftp"] = ftp
            __props__.__dict__["http"] = http
            __props__.__dict__["mapi"] = mapi
            __props__.__dict__["name"] = name
            __props__.__dict__["tcp"] = tcp
            __props__.__dict__["transparent"] = transparent
            __props__.__dict__["vdomparam"] = vdomparam
        super(WanoptProfile, __self__).__init__(
            'fortios:index/wanoptProfile:WanoptProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_group: Optional[pulumi.Input[str]] = None,
            cifs: Optional[pulumi.Input[pulumi.InputType['WanoptProfileCifsArgs']]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            ftp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileFtpArgs']]] = None,
            http: Optional[pulumi.Input[pulumi.InputType['WanoptProfileHttpArgs']]] = None,
            mapi: Optional[pulumi.Input[pulumi.InputType['WanoptProfileMapiArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tcp: Optional[pulumi.Input[pulumi.InputType['WanoptProfileTcpArgs']]] = None,
            transparent: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WanoptProfile':
        """
        Get an existing WanoptProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input[pulumi.InputType['WanoptProfileCifsArgs']] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[pulumi.InputType['WanoptProfileFtpArgs']] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[pulumi.InputType['WanoptProfileHttpArgs']] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input[pulumi.InputType['WanoptProfileMapiArgs']] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[pulumi.InputType['WanoptProfileTcpArgs']] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WanoptProfileState.__new__(_WanoptProfileState)

        __props__.__dict__["auth_group"] = auth_group
        __props__.__dict__["cifs"] = cifs
        __props__.__dict__["comments"] = comments
        __props__.__dict__["ftp"] = ftp
        __props__.__dict__["http"] = http
        __props__.__dict__["mapi"] = mapi
        __props__.__dict__["name"] = name
        __props__.__dict__["tcp"] = tcp
        __props__.__dict__["transparent"] = transparent
        __props__.__dict__["vdomparam"] = vdomparam
        return WanoptProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> pulumi.Output[str]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @property
    @pulumi.getter
    def cifs(self) -> pulumi.Output['outputs.WanoptProfileCifs']:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def ftp(self) -> pulumi.Output['outputs.WanoptProfileFtp']:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @property
    @pulumi.getter
    def http(self) -> pulumi.Output['outputs.WanoptProfileHttp']:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @property
    @pulumi.getter
    def mapi(self) -> pulumi.Output['outputs.WanoptProfileMapi']:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tcp(self) -> pulumi.Output['outputs.WanoptProfileTcp']:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @property
    @pulumi.getter
    def transparent(self) -> pulumi.Output[str]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
