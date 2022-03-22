# CHANGELOG

## v1.0.2 (March, 2022)

### Features:

* added new property `BackupLocation` on `CreateClusterProperties` structure
* added new query parameter `Direction` on `ApiClusterLogsGetRequest` structure

### Enhancements:

* improved error handling while decoding API Response
* added constructors for structures in order to support easier initialization
* added utility functions to easily convert pointers to values of basic types

## v1.0.1 (February, 2022)

### Enhancements:

* added support for `eu-south-2`, `eu-central-2` for `BackupLocation`

### Dependency-updates:

* updated `golang.org/x/oauth2` go import to the latest version

## v1.0.0 (January, 2022)

### Features:

* first release 🎉
