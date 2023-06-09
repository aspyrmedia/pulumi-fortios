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

__all__ = [
    'GetSystemDnsResult',
    'AwaitableGetSystemDnsResult',
    'get_system_dns',
    'get_system_dns_output',
]

@pulumi.output_type
class GetSystemDnsResult:
    """
    A collection of values returned by getSystemDns.
    """
    def __init__(__self__, alt_primary=None, alt_secondary=None, cache_notfound_responses=None, dns_cache_limit=None, dns_cache_ttl=None, dns_over_tls=None, domains=None, fqdn_cache_ttl=None, fqdn_min_refresh=None, id=None, interface=None, interface_select_method=None, ip6_primary=None, ip6_secondary=None, log=None, primary=None, protocol=None, retry=None, secondary=None, server_hostnames=None, server_select_method=None, source_ip=None, ssl_certificate=None, timeout=None, vdomparam=None):
        if alt_primary and not isinstance(alt_primary, str):
            raise TypeError("Expected argument 'alt_primary' to be a str")
        pulumi.set(__self__, "alt_primary", alt_primary)
        if alt_secondary and not isinstance(alt_secondary, str):
            raise TypeError("Expected argument 'alt_secondary' to be a str")
        pulumi.set(__self__, "alt_secondary", alt_secondary)
        if cache_notfound_responses and not isinstance(cache_notfound_responses, str):
            raise TypeError("Expected argument 'cache_notfound_responses' to be a str")
        pulumi.set(__self__, "cache_notfound_responses", cache_notfound_responses)
        if dns_cache_limit and not isinstance(dns_cache_limit, int):
            raise TypeError("Expected argument 'dns_cache_limit' to be a int")
        pulumi.set(__self__, "dns_cache_limit", dns_cache_limit)
        if dns_cache_ttl and not isinstance(dns_cache_ttl, int):
            raise TypeError("Expected argument 'dns_cache_ttl' to be a int")
        pulumi.set(__self__, "dns_cache_ttl", dns_cache_ttl)
        if dns_over_tls and not isinstance(dns_over_tls, str):
            raise TypeError("Expected argument 'dns_over_tls' to be a str")
        pulumi.set(__self__, "dns_over_tls", dns_over_tls)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if fqdn_cache_ttl and not isinstance(fqdn_cache_ttl, int):
            raise TypeError("Expected argument 'fqdn_cache_ttl' to be a int")
        pulumi.set(__self__, "fqdn_cache_ttl", fqdn_cache_ttl)
        if fqdn_min_refresh and not isinstance(fqdn_min_refresh, int):
            raise TypeError("Expected argument 'fqdn_min_refresh' to be a int")
        pulumi.set(__self__, "fqdn_min_refresh", fqdn_min_refresh)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if interface_select_method and not isinstance(interface_select_method, str):
            raise TypeError("Expected argument 'interface_select_method' to be a str")
        pulumi.set(__self__, "interface_select_method", interface_select_method)
        if ip6_primary and not isinstance(ip6_primary, str):
            raise TypeError("Expected argument 'ip6_primary' to be a str")
        pulumi.set(__self__, "ip6_primary", ip6_primary)
        if ip6_secondary and not isinstance(ip6_secondary, str):
            raise TypeError("Expected argument 'ip6_secondary' to be a str")
        pulumi.set(__self__, "ip6_secondary", ip6_secondary)
        if log and not isinstance(log, str):
            raise TypeError("Expected argument 'log' to be a str")
        pulumi.set(__self__, "log", log)
        if primary and not isinstance(primary, str):
            raise TypeError("Expected argument 'primary' to be a str")
        pulumi.set(__self__, "primary", primary)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if retry and not isinstance(retry, int):
            raise TypeError("Expected argument 'retry' to be a int")
        pulumi.set(__self__, "retry", retry)
        if secondary and not isinstance(secondary, str):
            raise TypeError("Expected argument 'secondary' to be a str")
        pulumi.set(__self__, "secondary", secondary)
        if server_hostnames and not isinstance(server_hostnames, list):
            raise TypeError("Expected argument 'server_hostnames' to be a list")
        pulumi.set(__self__, "server_hostnames", server_hostnames)
        if server_select_method and not isinstance(server_select_method, str):
            raise TypeError("Expected argument 'server_select_method' to be a str")
        pulumi.set(__self__, "server_select_method", server_select_method)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if ssl_certificate and not isinstance(ssl_certificate, str):
            raise TypeError("Expected argument 'ssl_certificate' to be a str")
        pulumi.set(__self__, "ssl_certificate", ssl_certificate)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="altPrimary")
    def alt_primary(self) -> str:
        """
        Alternate primary DNS server. (This is not used as a failover DNS server.)
        """
        return pulumi.get(self, "alt_primary")

    @property
    @pulumi.getter(name="altSecondary")
    def alt_secondary(self) -> str:
        """
        Alternate secondary DNS server. (This is not used as a failover DNS server.)
        """
        return pulumi.get(self, "alt_secondary")

    @property
    @pulumi.getter(name="cacheNotfoundResponses")
    def cache_notfound_responses(self) -> str:
        """
        Enable/disable response from the DNS server when a record is not in cache.
        """
        return pulumi.get(self, "cache_notfound_responses")

    @property
    @pulumi.getter(name="dnsCacheLimit")
    def dns_cache_limit(self) -> int:
        """
        Maximum number of records in the DNS cache.
        """
        return pulumi.get(self, "dns_cache_limit")

    @property
    @pulumi.getter(name="dnsCacheTtl")
    def dns_cache_ttl(self) -> int:
        """
        Duration in seconds that the DNS cache retains information.
        """
        return pulumi.get(self, "dns_cache_ttl")

    @property
    @pulumi.getter(name="dnsOverTls")
    def dns_over_tls(self) -> str:
        """
        Enable/disable/enforce DNS over TLS.
        """
        return pulumi.get(self, "dns_over_tls")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetSystemDnsDomainResult']:
        """
        DNS search domain list separated by space (maximum 8 domains)
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="fqdnCacheTtl")
    def fqdn_cache_ttl(self) -> int:
        """
        FQDN cache time to live in seconds (0 - 86400, default = 0).
        """
        return pulumi.get(self, "fqdn_cache_ttl")

    @property
    @pulumi.getter(name="fqdnMinRefresh")
    def fqdn_min_refresh(self) -> int:
        """
        FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
        """
        return pulumi.get(self, "fqdn_min_refresh")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> str:
        """
        Specify how to select outgoing interface to reach server.
        """
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter(name="ip6Primary")
    def ip6_primary(self) -> str:
        """
        Primary DNS server IPv6 address.
        """
        return pulumi.get(self, "ip6_primary")

    @property
    @pulumi.getter(name="ip6Secondary")
    def ip6_secondary(self) -> str:
        """
        Secondary DNS server IPv6 address.
        """
        return pulumi.get(self, "ip6_secondary")

    @property
    @pulumi.getter
    def log(self) -> str:
        """
        Local DNS log setting.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def primary(self) -> str:
        """
        Primary DNS server IP address.
        """
        return pulumi.get(self, "primary")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        DNS protocols.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def retry(self) -> int:
        """
        Number of times to retry (0 - 5).
        """
        return pulumi.get(self, "retry")

    @property
    @pulumi.getter
    def secondary(self) -> str:
        """
        Secondary DNS server IP address.
        """
        return pulumi.get(self, "secondary")

    @property
    @pulumi.getter(name="serverHostnames")
    def server_hostnames(self) -> Sequence['outputs.GetSystemDnsServerHostnameResult']:
        """
        DNS server host name list. The structure of `server_hostname` block is documented below.
        """
        return pulumi.get(self, "server_hostnames")

    @property
    @pulumi.getter(name="serverSelectMethod")
    def server_select_method(self) -> str:
        """
        Specify how configured servers are prioritized.
        """
        return pulumi.get(self, "server_select_method")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        IP address used by the DNS server as its source IP.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslCertificate")
    def ssl_certificate(self) -> str:
        """
        Name of local certificate for SSL connections.
        """
        return pulumi.get(self, "ssl_certificate")

    @property
    @pulumi.getter
    def timeout(self) -> int:
        """
        DNS query timeout interval in seconds (1 - 10).
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemDnsResult(GetSystemDnsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemDnsResult(
            alt_primary=self.alt_primary,
            alt_secondary=self.alt_secondary,
            cache_notfound_responses=self.cache_notfound_responses,
            dns_cache_limit=self.dns_cache_limit,
            dns_cache_ttl=self.dns_cache_ttl,
            dns_over_tls=self.dns_over_tls,
            domains=self.domains,
            fqdn_cache_ttl=self.fqdn_cache_ttl,
            fqdn_min_refresh=self.fqdn_min_refresh,
            id=self.id,
            interface=self.interface,
            interface_select_method=self.interface_select_method,
            ip6_primary=self.ip6_primary,
            ip6_secondary=self.ip6_secondary,
            log=self.log,
            primary=self.primary,
            protocol=self.protocol,
            retry=self.retry,
            secondary=self.secondary,
            server_hostnames=self.server_hostnames,
            server_select_method=self.server_select_method,
            source_ip=self.source_ip,
            ssl_certificate=self.ssl_certificate,
            timeout=self.timeout,
            vdomparam=self.vdomparam)


def get_system_dns(vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemDnsResult:
    """
    Use this data source to get information on fortios system dns


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemDns:getSystemDns', __args__, opts=opts, typ=GetSystemDnsResult).value

    return AwaitableGetSystemDnsResult(
        alt_primary=__ret__.alt_primary,
        alt_secondary=__ret__.alt_secondary,
        cache_notfound_responses=__ret__.cache_notfound_responses,
        dns_cache_limit=__ret__.dns_cache_limit,
        dns_cache_ttl=__ret__.dns_cache_ttl,
        dns_over_tls=__ret__.dns_over_tls,
        domains=__ret__.domains,
        fqdn_cache_ttl=__ret__.fqdn_cache_ttl,
        fqdn_min_refresh=__ret__.fqdn_min_refresh,
        id=__ret__.id,
        interface=__ret__.interface,
        interface_select_method=__ret__.interface_select_method,
        ip6_primary=__ret__.ip6_primary,
        ip6_secondary=__ret__.ip6_secondary,
        log=__ret__.log,
        primary=__ret__.primary,
        protocol=__ret__.protocol,
        retry=__ret__.retry,
        secondary=__ret__.secondary,
        server_hostnames=__ret__.server_hostnames,
        server_select_method=__ret__.server_select_method,
        source_ip=__ret__.source_ip,
        ssl_certificate=__ret__.ssl_certificate,
        timeout=__ret__.timeout,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_dns)
def get_system_dns_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemDnsResult]:
    """
    Use this data source to get information on fortios system dns


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
