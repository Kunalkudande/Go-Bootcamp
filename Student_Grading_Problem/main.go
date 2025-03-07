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

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	var students []student
	file, err := os.Open(filePath);
	if err != nil{
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()

	if err != nil{
		fmt.Println("Error: ", err)
	}
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

func createStudent(value[] string) (student, error){
	var st student
	st.firstName = value[0];
	st.lastName = value[1];
	st.university = value[2];

	test1, err1 := strconv.Atoi(value[3])
	test2, err2 := strconv.Atoi(value[4])
	test3, err3 := strconv.Atoi(value[5])
	test4, err4 := strconv.Atoi(value[6])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil{
		return student{}, errors.New("Invalid")
	}

	st.test1Score = test1
	st.test2Score = test2
	st.test3Score = test3
	st.test4Score = test4

	return st, nil
}

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

		graded = append(graded, studentStat{
			student: st,
			finalScore: avg,
			grade: grade,
		})
	}
	return graded
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := gradedStudents[0] 

	for _, s := range gradedStudents{
		if s.finalScore > topper.finalScore{
			topper = s
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperMap := make(map[string]studentStat)
	for _, s := range gs {
		currentTopper, exists := topperMap[s.university]
		if !exists || s.finalScore > currentTopper.finalScore {
			topperMap[s.university] = s
		}
	}
	return topperMap
}
