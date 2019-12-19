---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault"
sidebar_current: "docs-oci-resource-kms-vault"
description: |-
  Provides the Vault resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_vault
This resource provides the Vault resource in Oracle Cloud Infrastructure Kms service.

Creates a new vault. The type of vault you create determines key placement, pricing, and 
available options. Options include storage isolation, a dedicated service endpoint instead 
of a shared service endpoint for API calls, and either a dedicated hardware security module 
(HSM) or a multitenant HSM.

As a provisioning operation, this call is subject to a Key Management limit that applies to 
the total number of requests across all provisioning write operations. Key Management might 
throttle this call to reject an otherwise valid request when the total rate of provisioning 
write operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_vault" "test_vault" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.vault_display_name}"
	vault_type = "${var.vault_vault_type}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create this vault.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vault_type` - (Required) The type of vault to create. Each type of vault stores the key with different degrees of isolation and has different options and pricing. 
* `time_of_deletion` - (Optional) (Updatable) An optional property for the deletion time of the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

Vaults can be imported using the `id`, e.g.

```
$ terraform import oci_kms_vault.test_vault "id"
```

