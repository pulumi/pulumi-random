# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RandomPasswordArgs', 'RandomPassword']

@pulumi.input_type
class RandomPasswordArgs:
    def __init__(__self__, *,
                 length: pulumi.Input[int],
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lower: Optional[pulumi.Input[bool]] = None,
                 min_lower: Optional[pulumi.Input[int]] = None,
                 min_numeric: Optional[pulumi.Input[int]] = None,
                 min_special: Optional[pulumi.Input[int]] = None,
                 min_upper: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[bool]] = None,
                 numeric: Optional[pulumi.Input[bool]] = None,
                 override_special: Optional[pulumi.Input[str]] = None,
                 special: Optional[pulumi.Input[bool]] = None,
                 upper: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RandomPassword resource.
        :param pulumi.Input[int] length: The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[bool] lower: Include lowercase alphabet characters in the result. Default value is `true`.
        :param pulumi.Input[int] min_lower: Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_numeric: Minimum number of numeric characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_special: Minimum number of special characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_upper: Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[bool] number: Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        :param pulumi.Input[bool] numeric: Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        :param pulumi.Input[bool] special: Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        :param pulumi.Input[bool] upper: Include uppercase alphabet characters in the result. Default value is `true`.
        """
        pulumi.set(__self__, "length", length)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if lower is not None:
            pulumi.set(__self__, "lower", lower)
        if min_lower is not None:
            pulumi.set(__self__, "min_lower", min_lower)
        if min_numeric is not None:
            pulumi.set(__self__, "min_numeric", min_numeric)
        if min_special is not None:
            pulumi.set(__self__, "min_special", min_special)
        if min_upper is not None:
            pulumi.set(__self__, "min_upper", min_upper)
        if number is not None:
            warnings.warn("""**NOTE**: This is deprecated, use `numeric` instead.""", DeprecationWarning)
            pulumi.log.warn("""number is deprecated: **NOTE**: This is deprecated, use `numeric` instead.""")
        if number is not None:
            pulumi.set(__self__, "number", number)
        if numeric is not None:
            pulumi.set(__self__, "numeric", numeric)
        if override_special is not None:
            pulumi.set(__self__, "override_special", override_special)
        if special is not None:
            pulumi.set(__self__, "special", special)
        if upper is not None:
            pulumi.set(__self__, "upper", upper)

    @property
    @pulumi.getter
    def length(self) -> pulumi.Input[int]:
        """
        The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: pulumi.Input[int]):
        pulumi.set(self, "length", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter
    def lower(self) -> Optional[pulumi.Input[bool]]:
        """
        Include lowercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "lower")

    @lower.setter
    def lower(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lower", value)

    @property
    @pulumi.getter(name="minLower")
    def min_lower(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_lower")

    @min_lower.setter
    def min_lower(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_lower", value)

    @property
    @pulumi.getter(name="minNumeric")
    def min_numeric(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of numeric characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_numeric")

    @min_numeric.setter
    def min_numeric(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_numeric", value)

    @property
    @pulumi.getter(name="minSpecial")
    def min_special(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of special characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_special")

    @min_special.setter
    def min_special(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_special", value)

    @property
    @pulumi.getter(name="minUpper")
    def min_upper(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_upper")

    @min_upper.setter
    def min_upper(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_upper", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""**NOTE**: This is deprecated, use `numeric` instead.""")
    def number(self) -> Optional[pulumi.Input[bool]]:
        """
        Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def numeric(self) -> Optional[pulumi.Input[bool]]:
        """
        Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        """
        return pulumi.get(self, "numeric")

    @numeric.setter
    def numeric(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "numeric", value)

    @property
    @pulumi.getter(name="overrideSpecial")
    def override_special(self) -> Optional[pulumi.Input[str]]:
        """
        Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        """
        return pulumi.get(self, "override_special")

    @override_special.setter
    def override_special(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_special", value)

    @property
    @pulumi.getter
    def special(self) -> Optional[pulumi.Input[bool]]:
        """
        Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        """
        return pulumi.get(self, "special")

    @special.setter
    def special(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "special", value)

    @property
    @pulumi.getter
    def upper(self) -> Optional[pulumi.Input[bool]]:
        """
        Include uppercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "upper")

    @upper.setter
    def upper(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "upper", value)


@pulumi.input_type
class _RandomPasswordState:
    def __init__(__self__, *,
                 bcrypt_hash: Optional[pulumi.Input[str]] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 lower: Optional[pulumi.Input[bool]] = None,
                 min_lower: Optional[pulumi.Input[int]] = None,
                 min_numeric: Optional[pulumi.Input[int]] = None,
                 min_special: Optional[pulumi.Input[int]] = None,
                 min_upper: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[bool]] = None,
                 numeric: Optional[pulumi.Input[bool]] = None,
                 override_special: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[str]] = None,
                 special: Optional[pulumi.Input[bool]] = None,
                 upper: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering RandomPassword resources.
        :param pulumi.Input[str] bcrypt_hash: A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcrypt_hash` will contain a hash of the first 72 bytes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[int] length: The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        :param pulumi.Input[bool] lower: Include lowercase alphabet characters in the result. Default value is `true`.
        :param pulumi.Input[int] min_lower: Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_numeric: Minimum number of numeric characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_special: Minimum number of special characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_upper: Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[bool] number: Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        :param pulumi.Input[bool] numeric: Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        :param pulumi.Input[str] result: The generated random string.
        :param pulumi.Input[bool] special: Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        :param pulumi.Input[bool] upper: Include uppercase alphabet characters in the result. Default value is `true`.
        """
        if bcrypt_hash is not None:
            pulumi.set(__self__, "bcrypt_hash", bcrypt_hash)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if length is not None:
            pulumi.set(__self__, "length", length)
        if lower is not None:
            pulumi.set(__self__, "lower", lower)
        if min_lower is not None:
            pulumi.set(__self__, "min_lower", min_lower)
        if min_numeric is not None:
            pulumi.set(__self__, "min_numeric", min_numeric)
        if min_special is not None:
            pulumi.set(__self__, "min_special", min_special)
        if min_upper is not None:
            pulumi.set(__self__, "min_upper", min_upper)
        if number is not None:
            warnings.warn("""**NOTE**: This is deprecated, use `numeric` instead.""", DeprecationWarning)
            pulumi.log.warn("""number is deprecated: **NOTE**: This is deprecated, use `numeric` instead.""")
        if number is not None:
            pulumi.set(__self__, "number", number)
        if numeric is not None:
            pulumi.set(__self__, "numeric", numeric)
        if override_special is not None:
            pulumi.set(__self__, "override_special", override_special)
        if result is not None:
            pulumi.set(__self__, "result", result)
        if special is not None:
            pulumi.set(__self__, "special", special)
        if upper is not None:
            pulumi.set(__self__, "upper", upper)

    @property
    @pulumi.getter(name="bcryptHash")
    def bcrypt_hash(self) -> Optional[pulumi.Input[str]]:
        """
        A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcrypt_hash` will contain a hash of the first 72 bytes.
        """
        return pulumi.get(self, "bcrypt_hash")

    @bcrypt_hash.setter
    def bcrypt_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bcrypt_hash", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter
    def length(self) -> Optional[pulumi.Input[int]]:
        """
        The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        """
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "length", value)

    @property
    @pulumi.getter
    def lower(self) -> Optional[pulumi.Input[bool]]:
        """
        Include lowercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "lower")

    @lower.setter
    def lower(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lower", value)

    @property
    @pulumi.getter(name="minLower")
    def min_lower(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_lower")

    @min_lower.setter
    def min_lower(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_lower", value)

    @property
    @pulumi.getter(name="minNumeric")
    def min_numeric(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of numeric characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_numeric")

    @min_numeric.setter
    def min_numeric(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_numeric", value)

    @property
    @pulumi.getter(name="minSpecial")
    def min_special(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of special characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_special")

    @min_special.setter
    def min_special(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_special", value)

    @property
    @pulumi.getter(name="minUpper")
    def min_upper(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_upper")

    @min_upper.setter
    def min_upper(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_upper", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""**NOTE**: This is deprecated, use `numeric` instead.""")
    def number(self) -> Optional[pulumi.Input[bool]]:
        """
        Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def numeric(self) -> Optional[pulumi.Input[bool]]:
        """
        Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        """
        return pulumi.get(self, "numeric")

    @numeric.setter
    def numeric(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "numeric", value)

    @property
    @pulumi.getter(name="overrideSpecial")
    def override_special(self) -> Optional[pulumi.Input[str]]:
        """
        Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        """
        return pulumi.get(self, "override_special")

    @override_special.setter
    def override_special(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_special", value)

    @property
    @pulumi.getter
    def result(self) -> Optional[pulumi.Input[str]]:
        """
        The generated random string.
        """
        return pulumi.get(self, "result")

    @result.setter
    def result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "result", value)

    @property
    @pulumi.getter
    def special(self) -> Optional[pulumi.Input[bool]]:
        """
        Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        """
        return pulumi.get(self, "special")

    @special.setter
    def special(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "special", value)

    @property
    @pulumi.getter
    def upper(self) -> Optional[pulumi.Input[bool]]:
        """
        Include uppercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "upper")

    @upper.setter
    def upper(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "upper", value)


class RandomPassword(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 lower: Optional[pulumi.Input[bool]] = None,
                 min_lower: Optional[pulumi.Input[int]] = None,
                 min_numeric: Optional[pulumi.Input[int]] = None,
                 min_special: Optional[pulumi.Input[int]] = None,
                 min_upper: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[bool]] = None,
                 numeric: Optional[pulumi.Input[bool]] = None,
                 override_special: Optional[pulumi.Input[str]] = None,
                 special: Optional[pulumi.Input[bool]] = None,
                 upper: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Identical to random_string.

        This resource *does* use a cryptographic random number generator.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        password = random.RandomPassword("password",
            length=16,
            special=True,
            override_special="!#$%&*()-_=+[]{}<>:?")
        example = aws.rds.Instance("example",
            instance_class=aws.rds.InstanceType.T3_MICRO,
            allocated_storage=64,
            engine="mysql",
            username="someone",
            password=password.result)
        ```

        ## Import

        You can import external passwords into your Pulumi programs as follows:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[int] length: The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        :param pulumi.Input[bool] lower: Include lowercase alphabet characters in the result. Default value is `true`.
        :param pulumi.Input[int] min_lower: Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_numeric: Minimum number of numeric characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_special: Minimum number of special characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_upper: Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[bool] number: Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        :param pulumi.Input[bool] numeric: Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        :param pulumi.Input[bool] special: Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        :param pulumi.Input[bool] upper: Include uppercase alphabet characters in the result. Default value is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RandomPasswordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Identical to random_string.

        This resource *does* use a cryptographic random number generator.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        password = random.RandomPassword("password",
            length=16,
            special=True,
            override_special="!#$%&*()-_=+[]{}<>:?")
        example = aws.rds.Instance("example",
            instance_class=aws.rds.InstanceType.T3_MICRO,
            allocated_storage=64,
            engine="mysql",
            username="someone",
            password=password.result)
        ```

        ## Import

        You can import external passwords into your Pulumi programs as follows:

        :param str resource_name: The name of the resource.
        :param RandomPasswordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RandomPasswordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 lower: Optional[pulumi.Input[bool]] = None,
                 min_lower: Optional[pulumi.Input[int]] = None,
                 min_numeric: Optional[pulumi.Input[int]] = None,
                 min_special: Optional[pulumi.Input[int]] = None,
                 min_upper: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[bool]] = None,
                 numeric: Optional[pulumi.Input[bool]] = None,
                 override_special: Optional[pulumi.Input[str]] = None,
                 special: Optional[pulumi.Input[bool]] = None,
                 upper: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RandomPasswordArgs.__new__(RandomPasswordArgs)

            __props__.__dict__["keepers"] = keepers
            if length is None and not opts.urn:
                raise TypeError("Missing required property 'length'")
            __props__.__dict__["length"] = length
            __props__.__dict__["lower"] = lower
            __props__.__dict__["min_lower"] = min_lower
            __props__.__dict__["min_numeric"] = min_numeric
            __props__.__dict__["min_special"] = min_special
            __props__.__dict__["min_upper"] = min_upper
            __props__.__dict__["number"] = number
            __props__.__dict__["numeric"] = numeric
            __props__.__dict__["override_special"] = override_special
            __props__.__dict__["special"] = special
            __props__.__dict__["upper"] = upper
            __props__.__dict__["bcrypt_hash"] = None
            __props__.__dict__["result"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["bcryptHash", "result"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RandomPassword, __self__).__init__(
            'random:index/randomPassword:RandomPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bcrypt_hash: Optional[pulumi.Input[str]] = None,
            keepers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            length: Optional[pulumi.Input[int]] = None,
            lower: Optional[pulumi.Input[bool]] = None,
            min_lower: Optional[pulumi.Input[int]] = None,
            min_numeric: Optional[pulumi.Input[int]] = None,
            min_special: Optional[pulumi.Input[int]] = None,
            min_upper: Optional[pulumi.Input[int]] = None,
            number: Optional[pulumi.Input[bool]] = None,
            numeric: Optional[pulumi.Input[bool]] = None,
            override_special: Optional[pulumi.Input[str]] = None,
            result: Optional[pulumi.Input[str]] = None,
            special: Optional[pulumi.Input[bool]] = None,
            upper: Optional[pulumi.Input[bool]] = None) -> 'RandomPassword':
        """
        Get an existing RandomPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bcrypt_hash: A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcrypt_hash` will contain a hash of the first 72 bytes.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] keepers: Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        :param pulumi.Input[int] length: The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        :param pulumi.Input[bool] lower: Include lowercase alphabet characters in the result. Default value is `true`.
        :param pulumi.Input[int] min_lower: Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_numeric: Minimum number of numeric characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_special: Minimum number of special characters in the result. Default value is `0`.
        :param pulumi.Input[int] min_upper: Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        :param pulumi.Input[bool] number: Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        :param pulumi.Input[bool] numeric: Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        :param pulumi.Input[str] result: The generated random string.
        :param pulumi.Input[bool] special: Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        :param pulumi.Input[bool] upper: Include uppercase alphabet characters in the result. Default value is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RandomPasswordState.__new__(_RandomPasswordState)

        __props__.__dict__["bcrypt_hash"] = bcrypt_hash
        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["length"] = length
        __props__.__dict__["lower"] = lower
        __props__.__dict__["min_lower"] = min_lower
        __props__.__dict__["min_numeric"] = min_numeric
        __props__.__dict__["min_special"] = min_special
        __props__.__dict__["min_upper"] = min_upper
        __props__.__dict__["number"] = number
        __props__.__dict__["numeric"] = numeric
        __props__.__dict__["override_special"] = override_special
        __props__.__dict__["result"] = result
        __props__.__dict__["special"] = special
        __props__.__dict__["upper"] = upper
        return RandomPassword(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bcryptHash")
    def bcrypt_hash(self) -> pulumi.Output[str]:
        """
        A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcrypt_hash` will contain a hash of the first 72 bytes.
        """
        return pulumi.get(self, "bcrypt_hash")

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter
    def length(self) -> pulumi.Output[int]:
        """
        The length of the string desired. The minimum value for length is 1 and, length must also be >= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
        """
        return pulumi.get(self, "length")

    @property
    @pulumi.getter
    def lower(self) -> pulumi.Output[bool]:
        """
        Include lowercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "lower")

    @property
    @pulumi.getter(name="minLower")
    def min_lower(self) -> pulumi.Output[int]:
        """
        Minimum number of lowercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_lower")

    @property
    @pulumi.getter(name="minNumeric")
    def min_numeric(self) -> pulumi.Output[int]:
        """
        Minimum number of numeric characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_numeric")

    @property
    @pulumi.getter(name="minSpecial")
    def min_special(self) -> pulumi.Output[int]:
        """
        Minimum number of special characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_special")

    @property
    @pulumi.getter(name="minUpper")
    def min_upper(self) -> pulumi.Output[int]:
        """
        Minimum number of uppercase alphabet characters in the result. Default value is `0`.
        """
        return pulumi.get(self, "min_upper")

    @property
    @pulumi.getter
    @_utilities.deprecated("""**NOTE**: This is deprecated, use `numeric` instead.""")
    def number(self) -> pulumi.Output[bool]:
        """
        Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
        """
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def numeric(self) -> pulumi.Output[bool]:
        """
        Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
        """
        return pulumi.get(self, "numeric")

    @property
    @pulumi.getter(name="overrideSpecial")
    def override_special(self) -> pulumi.Output[Optional[str]]:
        """
        Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
        """
        return pulumi.get(self, "override_special")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        """
        The generated random string.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def special(self) -> pulumi.Output[bool]:
        """
        Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
        """
        return pulumi.get(self, "special")

    @property
    @pulumi.getter
    def upper(self) -> pulumi.Output[bool]:
        """
        Include uppercase alphabet characters in the result. Default value is `true`.
        """
        return pulumi.get(self, "upper")

