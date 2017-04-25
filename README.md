# my-mutex

This is codes for study of exclusive controlling.

## Description

About each directories:

- dekkers: Exclusive controlled increment by [Dekker's algorithm](https://en.wikipedia.org/wiki/Dekker%27s_algorithm). But it's not work correctly because Go is CPU's out-of-order execution.
- llsc: Exclusive cotrolled increment by LLSC instruction of MIPS.
- llscmutex: Mutex by LLSC instruction of MIPS.
- race: Simplest implementation of increment. Causes race-condition.
- tas: Mutex by Test-And-Set implementation. Internally using LLSC instruction of MIPS.

Some code uses MIPS special instruction so them tried on MIPS architecture or emulator.
