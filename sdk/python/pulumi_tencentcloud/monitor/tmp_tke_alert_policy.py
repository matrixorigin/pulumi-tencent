# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TmpTkeAlertPolicyArgs', 'TmpTkeAlertPolicy']

@pulumi.input_type
class TmpTkeAlertPolicyArgs:
    def __init__(__self__, *,
                 alert_rule: pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs'],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TmpTkeAlertPolicy resource.
        :param pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs'] alert_rule: Alarm notification channels.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        pulumi.set(__self__, "alert_rule", alert_rule)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="alertRule")
    def alert_rule(self) -> pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs']:
        """
        Alarm notification channels.
        """
        return pulumi.get(self, "alert_rule")

    @alert_rule.setter
    def alert_rule(self, value: pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs']):
        pulumi.set(self, "alert_rule", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _TmpTkeAlertPolicyState:
    def __init__(__self__, *,
                 alert_rule: Optional[pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpTkeAlertPolicy resources.
        :param pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs'] alert_rule: Alarm notification channels.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        if alert_rule is not None:
            pulumi.set(__self__, "alert_rule", alert_rule)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="alertRule")
    def alert_rule(self) -> Optional[pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs']]:
        """
        Alarm notification channels.
        """
        return pulumi.get(self, "alert_rule")

    @alert_rule.setter
    def alert_rule(self, value: Optional[pulumi.Input['TmpTkeAlertPolicyAlertRuleArgs']]):
        pulumi.set(self, "alert_rule", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class TmpTkeAlertPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_rule: Optional[pulumi.Input[pulumi.InputType['TmpTkeAlertPolicyAlertRuleArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a TmpTkeAlertPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TmpTkeAlertPolicyAlertRuleArgs']] alert_rule: Alarm notification channels.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpTkeAlertPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a TmpTkeAlertPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TmpTkeAlertPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpTkeAlertPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_rule: Optional[pulumi.Input[pulumi.InputType['TmpTkeAlertPolicyAlertRuleArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TmpTkeAlertPolicyArgs.__new__(TmpTkeAlertPolicyArgs)

            if alert_rule is None and not opts.urn:
                raise TypeError("Missing required property 'alert_rule'")
            __props__.__dict__["alert_rule"] = alert_rule
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(TmpTkeAlertPolicy, __self__).__init__(
            'tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_rule: Optional[pulumi.Input[pulumi.InputType['TmpTkeAlertPolicyAlertRuleArgs']]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'TmpTkeAlertPolicy':
        """
        Get an existing TmpTkeAlertPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TmpTkeAlertPolicyAlertRuleArgs']] alert_rule: Alarm notification channels.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpTkeAlertPolicyState.__new__(_TmpTkeAlertPolicyState)

        __props__.__dict__["alert_rule"] = alert_rule
        __props__.__dict__["instance_id"] = instance_id
        return TmpTkeAlertPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertRule")
    def alert_rule(self) -> pulumi.Output['outputs.TmpTkeAlertPolicyAlertRule']:
        """
        Alarm notification channels.
        """
        return pulumi.get(self, "alert_rule")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")
