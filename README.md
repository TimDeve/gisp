# Gisp

Gisp is a basic lisp implemented in Golang. Inspired by this
[article](https://m.stopa.io/risp-lisp-in-rust-90a0dad5b116).


## Task runner

The `lo` bash script contains the different tasks. `rg` and `entr` are
necessary for the *:watch tasks. Usage is as follow:

```sh
./lo <name of the task>
```


## Build

Test and build:

```sh
./lo
```

Build:

```sh
./lo build
```

Test:

```sh
./lo test
```


## Development

Run:

```sh
./lo run
```

Watch for file change and run:

```sh
./lo run:watch
```

Watch for file change and test:

```sh
./lo test:watch
```
