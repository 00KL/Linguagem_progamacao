/*
	Trabalho 1 - Linguagens de Progamação
	Aluno - Lucas Moraes Soares de Souza
*/
package main

import "flag"
import "os"

func main(){
	//fmt.Print(os.Args[1])
	// Entrada de dados
	arqPont := flag.String("pontos", os.Args[1], "coordenadas")
	arqDist := flag.String("distancia", os.Args[2], "distancia")
	

	pontos, dist := arqAnalise(*arqDist, *arqPont) //Tratamento de dados
	SSE, grupos := algoritmoLider(dist, pontos) // Algoritmo lider
	
	saida(SSE, grupos)
}