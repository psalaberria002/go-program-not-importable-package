# Reproduction for "import xxx is a program, not an importable package"

I am trying to compile a binary from source. This binary comes from a main package.

Previous to Go 1.21 (or 1.22, I am not sure) we could defined unused imports
in any go source file as follows:

```
package hack

// This is a hack :( for making go mod tidy not remove these dependencies from go.mod
// We want to build @com_github_jimmidyson_configmap_reload//:configmap-reload from source.
import (
	_ "github.com/jimmidyson/configmap-reload"
)
```

We could still `go build ./...` without problems. Now that isn't working anymore,
and we get "import xxx is a program, not an importable package" errors.

I created three folders with slightly different setups:

## [fails-with-bazel-and-plain-go](./fails-with-bazel-and-plain-go)

```
$ bazel build //fails-with-bazel-and-plain-go/...
ERROR: /home/paulsmsm/cognite/go-program-not-importable-package/fails-with-bazel-and-plain-go/BUILD.bazel:3:11: in go_library rule //fails-with-bazel-and-plain-go:fails-with-bazel-and-plain-go_lib:
Traceback (most recent call last):
        File "/home/paulsmsm/.cache/bazel/_bazel_paulsmsm/b371747d64acae5f2bcd5da51b9e04e9/external/rules_go~/go/private/rules/library.bzl", line 39, column 34, in _go_library_impl
                source = go.library_to_source(go, ctx.attr, library, ctx.coverage_instrumented())
        File "/home/paulsmsm/.cache/bazel/_bazel_paulsmsm/b371747d64acae5f2bcd5da51b9e04e9/external/rules_go~/go/private/context.bzl", line 283, column 26, in _library_to_source
                _check_binary_dep(go, dep, "deps")
        File "/home/paulsmsm/.cache/bazel/_bazel_paulsmsm/b371747d64acae5f2bcd5da51b9e04e9/external/rules_go~/go/private/context.bzl", line 346, column 13, in _check_binary_dep
                fail("rule {rule} depends on executable {dep} via {edge}. This is not safe for cross-compilation. Depend on go_library instead.".format(
Error in fail: rule @@//fails-with-bazel-and-plain-go:fails-with-bazel-and-plain-go_lib depends on executable @@gazelle~~go_deps~com_github_jimmidyson_configmap_reload//:configmap-reload via deps. This is not safe for cross-compilation. Depend on go_library instead.
ERROR: /home/paulsmsm/cognite/go-program-not-importable-package/fails-with-bazel-and-plain-go/BUILD.bazel:3:11: Analysis of target '//fails-with-bazel-and-plain-go:fails-with-bazel-and-plain-go_lib' failed
ERROR: Analysis of target '//fails-with-bazel-and-plain-go:fails-with-bazel-and-plain-go_lib' failed; build aborted
INFO: Elapsed time: 0.079s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
ERROR: Build did NOT complete successfully
```

```
$ go build -o out ./fails-with-bazel-and-plain-go/...
fails-with-bazel-and-plain-go/main.go:6:2: import "github.com/jimmidyson/configmap-reload" is a program, not an importable package
```

## [works-with-bazel-but-not-with-plain-go](./works-with-bazel-but-not-with-plain-go)

```
$ bazel build //works-with-bazel-but-not-with-plain-go/...
INFO: Analyzed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets...
INFO: Elapsed time: 0.078s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
```

```
$ go build -o out ./works-with-bazel-but-not-with-plain-go/...
works-with-bazel-but-not-with-plain-go/hack/hack.go:6:2: import "github.com/jimmidyson/configmap-reload" is a program, not an importable package
```

## [works-with-bazel-and-plain-go](./works-with-bazel-and-plain-go)

```
$ bazel build //works-with-bazel-and-plain-go/...
INFO: Analyzed target //works-with-bazel-and-plain-go:works-with-bazel-and-plain-go (1 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //works-with-bazel-and-plain-go:works-with-bazel-and-plain-go up-to-date:
  bazel-bin/works-with-bazel-and-plain-go/works-with-bazel-and-plain-go.a
INFO: Elapsed time: 0.116s, Critical Path: 0.02s
INFO: 2 processes: 1 internal, 1 linux-sandbox.
INFO: Build completed successfully, 2 total actions
```

```
$ go build -o out ./works-with-bazel-and-plain-go/...
$ ./out
Hello two
```
