package main

import(
    "net"
    "encoding/json"
    "io"
)

type workspace struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Monitor string `json:"monitor"`
    MonitorID int `json:"monitorID"`
    Windows int `json:"windows"`
    HasFullScreen bool `json:"hasfullscreen"`
    LastWindow string `json:"lastwindow"`
    LastWindowTitle string `json:"lastwindowtitle"`
}
func SwitchWorkspaceAndMove(connHost string, workspace string){
    send(connHost, "dispatch movetoworkspace " + workspace)
}
func SwitchWorkspace(connHost string, workspace string){
    send(connHost, "dispatch workspace " + workspace)
}
func MoveFocus(connHost string, direction string){
    send(connHost, "dispatch hy3:movefocus " + direction)
}
func MoveWindowWithFocus(connHost string, direction string){
    send(connHost, "dispatch hy3:movewindow " + direction)
}
func ResizeActive(connHost string, direction string){
    send(connHost, "dispatch resizeactive " + direction)
}

func GetWorkspaces(connHost string) []workspace{
    workspaces := sendRecWorkspaces(connHost, "j/workspaces")
    return workspaces
}
func GetCurrentWorkspace(connHost string) workspace{
    var currentWorkspace workspace = sendRecWorkspace(connHost, "j/activeworkspace")
    return currentWorkspace
}

func send(connHost string, command string) {
    connection, err := net.Dial("unix", connHost)
    if(err != nil){
	panic(err)
    }
    defer connection.Close()
    _, err = connection.Write([]byte(command))
    if(err != nil) {
	panic(err)
    }
}

func sendRecWorkspace(connHost string, command string) workspace{
    connection, err := net.Dial("unix", connHost)
    if(err != nil){
	panic(err)
    }
    defer connection.Close()
    _, _ = connection.Write([]byte(command))
    data, _ := io.ReadAll(connection)
    var currentWorkspace workspace
    err = json.Unmarshal(data, &currentWorkspace)
    if err != nil {
	panic(err)
    }
    return currentWorkspace
}
func sendRecWorkspaces(connHost string, command string) []workspace{
    connection, err := net.Dial("unix", connHost)
    if(err != nil){
	panic(err)
    }
    defer connection.Close()
    _, _ = connection.Write([]byte(command))
    data, err := io.ReadAll(connection)
    var currentWorkspace []workspace
    err = json.Unmarshal(data, &currentWorkspace)
    if err != nil {
	panic(err)
    }
    return currentWorkspace
}
