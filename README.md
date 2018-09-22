# Nevis
[![codecov](https://codecov.io/gh/alisdairrankine/nevis/branch/master/graph/badge.svg)](https://codecov.io/gh/alisdairrankine/nevis)
[![CircleCI](https://circleci.com/gh/alisdairrankine/nevis.svg?style=svg)](https://circleci.com/gh/alisdairrankine/nevis)

## Overview
Nevis is a frontend framework written in go and is, at it's core, a virtual DOM library.
Nevis exists only as a view layer at the moment, although state management is something that
can  be potentially explored.

Nevis currently supports rendering to the dom with gopherjs, and to a string for pre-rendering.
WASM is certainly a possibility.

Convenience methods are provided for HTML and SVG Elements, but the whole framework very basic for now.

## How to use
I wouldn't advise using it, but see the examples dir for a feeling of how it's put together.

## License

MIT License

## WARNING!!

This is highly experimental and cannot be trusted. It isn't being used by anyone, and doesn't really
have a purpose.

There are basically no tests, no comments. DOM Events are incomplete.