//formulas :
package main 

import "fmt"
			
func main(){

// // área del cuadrado en Go :
// const lado int=20
// areaCuadrado:= lado*lado
// fmt.Println(`El área del cuadrado es : `, areaCuadrado)

// // Rectangulo
// const base int=40
// const alturaRec int=80
// areaRectangulo:= base*alturaRec
// fmt.Println(`El área del rectangulo es : `, areaRectangulo)

// //Trapecio
// const baseTrapecio int=40
// const alturaTrapecio int=80
// areaTrapecio:= (baseTrapecio+alturaTrapecio)*alturaTrapecio/2
// fmt.Println(`El área del trapecio es : `, areaTrapecio)

// //círculo
// const radio float64 =10
// const pi float64=3.14159265359
// areaCirculo:= pi*radio*radio
// fmt.Println(`El área del círculo es : `, areaCirculo)
	
// //concatenación de variables...
// var message string ="hola"
// var message1 string ="mundo"

// fmt.Println(message + message1) //print en pantalla...

var nombre= "Platzi"
var cursos= 500
// fmt.Printf( "%s  tiene mas de %d cursos\n", nombre , cursos)
//impresión en pantalla:  Platzi  tiene mas de 500 cursos

//Sprintf
message := fmt.Sprintf( "%s  tiene mas de %d cursos\n", nombre , cursos)
var  mensaje string= message

fmt.Println(mensaje)
}



