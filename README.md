# HAL Benchmarks

The purpose of this repository is to make performance tests on how different setups (pure C calls, Go + Cgo calls, ...) perform to call the Semtech `lora_gateway` HAL. The calls are configured for a **Multitech Conduit** with **FTDI link**, on the **Europe LoRaWAN frequency plans**.

It will contain 3 different tests:

+ Pure C call
+ Go + Cgo call (located in the [go](go/) folder)
+ Rust call

## Generate benchmark binaries

The easiest way to generate the benchmark binaries is to use [GitLab CI](https://gitlab.com), which works just like Travis CI if you've ever used it. Create your fork of the repo on GitLab, enable CI for this repo, enable a [Runner](https://docs.gitlab.com/runner/), then push the repo to GitLab - your Runner will execute the [GitLab CI file](.gitlab-ci.yml) and generate the binaries for you.

## Results

### Benchmark 1

The first benchmark consists of:

* Configuration of the SX1301 chip: board configuration, channel configuration... with a **hardcoded configuration**.
* Starting retrieving uplinks for 2 minutes: we're counting the number of *cycles*, with 1 cycle = 10 uplink retrievals.
* The score of a test is the number of cycles executed over a course of two minutes.

We went for this setup because the tests were made in an environment where LoRa packets were emitted by a node, every ~5-10 seconds - going for 2 minutes would give us the right balance. The *cycles* system is to reduce the overhead of checking time/channel for every `lgw_receive` call.

|Setup|Score|
|-------|-------|
|Go|4567 cycles executed|
|C||
|Rust||
