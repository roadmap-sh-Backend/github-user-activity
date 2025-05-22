package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatalf("LACKS COMMAND")
	}

	if args[1] != "github-activity" {
		log.Fatalf("INVALID COMMAND: %v", args[1])
	}

	username := args[2]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR MAKING HTTP REQUEST: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ERROR READING RESPONSE BODY: %v", err)
	}

	var ges []GithubEvent
	err = json.Unmarshal(body, &ges)
	if err != nil {
		log.Fatalf("ERROR READING RESPONSE BODY: %v", err)
	}

	eventType := AllEvent
	if len(args) > 3 {
		eventType = GithubEventType(args[3])
		log.Printf("Github Event %s Type", eventType)
	}

	result := Result{
		Total: 0,
		Data:  []ResultData{},
	}

	mapEvent := make(map[string][]any, 0)
	for _, ge := range ges {
		b, err := json.Marshal(ge.Payload)
		if err != nil {
			log.Fatalf("ERROR RE-MARSHALL PUSH PAYLOAD: %v", err)
		}

		switch ge.Type {
		case string(PushEvent):
			var payload PushEventPayload
			err = json.Unmarshal(b, &payload)
			if err != nil {
				log.Fatalf("ERROR UNMARSHALL PUSH PAYLOAD: %v", err)
			}

			ge.Payload = payload

			mapEvent[ge.Type] = append(mapEvent[ge.Type], ge)
		case string(CreateEvent):
			var payload CreateEventPayload
			err = json.Unmarshal(b, &payload)
			if err != nil {
				log.Fatalf("ERROR UNMARSHALL CREATE PAYLOAD: %v", err)
			}

			ge.Payload = payload

			mapEvent[ge.Type] = append(mapEvent[ge.Type], ge)

			//ADD MORE GITHUB EVENT TYPE CASE
		default:
			slog.Error("UNKNOWN EVENT TYPE")
			continue
		}
	}

	if eventType == AllEvent {
		for eventName, groupedEvents := range mapEvent {
			resultData := ResultData{
				Type: eventName,
			}
			for _, data := range groupedEvents {
				result.Total++
				resultData.Items = append(resultData.Items, data)
			}

			resultData.Size = len(resultData.Items)
			result.Data = append(result.Data, resultData)
		}
	} else {
		resultData := ResultData{
			Type: string(eventType),
		}
		for _, data := range mapEvent[string(eventType)] {
			result.Total++
			resultData.Items = append(resultData.Items, data)
		}

		resultData.Size = len(resultData.Items)
		result.Data = append(result.Data, resultData)
	}

	slog.Info("Results",
		"Total", result.Total,
		"Type", string(eventType),
		"Data", result.Data)
}
