---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_key_version"
sidebar_current: "docs-oci-resource-kms-key_version"
description: |-
  Provides the Key Version resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_key_version
This resource provides the Key Version resource in Oracle Cloud Infrastructure Kms service.

Generates a new [KeyVersion](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/KeyVersion/) resource that provides new cryptographic
material for a master encryption key. The key must be in an ENABLED state to be rotated.

As a management operation, this call is subject to a Key Management limit that applies to the total number 
of requests across all  management write operations. Key Management might throttle this call to reject an 
otherwise valid request when the total rate of management write operations exceeds 10 requests per second 
for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_key_version" "test_key_version" {
	#Required
	key_id = "${oci_kms_key.test_key.id}"
	management_endpoint = "${var.key_version_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `key_id` - (Required) The OCID of the key.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
* `time_of_deletion` - (Optional) (Updatable) An optional property for the deletion time of the key version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key version.
* `id` - The OCID of the key version.
* `key_id` - The OCID of the master encryption key associated with this key version.
* `state` - The key version's current state.  Example: `ENABLED` 
* `time_created` - The date and time this key version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: "2018-04-03T21:10:29.600Z" 
* `time_of_deletion` - An optional property to indicate when to delete the key version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key version.

## Import

KeyVersions can be imported using the `id`, e.g.

```
$ terraform import oci_kms_key_version.test_key_version "managementEndpoint/{managementEndpoint}/keys/{keyId}/keyVersions/{keyVersionId}" 
```

