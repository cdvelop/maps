package maps_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/maps"
)

func TestBuildMapForm(t *testing.T) {
	// Estructura de ejemplo
	type EjemploStruct struct {
		Name   string
		Edad   int
		Altura float32
		activo bool // campo en minúscula no accesible para modificar
		Correo string
		Data   map[string]string // campo no valido
	}

	ejemplo := &EjemploStruct{
		Name:   "Juan",
		Edad:   30,
		Altura: 1.75,
		activo: true, // se ingreso el valor por error pero no debe aparecer en el map formulario
		Correo: "juan@example.com",
	}

	// Llamada a la función que se está probando
	result, err := maps.BuildFormString(ejemplo)

	// Verificación de errores
	if err != "" {
		t.Errorf("error TestBuildMapForm\nSe esperaba error nulo, pero se recibió error: %v", err)
		log.Fatalln()
	}

	// Verificación de que el mapa tiene la longitud esperada
	if len(result) != 4 {
		t.Errorf("error TestBuildMapForm\nLongitud del mapa incorrecta. Se esperaba 4, pero se recibió %d", len(result))
		log.Fatalln()
	}

	// Verificación de valores en el mapa
	expected := map[string]string{
		"name":   "Juan",
		"edad":   "30",
		"altura": "1.75",
		"correo": "juan@example.com",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("error TestBuildMapForm.\n-se esperaba:\n%v\n-pero se obtuvo:%v", expected, result)
		log.Fatalln()
	}
}
