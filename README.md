## Spin Terminal

This spin plugin provides an easy way to setup toolchain required for building and running your spin apps.

### Prerequisite:

Install dagger - follow instructions here https://docs.dagger.io/install

### Usage:

Create a file `.toolchains` in your Spin App directory

e.g.

```
spin=3.0.0
golang=1.23.2
tinygo=0.34.0
```

and then run `dagger call -m github.com/rajatjindal/daggerverse/wasi@main --source=. terminal`