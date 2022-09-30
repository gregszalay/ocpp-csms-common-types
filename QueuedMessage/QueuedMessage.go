package QueuedMessage

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type QueuedMessage struct {
	MessageId string
	DeviceId  string
	Payload   interface{}
}

func (j *QueuedMessage) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(b), &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	// MessageTypeId integer This is a Message Type Number which is used to identify the type of the message.
	message_id, err_message_id := raw["MessageId"].(string)
	if !err_message_id || message_id == "" {
		return fmt.Errorf("QueuedMessage data[0] is not a Number")
	}
	// Device Id
	device_id, err_device_id := raw["DeviceId"].(string)
	if !err_device_id || device_id == "" {
		return fmt.Errorf("QueuedMessage data[1] is not a string")
	}
	// Payload
	payload, payload_err := raw["Payload"].(map[string]interface{})
	if !payload_err || payload == nil {
		return fmt.Errorf("QueuedMessage data[2] is not a map[string]interface{}")
	}
	*j = QueuedMessage{
		MessageId: message_id,
		DeviceId:  device_id,
		Payload:   payload,
	}
	return nil
}

func (c *QueuedMessage) Marshal() []byte {
	result, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Could not marshal CALL message: %s\n", err)
		return []byte("")
	}
	return result
}

func (c *QueuedMessage) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}

func (c *QueuedMessage) GetPayloadAsJSON() []byte {
	if c == nil {
		fmt.Println("CALL object is empty")
		return []byte("")
	}
	// Re-marshal payload only
	re_marshalled_payload, re_marshall_err := json.MarshalIndent(c.Payload, "", " ")
	if re_marshall_err != nil {
		fmt.Println("Failed to remarshall CALL pazload to json")
		return []byte("")
	}

	return re_marshalled_payload
}
