// área del cuadrado en Go :
package main

import "fmt"

func main(){
	const baseCuadrado int =10
	areaCuadrado:= baseCuadrado*baseCuadrado
	   fmt.Println(`El área del cuadrado es : `, areaCuadrado)
// operadores --------->
	   x:=50
	   y:=10
//suma
	   z:=x+y
	   fmt.Println(`La suma de x+y es : `, z)
//resta
	   z=x-y
	   fmt.Println(`La resta de x-y es : `, z)
//multiplicacion
	   z=x*y
	   fmt.Println(`La multiplicacion de x*y es : `, z)
//division
	   z=x/y
	   fmt.Println(`La division de x/y es : `, z)
}
