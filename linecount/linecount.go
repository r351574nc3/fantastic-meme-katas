package linecount

import (
	"bufio"
	"os"
	"crypto/md5"
)

type Sentence struct {
	text string
	count int
}

func CountLinesIn(path string) (map[string]Sentence, error) {
	var sentences map[string]Sentence
	// var h hash.Hash

	h := md5.New()
	sentences = make(map[string]Sentence)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	[]Sentence 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := []byte(scanner.Text())
		// lines = append(lines, scanner.Text())
		
		checksum := string(h.Sum(data))
		if val, ok := sentences[checksum]; ok {
			sentence := sentences[checksum]
			sentence.count = val.count + 1 
		} else {
			sentences[checksum] = Sentence{ string(data), 1 }
		}
	}
	
	By(frequency).Sort(sentences)
	
	return sentences, scanner.Err()
}
