package main

import (
    "os"
    "math"
    "github.com/orsinium-labs/gamepad"
    "strconv"
    "image"
    "time"
)


func main(){
    connHost := os.Getenv("XDG_RUNTIME_DIR") + "/hypr/" + os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") + "/.socket.sock"
    
    //fmt.Println(GetWorkspaces(connHost))
    //fmt.Println(GetCurrentWorkspace(connHost))
    //SwitchWorkspace(connHost, "3")
    gamepad, err := gamepad.NewGamepad(1)
    if err != nil {
	println("likely no controller connected")
	return
    }
    resetLB := true
    resetRB := true
    triggerDeadzone := 20 
    stickDeadzone := 20
    var rightStick image.Point
    for true {
	state, _ := gamepad.State()
    
	if state.RB() && resetRB && state.RT() > triggerDeadzone{
	    SwitchWorkspaceAndMove(connHost, strconv.Itoa(GetCurrentWorkspace(connHost).ID+1))
	    resetRB = false
	} else if state.RB() && resetRB{
	    SwitchWorkspace(connHost, strconv.Itoa(GetCurrentWorkspace(connHost).ID+1))
	    resetRB = false
	} else if !state.RB(){
	    resetRB = true
	}
	if state.LB() && resetLB && state.RT() > triggerDeadzone{
	    SwitchWorkspaceAndMove(connHost, strconv.Itoa(GetCurrentWorkspace(connHost).ID-1))
	    resetLB = false
	} else if state.LB() && resetLB{
	    SwitchWorkspace(connHost, strconv.Itoa(GetCurrentWorkspace(connHost).ID-1))
	    resetLB = false
	} else if !state.LB(){
	    resetLB = true
	}


	//window movement
	rightStick = state.RS();
	if rightStick.X > stickDeadzone && state.RT() > triggerDeadzone {
	    MoveWindowWithFocus(connHost, "r")
	    time.Sleep(250*time.Millisecond)
	} else if rightStick.X > stickDeadzone {
	    MoveFocus(connHost, "r")
	    time.Sleep(250*time.Millisecond)
	}
	if rightStick.X < -stickDeadzone && state.RT() > triggerDeadzone {
	    MoveWindowWithFocus(connHost, "l")
	    time.Sleep(250*time.Millisecond)
	} else if rightStick.X < -stickDeadzone {
	    MoveFocus(connHost, "l")
	    time.Sleep(250*time.Millisecond)
	}
	if rightStick.Y > stickDeadzone && state.RT() > triggerDeadzone {
	    MoveWindowWithFocus(connHost, "d")
	    time.Sleep(250*time.Millisecond)
	} else if rightStick.Y > stickDeadzone {
	    MoveFocus(connHost, "d")
	    time.Sleep(250*time.Millisecond)
	} 
	if rightStick.Y < -stickDeadzone && state.RT() > triggerDeadzone {
	    MoveWindowWithFocus(connHost, "u")
	    time.Sleep(250*time.Millisecond)
	} else if rightStick.Y < -stickDeadzone {
	    MoveFocus(connHost, "u")
	    time.Sleep(250*time.Millisecond)
	}
	leftStick := state.LS()
	if ((int(math.Abs(float64(leftStick.Y))) > stickDeadzone || int(math.Abs(float64(leftStick.X))) > stickDeadzone) && state.RT() > triggerDeadzone) {
	    resizeParams := strconv.Itoa(int(leftStick.X/10))+ " "+ strconv.Itoa(int(leftStick.Y/10)*-1)
	    ResizeActive(connHost, resizeParams)
	    time.Sleep(25*time.Millisecond)
	}





    }
}
