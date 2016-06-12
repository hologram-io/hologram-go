package HologramGo

import (
	"fmt"
	"os"
	"strconv"
)

type Device map[string]interface{}

// Devices is just a list of Device(s).
type Devices []Device

// TODO: Fixes Devices return type.
// EFFECTS: Returns device details.
func GetDevices() Device {

	req := createGetRequest("/devices")

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
	return payload["data"].(map[string]interface{})
}

// REQUIRES: a device id.
// EFFECTS: Returns device details.
func GetDevice(deviceid int) Device {

	req := createGetRequest("/devices/" + strconv.Itoa(deviceid))

	resp, err := SendRequest(req)

	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	// error handling
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return payload["data"].(map[string]interface{})
}

// REQUIRES: A sim number.
// EFFECTS: Claim ownership and activate the given device.
func ClaimOwnershipAndActiveDevice(simnumber int) Device {

	var params Parameters
	req := createPostRequest("/cellular/sim_" + strconv.Itoa(simnumber), params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
	return payload["data"].(map[string]interface{})
}

// REQUIRES: A device id.
// EFFECTS: Purchase and assign a phone number to the device.
func PurchaseAndAssignPhoneNumberToDevice(deviceid int) Device {

	var params Parameters
	req := createPostRequest("/devices/" + strconv.Itoa(deviceid) + "/addnumber", params)

	resp, err := SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	var payload = Placeholder{}
	err = resp.Parse(&payload)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
	return payload["data"].(map[string]interface{})
}

// EFFECTS: Returns the user id.
func (device Device) GetDeviceUserId() float64 {
	return device["userid"].(float64)
}

// EFFECTS: Returns the device name.
func (device Device) GetDeviceName() string {
	return device["name"].(string)
}

// TODO: Handle as Time type
func (device Device) GetWhenCreated() string {
	return device["whencreated"].(string)
}

// EFFECTS: Returns a phone number.
func (device Device) GetPhoneNumber() string {
	return device["phonenumber"].(string)
}

// EFFECTS: Returns true if it is tunnelable.
func (device Device) GetTunnelable() bool {
	return device["tunnelable"].(bool)
}

// TODO: links/cellular
