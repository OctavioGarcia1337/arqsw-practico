package main

import (
	"fmt"

	"github.com/OctavioGarcia1337/arq-software/ej-go-division/div"
)

func main() {
	var numerador float32 = 23
	var divisor float32 = 0

	fmt.Println(div.Division(numerador, divisor))

}

/* //NO LO PROBE, SOLO LO COPIE PARA QUE ME QUEDE GUARDADO Y TESTEARLO EN CASA.
var err error
if err!= nil{
	fmt.Println("Errors:", err.Error())
	return
}
fmt.Println("Division:", divisor)
*/
