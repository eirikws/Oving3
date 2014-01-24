package main
import(
    "net"
    "fmt"
    "runtime"
)
const ipAdrServ="129.241.187.161"
const myPort="12001"
const ipAddrThisPC="129.241.187.147"
const serverPort="34933"


func main(){
    runtime.GOMAXPROCS(runtime.NumCPU())
    go server()
	serverAddr, err := net.ResolveTCPAddr("tcp",ipAdrServ+":"+serverPort)
	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {
		// handle error
		}
	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)
	msg:= "Connect to: "+ipAddrThisPC+":"+myPort+"\x00"
	con.Write([]byte(msg))
	
//	stop :=0
//	for stop !=-1{
//		input := ""
//		fmt.Scanf("%s",&input)
//		if input=="stop"{
//			stop=-1
//		}
//		con.Write([]byte(input+"\x00"))
//		con.Read(cd)
//		fmt.Printf("%s\n",cd)
//	}
}


func server(){
    in,err:=net.Listen("tcp",":"+myPort)
    if err !=nil{
    }
    msg:=make([]byte,1024)
    for {
        conn,err:=in.Accept()
        if err!=nil{
        }
        conn.Read(msg)
        fmt.Printf("%s",msg)
    }
}

