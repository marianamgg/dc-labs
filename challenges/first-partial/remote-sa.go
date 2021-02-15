package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"math"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	var area float64
  lados := []float64{}

  for i := 0; i < len(points)-1; i++ {
		var distanciaN float64
		distanciaN = math.Hypot(points[i+1].X-points[i].X, points[i+1].Y-points[i].Y)
		lados = append(lados, distanciaN)
	}
  lados = append(lados, math.Hypot(points[len(points)-1].X-points[0].X, points[len(points)-1].Y-points[0].Y))

  switch len(points){
    case 3:
      //triangulo
      var sPerimetro, altura float64
      sPerimetro = (lados[0]+lados[1]+lados[2]) / 2
      altura = (2 / lados[1]) * math.Sqrt(sPerimetro * (sPerimetro - lados[1]) * (sPerimetro - lados[2]) * (sPerimetro - lados[0]))
      area = lados[1] * altura /2
    case 4:
      //dos tringulos
      var sPerimetro1, sPerimetro2, altura1, altura2 float64
      sPerimetro1 = (lados[0]+lados[1]+ math.Hypot(points[2].X - points[0].X, points[2].Y - points[0].Y)) / 2
      altura1 = (2 / lados[1]) * math.Sqrt(sPerimetro1 * (sPerimetro1 - lados[1]) * (sPerimetro1 - math.Hypot(points[2].X - points[0].X, points[2].Y - points[0].Y)) * (sPerimetro1 - lados[0]))
      //segundo triangulo
      sPerimetro2 = (lados[2] + lados[3] + math.Hypot(points[2].X - points[0].X, points[2].Y - points[0].Y)) / 2
      altura2 = (2 / lados[2]) * math.Sqrt(sPerimetro2 * (sPerimetro2 - lados[3]) * (sPerimetro2 - lados[2]) * (sPerimetro2 - math.Hypot(points[2].X - points[0].X, points[2].Y - points[0].Y)))
      //fmt.Println(altura2 * lados[2] / 2)

      area = (lados[1] * altura1 / 2) + (lados[2] * altura2 / 2)
    case 5:
      //tres triangulos
      var sPerimetro1, sPerimetro2, sPerimetro3, altura1, altura2, altura3, ladoI1, ladoI2 float64
      ladoI1 = math.Hypot(points[2].X - points[0].X, points[2].Y - points[0].Y)
      ladoI2 = math.Hypot(points[4].X - points[2].X, points[4].Y - points[2].Y)

      //primer triangulo
      sPerimetro1 = (lados[0]+lados[1]+ ladoI1) / 2
      altura1 = (2 / lados[1]) * math.Sqrt(sPerimetro1 * (sPerimetro1 - lados[1]) * (sPerimetro1 - ladoI1) * (sPerimetro1 - lados[0]))

      //segundo triangulo
      sPerimetro2 = (lados[4] + ladoI1 + ladoI2) / 2
      altura2 = (2 / lados[4]) * math.Sqrt(sPerimetro2 * (sPerimetro2 - lados[4]) * (sPerimetro2 - ladoI2) * (sPerimetro2 - ladoI1))

      //tercer triangulo
      sPerimetro3 = (ladoI2 + lados[2] + lados[3]) / 2
      altura3 = (2 / lados[2]) * math.Sqrt(sPerimetro3 * (sPerimetro3 - lados[3]) * (sPerimetro3 - lados[2]) * (sPerimetro3 - ladoI2))

      area = (lados[1] * altura1 / 2) + (lados[4] * altura2 / 2) + (lados[3] * altura3 / 2)
  }
  return area
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	// Your code goes here
	var perimetro float64
	perimetro = 0

	for i := 0; i < len(points)-1; i++ {
		var distanciaN float64
		distanciaN = math.Hypot(points[i+1].X-points[i].X, points[i+1].Y-points[i].Y)
		perimetro += distanciaN
	}
	perimetro += math.Hypot(points[len(points)-1].X-points[0].X, points[len(points)-1].Y-points[0].Y)

	return perimetro
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	if len (vertices) > 2 {
		response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
	response += fmt.Sprintf(" - Area            : %v\n", area)
	} else{
		response += fmt.Sprintf("Error: [%v] vertices\n", len(vertices))
		response += fmt.Sprintf("ERROR - Your shape is not compliying with the minimum number of vertices")
	}
	

	// Send response to client
	fmt.Fprintf(w, response)
}
