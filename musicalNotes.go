//Given a list of frequencies (in Hz), write a function to determine the closest musical note for each frequency based on the A440 pitch standard. Extra credit: indicate if the note is flat or sharp! 

//Example: 

//> getNoteNames([440, 490, 524, 293.66])
//> ["This is a A", "This is a B, but it's flat", "This is a C, but it's sharp", "This is a D"]
// PS I couldn't figure out the logic for flat or sharp
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := []float64{440, 490, 524, 293.66}
	getNoteNames(input)
}

func makeNoteMap() map[int]string{
	var noteMap map[int]string
	noteMap = make(map[int]string)
	noteMap[0] = "This is C"
	noteMap[1] = "This is C, but it's sharp"
	noteMap[2] = "This is D"
	noteMap[3] = "This is E, but it's flat"
	noteMap[4] = "This is E"
	noteMap[5] = "This is F"
	noteMap[6] = "This is F, but it's sharp"
	noteMap[7] = "This is G"
	noteMap[8] = "This is A, but it's flat"
	noteMap[9] = "This is A"
	noteMap[10] = "This is B, but it's flat"
	noteMap[11] = "This is B"
	return noteMap
}

func makeGraph() [12][10]float64 {
	var graph [12][10]float64;
	graph[0][0] = 8.17;
	graph[1][0] = 8.66;
	graph[2][0] = 9.177024;
	graph[3][0] = 9.72;
	graph[4][0] = 10.30;
	graph[5][0] = 10.91;
	graph[6][0] = 11.56;
	graph[7][0] = 12.24;
	graph[8][0] = 12.97;
	graph[9][0] = 13.75;
	graph[10][0] = 14.56;
	graph[11][0] = 15.43;

	for i:=0 ; i<len(graph); i++ {
		for j:=1; j<len(graph[0]);j++ {
			prev := graph[i][j-1]
			graph[i][j] = prev*2	
		}
	}
	fmt.Println(graph[2])
	return graph
}

func computeDifferenceAndFindClosestNote(graph [12][10]float64, note float64) int {
	var key int
	var c float64 = 100000.0
	for i:=0 ; i< len(graph) ; i++ {
		for j:=0 ; j<len(graph[0]); j++ {
			diff := roundoff(graph[i][j]) - note
			if diff == 0 {
				// found exact note
				c = diff
				key= i
				return key
			}
			diff=math.Abs(diff)
			if(diff < c) {
				c = diff
				key=i
			}
		}
	}
	fmt.Println(c)
	return key
}

func getNoteNames(notes []float64) {
	graph := makeGraph()
	noteMap := makeNoteMap()
	for i:=0 ; i<len(notes);i++ {
		fmt.Println(noteMap[computeDifferenceAndFindClosestNote(graph, notes[i])])
	}
}

func roundoff(a float64) float64{
	sol, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",a),64)
	return sol
}
