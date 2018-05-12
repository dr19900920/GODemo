package main

func main() {
	cli := CLI{}
	//	cli.CreateBlockchain("Ivan")
	//	cli.send("Ivan", "Pedro", 6)
	cli.send("Pedro", "Helen", 2)
	cli.send("Ivan", "Helen", 2)
	cli.getBalance("Ivan")
	cli.getBalance("Pedro")
	cli.getBalance("Helen")

}
