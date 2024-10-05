## 0.7.6 (October 4, 2024)
**BUG FIXES**
* Fix a panic caused by incorrect type assertion with projects field on AWS and Azure Push Action resources.

## 0.7.4 (April 24, 2024)
**BUG FIXES**
* An update to handle the underlying API change with push actions now requiring one of 3 (instead of 2) parameters to be true: `include_parameters`, `include_secrets` and `include_templates`"

## 0.7.3 (April 19, 2024)
**BUG FIXES**
* Fix occasional race conditions with nested project deletes

## 0.7.2 (April 10, 2024)
**ENHANCEMENTS**
* Add support for resource tags in integrations

## 0.7.1 (February 6, 2024)
**ENHANCEMENTS**
* Log resource IDs when read lookups fail

## 0.7.0 (December 17, 2023)
**BUG FIXES**
* Fixes an issue where template data sources were ignoring the environment and using the default environment.

**ENHANCEMENTS**
* Added support for the pattern field in project resources
* Added basic doc with concrete examples
* Updated the Go version to 1.21.5. Also using updated the Go OpenAPI bindings for the CloudTruth API

## 0.5.11 (July 29, 2023)
**BUG FIXES**
* Fixes an issue where the AWS integration resource was not outputting the external AWS ID correctly

## 0.5.10 (July 28, 2023)
**BUG FIXES**
* Fixes a panic when a push action references an nonexistent tag in an environment that exists.

## 0.5.9 (July 3, 2023)
**ENHANCEMENTS**
* Added a flag to make it easy to skip tests when running against a self-hosted/local instance of CloudTruth. Also included
- a bunch of Dependabot updates.

## 0.5.5 (March 31, 2023)
**ENHANCEMENTS**
* Updated the cloudtruth_parameter_value to handle changes to the environment property, which effectively removes a parameter
value from the source environment and recreates it in the target environment.

## 0.5.2 (February 10, 2023)
**ENHANCEMENTS**
* Upgraded from Go 1.19 -> 1.20

**BUG FIXES**
* Fixed an issue with resources without a `project` property when no default project is specified.

## 0.5.1 (December 5, 2022)
**ENHANCEMENTS**
* Improved error message with parameter value operations which fail when the containing parameter cannot be found.

## 0.5.0 (September 13, 2022)
**FEATURES**
* Support for managing AWS and Azure push and import actions via the provider, see the new resource types: `cloudtruth_aws_import_action`, `cloudtruth_azure_import_action`, `cloudtruth_aws_push_action` and `cloudtruth_azure_push_action`.
* Support for managing AWS integrations, see the new resource type: `cloudtruth_aws_integration`.

**ENHANCEMENTS**
* Upgraded from Go 1.18 -> 1.19
* Better error handling, end users now see full API error messages with 4xx responses.
* All resource types which support rename updates in the UI can also be renamed from the provider.
* Improved retry logic for CloudTruth API calls.
* Some TypeList types were replaced with TypeSet types to avoid plans erroneously reporting changes when the sub-resource order is not stable.

**BUG FIXES**
* Improved read functionality for all resource types. Previously, resource reads were not always querying the backing API resources and therefore not reporting deltas in situations where there was resource drift.
* Numerous fixes to parameter and type rule handling.

## 0.4.0 (August 14, 2022)

**FEATURES**
* The cloudtruth_parameter resource now supports max/min length and regex rules for string types and max/min rules for integer types.
* The cloudtruth_parameter_value resource is a new resource which corresponds to an environment specific value for a parameter. The provider resources cloudtruth_parameter and
cloudtruth_parameter_value correspond to their counterparts in the CloudTruth API. Until now, a cloudtruth_parameter resource represented both a parameter and all of its contained values.

**ENHANCEMENTS**
* Better error messages for client side errors e.g. when the TF provider is attempting to create a resource which already exists in the CloudTruth project (name collision).

**DEPRECATIONS**
* The cloudtruth_parameter and cloudtruth_parameters data sources have been renamed to 'cloudtruth_parameter_value' and 'cloudtruth_parameter_values' to reflect that they are environment
specific. This change also maintains consistency with the recent change to break out the cloudtruth_parameter_value resource from the cloudtruth_parameter resource.

* The rules property in the cloudtruth_type resource will be changing to match the flattened rule properties that the cloudtruth_parameter resource uses. Instead of being a nested
list of objects, rules will be represented with max/min/regex fields where max/min are either string length limits for string based types or value limits for integer based types. Regex
rules are only applicable for string based types. As with parameters, boolean types do not currently support rules.

## 0.3.6 (July 27, 2022)

**FEATURES**
* cloudtruth_type resource for managing custom parameter types

**BUG FIXES**
* fixed a panic when client initialization requests return 4xx errors e.g. 401 Unauthorized

**ENHANCEMENTS**
* Parallelized cache loading during client initialization
* Improved retry logic

**DEPRECATIONS**
* The environment, value, dynamic, external, location and filter fields will be removed from the cloudtruth_parameter resource in the next minor point release to allow for the provider to support a new cloudtruth_parameter_value resource which, together with cloudtruth_parameter, will map directly to the backing API constructs and allow for a cleaner/decoupled implementation of additional Parameter and Parameter Value support.

## 0.3.5 (July 22, 2022)

**FEATURES**
* cloudtruth_user and cloudtruth_users (all org users) data sources.
* cloudtruth_access_grant resource for assigning owner/admin/contributor/viewer access to
a user or group (of users) for a target environment or project

## 0.3.3 (July 15, 2022)

**FEATURES**
* as_of and tag support for the cloudtruth_parameter data source.
* cloudtruth_group resource for users and service accounts

**ENHANCEMENTS**
* Retry logic for all API operations

## 0.3.1 (July 3, 2022)

**FEATURES**
* Support for tags as data sources and resources

**ENHANCEMENTS**
* Lint and format pre-commit hook
* Various doc updates

## 0.3.0 (June 28, 2022)

**BUG FIXES**
* https://github.com/cloudtruth/terraform-provider-cloudtruth/issues/8 - the provider handles nested environments and projects correctly now, allowing multiple levels of nesting in an initial apply.


## 0.2.0 (June 24, 2022)

**BUG FIXES**
* The provider returns an error now with `cloudtruth_parameter` and `cloudtruth_template` data sources which don't exist instead of panicking.

**FEATURES**
* Support for external parameters


## 0.1.0 (June 21, 2022)

**Initial release**

