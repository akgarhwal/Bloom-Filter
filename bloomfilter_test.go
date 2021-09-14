package bloomfilter

import "testing"

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)
	bf.Insert([]byte("foo"))
	bf.Insert([]byte("john"))
	bf.Insert([]byte("tom"))

	if bf.Lookup([]byte("tom")) == false {
		t.Errorf("BloomFilter: test failed for element : tom\n")
	}

	if bf.Lookup([]byte("john")) == false {
		t.Errorf("BloomFilter: test failed for element : john\n")
	}

	if bf.Lookup([]byte("foo")) == false {
		t.Errorf("BloomFilter: test failed for element : foo\n")
	}
}
