# top10url

A command-line utility that reads a file containing URLs and associated numerical values, and outputs the top 10 URLs with the largest values


## Features

### Min-Heap Usage

* A min-heap is used to keep track of the top 10 largest values efficiently.

* The smallest value in the heap can be replaced by a larger value when encountered.

* This ensures that only the largest 10 values are kept in the heap.

### Memory Efficiency

* Only the top 10 largest values and their associated URLs are kept in memory, making the program suitable for handling large files.

### File Reading

* The file is read line by line to avoid loading the entire file into memory at once.


> :warning: **TODO**: Future Improvements
> 
> Parallel Offset-Based Reading of File Chunks Using Goroutines
>   * Splitting the File into Chunks: Divide the file into smaller chunks that can be processed independently.
>   * Reading File in Sections: Use offsets to read only a part of the file at a time.


## Building, Running and Testing

### Build

Make sure you have [Go installed](https://golang.org/doc/install).


```sh
make build
```

*Built binary will be in `bin` directory*

### Run 

```sh
./bin/top10url -file=<file_path> [-top=<number_of_urls>]

```

`file_path` The absolute path to the input file containing URLs and values

`number_of_urls` (Optional) The number of top URLs to output. Defaults to 10


### Run unit tests

```sh
make tests
```

### Run e2e tests

```sh
make e2e
```

### Generate test datasets

*Generated test data will be in `bin` directory*

18G in18G.txt

1.7G in1_7G.txt

1.7M in1_7Mb.txt

```sh
make gendata
```
Run `top10url`

```sh
./bin/top10url -file ./bin/in18G.txt -top 10
```

### Clean 

```sh
make clean
```

## Credits

### RAM/CPU monitoring 

Simple terminal helper to log ram/cpu usage 

It can be run in a separate terminal session and will collect metrics for the running tool.

```
pip install psutil==6.0.0

python ./scripts/showusage.py
```
