# HAL Benchmarks

The purpose of this repository is to make performance tests on how different setups (pure C calls, Go + Cgo calls, ...) perform to call the Semtech `lora_gateway` HAL.

It contains 3 different tests:

+ Pure C call

+ Go + Cgo call

+ Rust call

Every program is:

+ Configuration of the SX1301 module

+ Loop of:

  + 10 receive

  + 1 send downlink

## Run benchmarks

+ Compile `libloragw`, copy it to the current folder
