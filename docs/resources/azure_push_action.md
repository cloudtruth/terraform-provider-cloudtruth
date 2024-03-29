---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudtruth_azure_push_action Resource - terraform-provider-cloudtruth"
subcategory: ""
description: |-
  A CloudTruth push action.
---

# cloudtruth_azure_push_action (Resource)

A CloudTruth push action.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `integration` (String) The name (using the format VAULT_NAME@TENANT_ID) of the CloudTruth integration corresponding to this import action
- `name` (String) The name of the push action
- `projects` (Set of String) The projects containing the parameters to pushed to the Azure destination
- `resource` (String) The mustache style resource string specifying the environment, project, and parameter
- `tags` (Set of String) Tags specified in the form 'environment_name:tag_name' indicating the sync point for parameters to be pushed (multiple tags allowed but only one per environment)

### Optional

- `coerce` (Boolean) Include secrets/parameters even if the upstream destination doesn't allow them, defaults to false
- `description` (String) A description of the push action
- `dry_run` (Boolean) When true, the action only reports what it would push without actually pushing changes, defaults to true
- `force` (Boolean) Allow CloudTruth to take ownership and overwrite any pre-existing items, defaults to false
- `local` (Boolean) Only send the parameters defined directly in the specified projects (not inherited), defaults to false
- `parameters` (Boolean) Include parameters (non-secrets) when pushing, defaults to true
- `secrets` (Boolean) Include secrets when pushing, defaults to true
- `templates` (Boolean) Include templates when pushing, defaults to true

### Read-Only

- `id` (String) The ID of this resource.
- `integration_id` (String) The ID of the CloudTruth integration corresponding to this import action
