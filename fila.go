package fila
import (
	"encoding/json"
)

type Postal struct {
	code int
}
type Address struct{
	street string
	number int
	postal Postal
}
func main() {
	adr := Address {
		street: "av jurema",
		number: 1,
		postal: Postal{
			code: 12,
		},
	}
	adrJson := json.Marshall(adr)
	if err != nil{
		println("deu erro")
	}
	println(adrJson)
}