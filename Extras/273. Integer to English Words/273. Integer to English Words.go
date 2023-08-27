package main

import "strings"

func main() {
	
}



func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	billion := num / 1000000000
	num %= 1000000000

	million := num / 1000000
	num %= 1000000

	thousand := num / 1000
	num %= 1000
	s := []string{} 
	if billion > 0 {
		s=append(s, numberToString(billion)...) 
		s=append(s, "Billion")
	}
	if million > 0 {
		s=append(s, numberToString(million)...) 
		s=append(s, "Million")	
	}
	if thousand > 0 {
		s=append(s, numberToString(thousand)...) 
		s=append(s, "Thousand")
	}
	if num > 0 {
		s=append(s, numberToString(num)...) 
	}
	return strings.Join(s, " ")
}


func numberToString(num int) []string {
	digitMap := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	tenMap := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy","Eighty", "Ninety"}
	teenMap := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen","Seventeen", "Eighteen", "Nineteen"}
	s := []string{}
	hundred := num/100 
	num %=100
	tens := num/10
	num %=10
	if hundred > 0 {
		s = append(s, digitMap[hundred], "Hundred")
	}
	if tens > 0 {
		if tens ==1 {
			s =append(s, teenMap[num])
		}else {
			s =append(s, tenMap[tens])
		}
	}
	if tens != 1 && num>0 {
		s =append(s, digitMap[num])
	}
	return s
}