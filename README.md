# imgconv

A tiny elven image conversion command line tool.


## Install

``` bash
$ git clone https://github.com/collinglass/imgconv
$ cd imgconv
$ go install
```


## Usage

``` bash
$ imgconv [OPTIONS] IMAGE [IMAGE...]
```

Make sure ```export PATH=$PATH:$GOPATH/bin``` is set.


##### Options
	-to=TYPE                       Choose output image type.
	                               Available: png, jpeg, gif.
	                               Default: png.


##### Image

Relative path to images to convert.


#### To Do

1. Support periods ```.``` in image names

2. Support Regex

3. Improve image quality (?)

