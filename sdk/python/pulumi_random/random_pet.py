# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['RandomPetArgs', 'RandomPet']

@pulumi.input_type
class RandomPetArgs:
    def __init__(__self__, *,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 length: Optional[pulumi.Input[_builtins.int]] = None,
                 prefix: Optional[pulumi.Input[_builtins.str]] = None,
                 separator: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a RandomPet resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[_builtins.int] length: The length (in words) of the pet name. Defaults to 2
        :param pulumi.Input[_builtins.str] prefix: A string to prefix the name with.
        :param pulumi.Input[_builtins.str] separator: The character to separate words in the pet name. Defaults to "-"
        """
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if length is not None:
            pulumi.set(__self__, "length", length)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if separator is not None:
            pulumi.set(__self__, "separator", separator)

    @_builtins.property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "keepers", value)

    @_builtins.property
    @pulumi.getter
    def length(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The length (in words) of the pet name. Defaults to 2
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "length", value)

    @_builtins.property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A string to prefix the name with.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "prefix", value)

    @_builtins.property
    @pulumi.getter
    def separator(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The character to separate words in the pet name. Defaults to "-"
        """
        return pulumi.get(self, "separator")

    @separator.setter
    def separator(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "separator", value)


@pulumi.input_type
class _RandomPetState:
    def __init__(__self__, *,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 length: Optional[pulumi.Input[_builtins.int]] = None,
                 prefix: Optional[pulumi.Input[_builtins.str]] = None,
                 separator: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering RandomPet resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[_builtins.int] length: The length (in words) of the pet name. Defaults to 2
        :param pulumi.Input[_builtins.str] prefix: A string to prefix the name with.
        :param pulumi.Input[_builtins.str] separator: The character to separate words in the pet name. Defaults to "-"
        """
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if length is not None:
            pulumi.set(__self__, "length", length)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if separator is not None:
            pulumi.set(__self__, "separator", separator)

    @_builtins.property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "keepers", value)

    @_builtins.property
    @pulumi.getter
    def length(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The length (in words) of the pet name. Defaults to 2
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "length", value)

    @_builtins.property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A string to prefix the name with.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "prefix", value)

    @_builtins.property
    @pulumi.getter
    def separator(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The character to separate words in the pet name. Defaults to "-"
        """
        return pulumi.get(self, "separator")

    @separator.setter
    def separator(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "separator", value)


@pulumi.type_token("random:index/randomPet:RandomPet")
class RandomPet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 length: Optional[pulumi.Input[_builtins.int]] = None,
                 prefix: Optional[pulumi.Input[_builtins.str]] = None,
                 separator: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        The resource `RandomPet` generates random pet names that are intended to be used as unique identifiers for other resources.

        This resource can be used in conjunction with resources that have the `create_before_destroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        # The following example shows how to generate a unique pet name
        # for an AWS EC2 instance that changes each time a new AMI id is
        # selected.
        server = random.RandomPet("server", keepers={
            "ami_id": ami_id,
        })
        server_instance = aws.ec2.Instance("server",
            tags={
                "Name": server.id.apply(lambda id: f"web-server-{id}"),
            },
            ami=server.keepers["amiId"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[_builtins.int] length: The length (in words) of the pet name. Defaults to 2
        :param pulumi.Input[_builtins.str] prefix: A string to prefix the name with.
        :param pulumi.Input[_builtins.str] separator: The character to separate words in the pet name. Defaults to "-"
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RandomPetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The resource `RandomPet` generates random pet names that are intended to be used as unique identifiers for other resources.

        This resource can be used in conjunction with resources that have the `create_before_destroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        # The following example shows how to generate a unique pet name
        # for an AWS EC2 instance that changes each time a new AMI id is
        # selected.
        server = random.RandomPet("server", keepers={
            "ami_id": ami_id,
        })
        server_instance = aws.ec2.Instance("server",
            tags={
                "Name": server.id.apply(lambda id: f"web-server-{id}"),
            },
            ami=server.keepers["amiId"])
        ```

        :param str resource_name: The name of the resource.
        :param RandomPetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RandomPetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 length: Optional[pulumi.Input[_builtins.int]] = None,
                 prefix: Optional[pulumi.Input[_builtins.str]] = None,
                 separator: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RandomPetArgs.__new__(RandomPetArgs)

            __props__.__dict__["keepers"] = keepers
            __props__.__dict__["length"] = length
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["separator"] = separator
        super(RandomPet, __self__).__init__(
            'random:index/randomPet:RandomPet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            length: Optional[pulumi.Input[_builtins.int]] = None,
            prefix: Optional[pulumi.Input[_builtins.str]] = None,
            separator: Optional[pulumi.Input[_builtins.str]] = None) -> 'RandomPet':
        """
        Get an existing RandomPet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[_builtins.int] length: The length (in words) of the pet name. Defaults to 2
        :param pulumi.Input[_builtins.str] prefix: A string to prefix the name with.
        :param pulumi.Input[_builtins.str] separator: The character to separate words in the pet name. Defaults to "-"
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RandomPetState.__new__(_RandomPetState)

        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["length"] = length
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["separator"] = separator
        return RandomPet(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @_builtins.property
    @pulumi.getter
    def length(self) -> pulumi.Output[_builtins.int]:
        """
        The length (in words) of the pet name. Defaults to 2
        """
        return pulumi.get(self, "length")

    @_builtins.property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A string to prefix the name with.
        """
        return pulumi.get(self, "prefix")

    @_builtins.property
    @pulumi.getter
    def separator(self) -> pulumi.Output[_builtins.str]:
        """
        The character to separate words in the pet name. Defaults to "-"
        """
        return pulumi.get(self, "separator")

