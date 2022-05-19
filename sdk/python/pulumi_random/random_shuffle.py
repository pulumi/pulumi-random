# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RandomShuffleArgs', 'RandomShuffle']

@pulumi.input_type
class RandomShuffleArgs:
    def __init__(__self__, *,
                 inputs: pulumi.Input[Sequence[pulumi.Input[str]]],
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 result_count: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RandomShuffle resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inputs: The list of strings to shuffle.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
               documentation](../index.html) for more information.
        :param pulumi.Input[int] result_count: The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
               elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
               more frequently than the number of items in the input list.
        :param pulumi.Input[str] seed: Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
               list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
               different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        pulumi.set(__self__, "inputs", inputs)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if result_count is not None:
            pulumi.set(__self__, "result_count", result_count)
        if seed is not None:
            pulumi.set(__self__, "seed", seed)

    @property
    @pulumi.getter
    def inputs(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of strings to shuffle.
        """
        return pulumi.get(self, "inputs")

    @inputs.setter
    def inputs(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "inputs", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
        documentation](../index.html) for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="resultCount")
    def result_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
        elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
        more frequently than the number of items in the input list.
        """
        return pulumi.get(self, "result_count")

    @result_count.setter
    def result_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "result_count", value)

    @property
    @pulumi.getter
    def seed(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
        list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
        different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        return pulumi.get(self, "seed")

    @seed.setter
    def seed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "seed", value)


@pulumi.input_type
class _RandomShuffleState:
    def __init__(__self__, *,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 result_count: Optional[pulumi.Input[int]] = None,
                 results: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 seed: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RandomShuffle resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inputs: The list of strings to shuffle.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
               documentation](../index.html) for more information.
        :param pulumi.Input[int] result_count: The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
               elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
               more frequently than the number of items in the input list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] results: Random permutation of the list of strings given in `input`.
        :param pulumi.Input[str] seed: Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
               list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
               different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        if inputs is not None:
            pulumi.set(__self__, "inputs", inputs)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if result_count is not None:
            pulumi.set(__self__, "result_count", result_count)
        if results is not None:
            pulumi.set(__self__, "results", results)
        if seed is not None:
            pulumi.set(__self__, "seed", seed)

    @property
    @pulumi.getter
    def inputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of strings to shuffle.
        """
        return pulumi.get(self, "inputs")

    @inputs.setter
    def inputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "inputs", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
        documentation](../index.html) for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="resultCount")
    def result_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
        elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
        more frequently than the number of items in the input list.
        """
        return pulumi.get(self, "result_count")

    @result_count.setter
    def result_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "result_count", value)

    @property
    @pulumi.getter
    def results(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Random permutation of the list of strings given in `input`.
        """
        return pulumi.get(self, "results")

    @results.setter
    def results(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "results", value)

    @property
    @pulumi.getter
    def seed(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
        list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
        different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        return pulumi.get(self, "seed")

    @seed.setter
    def seed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "seed", value)


class RandomShuffle(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 result_count: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The resource `RandomShuffle` generates a random permutation of a list of strings given as an argument.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        az = random.RandomShuffle("az",
            inputs=[
                "us-west-1a",
                "us-west-1c",
                "us-west-1d",
                "us-west-1e",
            ],
            result_count=2)
        example = aws.elb.LoadBalancer("example", availability_zones=[az.results])
        # ... and other aws_elb arguments ...
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inputs: The list of strings to shuffle.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
               documentation](../index.html) for more information.
        :param pulumi.Input[int] result_count: The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
               elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
               more frequently than the number of items in the input list.
        :param pulumi.Input[str] seed: Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
               list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
               different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RandomShuffleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The resource `RandomShuffle` generates a random permutation of a list of strings given as an argument.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        az = random.RandomShuffle("az",
            inputs=[
                "us-west-1a",
                "us-west-1c",
                "us-west-1d",
                "us-west-1e",
            ],
            result_count=2)
        example = aws.elb.LoadBalancer("example", availability_zones=[az.results])
        # ... and other aws_elb arguments ...
        ```

        :param str resource_name: The name of the resource.
        :param RandomShuffleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RandomShuffleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 result_count: Optional[pulumi.Input[int]] = None,
                 seed: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RandomShuffleArgs.__new__(RandomShuffleArgs)

            if inputs is None and not opts.urn:
                raise TypeError("Missing required property 'inputs'")
            __props__.__dict__["inputs"] = inputs
            __props__.__dict__["keepers"] = keepers
            __props__.__dict__["result_count"] = result_count
            __props__.__dict__["seed"] = seed
            __props__.__dict__["results"] = None
        super(RandomShuffle, __self__).__init__(
            'random:index/randomShuffle:RandomShuffle',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            inputs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            result_count: Optional[pulumi.Input[int]] = None,
            results: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            seed: Optional[pulumi.Input[str]] = None) -> 'RandomShuffle':
        """
        Get an existing RandomShuffle resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] inputs: The list of strings to shuffle.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
               documentation](../index.html) for more information.
        :param pulumi.Input[int] result_count: The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
               elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
               more frequently than the number of items in the input list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] results: Random permutation of the list of strings given in `input`.
        :param pulumi.Input[str] seed: Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
               list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
               different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RandomShuffleState.__new__(_RandomShuffleState)

        __props__.__dict__["inputs"] = inputs
        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["result_count"] = result_count
        __props__.__dict__["results"] = results
        __props__.__dict__["seed"] = seed
        return RandomShuffle(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def inputs(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of strings to shuffle.
        """
        return pulumi.get(self, "inputs")

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
        documentation](../index.html) for more information.
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter(name="resultCount")
    def result_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some
        elements will be excluded from the result. If more items are requested, items will be repeated in the result but not
        more frequently than the number of items in the input list.
        """
        return pulumi.get(self, "result_count")

    @property
    @pulumi.getter
    def results(self) -> pulumi.Output[Sequence[str]]:
        """
        Random permutation of the list of strings given in `input`.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter
    def seed(self) -> pulumi.Output[Optional[str]]:
        """
        Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the
        list. **Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across
        different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.
        """
        return pulumi.get(self, "seed")

