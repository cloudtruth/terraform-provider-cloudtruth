## 0.3.6 (July 27, 2022)

FEATURES:
* cloudtruth_type resource for managing custom parameter types

BUG FIXES:
* fixed a panic when client initialization requests return 4xx errors e.g. 401 Unauthorized

ENHANCEMENTS:
* Parallelized cache loading during client initialization
* Improved retry logic

DEPRECATIONS:
* The environment, value, dynamic, external, location and filter fields will be removed from the cloudtruth_parameter resource in the next minor point release to allow for the provider to support a new cloudtruth_parameter_value resource which, together with cloudtruth_parameter, will map directly to the backing API constructs and allow for a cleaner/decoupled implementation of additional Parameter and Parameter Value support.

## 0.3.5 (July 22, 2022)

FEATURES:
* cloudtruth_user and cloudtruth_users (all org users) data sources.
* cloudtruth_access_grant resource for assigning owner/admin/contributor/viewer access to
a user or group (of users) for a target environment or project

## 0.3.3 (July 15, 2022)

FEATURES:

* as_of and tag support for the cloudtruth_parameter data source.
* cloudtruth_group resource for users and service accounts

ENHANCEMENTS:

* Retry logic for all API operations

## 0.3.1 (July 3, 2022)

FEATURES:

* Support for tags as data sources and resources

ENHANCEMENTS:

* Lint and format pre-commit hook
* Various doc updates

## 0.3.0 (June 28, 2022)

BUG FIXES:

* https://github.com/cloudtruth/terraform-provider-cloudtruth/issues/8 - the provider handles nested environments and projects correctly now, allowing multiple levels of nesting in an initial apply.


## 0.2.0 (June 24, 2022)

BUG FIXES:

* The provider returns an error now with `cloudtruth_parameter` and `cloudtruth_template` data sources which don't exist instead of panicking.

FEATURES:

* Support for external parameters


## 0.1.0 (June 21, 2022)

Initial release

