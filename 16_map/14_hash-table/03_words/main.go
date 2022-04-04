package main

import "fmt"

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([][]string,12)
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word,12)
	}

	
}


func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comment out the above, then uncomment the below
	// a more uneven distribution
	// return len(word) % buckets
}