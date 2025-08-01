# Changelog

## [0.3.4](https://github.com/andygodish/yamlres/compare/v0.3.3...v0.3.4) (2025-08-01)


### Features

* added fields to Basics schema in backend, refactored frontend to use a basicInfo component ([357fb0b](https://github.com/andygodish/yamlres/commit/357fb0b0804149bc8253e78d1cffe7d8a2b4ed5e))
* added root dockerfile that builds and deploys both ui and backend services ([ba12bc2](https://github.com/andygodish/yamlres/commit/ba12bc250bafe4314a9ae81e4504548c3ea57cea))
* added semver display to the UI ([3887b9b](https://github.com/andygodish/yamlres/commit/3887b9bac3a304e500341681bc57f5b5d517c1fe))
* added Vue app for UI ([dc12464](https://github.com/andygodish/yamlres/commit/dc1246432d7a197992d676a4e7baf508a7bf8753))
* backend service conting endpoint for getting resume from fs ([0cbc8dc](https://github.com/andygodish/yamlres/commit/0cbc8dc3a2a870e5773f073c7ce508d5596cc758))
* force release ([fd41d36](https://github.com/andygodish/yamlres/commit/fd41d36290c8f3142961b8d79941e0099beea2a7))
* trigger release-please ([369cb43](https://github.com/andygodish/yamlres/commit/369cb4340fdf1f2d28f4c5d19ba3bb7ac6848315))


### Bug Fixes

* added support for multi-arch image builds ([90569db](https://github.com/andygodish/yamlres/commit/90569dbaaf189a253ca1127f324c50d44ca9d4dc))
* **deps:** update dependency axios to v1.11.0 ([c631948](https://github.com/andygodish/yamlres/commit/c6319483da5d844f80e45170c18463fc40542f9e))
* **deps:** update dependency axios to v1.11.0 ([025ad7a](https://github.com/andygodish/yamlres/commit/025ad7a17a2a870ce722fd88ecba9c01d91a8c15))
* **deps:** update dependency core-js to v3.44.0 ([52c1eb0](https://github.com/andygodish/yamlres/commit/52c1eb093bd24253360b56c82c7ecb26c1381e30))
* **deps:** update dependency core-js to v3.44.0 ([fba6b80](https://github.com/andygodish/yamlres/commit/fba6b808a8e0611be8070ce3c0ccebbedc755160))
* **deps:** update dependency vue to v3.5.18 ([fc32800](https://github.com/andygodish/yamlres/commit/fc32800541b7f37e0f2c5591eb21ca34ef5c4d79))
* **deps:** update dependency vue to v3.5.18 ([61bbf16](https://github.com/andygodish/yamlres/commit/61bbf160cdce0c385192e791127b151c98ea928c))
* downgrading eslint due to internal package conflicts ([110c7d1](https://github.com/andygodish/yamlres/commit/110c7d1e5dfbc44f90e2af7a3bc87b756fa77d01))
* dummy package.json in root ([d36983c](https://github.com/andygodish/yamlres/commit/d36983c43ddff0dab52b85ed7b7495f1acb12b0f))
* removing package name ([ea30352](https://github.com/andygodish/yamlres/commit/ea303527a3ba31a8de0e0a9a7bef1cd3d08d9053))
* rp config in actions ([9d3b0d0](https://github.com/andygodish/yamlres/commit/9d3b0d04aa921767dc67d3439005ad510d889aed))
* trying to remove name from dummy package.json in root dir ([4ba26cb](https://github.com/andygodish/yamlres/commit/4ba26cb5c1ed334bccd5908b3973ea745f2996f2))


### Miscellaneous Chores

* release 0.1.0 ([bb11016](https://github.com/andygodish/yamlres/commit/bb110161e885f84a3de4412fd6e5de9e615ed160))
* release 0.1.0 ([08e4e34](https://github.com/andygodish/yamlres/commit/08e4e344ee6eb0ba9fd9fc9767ff911fddc6d5d1))
* trigger release-please ([0148dd9](https://github.com/andygodish/yamlres/commit/0148dd904501569d95aea9e2e8ba0504493a5b71))
* trigger release-please ([cd8db51](https://github.com/andygodish/yamlres/commit/cd8db5102a6316ce144403fd60be6454fda956eb))
* trigger release-please ([10925ab](https://github.com/andygodish/yamlres/commit/10925abf83bb79760c198c8a308fc693ed8031a3))
* trigger release-please ([ec32dbe](https://github.com/andygodish/yamlres/commit/ec32dbe9862dca10cba022f72a1de6c8fa143f03))
* trigger release-please ([c2c0a11](https://github.com/andygodish/yamlres/commit/c2c0a111a2d829ff983cdbfbd8bcb4c3d1e44f16))
* trigger release-please ([3f6de1d](https://github.com/andygodish/yamlres/commit/3f6de1d5c0bd6d71f5e0bc813912d69fad533e3c))
* trigger release-please ([d78ac4f](https://github.com/andygodish/yamlres/commit/d78ac4f99bd15bda6577dcf8b29e1fc2564ce944))
* trigger release-please ([11a9789](https://github.com/andygodish/yamlres/commit/11a978999b617339069d2ebc9f8d32a98bdab1cd))
* trigger release-please ([ee9e5ef](https://github.com/andygodish/yamlres/commit/ee9e5effb0bfdbfc4132e3f29a9cfadb011b923f))
---
> Note on versions 0.3.0-0.3.3: These releases contain duplicate changelog entries due to a migration from manual release management to automated release-please tooling. During this transition period, we reconfigured our CI/CD pipeline, updated release configurations, and resolved compatibility issues between release-please and our existing tag structure. The redundant entries reflect the same underlying changes being processed multiple times as we stabilized the automation. Starting with v0.3.4, the changelog reflects clean, automated release tracking.
---
## [0.2.0](https://github.com/andygodish/yamlres/compare/v0.1.0...v0.2.0) (2025-04-29)


### Features

* added semver display to the UI ([3887b9b](https://github.com/andygodish/yamlres/commit/3887b9bac3a304e500341681bc57f5b5d517c1fe))


### Miscellaneous Chores

* release 0.1.0 ([bb11016](https://github.com/andygodish/yamlres/commit/bb110161e885f84a3de4412fd6e5de9e615ed160))
* trigger release-please ([11a9789](https://github.com/andygodish/yamlres/commit/11a978999b617339069d2ebc9f8d32a98bdab1cd))
* trigger release-please ([ee9e5ef](https://github.com/andygodish/yamlres/commit/ee9e5effb0bfdbfc4132e3f29a9cfadb011b923f))

## 0.1.0 (2025-04-29)


### Features

* added root dockerfile that builds and deploys both ui and backend services ([ba12bc2](https://github.com/andygodish/yamlres/commit/ba12bc250bafe4314a9ae81e4504548c3ea57cea))
* added Vue app for UI ([dc12464](https://github.com/andygodish/yamlres/commit/dc1246432d7a197992d676a4e7baf508a7bf8753))
* backend service conting endpoint for getting resume from fs ([0cbc8dc](https://github.com/andygodish/yamlres/commit/0cbc8dc3a2a870e5773f073c7ce508d5596cc758))


### Bug Fixes

* added support for multi-arch image builds ([90569db](https://github.com/andygodish/yamlres/commit/90569dbaaf189a253ca1127f324c50d44ca9d4dc))


### Miscellaneous Chores

* release 0.1.0 ([08e4e34](https://github.com/andygodish/yamlres/commit/08e4e344ee6eb0ba9fd9fc9767ff911fddc6d5d1))
