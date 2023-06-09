# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemAliasResult',
    'AwaitableGetSystemAliasResult',
    'get_system_alias',
    'get_system_alias_output',
]

@pulumi.output_type
class GetSystemAliasResult:
    """
    A collection of values returned by getSystemAlias.
    """
    def __init__(__self__, command=None, id=None, name=None, vdomparam=None):
        if command and not isinstance(command, str):
            raise TypeError("Expected argument 'command' to be a str")
        pulumi.set(__self__, "command", command)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def command(self) -> str:
        """
        Command list to execute.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Alias command name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemAliasResult(GetSystemAliasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAliasResult(
            command=self.command,
            id=self.id,
            name=self.name,
            vdomparam=self.vdomparam)


def get_system_alias(name: Optional[str] = None,
                     vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAliasResult:
    """
    Use this data source to get information on an fortios system alias


    :param str name: Specify the name of the desired system alias.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAlias:getSystemAlias', __args__, opts=opts, typ=GetSystemAliasResult).value

    return AwaitableGetSystemAliasResult(
        command=__ret__.command,
        id=__ret__.id,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_alias)
def get_system_alias_output(name: Optional[pulumi.Input[str]] = None,
                            vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemAliasResult]:
    """
    Use this data source to get information on an fortios system alias


    :param str name: Specify the name of the desired system alias.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
