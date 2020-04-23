# gazelle:prefix github.com/datawire/ambassador
# gazelle:build_file_name BUILD
# gazelle:proto disable_global
# gazelle:external vendored
# gazelle:exclude cxx
# gazelle:exclude build-aux
# gazelle:exclude python

load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
#load("@io_bazel_rules_docker//python:image.bzl", "py_layer", "py_image")
#load("@io_bazel_rules_docker//python3:image.bzl", "py3_image")
load(":util.bzl", "py_image")

gazelle(name = "gazelle")

# ambassador ###################################################################

py_image(base="@alpine_glibc//image", name=".ambassador.stage0", binary="@ambassador_python//ambassador_cli:grab_snapshots")
py_image(base=".ambassador.stage0", name=".ambassador.stage1", binary="@ambassador_python//ambassador_cli:ert")
py_image(base=".ambassador.stage1", name=".ambassador.stage2", binary="@ambassador_python//ambassador_cli:mockery")
py_image(base=".ambassador.stage2", name=".ambassador.stage3", binary="@ambassador_python//ambassador_cli:ambassador")
py_image(base=".ambassador.stage3", name=".ambassador.stage4", binary="@ambassador_python//ambassador_diag:diagd")
py_image(base=".ambassador.stage4", name=".ambassador.stage5", binary="@ambassador_python//:post_update")
py_image(base=".ambassador.stage5", name=".ambassador.stage6", binary="@ambassador_python//:kubewatch")
py_image(base=".ambassador.stage6", name=".ambassador.stage7", binary="@ambassador_python//:watch_hook")
go_image(base=".ambassador.stage7", name=".ambassador.stage8", binary="//cmd/watt:watt")
go_image(base=".ambassador.stage8", name=".ambassador.stage9", binary="//cmd/kubestatus:kubestatus")

container_image(
    base = ".ambassador.stage9",
    name = "ambassador",
    workdir = "/ambassador",
    entrypoint = None,
)

# kat-client ###################################################################

go_image(
    name = ".kat-client.stage0",
    base = "@alpine_glibc//image",
    binary = "//cmd/kat-client:kat-client",
)

container_image(
    name = "kat-client",
    base = ".kat-client.stage0",
    # symlinks to add
    symlinks = {
        "/usr/local/bin/kat-client": "/app/cmd/kat-client/kat-client",
        "/work/kat_client": "/usr/local/bin/kat-client",
    },
    # runtime info
    workdir = "/work",
    entrypoint = None,
    cmd = [
        "sleep",
        "3600",
    ],
)

# kat-server ###################################################################

go_image(
    name = ".kat-server.stage0",
    base = "@alpine_glibc//image",
    binary = "//cmd/kat-server:kat-server",
)

container_image(
    name = "kat-server",
    base = ".kat-server.stage0",
    # files to add
    directory = "/work",
    mode = "0o644",
    files = [
        ":server.crt",
        ":server.key",
    ],
    # symlinks to add
    symlinks = {
        "/usr/local/bin/kat-server": "/app/cmd/kat-server/kat-server",
    },
    # runtime info
    workdir = "/work",
    entrypoint = None,
    cmd = ["kat-server"],
)
