package bloomfilter

type BloomFilter interface {
	Insert([]byte)
	Lookup([]byte) bool
}

type bloomFilter struct {
	m        uint64
	k        uint16
	bitArray []bool
}

// update this func to take input as totalNumber uint32, false positive prob rate
func NewBloomFilter(m uint64, k uint16) BloomFilter {
	return &bloomFilter{
		m:        m,
		k:        k,
		bitArray: make([]bool, 10000),
	}
}

func (*bloomFilter) Insert(element []byte) {

}

func (*bloomFilter) Lookup(element []byte) bool {
	return true
}
