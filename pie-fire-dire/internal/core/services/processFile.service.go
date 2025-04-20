package services

import (
	"bufio"
	"fmt"
	domain "newnok6/logic-test/pie-fire-dire/internal/core/domain"
	"os"
	"regexp"
	"sync"
)

type processFileService struct {
	meta domain.FileMeta
}

func NewProcessFileService(meta domain.FileMeta) ProcessFileService {
	return &processFileService{meta: meta}
}

func (p *processFileService) GetFileName() string {
	return p.meta.FileName
}

func (p *processFileService) GetFilePath() string {
	return p.meta.FilePath
}

func (p *processFileService) GetMeatList() (map[string]uint32, error) {
	fileLocation := p.meta.FilePath + "/" + p.meta.FileName
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	jobs := make(chan []string, 100)
	results := make(chan map[string]uint32, 100)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go countWords(i, jobs, results, &wg) // FAN-OUT the line to multiple workers
	}

	scanner := bufio.NewScanner(file) // Scan line by line for supporting large files in the future
	for scanner.Scan() {
		line := scanner.Text()
		filteredWords := makeKeySearchFromLine(line)
		if len(filteredWords) > 0 {
			jobs <- filteredWords // send line to single channel
		}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	finalResult := make(map[string]uint32)

	for result := range results {
		for word, count := range result {
			finalResult[word] += count // FAN-IN the result to a single map
		}
	}

	fmt.Println("Final Result:", finalResult)

	return finalResult, nil
}

func countWords(workerId int, jobs <-chan []string, result chan<- map[string]uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	for words := range jobs {
		keyMap := make(map[string]uint32)
		for _, word := range words {
			keyMap[word]++
			//fmt.Println("Worker ID:", workerId, "Processing paragraph:", word)
		}
		result <- keyMap // send the result to the result channel
	}

}

func makeKeySearchFromLine(paragraph string) []string {
	keyList := removeUnusedSymbol(paragraph)
	return keyList
}

func removeUnusedSymbol(paragraph string) []string {
	// Use regular expression to split by ".", "..", space, and ","
	re := regexp.MustCompile(`\.\.|\s|\.|,`)
	result := re.Split(paragraph, -1)

	// Remove empty strings if any
	filtered := make([]string, 0, len(result))
	for _, r := range result {
		if r != "" {
			filtered = append(filtered, r)
		}
	}

	return filtered
}
