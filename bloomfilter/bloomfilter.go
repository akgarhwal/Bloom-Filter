package bloomfilter

import (
	"hash/fnv"
	"math"
)

type BloomFilter interface {
	Insert([]byte)
	Lookup([]byte) bool
}

type bloomFilter struct {
	n        uint64
	m        uint32
	k        int
	bitArray []bool
}

// create a new bloom filter
func NewBloomFilter(n uint64, falsePositiveProb float64) BloomFilter {

	bf := bloomFilter{}

	bf.n = n
	bf.estimateM(n, falsePositiveProb)
	bf.estimateK(n)
	bf.bitArray = make([]bool, bf.m)

	return &bf
}

// Insert(x) : To insert an element in the Bloom Filter.
func (bf *bloomFilter) Insert(element []byte) {
	bf.insert(element)
}

// Lookup(x) : to check whether an element is already present in Bloom Filter
// if element is present it will return true with a false positive probability.
// if element is not present it will return false
func (bf *bloomFilter) Lookup(element []byte) bool {
	return bf.lookup(element)
}

func (bf *bloomFilter) insert(element []byte) {
	for i := 1; i <= bf.k; i++ {
		hash := bf.calculateHash(element)
		index := bf.calculateIndex(hash, uint(i))
		bf.bitArray[index] = true
	}
}

func (bf *bloomFilter) lookup(element []byte) bool {
	for i := 1; i <= bf.k; i++ {
		hash := bf.calculateHash(element)
		index := bf.calculateIndex(hash, uint(i))
		// found one bit which is not set to `true` so return false
		if !bf.bitArray[index] {
			return false
		}
	}
	// all bits are set to `true` so return true
	return true
}

func (bf *bloomFilter) calculateHash(data []byte) uint {
	hash := fnv.New32()
	hash.Write(data)
	return uint(hash.Sum32())
}

func (bf *bloomFilter) calculateIndex(offset uint, hash uint) uint {
	return (hash + offset) % uint(bf.m)
}

func (bf *bloomFilter) estimateM(n uint64, falsePositiveProb float64) {
	//m = -1 * (n * lnP)/(ln2)^2
	nf := float64(n)
	ln2 := math.Log(2)
	num := nf * math.Log(falsePositiveProb)
	deno := math.Pow(ln2, 2)

	bf.m = uint32(-1 * num / deno)
}

func (bf *bloomFilter) estimateK(n uint64) {
	//k = m/n * ln2
	nf := float64(n)
	ln2 := math.Log(2)
	bf.k = int(math.Ceil(float64(bf.m) / nf * ln2))
}
