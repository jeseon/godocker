package main


func main() {
	r := NewServer()
	r.Use(Logging)
	r.Run()
}