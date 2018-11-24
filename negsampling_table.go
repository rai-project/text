package text

import (
	"fmt"
	"math"
)


func create_negsamplingunigram(id2freq []uint64, id2freq_size uint64, power float64, tsize uint64) []uint64 {
	/*
		Function:
		Takes in a id to frequency table and returns a unigram table from which negative samples can be
		sampled uniformly
		Arg:
		id2freq - id to frequency table, where id can be vocabulary or any other possible sample space
		id2freq_size - size of the id2freq table
		power - exponent for negative sampling
		tsize - size of the returned unigram table
	*/

	var unigramtable = make([]uint64, tsize)
	var train_words_pow float64 = 0.0
	var i uint64
	for i = 0; i < id2freq_size; i++ {
	        train_words_pow += math.Pow(float64(id2freq[i]), power)
	}
	fmt.Println(train_words_pow)
	var elem uint64 = 0
	var d1 float64 = math.Pow(float64(id2freq[elem]), power) / train_words_pow
	var a uint64
	for a = 0; a < tsize; a++ {
		fmt.Println(d1)
		unigramtable[a] = elem
		if (float64(a+1) / float64(tsize)) > d1 {
			elem += 1
			d1 += math.Pow(float64(id2freq[elem]), power) / train_words_pow
		}
		if elem >= id2freq_size {
			elem = id2freq_size - 1
		}
	}
	return unigramtable
}

// test
/*
func main() {
	id2freq := []uint64{10, 20, 30, 40}
	var id2freq_size uint64 = 4
	var tsize uint64  = 10
	var power float64 = 1.0
	utable := create_negsamplingunigram(id2freq, id2freq_size, power, tsize)
	var i uint64
	for i = 0; i < tsize; i++ {
		fmt.Println(" number ", i, " - ", utable[i])
	}
}
*/
