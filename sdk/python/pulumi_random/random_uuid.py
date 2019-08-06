# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class RandomUuid(pulumi.CustomResource):
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new uuid to be generated. See
    the main provider documentation for more information.
    """
    result: pulumi.Output[str]
    """
    The generated uuid presented in string format.
    """
    def __init__(__self__, resource_name, opts=None, keepers=None, __name__=None, __opts__=None):
        """
        The resource `random_uuid` generates random uuid string that is intended to be
        used as unique identifiers for other resources.
        
        This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
        for use with services needed a unique string identifier.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new uuid to be generated. See
               the main provider documentation for more information.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/uuid.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['keepers'] = keepers

        __props__['result'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RandomUuid, __self__).__init__(
            'random:index/randomUuid:RandomUuid',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

