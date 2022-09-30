package QueuedError

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type QueuedError struct {
	MessageId        string
	DeviceId         string
	ErrorCode        string
	ErrorDescription string
	ErrorDetails     string
}

func (j *QueuedError) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(b), &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	// MessageTypeId integer This is a Message Type Number which is used to identify the type of the message.
	message_id, err_message_id := raw["MessageId"].(string)
	if !err_message_id || message_id == "" {
		return fmt.Errorf("QueuedError data[0] is not a Number")
	}
	// Device Id
	device_id, err_device_id := raw["DeviceId"].(string)
	if !err_device_id || device_id == "" {
		return fmt.Errorf("QueuedError data[1] is not a string")
	}

	error_code, err_error_code := raw["ErrorCode"].(string)
	if !err_error_code || error_code == "" {
		return fmt.Errorf("QueuedError data[2] is not a string")
	}

	error_description, err_error_description := raw["ErrorDescription"].(string)
	if !err_error_description || error_description == "" {
		return fmt.Errorf("QueuedError data[3] is not a string")
	}

	error_details, err_error_details := raw["ErrorDetails"].(string)
	if !err_error_details || error_details == "" {
		return fmt.Errorf("QueuedError data[4] is not a string")
	}

	*j = QueuedError{
		MessageId:        message_id,
		DeviceId:         device_id,
		ErrorCode:        error_code,
		ErrorDescription: error_description,
		ErrorDetails:     error_details,
	}
	return nil
}

func (c *QueuedError) Marshal() []byte {
	result, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Could not marshal QueuedError message: %s\n", err)
		return []byte("")
	}
	return result
}

func (c *QueuedError) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}
