package main
import "github.com/joho/godotenv"

func main(){
	if error := godotenv.Load();err !== nil{
		panic(err)
	}
}