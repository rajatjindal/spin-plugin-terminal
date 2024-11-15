## spin terminal

A spin plugin to provide a terminal (powered by containers and orchestrated by [dagger](https://dagger.io)) with the toolchains required to run a project.

### Install:

```
spin plugin install -u https://github.com/rajatjindal/spin-plugin-terminal/releases/download/canary/terminal.json
```

### Get a terminal with the requested toolchain
Create a file `.toolchains` in your Spin App directory

e.g.

```
spin=3.0.0
golang=1.23.2
tinygo=0.34.0
```

and then run `spin terminal`

### Build and export
Create a file `.toolchains` in your Spin App directory

e.g.

```
spin=3.0.0
golang=1.23.2
tinygo=0.34.0
```

and then run `spin terminal buildx`

This will build the app and export the artifacts back to your host directory.
