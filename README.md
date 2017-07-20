# HAL Benchmarks

The purpose of this repository is to make performance tests on how different setups (pure C calls, Go + Cgo calls, ...) perform to call the Semtech `lora_gateway` HAL. The calls are configured for a **Multitech Conduit** with **FTDI link**, on the **Europe LoRaWAN frequency plans**.

It will contain 3 different tests:

+ Pure C call
+ Go + Cgo call
+ Rust call

## Generate benchmark binaries

The easiest way to generate the benchmark binaries is to use [GitLab CI](https://gitlab.com), which works just like Travis CI if you've ever used it. Create your fork of the repo on GitLab, enable CI for this repo, enable a [Runner](https://docs.gitlab.com/runner/), then push the repo to GitLab - your Runner will execute the [GitLab CI file](.gitlab-ci.yml) and generate the binaries for you.

## Results

### Benchmark 1
