package bitmap

type Bitmap struct {
	// This is a array of word, one word is equal 4 bytes
	words []uint
}

const (
	bitNum = 32 << ((^uint(0)) >> 63) // computer is 32 bits or 64 bits
)

// @title Add
// @description add the element to the set
// @param x int
// @return null
func (s *Bitmap) Add(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = s.words[word] | 1<<bit
}

// @title Remove
// @description remove the element in the set
// @param x int
// @return null
func (s *Bitmap) Remove(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

// @title Clear
// @description make the bitmap empty
// @param null
// @return null
func (s *Bitmap) Clear() {
	s.words = append([]uint{})
}

// @title Copy
// @description copy the bitmap to a new bitmap
// @param null
// @return *Bitmap
func (s *Bitmap) Copy() *Bitmap {
	intSet := &Bitmap{
		words: []uint{},
	}
	for _, value := range s.words {
		intSet.words = append(intSet.words, value)
	}
	return intSet
}

// @title UnionWith
// @description whether the element already exists in the set
// @param x int
// @return bool
func (s *Bitmap) Has(x int) bool {
	word, bit := x/bitNum, uint(x%bitNum)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// @title UnionWith
// @description return the every element in the set
// @param null
// @return []int
func (s *Bitmap) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, bitNum*i+j)
			}
		}
	}
	return elems
}

// @title UnionWith
// @description calculate the unionSet
// @param t *Bitmap
// @return null
func (s *Bitmap) UnionWith(t *Bitmap) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] | word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// @title IntersectWith
// @description calculate the interSet
// @param t *Bitmap
// @return null
func (s *Bitmap) IntersectWith(t *Bitmap) {
	for i, word := range t.words {
		if i >= len(s.words) {
			continue
		}
		s.words[i] &= word
	}
}

// @title DifferenceWith
// @description calculate the differenceSet
// @param t *Bitmap
// @return null
func (s *Bitmap) DifferenceWith(t *Bitmap) {
	t1 := t.Copy()
	t1.IntersectWith(s)
	for i, word := range t1.words {
		if i < len(s.words) {
			s.words[i] ^= word
		}
	}
}
