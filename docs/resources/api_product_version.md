---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stevan_api_product_version Resource - stevan"
subcategory: ""
description: |-
  
---

# stevan_api_product_version (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The version of the API product

### Optional

- `api_product_id` (String) The API product identifier
- `created_at` (String) An ISO-8601 timestamp representation of entity creation date.
- `deprecated` (Boolean) Indicates if this API product version is deprecated
- `gateway_service` (Attributes) gateway_service (see [below for nested schema](#nestedatt--gateway_service))
- `id` (String) The API product version identifier.
- `notify` (Boolean) When set to `true`, and all the following conditions are true:- version of the API product deprecation has changed from `false` -> `true`- version of the API product is publishedthen consumers of the now deprecated verion of the API product will be notified.
- `publish_status` (String) The publish status of the API product version
- `updated_at` (String) An ISO-8601 timestamp representation of entity update date.

<a id="nestedatt--gateway_service"></a>
### Nested Schema for `gateway_service`

Required:

- `control_plane_id` (String) The identifier of the control plane that the gateway service resides in
- `id` (String) The identifier of a gateway service associated with the version of the API product.