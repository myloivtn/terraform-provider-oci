---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_generated_key"
sidebar_current: "docs-oci-resource-kms-generated_key"
description: |-
  Provides the Generated Key resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_generated_key
This resource provides the Generated Key resource in Oracle Cloud Infrastructure Kms service.

Generates a key that you can use to encrypt or decrypt data.


## Example Usage

```hcl
resource "oci_kms_generated_key" "test_generated_key" {
	#Required
	crypto_endpoint = "${var.generated_key_crypto_endpoint}"
	include_plaintext_key = "${var.generated_key_include_plaintext_key}"
	key_id = "${oci_kms_key.test_key.id}"
	key_shape {
		#Required
		algorithm = "${var.generated_key_key_shape_algorithm}"
		length = "${var.generated_key_key_shape_length}"
	}

	#Optional
	associated_data = "${var.generated_key_associated_data}"
	logging_context = "${var.generated_key_logging_context}"
}
```

## Argument Reference

The following arguments are supported:

* `associated_data` - (Optional) Information that can be used to provide an encryption context for the encrypted data.  The length of the string representation of the associated data must be fewer than 4096  characters. 
* `crypto_endpoint` - (Required) The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. see Vault Crypto endpoint.
* `include_plaintext_key` - (Required) If true, the generated key is also returned unencrypted.
* `key_id` - (Required) The OCID of the master encryption key to encrypt the generated data encryption key with.
* `key_shape` - (Required) 
	* `algorithm` - (Required) The algorithm used by a key's key versions to encrypt or decrypt.
	* `length` - (Required) The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `logging_context` - (Optional) Information that provides context for audit logging. You can provide this additional  data by formatting it as key-value pairs to include in audit logs when audit logging is enabled. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ciphertext` - The encrypted data encryption key generated from a master encryption key.
* `plaintext` - The plaintext data encryption key, a base64-encoded sequence of random bytes, which is  included if the [GenerateDataEncryptionKey](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/GeneratedKey/GenerateDataEncryptionKey)  request includes the `includePlaintextKey` parameter and sets its value to "true".
* `plaintext_checksum` - The checksum of the plaintext data encryption key, which is included if the  [GenerateDataEncryptionKey](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/GeneratedKey/GenerateDataEncryptionKey)  request includes the `includePlaintextKey` parameter and sets its value to "true".

## Import

Import is not supported for this resource.

