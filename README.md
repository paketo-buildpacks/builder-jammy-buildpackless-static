# Paketo Jammy Buildpackless Static Builder

## `paketobuildpacks/builder-jammy-buildpackless-static`

This builder uses the [Paketo Jammy Static
Stack](https://github.com/paketo-buildpacks/jammy-static-stack) (Ubuntu Jammy
Jellyfish build image, distroless-like run image) and contains **no buildpacks
nor order groups**. To use this builder, you must specify buildpacks at build
time using whatever mechanisms your CNB platform of choice offers.

For example, with the `pack` CLI, use `--buildpack` as follows:
```
pack build <app-name> \
           --buildpack <some-buildpack> \
           --builder paketobuildpacks/builder-jammy-buildpackless-static:latest
```

To see which versions of build and run images, buildpacks, and the lifecycle
that are contained within a given builder version, see the
[Releases](https://github.com/paketo-buildpacks/builder-jammy-buildpackless-static/releases) on this
repository. This information is also available in the `builder.toml`.

For more information about this builder and how to use it, visit the [Paketo
builder documentation](https://paketo.io/docs/builders/).  To learn about the
stack included in this builder, visit the [Paketo stack
documentation](https://paketo.io/docs/stacks/) and the [Paketo Jammy Static Stack
repo](https://github.com/paketo-buildpacks/jammy-static-stack).
