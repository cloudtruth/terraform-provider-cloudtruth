---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudtruth_user Data Source - terraform-provider-cloudtruth"
subcategory: ""
description: |-
  A CloudTruth user data source
---

# cloudtruth_user (Data Source)

A CloudTruth user data source



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `email` (String) The user's email, one of 'name' or 'email' is required"
- `name` (String) The CloudTruth user's name, one of 'name' or 'email' is required

### Read-Only

- `id` (String) The ID of this resource.
- `organization` (String) The name of the CloudTruth organization containging the user
- `role` (String) The user's role
- `type` (String) The user's type: interactive or service

