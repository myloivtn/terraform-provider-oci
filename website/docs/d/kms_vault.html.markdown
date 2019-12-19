---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault"
sidebar_current: "docs-oci-datasource-kms-vault"
description: |-
  Provides details about a specific Vault in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_vault
This data source provides details about a specific Vault resource in Oracle Cloud Infrastructure Kms service.

Gets the specified vault's configuration information.

As a provisioning operation, this call is subject to a Key Management limit that applies to 
the total number of requests across all provisioning read operations. Key Management might 
throttle this call to reject an otherwise valid request when the total rate of provisioning 
read operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
data "oci_kms_vault" "test_vault" {
	#Required
	vault_id = "${oci_kms_vault.test_vault.id}"
}
```

## Argument Reference

The following arguments are supported:

* `vault_id` - (Required) The OCID of the vault.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains a particular vault.
* `crypto_endpoint` - The service endpoint to perform cryptographic operations against. Cryptographic operations include  [Encrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/EncryptedData/Encrypt), [Decrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/DecryptedData/Decrypt),  and [GenerateDataEncryptionKey](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/GeneratedKey/GenerateDataEncryptionKey) operations.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the vault.
* `management_endpoint` - The service endpoint to perform management operations against. Management operations include "Create," "Update," "List," "Get," and "Delete" operations. 
* `state` - The vault's current state.  Example: `DELETED` 
* `time_created` - The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property to indicate when to delete the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `vault_type` - The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing. 

