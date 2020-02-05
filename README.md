# gogogamesense

A [Go](https://golang.org/) implementation of the JSON API for the [SteelSeries Engine 3 GameSense service](https://steelseries.com/developer).
It allows you to programmatically update supported SteelSeries devices' RGB effects.

Currently this library only supports controlling illumination effects.

## Current implementation

I only have a [SteelSeries Rival 310](https://steelseries.com/gaming-mice/rival-310) mouse which is an `rgb-2-zone` device and thus
all other devices are untested.

I have implemented setting single colors as well as blinking effects. Most of
the values are available, just completely untested. Gradient colors are
not implemented yet.

If you have other SteelSeries GameSense-compatible hardware, I'll gladly work
with you to get them working as well.

## Examples

See the examples folder for JSON and code samples

## Licence

MIT

## Contact

Twitter: [@donovansolms](https://twitter.com/donovansolms)
