# ccwc
Write Your Own wc Tool - https://codingchallenges.fyi/challenges/challenge-wc

### How to build and test 
> Binary will be generated in `./bin` directory 
```shell
make clean build
make test
```

### How to use 
```shell
./bin/ccwc --help
Custom wc cli tool

Usage:
  ccwc [flags]

Flags:
  -c, --bytes       Total number of bytes
  -h, --help        help for ccwc
  -l, --lines       Total number of lines
  -m, --multibyte   Total number of characters
  -v, --version     version for ccwc
  -w, --words       Total number of words

Example: 

./bin/ccwc cmd/testdata/test.txt

```