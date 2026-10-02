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
buf lint && buf build
```

3. generate swift/kotlin files

```bash
buf generate
```

## Generate code using docker
```bash
make
```

## Native NEAR sends (`NearSpecific`)

Native NEAR transfers carry `NearSpecific` in the `KeysignPayload.blockchain_specific` oneof (tag `16`). `nonce` (`uint64`) and `block_hash` (`bytes`) are initiator-frozen signing inputs — every co-signer must sign the exact same values or the ceremony fails. `gas_fee` is an unsigned decimal string in yoctoNEAR (1 NEAR = 10^24 yoctoNEAR) reserved upfront by the initiator for local fee display and balance checks; it is not a signed gas limit or a cap, since NEAR charges gas actually burnt. Receiver and amount stay in the canonical `KeysignPayload.to_address` / `to_amount` fields and are not duplicated here. Clients that predate this oneof tag do not support NEAR at all and cannot co-sign such a payload. This paragraph documents the wire contract only; it makes no production-readiness claim and implies no platform parity.
