package main 
//making a frequency dictionary

import (
	"fmt"
	"log"
	"regexp"
	"io/ioutil"
	"strings"
	"sort"
)

var filtered []string
var wordSlice []string
var uniqSlice []string
var count = 1

type finalStruct struct {
	word string
	frequency int
}
var finalSlice []finalStruct

func filter() {
	content, err := ioutil.ReadFile("dicGo1.log")
	if err != nil {
		log.Fatalf(err)
	}

	re := regexp.MustCompile("[a-zA-Z]{3,24}")

	filtered = re.FindAllString(string(content), -1)
}


func lowerAndSorting() {
	for _, word := range filtered {
		wordSlice = append(wordSlice, strings.ToLower(word))
	}

	sort.Strings(wordSlice)
}

func unification() {
	uniqSlice = append(uniqSlice, wordSlice[0]) //making some length
	for _, word := range wordSlice {
		if word != uniqSlice[len(uniqSlice) - 1] {
			uniqSlice = append(uniqSlice, word)
		}
	}
}

func counting() {
	//assign first elem of wordSlice to word value in finalStruct and also assign 0 to frequency value
	finalSlice = []finalStruct{{wordSlice[0], 0},} //присваиваем фс первый элемент со знач счетчика 0
	for i, word := range wordSlice { //берем слово из вСл
		for _, word2 := range finalSlice { //берем последн слово из финСл
			//if crrent word of wordSlice == word value of last elem of finalSlice
			if word == word2.word { //елси последн из фин == текущ из общСлайса
				count++
			//else if current word from wordSlice != last word from finalSlice && previous word from wordSlice == last word from finalSlice
			} else if word != finalSlice[len(finalSlice) - 1].word && wordSlice[i - 1] == finalSlice[len(finalSlice) - 1].word { //иначе если новое слово из общСл != последн слову из финСл && предыдущ слово из общСл == последнее слово из  финСл
				//assign count value to frequency of the last elem of finalSlice 
				finalSlice[len(finalSlice) - 1].frequency = count //присвоить frequency предыдущ
				//append new elem to finalSlice
				finalSlice = append(finalSlice, finalStruct{word: word}) //добавить  новый эл-т к финСл
				count = 1
			}
		}
	}
}

type ByFreq []finalStruct
func (a ByFreq) Len() int {return len(a)}
func (a ByFreq) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByFreq) Less(i, j int) bool {return a[i].frequency < a[j].frequency}


func main() {
	filter()

	lowerAndSorting()

	unification()

	counting()

	sort.Slice(finalSlice, func(i, j int) bool {
		return finalSlice[i].frequency > finalSlice[j].frequency
	})
 
	for _, word := range finalSlice {
		fmt.Println(word.word, word.frequency)
	}
}

//I use command << to write result in file
