***This tool's maturity is ALPHA. Breaking changes are likely.***

# Gomnia

[![Documentation][gomnia-godoc-svg]][gomnia-godoc]

Gomnia is a CLI tool designed to make using numerical and scientific algorithms
from the terminal productive and performant.

The tool prioritizes:

* Streaming/online versions of algorithms
* Reading from stdin and writing to stdout
* Defaulting to human readable i/o
* Using existing libraries

Gomnia aims to expose functionality normally only available from within a
specific language/runtime/library to the terminal (e.g. Octave, Matlab, R,
Scipy, Gonum, GSL, etc). Where possible it will use existing libraries to
fulfill that mission.

In particular, since Gomnia is written in Go it largely exposes functionality
found in [Gonum][gonum].

## Install

```sh
go get -u github.com/calebcase/gomnia
```

## Example Usage

```sh
$ gomnia generate exponential | gomnia limit --count 1000 | gomnia summarize histogram | gomnia plot histogram
0.0015282691959209727 : ████████████████████████████████████████▏ 570
0.7982194123321851    : ████████████████▌ 234
1.594910555468449     : ███████▍ 105
2.3916016986047133    : ███▊ 53
3.1882928417409775    : █▋ 22
3.9849839848772417    : ▋ 8
4.7816751280135055    : ▍ 4
5.57836627114977      : ▎ 3
6.375057414286034     : ▏ 0
7.171748557422298     : ▏ 1
```

The command generates 1000 samples from the exponential probability
distribution, summarizes the samples with a histogram, and then plots the
histogram on the terminal.

The original motivating use case for Gomnia was generating file sizes between
500 KB and 10 MB with a total size of 1 TB. We desired a distribution of file
sizes that emphasized larger files over smaller files.

```sh
$ gomnia generate --min 500 --max 10000 beta -a 5 -b 1 | gomnia limit --sum 1000000000 > filesizes.txt
$ wc -l filesizes.txt
121860 filesizes.txt
$ datamash sum 1 < filesizes.txt
999996684.48038
```

Now we have 121860 file sizes between 500 KB and 10 MB which sum up to just
under 1 TB.

We can get a sense of the distribution by looking at the histogram:

```sh
$ gomnia summarize histogram < filesizes.txt
575.3084551331915 60
1517.777528669692 260
2460.2466022061926 795
3402.715675742693 1869
4345.184749279193 3883
5287.653822815693 7434
6230.122896352194 12202
7172.591969888695 19767
8115.061043425195 30730
9057.530116961696 44860
$ gomnia summarize histogram < filesizes.txt | gomnia plot histogram
575.3084551331915  : ▏ 60
1517.777528669692  : ▎ 260
2460.2466022061926 : ▊ 795
3402.715675742693  : █▊ 1869
4345.184749279193  : ███▌ 3883
5287.653822815693  : ██████▊ 7434
6230.122896352194  : ███████████ 12202
7172.591969888695  : █████████████████▊ 19767
8115.061043425195  : ███████████████████████████▌ 30730
9057.530116961696  : ████████████████████████████████████████▏ 44860
$ gomnia summarize histogram < filesizes.txt | gomnia plot histogram --variant vertical
 44860 ┼                                        ╭──╮     
 42617 ┤                                        │  │     
 40374 ┤                                        │  │     
 38131 ┤                                        │  │     
 35888 ┤                                        │  │     
 33645 ┤                                        │  │     
 31402 ┤                                    ╭──╮│  │     
 29159 ┤                                    │  ││  │     
 26916 ┤                                    │  ││  │     
 24673 ┤                                    │  ││  │     
 22430 ┤                                    │  ││  │     
 20187 ┤                                ╭──╮│  ││  │     
 17944 ┤                                │  ││  ││  │     
 15701 ┤                                │  ││  ││  │     
 13458 ┤                                │  ││  ││  │     
 11215 ┤                            ╭──╮│  ││  ││  │     
  8972 ┤                            │  ││  ││  ││  │     
  6729 ┤                        ╭──╮│  ││  ││  ││  │     
  4486 ┤                    ╭──╮│  ││  ││  ││  ││  │     
  2243 ┤                ╭──╮│  ││  ││  ││  ││  ││  │     
     0 ┼────────────────╯  ╰╯  ╰╯  ╰╯  ╰╯  ╰╯  ╰╯  ╰──── 
```

---

[gomnia-godoc-svg]: https://godoc.org/github.com/calebcase/gomnia?status.svg
[gomnia-godoc]: https://godoc.org/github.com/calebcase/gomnia
[gonum]: https://github.com/gonum/gonum
