## 0.3.1 (July 3, 2022)

FEATURES:

* Support for tags as data sources and resources

ENHANCEMENTS:

* Lint and format pre-commit hook
* Various doc updates

## 0.3.0 (June 28, 2022)

BUG FIXES:

* https://github.com/cloudtruth/terraform-provider-cloudtruth/issues/8 - the provider handles nest environments and projects correctly now, allowing multiple levels of nesting in an initial apply.


## 0.2.0 (June 24, 2022)

BUG FIXES:

* The provider returns an error now with `cloudtruth_parameter` and `cloudtruth_template` data sources which don't exist instead of panicking.

FEATURES:

* Support for external parameters


## 0.1.0 (June 21, 2022)

Initial release
