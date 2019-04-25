package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
)

/*
	Analise dos dados extraidos dos arquivos dados como entrada 

	Com a distancia máxima dada como entrada o algoritmo forma grupos

	Entrada - Arquivos a serem analisados 
	saida - Distância e os Grupos 
*/

func arqAnalise(arqD, arqP string) ([][]float64, float64) {

	// Abrindo os arquivos para analise
	arqPontos, ERRO := os.Open(arqP) // Abre arquivo 
	
	if ERRO != nil{
		print("A leitura deu errado nas coordenadas.")
		os.Exit(3)
	}

	arqDist, ERRO := os.Open(arqD) // Abre arquivo 
	
	if ERRO != nil{
		print("A leitura deu errado na distancia.")
		os.Exit(3)
	}

	// Pegar valor da distancia 
	var dist float64
	fmt.Fscanf(arqDist, "%f", &dist)
	//fmt.Println(dist)
	// Pegando os pontos(linhas) e a quantidade de pontos(quantidade de linhas)
	linhas, quantLinhas := pegaLinhas(arqPontos)
	//fmt.Print(linhas)
	
	// Convertendo as linhas em pontos
	pontos := retornaPontos(linhas, quantLinhas)
	//fmt.Print(pontos)

	arqPontos.Close()
	arqDist.Close()
	return pontos, dist
}


/*
Pega um arquivo e retorna suas linhas e a quantidade de linhas 

entrada: O arquivo que deve ser lido 

saida: linhas separadas em strings e a quantidade de linhas 
*/

func pegaLinhas(arqPontos *os.File) ([]string, int) {
	// Determina quant
	scanner := bufio.NewScanner(arqPontos)
	var linhas []string 
	quantLinhas := 0 

	for scanner.Scan() {
		linhas = append(linhas, scanner.Text()) // pegando as linhas do arquivo 
		quantLinhas++ 
	}

	return linhas, quantLinhas
}

/*
Transforma as linhas em pontos 

entrada - slice de linhas e quantidade de linhas do arquivo 

saida - as coordenadas dos pontos 
*/
func retornaPontos(linhas []string , quantLinhas int) [][]float64 {

	// Determina a dimensão do plano
	dimensão := len(strings.Fields(linhas[0]))

	// Cria slice de pontos 
	pontos := make([][]float64, quantLinhas)
	for i := 0; i < quantLinhas; i++ {
		pontos[i] = make([]float64, dimensão)
	}

	// Pegar os pontos das linhas:
	for i := 0; i < quantLinhas ;  i++{
		
			coordenadas := strings.Fields(linhas[i]) // Separando coordenadas 

			for j := 0; j < dimensão;  j++{
				fmt.Sscanf(coordenadas[j], "%f", &pontos[i][j]) //Criando pontos com as coordenadas separadas
			}
	}

	return pontos
}

/*
Cria arquivos de saida com os dados analisados 
*/
func saida(SSE float64, grupos [][]int) {
	conteudo := fmt.Sprintf("%.4f", SSE) // Criando uma String com o conteudo de SSE
	ERRO := ioutil.WriteFile("result.txt", []byte(conteudo), 0666) //esvreve a soma em um arquivo (pq 0666?)


	if ERRO != nil{
		print("A escrita do SSE deu errado")
		os.Exit(3)
	}

	// Criando uma String com os grupos de pontos:
	var stringGrupos string 
	for i := 0; i < len(grupos); i++ {
		for j := 0; j < len(grupos[i]); j++ {
			stringGrupos += strconv.Itoa(grupos[i][j]) + " "
		}
		stringGrupos+="\n"
	}

	ERRO = ioutil.WriteFile("saida.txt", []byte(stringGrupos), 0666) // escreve os grupos num arquivo

	if ERRO != nil {
		print("A escrita dos grupos deu errado")
		os.Exit(3)
	}


}