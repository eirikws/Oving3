package main
import "net"
import "fmt"

func main(){
	ipAdr := "129.241.187.161"

	serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+"34933")
	con, err := net.DialTCP("tcp", nil, serverAddr);
	if err != nil {
		// handle error
		}
	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)
	
	msg:= "Connect to: 129.241.187.159:34933\x00"	
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
}
