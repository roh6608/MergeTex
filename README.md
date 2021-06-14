# MergeTex
A command line interface application that can merge .tex files together.

## Getting Started
To get started, the repository can be cloned or the binaries downloaded. To download the binaries go to the [releases](https://github.com/roh6608/MergeTex/releases) section. To clone the repository;

```
git clone https://github.com/roh6608/MergeTex
```
The binaries can then be re-compiled using the Makefile located in the root directory of the repository.

```
make compile
```
To install the binaries so they are accessable anywhere within the computer, follow the tutorial [here](https://golang.org/doc/tutorial/compile-install).

### Example
An example use of the MergeTex application is shown below. Where the ```-i``` flag signifies the input directory that contains the .tex files to be merged together. These files are to be ordered in the directory in the order they are to be merged. The ```-o``` flag signifies the output of the merged file.

```
./MergeTex -i ./directory -o merged.tex
```

### Options
MergeTex takes command line flag arguments, the arguments it taks can be listed by running;

```
./MergeTex -h
```
This will list the flags, the defaults and what arguments are expected.

## Authors
- https://github.com/roh6608
## Contributing

## Licence
This application is licenced under the Apache 2.0 licence. More information can be found in the [Licence](https://github.com/roh6608) file.
