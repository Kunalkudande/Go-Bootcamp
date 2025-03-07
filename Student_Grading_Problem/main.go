package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

// student struct to hold student data
type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

// studentStat struct extends student with final score and grade
type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// parseCSV reads the CSV file and returns a list of students
func parseCSV(filePath string) []student {
	var students []student
	// Open the CSV file
	file, err := os.Open(filePath);   
	if err != nil{
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	// Read all records from the CSV
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()

	if err != nil{
		fmt.Println("Error: ", err)
	}

	// Loop through the records (skip header)
	for i, n := range record{
		if i > 0{
			st, err := createStudent(n);
			if err != nil {
				log.Fatal(err)
			}
			students = append(students, st)
		}
	}
	return students
}

// createStudent creates a student from a CSV row (string array)
func createStudent(value[] string) (student, error){
	var st student
	st.firstName = value[0];
	st.lastName = value[1];
	st.university = value[2];

	// Convert string scores to integers
	test1, err1 := strconv.Atoi(value[3])
	test2, err2 := strconv.Atoi(value[4])
	test3, err3 := strconv.Atoi(value[5])
	test4, err4 := strconv.Atoi(value[6])

	// Return error if any score is invalid
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil{
		return student{}, errors.New("Invalid")
	}

	st.test1Score = test1
	st.test2Score = test2
	st.test3Score = test3
	st.test4Score = test4

	return st, nil
}

// calculateGrade calculates the average score and assigns grades
func calculateGrade(students []student) []studentStat {
	var graded []studentStat
	for _, st := range students{
		avg := float32(st.test1Score + st.test2Score + st.test3Score + st.test4Score) / 4

		var grade Grade
		if avg >= 70 {
			grade = A
		} else if avg < 70 && avg >= 50 {
			grade = B
		} else if avg < 50 && avg >= 35 {
			grade = C
		} else {
			grade = F
		}

		// Add studentStat to graded list
		graded = append(graded, studentStat{
			student: st,
			finalScore: avg,
			grade: grade,
		})
	}
	return graded
}


// findOverallTopper finds the student with the highest final score
func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := gradedStudents[0]  // Assume first student is the topper

	for _, s := range gradedStudents{
		if s.finalScore > topper.finalScore{
			topper = s
		}
	}
	return topper
}

// finds the top student in each university
func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperMap := make(map[string]studentStat)
	for _, s := range gs {
		currentTopper, exists := topperMap[s.university]
		// If no topper exists yet or current student has a higher score
		if !exists || s.finalScore > currentTopper.finalScore {
			topperMap[s.university] = s
		}
	}
	return topperMap
}
