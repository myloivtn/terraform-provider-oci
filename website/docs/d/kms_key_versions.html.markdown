---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_key_versions"
sidebar_current: "docs-oci-datasource-kms-key_versions"
description: |-
  Provides the list of Key Versions in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_key_versions
This data source provides the list of Key Versions in Oracle Cloud Infrastructure Kms service.

Lists all [KeyVersion](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/KeyVersion/) resources for the specified
master encryption key.

As a management operation, this call is subject to a Key Management limit that applies to the total number 
of requests across all management read operations. Key Management might throttle this call to reject an 
otherwise valid request when the total rate of management read operations exceeds 10 requests per second 
for a given tenancy.


## Example Usage

```hcl
data "oci_kms_key_versions" "test_key_versions" {
	#Required
	key_id = "${oci_kms_key.test_key.id}"
	management_endpoint = "${var.key_version_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `key_id` - (Required) The OCID of the key.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.


## Attributes Reference

The following attributes are exported:

* `key_versions` - The list of key_versions.

### KeyVersion Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key version.
* `id` - The OCID of the key version.
* `key_id` - The OCID of the master encryption key associated with this key version.
* `state` - The key version's current state.  Example: `ENABLED` 
* `key_version_id` - The OCID of the key version.
* `time_created` - The date and time this key version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: "2018-04-03T21:10:29.600Z" 
* `time_of_deletion` - An optional property to indicate when to delete the key version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key version.

