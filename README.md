# commondata

This repository is for the protobuf schema used by all projects that need to communicate with vultisig

## Buf

We use buf to manage protobuf files in this repo , if you don't have buf installed , please follow [install buf](https://buf.build/docs/installation) to install it locally

## How to make changes in this repo?

Update the proto files you want to change, but follow these steps to generate files

1. format proto files

```bash
buf format --write
```

2. lint proto files

```bash
buf lint && buf buil
```

3. generate swift/kotlin files

```bash
buf generate
```

## Generate code using docker
```bash
make
```
