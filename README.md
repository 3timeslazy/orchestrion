# Orchestrion

[![Documentation](https://img.shields.io/badge/documentation-datadoghq.dev/orchestrion-blue.svg?style=flat)](https://datadoghq.dev/orchestrion)
![Latest Release](https://img.shields.io/github/v/release/DataDog/orchestrion?display_name=tag&label=Latest%20Release)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/datadog/orchestrion)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/DataDog/orchestrion/badge)](https://scorecard.dev/viewer/?uri=github.com/DataDog/orchestrion)

Automatic compile-time instrumentation of Go code.

## Overview

[Orchestrion](https://en.wikipedia.org/wiki/Orchestrion) processes Go source code at compilation time and automatically
inserts instrumentation. This instrumentation produces Datadog APM traces from the instrumented code and supports
Datadog Application Security Management.

> [!IMPORTANT]
> Orchestrion is under active development. The supported features are rapidly growing, and the user experience may evolve
> with future releases.
>
> Should you encounter issues or a bug when using `orchestrion`, please report it in the [bug tracker][gh-issues].
>
> For support & general questions, you are welcome to use [GitHub discussions][gh-discussions]. You may also contact us
> privately via Datadog support.
>
> [gh-issues]: https://github.com/DataDog/orchestrion/issues/new/choose
> [gh-discussions]: https://github.com/DataDog/orchestrion/discussions

## Requirements

Orchestrion supports the two latest releases of Go, matching Go's [official release policy][go-releases]. It may
function correctly with older Go releases; but we will not be able to offer support for these if they don't.

In addition to this, Orchestrion only supports projects using [Go modules][go-modules].

[go-releases]: https://go.dev/doc/devel/release#policy
[go-modules]: https://pkg.go.dev/cmd/go#hdr-Modules__module_versions__and_more

## Getting started

1. Install Orchestrion:
    ```console
    $ go install github.com/DataDog/orchestrion@latest
    ```

2. <details><summary>Optional: project <tt>go.mod</tt> registration</summary>

      >  You can automatically add `orchestrion` to your project's dependencies by running:
      > ```console
      > $ orchestrion pin
      > ```
      > This will:
      > 1. Create a new `orchestrion.tool.go` file containing content similar to:
      >     ```go
      >     // Code generated by `orchestrion pin`; DO NOT EDIT.
      >
      >     // This file is generated by `orchestrion pin`, and is used to include a blank import of the
      >     // orchestrion package(s) so that `go mod tidy` does not remove the requirements from go.mod.
      >     // This file should be checked into source control.
      >
      >     //go:build tools
      >
      >     package tools
      >
      >     import _ "github.com/DataDog/orchestrion"
      >     ```
      > 2. Run `go get github.com/DataDog/orchstrion@<current-release>` to make sure the project version corresponds to the
      >    one currently being used
      > 3. Run `go mod tidy` to make sure your `go.mod` and `go.sum` files are up-to-date
      >
      > If you do not run this command, it will be done automatically when required. Once done, the version of `orchestrion`
      > used by this project can be controlled directly using the `go.mod` file, as you would control any other dependency.
    </details>

3. Prefix your `go` commands with `orchestrion`:
    ```console
    $ orchestrion go build .
    $ orchestrion go test -race ./...
    ```

    If you have not run `orchestrion pin`, you may see a message similar to the following appear, as `orchestrion pin`
    is automatically executed:
    ```
    ╭──────────────────────────────────────────────────────────────────────────────╮
    │                                                                              │
    │  Warning: github.com/DataDog/orchestrion is not present in your go.mod       │
    │  file.                                                                       │
    │  In order to ensure build reliability and reproductibility, orchestrion      │
    │  will now add itself in your go.mod file by:                                 │
    │                                                                              │
    │      1. creating a new file named orchestrion.tool.go                        │
    │      2. running go get github.com/DataDog/orchestrion@v0.7.0-dev.2           │
    │      3. running go mod tidy                                                  │
    │                                                                              │
    │  You should commit the resulting changes into your source control system.    │
    │                                                                              │
    ╰──────────────────────────────────────────────────────────────────────────────╯
    ```


    <details><summary>Alternative</summary>

    > _Orchestrion_ at the core is a standard Go toolchain `-toolexec` proxy. Instead of using `orchestrion go`, you can
    > also manually provide the `-toolexec` argument to `go` commands that accept it:
    > ```console
    > $ go build -toolexec 'orchestrion toolexec' .
    > $ go test -toolexec 'orchestrion toolexec' -race .
    > ```
    </details>

> The version of `orchestrion` used to compile your project is ultimately tracked in the `go.mod` file. You can manage
> it in the same way you manage any other dependency, and updating to the latest release is as simple as doing:
> ```console
> $ go get github.com/DataDog/orchestrion@latest
> ```

## Supported libraries

Orchestrion supports automatic tracing of the following libraries:

Library                                               | Since    | Notes
------------------------------------------------------|:--------:|------------------------------------------------------------
`database/sql`                                        | `v0.7.0` | [Aspect][db-sql]
`github.com/gin-gonic/gin`                            | `v0.7.0` | [Aspect][gin]
`github.com/go-chi/chi/v5`                            | `v0.7.0` | [Aspect][chi-v5]
`github.com/go-chi/chi`                               | `v0.7.0` | [Aspect][chi-v1]
`github.com/go-redis/redis/v7`                        | `v0.7.0` | [Aspect][go-redis-v7]
`github.com/go-redis/redis/v8`                        | `v0.7.0` | [Aspect][go-redis-v8]
`github.com/gofiber/fiber/v2`                         | `v0.7.0` | [Aspect][fiber-v2]
`github.com/gomodule/redigo/redis`                    | `v0.7.0` | [Aspect][redigo]
`github.com/gorilla/mux`                              | `v0.7.0` | [Aspect][gorilla] ([library-side][lib-side])
`github.com/jinzhu/gorm`                              | `v0.7.0` | [Aspect][jinzhu-gorm]
`github.com/labstack/echo/v4`                         | `v0.7.0` | [Aspect][echo]
`google.golang.org/grpc`                              | `v0.7.0` | [Aspect][grpc]
`gorm.io/gorm`                                        | `v0.7.0` | [Aspect][gorm]
`net/http`                                            | `v0.7.0` | [Client][net-http.client] ([library-side][lib-side]) / [Server][net-http.server]
`go.mongodb.org/mongo-driver/mongo`                   | `v0.7.3` | [Aspect][mongo]
`github.com/aws-sdk-go/aws`                           | `v0.7.4` | [Aspect][aws-sdk-go]
`github.com/hashicorp/vault`                          | `v0.7.4` | [Aspect][hashicorp-vault]
`github.com/IBM/sarama`                               | `v0.7.4` | [Aspect][ibm-sarama]
`github.com/Shopify/sarama`                           | `v0.7.4` | [Aspect][shopify-sarama]
`k8s.io/client-go`                                    | `v0.7.4` | [Aspect][k8s-client]
`log/slog`                                            | `v0.7.4` | [Aspect][log-slog]
`os`                                                  | `v0.8.0` | [Aspect][os]
`github.com/aws/aws-sdk-go-v2`                        | `v0.8.0` | [Aspect][aws-sdk-go-v2]
`github.com/redis/go-redis/v9`                        | `v0.8.0` | [Aspect][go-redis-v9]
`github.com/gocql/gocql`                              | `v0.8.0` | [Aspect][gocql]
`cloud.google.com/go/pubsub`                          | `v0.9.0` | [Aspect][pubsub] ([library-side][lib-side])
`github.com/99designs/gqlgen`                         | `v0.9.1` | [Aspect][gqlgen]
`github.com/redis/go-redis`                           | `v0.9.1` | [Aspect][go-redis]
`github.com/graph-gophers/graphql-go`                 | `v0.9.1` | [Aspect][graph-gophers]
`github.com/graphql-go/graphql`                       | `v0.9.1` | [Aspect][graphql]
`testing`                                             | `v0.9.3` | [Aspect][testing] with [Test Optimization][test-optimization]
`github.com/jackc/pgx`                                | `v0.9.4` | [Aspect][pgx]
`github.com/elastic/go-elasticsearch`                 | `v0.9.4` | [Aspect][elasticsearch]
`github.com/twitchtv/twirp`                           | `v0.9.4` | [Aspect][twirp]
`github.com/segmentio/kafka-go`                       | `v0.9.4` | [Aspect][segmentio-kafka-go] ([library-side][lib-side])
`github.com/confluentinc/confluent-kafka-go/kafka`    | `v0.9.4` | [Aspect][confluent-kafka-go-v1] ([library-side][lib-side])
`github.com/confluentinc/confluent-kafka-go/kafka/v2` | `v0.9.4` | [Aspect][confluent-kafka-go-v2] ([library-side][lib-side])
`github.com/julienschmidt/httprouter`                 | `v0.9.4` | [Aspect][httprouter] ([library-side][lib-side])
`github.com/sirupsen/logrus`                          | `v0.9.4` | [Aspect][logrus]

[lib-side]: #library-side

[db-sql]: https://datadoghq.dev/orchestrion/docs/built-in/stdlib/database-sql/
[gin]: https://datadoghq.dev/orchestrion/docs/built-in/http/gin/
[chi-v5]: https://datadoghq.dev/orchestrion/docs/built-in/http/chi/#use-v5-tracer-middleware
[chi-v1]: https://datadoghq.dev/orchestrion/docs/built-in/http/chi/#use-v1-tracer-middleware
[go-redis-v7]: https://datadoghq.dev/orchestrion/docs/built-in/databases/go-redis/#wrap-v7-client
[go-redis-v8]: https://datadoghq.dev/orchestrion/docs/built-in/databases/go-redis/#wrap-v8-client
[go-redis-v9]: https://datadoghq.dev/orchestrion/docs/built-in/databases/go-redis/#wrap-v9-client
[fiber-v2]: https://datadoghq.dev/orchestrion/docs/built-in/http/fiber/
[redigo]: https://datadoghq.dev/orchestrion/docs/built-in/databases/redigo/
[gorilla]: https://datadoghq.dev/orchestrion/docs/built-in/http/gorilla/
[jinzhu-gorm]: https://datadoghq.dev/orchestrion/docs/built-in/databases/gorm/#jinzhugorm
[echo]: https://datadoghq.dev/orchestrion/docs/built-in/http/echo/
[grpc]: https://datadoghq.dev/orchestrion/docs/built-in/grpc/
[gorm]: https://datadoghq.dev/orchestrion/docs/built-in/databases/gorm/#gormiogorm
[net-http.client]: https://datadoghq.dev/orchestrion/docs/built-in/stdlib/net-http.client/
[net-http.server]: https://datadoghq.dev/orchestrion/docs/built-in/stdlib/net-http.server/
[mongo]: https://datadoghq.dev/orchestrion/docs/built-in/databases/mongo/
[k8s-client]: https://datadoghq.dev/orchestrion/docs/built-in/k8s-client/
[hashicorp-vault]: https://datadoghq.dev/orchestrion/docs/built-in/api/vault/
[log-slog]: https://datadoghq.dev/orchestrion/docs/built-in/stdlib/slog/
[aws-sdk-go]: https://datadoghq.dev/orchestrion/docs/built-in/cloud/aws-sdk/
[aws-sdk-go-v2]: https://datadoghq.dev/orchestrion/docs/built-in/cloud/aws-sdk-v2/
[ibm-sarama]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/ibm_sarama/
[shopify-sarama]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/shopify_sarama/
[os]: https://datadoghq.dev/orchestrion/docs/built-in/stdlib/ossec/
[gocql]: https://datadoghq.dev/orchestrion/docs/built-in/databases/gocql/
[pubsub]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/gcp_pubsub/
[gqlgen]: https://datadoghq.dev/orchestrion/docs/built-in/graphql/gqlgen/
[go-redis]: https://datadoghq.dev/orchestrion/docs/built-in/databases/go-redis/#wrap-v0-client
[graph-gophers]: https://datadoghq.dev/orchestrion/docs/built-in/graphql/graph-gophers/
[graphql]: https://datadoghq.dev/orchestrion/docs/built-in/graphql/graphql-go/
[testing]: https://datadoghq.dev/orchestrion/docs/built-in/civisibility/testing/
[pgx]: https://datadoghq.dev/orchestrion/docs/built-in/databases/pgx
[elasticsearch]: https://datadoghq.dev/orchestrion/docs/built-in/databases/go-elasticsearch/
[twirp]: https://datadoghq.dev/orchestrion/docs/built-in/rpc/twirp/
[segmentio-kafka-go]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/segmentio_kafka_v0/
[confluent-kafka-go-v1]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/confluentinc_kafka/#inject-kafka-library-version-v1
[confluent-kafka-go-v2]: https://datadoghq.dev/orchestrion/docs/built-in/datastreams/confluentinc_kafka/#inject-kafka-library-version-v2
[httprouter]: https://datadoghq.dev/orchestrion/docs/built-in/http/julienschmidt_httprouter/
[logrus]: https://datadoghq.dev/orchestrion/docs/built-in/logs/logrus/

[test-optimization]: https://docs.datadoghq.com/tests/

### Library Side

Most integrations are added by orchestrion at the call site, making it possible to use the [`//orchestrion:ignore`
directive][orchestrion-ignore] to locally opt out of instrumentation for a specific instance of a component.

[orchestrion-ignore]: https://datadoghq.dev/orchestrion/docs/custom-trace/#prevent-instrumentation-of-a-section-of-code

Some integrations are however injected directly into the library (library-side instrumentation, also called callee-side
instrumentation), and are hence always active and cannot be locally opted out. If you have a use-case where you need to
locally opt out of a library-side instrumentation, please let us know about it by filing a [GitHub issue][new-gh-issue].

[new-gh-issue]: https://github.com/DataDog/orchestrion/issues/new/choose

## Troubleshooting

If you run into issues when using `orchestrion` please make sure to collect all relevant details about your setup in
order to help us identify (and ideally reproduce) the issue. The version of orchestrion (which can be obtained from
`orchestrion version`) as well as of the go toolchain (from `go version`) are essential and must be provided with any
bug report.

You can inspect everything Orchestrion is doing by adding the `-work` argument to your `go build` command; when doing so
the build will emit a `WORK=` line pointing to a working directory that is retained after the build is finished. The
contents of this directory contains all updated source code Orchestrion produced and additional metadata that can help
diagnosing issues.

## More information

Orchestrion's documentation can be found at [datadoghq.dev/orchestrion](https://datadoghq.dev/orchestrion); in
particular:
- the [user guide](https://datadoghq.dev/orchestrion/docs/) provides information about available configuration, and how
  to customize the traces produced by your application;
- the [contributor guide](https://datadoghq.dev/orchestrion/contributing/) provides more detailed information about how
  orchestrion works and how to contribute new instrumentation to it.
