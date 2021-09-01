# coding=utf-8
# *** WARNING: this file was generated by . ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
import pulumi_py_tests

__all__ = [
    'ListStorageAccountKeysResult',
    'AwaitableListStorageAccountKeysResult',
    'list_storage_account_keys',
    'list_storage_account_keys_output',
]

@pulumi.output_type
class ListStorageAccountKeysResult:
    """
    The response from the ListKeys operation.
    """
    def __init__(__self__, keys=None):
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)

    @property
    @pulumi.getter
    def keys(self) -> Sequence['pulumi_py_tests.codegentest.outputs.StorageAccountKeyResponse']:
        """
        Gets the list of storage account keys and their properties for the specified storage account.
        """
        return pulumi.get(self, "keys")


class AwaitableListStorageAccountKeysResult(ListStorageAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListStorageAccountKeysResult(
            keys=self.keys)


def list_storage_account_keys(account_name: Optional[str] = None,
                              expand: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListStorageAccountKeysResult:
    """
    The response from the ListKeys operation.
    API Version: 2021-02-01.


    :param str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param str expand: Specifies type of the key to be listed. Possible value is kerb.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:codegentest:listStorageAccountKeys', __args__, opts=opts, typ=ListStorageAccountKeysResult).value

    return AwaitableListStorageAccountKeysResult(
        keys=__ret__.keys)


@_utilities.lift_output_func(list_storage_account_keys)
def list_storage_account_keys_output(account_name: Optional[pulumi.Input[str]] = None,
                                     expand: Optional[pulumi.Input[Optional[str]]] = None,
                                     resource_group_name: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListStorageAccountKeysResult]:
    """
    The response from the ListKeys operation.
    API Version: 2021-02-01.


    :param str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param str expand: Specifies type of the key to be listed. Possible value is kerb.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    ...
