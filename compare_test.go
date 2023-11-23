package maps_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/maps"
)

func TestAreSliceMapsIdentical(t *testing.T) {
	// Caso 1: Slices de mapas idénticos
	slice1 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice2 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	if !maps.AreSliceMapsIdentical(slice1, slice2) {
		t.Errorf("Los slices de mapas deberían ser idénticos, pero no lo son.")
	}

	// Caso 2: Slices de mapas con diferentes valores
	slice3 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice4 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "5"},
	}
	if maps.AreSliceMapsIdentical(slice3, slice4) {
		t.Errorf("Los slices de mapas deberían tener diferentes valores, pero son idénticos.")
	}

	// Caso 3: Slices de mapas con diferentes claves
	slice5 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice6 := []map[string]string{
		{"a": "1", "b": "2"},
		{"e": "3", "d": "4"},
	}
	if maps.AreSliceMapsIdentical(slice5, slice6) {
		t.Errorf("Los slices de mapas deberían tener diferentes claves, pero son idénticos.")
	}

	// Caso 4: Slices de mapas idénticos con diferente orden
	slice7 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice8 := []map[string]string{
		{"c": "3", "d": "4"},
		{"a": "1", "b": "2"},
	}
	if !maps.AreSliceMapsIdentical(slice7, slice8) {
		fmt.Println("Caso 4 Los slices de mapas deberían ser idénticos, pero no lo son.")
		log.Fatal()
	}

	// Caso 5: Slices de mapas idénticos con diferente orden en valores
	slice9 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice10 := []map[string]string{
		{"d": "4", "c": "3"},
		{"b": "2", "a": "1"},
	}
	if !maps.AreSliceMapsIdentical(slice9, slice10) {
		fmt.Println("Caso 5 Los slices de mapas deberían ser idénticos, pero no lo son.")
		log.Fatal()
	}

}
