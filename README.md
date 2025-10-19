## random

cli tool to generate random bytes.

## install

`go install github.com/tiredkangaroo/random@latest`

## options:

`--output` and `-o`: specify the output encoding format.

supports: `base32`, `base64`, `hex`, `raw`.
default: `hex`.

`--length` and `-l`: specify the number of the random bytes to generate. this will not be the length of the output, as it depends on the encoding format. must be a positive integer.
default: 16.

`--newline` and `-n`: add a newline at the end of the output.
default: `false`.

`--repeat` and `-r`: repeat the output multiple times. must be a positive integer.
default: 0 (only once).

`--help` and `-h`: show the help message.
