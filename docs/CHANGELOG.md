# CHANGELOG

## v1.1.2 (July, 2023)

### Features

* Added functionality for `users` endpoint
* Added functionality for `databases` endpoint

## v1.1.1 (March, 2023)

### Features

* added new state for clusters: `DEGRADED`
* added new parameters: `limit` and `offset`

## v1.0.5 (January, 2023)

### Enhancements:
* Added logger and logger level

## v1.0.4 (July, 2022)

### Features
* Added Size and Location to ClusterBackup

## v1.0.3 (May, 2022)

### Enhancements:
* `location` and `backup_location` parameters on `CreateClusterProperties` are now strings
  * `Location` and `BackupLocation` models are now removed

### Features
* **new values** for `storage_type` parameter: _**SSD_STANDARD**, **SSD_PREMIUM**_. Value **_SSD_** is deprecated. Use the equivalent **_SSD_PREMIUM_** instead.
* Added option to do certificate pinning by setting `IONOS_PINNED_CERT` env variable to be the public sha256 fingerprint of the certificate to be pinned.

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
