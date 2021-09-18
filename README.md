# Bloom-Filter

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/akgarhwal/BloomFilter/main/LICENSE)   [![GoDoc](https://godoc.org/github.com/akgarhwal/bloomfilter?status.svg)](https://pkg.go.dev/github.com/akgarhwal/bloomfilter)


A [Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set.

An empty Bloom filter is a bit array of <b>m</b> bits, all set to 0. There must also be <b>k</b> different hash functions defined, each of which maps or hashes some set element to one of the <b>m</b> array positions, generating a uniform random distribution.

```
m -> size of bit array
k -> number of hash function
n -> estimated number of element to add in bloom filter
```

[False positive](https://en.wikipedia.org/wiki/Type_I_and_type_II_errors) matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set".


## Supported Operations
```
  insert(x) : To insert an element in the Bloom Filter.
  lookup(x) : to check whether an element is already present 
              in Bloom Filter with a false positive probability.
```
<b>NOTE:</b> <i>Delete</i> operation can be supported by using [Counting Bloom filter](https://en.wikipedia.org/wiki/Counting_Bloom_filter).

## Install

```
go get github.com/akgarhwal/bloomfilter
```

## Usage

```go
// create a new bloom filter 
// with 100 elements to be inserted and 
// false positive probability of 0.1
bf := bloomfilter.NewBloomFilter(100, 0.1)
bf.Insert([]byte("sachin"))
bf.Insert([]byte("rohit"))
bf.Insert([]byte("virat"))
bf.Insert([]byte("msdhoni"))

// new user name to check
newUserName := "msdhoni"

if bf.Lookup([]byte(newUserName)) {
  fmt.Println(newUserName, " is present")
} else {
  fmt.Println(newUserName, " is NOT present")
}
```

## Probability of False Positive
If <b>m</b> is size of bit array and <b>n</b> is number of elements to be inserted and <b>k</b> is number of hash functions, then the probability of false positive ε can be calculated as:<br/>
<img width="198" alt="Screenshot 2021-09-14 at 1 56 10 PM" src="https://user-images.githubusercontent.com/20686129/133223233-d695899c-fb09-4fd7-a725-3a4ce5071e07.png">

## Size of Bit Array:
If expected number of elements <b>n</b> is known and desired false positive probability is <b>ε</b> then the size of bit array:<br/> 
<img width="146" alt="Screenshot 2021-09-14 at 2 00 54 PM" src="https://user-images.githubusercontent.com/20686129/133223870-a8128bab-6ec1-4162-9099-2155420fe47f.png">

## Optimal number of hash functions
The number of hash functions <b>k</b> must be a positive integer. If <b>m</b> is size of bit array and <b>n</b> is number of elements to be inserted, then <b>k</b> can be calculated as:<br>
<img width="138" alt="Screenshot 2021-09-14 at 2 05 38 PM" src="https://user-images.githubusercontent.com/20686129/133224561-82cb32c3-9382-40df-a96c-56d7b4a6e753.png">



## License
This package is licensed under MIT license. See LICENSE for details.
