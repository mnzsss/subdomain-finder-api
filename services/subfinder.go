package services

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func Subfinder(domain string) (any, error) {
	subfinderOpts := &runner.Options{
		Threads:            10,     // Thread controls the number of threads to use for active enumerations
		Timeout:            3 * 60, // 3 mins
		MaxEnumerationTime: 10,     // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		JSON:               true,   // JSON specifies whether to use JSON for output
	}

	// disable timestamps in logs / configure logger
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		log.Fatalf("failed to create subfinder runner: %v", err)
		return nil, err
	}

	output := &bytes.Buffer{}
	// To run subdomain enumeration on a single domain
	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		log.Fatalf("failed to enumerate single domain: %v", err)
		return nil, err
	}

	result, err := ndjsonToJSON(output.String())
	if err != nil {
		fmt.Println("Error on convert to json:", err)
		return nil, err
	}

	return result, nil
}

func ndjsonToJSON(ndjson string) ([]interface{}, error) {
	var result []interface{}

	scanner := bufio.NewScanner(strings.NewReader(ndjson))
	for scanner.Scan() {
		line := scanner.Text()

		var obj interface{}
		err := json.Unmarshal([]byte(line), &obj)
		if err != nil {
			return nil, err
		}

		result = append(result, obj)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
