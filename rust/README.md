# Rust

## Installation

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## Check the installation

```bash
❯ cargo version
cargo 1.52.0 (69767412a 2021-04-21)
```

## Getting started

### Generating a new project

```bash
❯ cargo new hello-test
     Created binary (application) `hello-test` package

❯ tree .
.
└── hello-test
    ├── Cargo.toml
    └── src
        └── main.rs

2 directories, 2 files
```

### Execute the file

```bash
❯ cd hello-test

❯ cargo run
   Compiling hello-test v0.1.0 (/Users/kanata-miyahana/repos/rust/hello-world/hello-test)
    Finished dev [unoptimized + debuginfo] target(s) in 1.41s
     Running `target/debug/hello-test`
Hello, world!
```

### Check generated files

```bash
❯ tree .
.
└── hello-test
    ├── Cargo.lock
    ├── Cargo.toml
    ├── src
    │   └── main.rs
    └── target
        ├── CACHEDIR.TAG
        └── debug
            ├── build
            ├── deps
            │   ├── hello_test
            │   ├── hello_test.d
            │   └── hello_test.dSYM
            │       └── Contents
            │           ├── Info.plist
            │           └── Resources
            │               └── DWARF
            │                   └── hello_test
            ├── examples
            ├── hello-test
            ├── hello-test.d
            ├── hello-test.dSYM -> deps/hello_test.dSYM
            └── incremental

13 directories, 10 files
```
