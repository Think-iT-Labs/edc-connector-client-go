<div align="center">
  <h1>EDC Connector Client 👩‍🚀</h1>
  <p>
    <b>
      A HTTP client to communicate with the <a href="https://github.com/eclipse-edc/Connector">EDC Connector</a>. Written in Golang.
    </b>
  </p>
  <sub>
    Built with ❤️ at <a href="https://think-it.io">Think-it</a>.
  </sub>
</div>

## Abstract

The [**EDC Connector**](https://github.com/eclipse-edc/Connector) is a framework for a sovereign, inter-organizational
data exchange. It provides _low-level_ primitives to allow network participants to expose and consume offers.
The _Connector_ does so by providing an extensive HTTP API documented via
[OpenAPI specification](https://github.com/eclipse-edc/Connector/blob/a6fdb2a0b4360629ec562e11ae19d077162200b7/resources/openapi/openapi.yaml).

This project aims to increase the level of abstraction, bringing the _low-level_ HTTP API to _mid-level_
developers by providing an HTTP Client which is thoroughly tested and fully type-safe.

> Similarly to the **EDC Connector**, this library is at its early stage.
> It aims to maintain compatibility with the latest version of the _Connector_, currently `0.2.0`.
> The compatibility is reflected in the library's versioning.

## Usage

You can refer to the code under `examples/` directory for code snippets on how to use the SDK for various operations.

## Development

The development of this library is following a TDD approach. At the moment it
doesn't cover all EDC Connector capabilities, but aims to do so.

`docker-compose` is used to run the development environment. It runs two
connectors with capabilities described in the
[gradle configuration](connector/build.gradle.kts) file.

Please, adhere to the [CONTRIBUTING](CONTRIBUTING.md) guidelines when suggesting
changes in this repository.

## Disclaimer
**Please Note:** This project is currently not actively maintained. While the existing codebase is provided as-is, it may not receive updates, bug fixes, or support in the foreseeable future. Feel free to use, modify, and distribute the code according to the project's license, but please be aware that there might be unresolved issues or compatibility concerns with newer software dependencies.

## License

Copyright 2022-2023 Think.iT GmbH.

Licensed under the [Apache License, Version 2.0](LICENSE). Files in the project
may not be copied, modified, or distributed except according to those terms.
