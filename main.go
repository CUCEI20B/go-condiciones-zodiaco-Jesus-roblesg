package main

import "fmt"

func main()  {
	var day int
	var month int
	fmt.Scan(&day)
	fmt.Scan(&month)

	switch month {
	case 1:
		switch {
		case day > 0 && day < 21:
			fmt.Print("capricornio")
		case day >= 21:
			fmt.Print("acuario")
		}
	case 2:
		switch {
		case day > 0 && day < 19:
			fmt.Print("acuario")
		case day >= 19:
			fmt.Print("picis")
		}
	case 3:
		switch {
		case day > 0 && day < 21:
			fmt.Print("picis")
		case day >= 21:
			fmt.Print("aries")
		}
	case 4:
		switch {
		case day > 0 && day < 21:
			fmt.Print("aries")
		case day >= 21:
			fmt.Print("tauro")
		}
	case 5:
		switch {
		case day > 0 && day < 21:
			fmt.Print("tauro")
		case day >= 21:
			fmt.Print("geminis")
		}
	case 6:
		switch {
		case day > 0 && day < 22:
			fmt.Print("geminis")
		case day >= 22:
			fmt.Print("cancer")
		}
	case 7:
		switch {
		case day > 0 && day < 23:
			fmt.Print("cancer")
		case day >= 23:
			fmt.Print("leo")
		}
	case 8:
		switch {
		case day > 0 && day < 23:
			fmt.Print("leo")
		case day >= 23:
			fmt.Print("virgo")
		}
	case 9:
		switch {
		case day > 0 && day < 23:
			fmt.Print("virgo")
		case day >= 23:
			fmt.Print("libra")
		}
	case 10:
		switch {
		case day > 0 && day < 23:
			fmt.Print("libra")
		case day >= 23:
			fmt.Print("escorpio")
		}
	case 11:
		switch {
		case day > 0 && day < 23:
			fmt.Print("escorpio")
		case day >= 23:
			fmt.Print("sagitario")
		}
	case 12:
		switch {
		case day > 0 && day < 22:
			fmt.Print("sagitario")
		case day >= 22:
			fmt.Print("capricornio")
		}
	}
}