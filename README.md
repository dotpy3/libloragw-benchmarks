# HAL Benchmarks

The purpose of this repository is to make performance tests on how different setups (pure C calls, Go + Cgo calls, ...) perform to call the Semtech `lora_gateway` HAL. The calls are configured for a **Multitech Conduit** with **FTDI link**, on the **Europe LoRaWAN frequency plans**.

It will contain 3 different tests:

+ Pure C call
+ Go + Cgo call
+ Rust call

Every program is:

+ Configuration of the SX1301 module
+ Loop of:
  + 10 receive
  + 1 send downlink

## Run benchmarks

1. Move to the directory of the repository, and download the `libloragw` source code: `git clone git@github.com:TheThingsNetwork/lora_gateway.git`
