package main

import (
        "net"
        "fmt"
        "runtime"
)
const MY_IP= "129.241.187.147"
const TARGET_IP = "129.241.187.255"
const TARGET_PORT = "20011"
const LISTEN_PORT = "30011"

func ConnectTo(ipAdr string, port string){
        serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
        con, err := net.DialUDP("udp", nil, serverAddr)
        
        if err != nil {
                fmt.Println("fail")
        }
        
        stop :=0
                
        
        for stop !=-1{
                input := ""
                fmt.Scanf("%s",&input)
                if input=="stop"{
                        stop=-1
                }
                con.Write([]byte(input+"\x00"))
                
        }
}
func ListenerCon(ipAdr string, port string){
        serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
        psock, err := net.ListenUDP("udp4", serverAddr)
        
        if err != nil { return }
        buf := make([]byte,1024)
 
          for {                 
                    if err != nil { return }
                    _, remoteAddr, _ := psock.ReadFromUDP(buf)
                    if remoteAddr.IP.String() != MY_IP {
                        fmt.Printf("%s\n",buf)
                }
         }        
                
}

func main() {
        runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...

        go ConnectTo(TARGET_IP,TARGET_PORT)
        go ListenerCon(TARGET_IP, TARGET_PORT)
        
        for true{
        }
}
