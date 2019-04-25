package main

import "math"


/*
Cacula distancia entre dois pontos 

entrada - dois vetores de floats 

saida - um valor float 
*/
func distanciaEntrePontos(x, x0 []float64) float64 {
	var distancia float64 = 0
	for i := 0; i < len(x); i++ {
		distancia += (x[i] - x0[i]) * (x[i] - x0[i])
	}
	distancia = math.Sqrt(distancia)

	return distancia
}


/*
Calcula centro de massa do grupo 

entrada - slice de pontos, um grupo

saida - o ponto central do grupo 
*/
func centroDeMassa(pontos [][]float64, grupo []int) []float64 {
	dimensao := len(pontos[0]) // A dimensão do ponto 
	tamGrupo := len(grupo) // quantidade de pontos no grupo 
	centro := make([]float64, dimensao) // criando varivel de centro de massa

	// incializando variavel centro de massa
	for i := 0; i < dimensao; i++ {
		centro[i] = 0
	}

	// Calcula o centro de massa do grupo
	for i := 0; i < tamGrupo; i++ {
		ponto := pontos[grupo[i] - 1] // Pegando ponto de um grupo

		// Somando as coordenado a coordenada ao centro de massa 
		for j := 0; j < dimensao; j++ {
			centro[j] += ponto[j] / float64(tamGrupo) 
		}
	}
	return centro
}

/*
Calcula a soma euclidiana das distancias entre os grupos 

entrada - as slices de grupos 

saida - SSE
*/
func calcSSE(pontos [][]float64, grupos [][]int ) float64 {
	var SSE float64 = 0

	for i := 0; i < len(grupos); i++ {

		centro := centroDeMassa(pontos, grupos[i])
		for j := 0; j < len(grupos[i]); j++ {
			SSE += distanciaEntrePontos(pontos[grupos[i][j] - 1], centro) * distanciaEntrePontos(pontos[grupos[i][j]-1], centro)

		}
	}

	return SSE
}


/*
Calcula a soma da distância euclidiana e os pontos que formam cada grupo

entrada - distância máxima e o a slice de pontos 

saida - soma euclidana e os grupos 

*A primeira posição de cada grupo será seu número
*/

func algoritmoLider(dist float64, pontos [][]float64) (float64, [][]int) {

	// Criando grupos 
	var grupos [][]int 
	var grupo0 []int 

	// inicia grupos 
	grupo0 = append(grupo0, 1)
	grupos = append(grupos, grupo0)

	// Percorrer todos os pontos
	for i := 1; i < len(pontos); i++ {
		lider := true

		// checando se o ponto atual é um lider
		for j := 0; j < len(grupos); j++ {
			if distanciaEntrePontos(pontos[grupos[j][0] - 1], pontos[i]) <= dist {
				//caso tenha uma distancia menor q a dada como entrada em relação
				// a um outro ponto q ja seja lider o novo ponto será colocado em um 
				//grupo ja criado 
				grupos[j] = append(grupos[j], i+1)
				lider = false
				break
			}
		}

		//Criando um novo grupo para um novo lider 
		if lider {
			var grupo []int
			grupo = append(grupo, i+1)
			grupos = append(grupos, grupo)
		}

	}

	sse := calcSSE(pontos, grupos)

	return sse, grupos

}