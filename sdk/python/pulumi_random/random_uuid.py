# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities, tables

class RandomUuid(pulumi.CustomResource):
    """
    The resource `random_uuid` generates random uuid string that is intended to be
    used as unique identifiers for other resources.
    
    This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
    for use with services needed a unique string identifier.
    
    """
    def __init__(__self__, __name__, __opts__=None, keepers=None):
        """Create a RandomUuid resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['keepers'] = keepers

        __props__['result'] = None

        super(RandomUuid, __self__).__init__(
            'random:index/randomUuid:RandomUuid',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

