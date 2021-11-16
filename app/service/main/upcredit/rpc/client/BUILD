package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

go_test(
    name = "go_default_test",
    srcs = ["up_test.go"],
    embed = [":go_default_library"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = ["up.go"],
    importmap = "github.com/namelessup/bilibili/app/service/main/upcredit/rpc/client",
    importpath = "github.com/namelessup/bilibili/app/service/main/upcredit/rpc/client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/net/rpc:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
