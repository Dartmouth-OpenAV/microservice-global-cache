package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

func setState(socketKey string, port string, state string) (string, error) {
	function := "setState"

	if port == "1" {
		port = "1:1"
	} else if port == "2" {
		port = "1:2"
	} else if port == "3" {
		port = "1:3"
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized port value: " + port)
		framework.AddToErrors(socketKey, errMsg)
		return port, errors.New(errMsg )
	}

	if state == `"on"` {
		state = "1"
	} else if state == `"off"` {
		state = "0" 
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized state value: " + state)
		framework.AddToErrors(socketKey, errMsg)
		return state, errors.New(errMsg )
	}

	baseStr := "setstate"
	commandStr := baseStr + "," + port + "," + state
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent := framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp := framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	// If we got here, the response was good, so successful return with the state indication
	return "ok", nil
}


func getState(socketKey string, port string) (string, error) {
	function := "getState"

	if port == "1" {
		port = "1:1"
	} else if port == "2" {
		port = "1:2"
	} else if port == "3" {
		port = "1:3"
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized port value: " + port)
		framework.AddToErrors(socketKey, errMsg)
		return port, errors.New(errMsg )
	}
	
	baseStr := "getstate"
	commandStr := baseStr + "," + port
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent := framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp := framework.ReadLineFromSocket( socketKey )
	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	// if len(respMute) != 2 {
	// 	errMsg := fmt.Sprintf(function+" - fmcx3rsd wrong number of tokens in response: %s", string(resp))
	// 	framework.AddToErrors(socketKey, errMsg)
	// 	return string(resp), errors.New(errMsg)
	// }

	value := "unknown"
	splitResp := strings.Split(resp, ",")
	if splitResp[2] == "0" {
		value = `"off"`
	} else if splitResp[2] == "1" {
		value = `"on"`
	} else { // not a legal value
		errMsg := function + " - unrecognized state returned: " + resp + " is not a legal value\n"
		framework.AddToErrors(socketKey, errMsg)
		return value, errors.New(errMsg)

	}

	// If we got here, the response was good, so successful return with the state indication
	return value, nil
}


func setTrigger(socketKey string, port string, sleep string) (string, error) {
	function := "setTrigger"
	// framework.Log( "[" + state + "]" )
	// framework.Log( "[" + `"true"` + "]" )

	if port == "1" {
		port = "1:1"
	} else if port == "2" {
		port = "1:2"
	} else if port == "3" {
		port = "1:3"
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized port value: " + port)
		framework.AddToErrors(socketKey, errMsg)
		return port, errors.New(errMsg )
	}


	// step 1 - we pull the trigger: "on" position
	baseStr := "setstate"
	commandStr := baseStr + "," + port + ",1"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent := framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp := framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	// step 2 - sleep for 1
	intVar, err := strconv.Atoi(sleep)
	if( err!=nil ) {
		errMsg := fmt.Sprintf(function+" - can't parse sleep time to int: "+sleep)
		framework.AddToErrors(socketKey, errMsg)
	}
	if( intVar>5 ) {
		errMsg := fmt.Sprintf(function+" - sleep time can't be greater than 5")
		framework.AddToErrors(socketKey, errMsg)
		intVar = 5
	}
	if( intVar<1 ) {
		errMsg := fmt.Sprintf(function+" - sleep time can't be less than 1")
		framework.AddToErrors(socketKey, errMsg)
		intVar = 1
	}
	time.Sleep( time.Duration(intVar) * time.Second )


	// step 3 - back to "off"
	baseStr = "setstate"
	commandStr = baseStr + "," + port + ",0"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent = framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp = framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	// If we got here, the response was good, so successful return with the state indication
	return "ok", nil
}

func setTimedtrigger(socketKey string, port1 string, port2 string, sleep string) (string, error) {
	function := "setTimedtrigger"

	// step 1 - run the setTrigger function which turns on trigger, sleeps for 1, then turns off trigger
	setTrigger(socketKey, port1, "1")
	// framework.Log( "[" + state + "]" )
	// framework.Log( "[" + `"true"` + "]" )

	if port1 == "1" {
		port1 = "1:1"
	} else if port1 == "2" {
		port1 = "1:2"
	} else if port1 == "3" {
		port1 = "1:3"
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized port1 value: " + port1)
		framework.AddToErrors(socketKey, errMsg)
		return port1, errors.New(errMsg )
	}

	if port2 == "1" {
		port2 = "1:1"
	} else if port2 == "2" {
		port2 = "1:2"
	} else if port2 == "3" {
		port2 = "1:3"
	} else {
		errMsg := fmt.Sprintf(function + " - unrecognized port2 value: " + port2)
		framework.AddToErrors(socketKey, errMsg)
		return port2, errors.New(errMsg )
	}

	// step 2 - sleep for length of floatVar
	floatVar, err := strconv.ParseFloat(sleep, 64)
	if( err!=nil ) {
		errMsg := fmt.Sprintf(function+" - can't parse sleep time to int: "+sleep)
		framework.AddToErrors(socketKey, errMsg)
	}
	if( floatVar>30 ) {
		errMsg := fmt.Sprintf(function+" - sleep time can't be greater than 30")
		framework.AddToErrors(socketKey, errMsg)
		floatVar = 30
	}
	if( floatVar<1 ) {
		errMsg := fmt.Sprintf(function+" - sleep time can't be less than 1")
		framework.AddToErrors(socketKey, errMsg)
		floatVar = 1
	}
	time.Sleep( time.Duration(floatVar - 1) * time.Second)

	// step 3 - then we pull the trigger: "on" position for both ports at the same time to "stop" the screen
	baseStr := "setstate"
	commandStr := baseStr + "," + port1 + ",1"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent := framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp := framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	
	baseStr = "setstate"
	commandStr = baseStr + "," + port2 + ",1"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent = framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp = framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	time.Sleep(time.Duration(1)*time.Second)

	// step 4 - back to "off"
	baseStr = "setstate"
	commandStr = baseStr + "," + port1 + ",0"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent = framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp = framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	
	baseStr = "setstate"
	commandStr = baseStr + "," + port2 + ",0"
	// resp, err := SendCommand(socketKey, commandStr, baseStr)
	sent = framework.WriteLineToSocket( socketKey, commandStr + string(0xD) + string(0xA) )
	//framework.Log(fmt.Sprintf(function + " - ot resp: " + string(resp) + "\n")
	if !sent {
		errMsg := fmt.Sprintf(function+" - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	resp = framework.ReadLineFromSocket( socketKey )

	framework.Log(fmt.Sprintf(function + " - resp: " + string(resp) + "\n"))

	if( resp!=commandStr ) {
		errMsg := fmt.Sprintf(function+" - unrecognized response to state command: " + resp )
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	// If we got here, the response was good, so successful return with the state indication
	return "ok", nil
}