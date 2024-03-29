---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudtruth_users Data Source - terraform-provider-cloudtruth"
subcategory: ""
description: |-
  A data source representing all user accounts in a CloudTruth organization, sorted by name and optionally filtered by type (interactive/service/all)
---

# cloudtruth_users (Data Source)

A data source representing all user accounts in a CloudTruth organization, sorted by name and optionally filtered by type (interactive/service/all)



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `type` (String) The type of user account to return: interactive, service or all, default: interactive

### Read-Only

- `id` (String) The ID of this resource.
- `organization` (String) The CloudTruth organization
- `users` (List of Object) The CloudTruth user's name, one of 'name' or 'email' is required (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `email` (String)
- `name` (String)
- `organization` (String)
- `role` (String)
- `type` (String)
