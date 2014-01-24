package main
import "net"
import "fmt"

\\test


func main(){
    runtime.GOMAXPROX(runtime.NumCPU())
	ipAdrServ := "129.241.187.161"

	serverAddr, err := net.ResolveTCPAddr("tcp",ipAdrServ+":"+"34933")
	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {
		// handle error
		}
	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)
	
	msg:= "Connect to: 129.241.187.159:12001\x00"	
	con.Write([]byte(msg))
	
	stop :=0
	for stop !=-1{
		input := ""
		fmt.Scanf("%s",&input)
		if input=="stop"{
			stop=-1
		}
		con.Write([]byte(input+"\x00"))
		con.Read(cd)
		fmt.Printf("%s",cd)
	}
	server()
}


func server(){
    servAddr,err:=net.ResolveTCPAddr("tcp",
    ln,err:=net.ListenTCP("tcp",nil)
    if err !=nil{
    }
    for {
        conn,err:=ln.Accept()
        go handleconnection(conn)
    }
}

func handleconnection(client net.Conn){
    client.Write([]byte("hei der!!\n\x00"))
}
