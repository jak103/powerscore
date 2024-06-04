# Backend

## Data Ingest Simulator

### Serial Port Emulator

You can use any method to emulate a serial port that works for you, but the following steps worked for @Potato-Man114.

1. Get the free 2-week trial of [Virual Serial Port Driver Pro](https://www.virtual-serial-port.org/).
1. Setup a "loopback" port on COM4.
    1. Ensure `Baudrate Emulation` is Enabled

The serial port should now be setup for local simulation.

## Serial Data Format

[Format Speccification](FormatSpec.md)