package maps

import (
	"fmt"
	"log"
	"testing"
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
	if !AreSliceMapsIdentical(slice1, slice2) {
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
	if AreSliceMapsIdentical(slice3, slice4) {
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
	if AreSliceMapsIdentical(slice5, slice6) {
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
	if !AreSliceMapsIdentical(slice7, slice8) {
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
	if !AreSliceMapsIdentical(slice9, slice10) {
		fmt.Println("Caso 5 Los slices de mapas deberían ser idénticos, pero no lo son.")
		log.Fatal()
	}

}

func TestCompareMaps(t *testing.T) {
	tests := []struct {
		name     string
		mapA     map[string]string
		mapB     map[string]string
		expected string
	}{
		{
			name:     "EqualMaps",
			mapA:     map[string]string{"key1": "value1", "key2": "value2"},
			mapB:     map[string]string{"key1": "value1", "key2": "value2"},
			expected: "",
		},
		{
			name:     "KeysNotFound",
			mapA:     map[string]string{"key1": "value1", "key2": "value2"},
			mapB:     map[string]string{"key1": "value1", "key3": "value3"},
			expected: mapKeyNotFound("key2", "B") + ", " + mapKeyNotFound("key3", "A"),
		},
		{
			name:     "DifferentValues",
			mapA:     map[string]string{"key1": "value1", "key2": "value2"},
			mapB:     map[string]string{"key1": "value1", "key2": "value_different"},
			expected: mapDifValue("key2", "value2", "value_different"),
		},
		{
			name:     "DifferentSize",
			mapA:     map[string]string{"key1": "value1"},
			mapB:     map[string]string{"key1": "value1", "key2": "value_different"},
			expected: mapKeyNotFound("key2", "A"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CompareMaps(test.mapA, test.mapB)
			if result != test.expected {
				t.Errorf("\n- Expected:\n%q\n- But got:\n%q", test.expected, result)
			}
		})
	}
}
