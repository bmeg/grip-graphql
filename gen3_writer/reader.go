package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"google.golang.org/protobuf/types/known/structpb"
)

func StreamEdgesFromReader(reader io.Reader, workers int) (chan *gripql.Edge, error) {
	if workers < 1 {
		workers = 1
	}
	if workers > 99 {
		workers = 99
	}
	lineChan := processReader(reader, workers)

	edgeChan := make(chan *gripql.Edge, workers)
	var wg sync.WaitGroup

	jum := gripql.NewFlattenMarshaler()
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for line := range lineChan {
				e := &gripql.Edge{}
				err := jum.Unmarshal([]byte(line), e)
				if err != nil {
					log.WithFields(log.Fields{"error": err}).Errorf("Unmarshaling edge: %s", line)

				} else {
					edgeChan <- e
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(edgeChan)
	}()

	return edgeChan, nil
}

func StreamVerticesFromReader(reader io.Reader, workers int) (chan *gripql.Vertex, error) {
	if workers < 1 {
		workers = 1
	}
	if workers > 99 {
		workers = 99
	}
	lineChan := processReader(reader, workers)

	vertChan := make(chan *gripql.Vertex, workers)
	var wg sync.WaitGroup

	jum := gripql.NewFlattenMarshaler()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for line := range lineChan {
				v := &gripql.Vertex{}
				err := jum.Unmarshal([]byte(line), v)

				if err != nil {
					log.WithFields(log.Fields{"error": err}).Errorf("Unmarshaling vertex: %s", line)
				} else {
					vertChan <- v
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(vertChan)
	}()

	return vertChan, nil
}

func streamJsonFromReader(reader io.Reader, graph string, mapExtraArgs map[string]any, workers int) (chan *gripql.RawJson, chan string) {
	lineChan := processReader(reader, workers)

	vertChan := make(chan *gripql.RawJson, workers)
	warnings := make(chan string, workers)
	var wg sync.WaitGroup
	jum := gripql.NewFlattenMarshaler()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lineChan {
				if !json.Valid([]byte(line)) {
					log.WithFields(log.Fields{"line": line}).Errorf("Skipping invalid JSON line")
					warnings <- fmt.Sprintf("Invalid Json: %s", line)
					continue
				}
				rawData := &gripql.RawJson{
					Data:      &structpb.Struct{},
					ExtraArgs: &structpb.Struct{},
					Graph:     graph,
				}
				err := jum.Unmarshal([]byte(line), rawData.Data)
				if err != nil {
					log.WithFields(log.Fields{"error": err}).Errorf("Unmarshaling into rawData.Data: %s", line)
					warnings <- fmt.Sprintf("Error: %v when unmarshaling: %s", err, line)
					continue
				}
				stuctpbExtraArgs, err := structpb.NewStruct(mapExtraArgs)
				if err != nil {
					log.WithFields(log.Fields{"error": err}).Errorf("Creating new StructPB rawData.ExtraArgs: %s", mapExtraArgs)
					warnings <- fmt.Sprintf("Error: %v when creating new struct for extra args: %s", err, mapExtraArgs)
					continue
				}
				rawData.ExtraArgs = stuctpbExtraArgs
				vertChan <- rawData
			}
		}()
	}

	go func() {
		wg.Wait()
		close(vertChan)
		close(warnings)
	}()

	return vertChan, warnings
}

func processReader(reader io.Reader, chansize int) <-chan string {
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	maxCapacity := 16 * 1024 * 1024
	scanner.Buffer(buf, maxCapacity)

	lineChan := make(chan string, chansize)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			lineChan <- line
		}
		if err := scanner.Err(); err != nil {
			log.WithFields(log.Fields{"error reading from reader: ": err})
		}
		close(lineChan)
	}()

	return lineChan
}
