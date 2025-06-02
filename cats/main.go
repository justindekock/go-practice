package main

import (
	"fmt"
	"time"
)

type cat struct {
        name string
        age uint8
        birthday time.Time
        sex string
        color string
        is_nice bool
    }

func (c *cat) increaseAge() {
        if time.Now().Month() != c.birthday.Month() {
            fmt.Printf("It is NOT your birthday %s. Go away\n\n", c.name)
            return
        }
        c.age = c.age + 1
        fmt.Printf("Happy birthday %s!!! You are now %d\n\n", c.name, c.age)
    }

func main() {
    my_cats := [3]cat {
        {name:"Rosie", age:5, birthday: time.Date(2020, time.June, 2, 2, 0, 0, 0, time.UTC), sex:"F", color:"Tabby", is_nice:true},
        {name:"Milk", age:3, birthday: time.Date(2020, time.April, 15, 0, 0, 0, 0, time.UTC), sex:"M", color:"White & Black", is_nice:false},
        {name:"Guinness", age:2, birthday: time.Date(2020, time.October, 31, 0, 0, 0, 0, time.UTC), sex:"M", color:"Black", is_nice:true},
    }

    for i:=range len(my_cats) {
        var pn string
        var nice string
        var bty string
        if my_cats[i].sex == "F" {
            pn = "She"
        } else {
            pn = "He"
        }

        if my_cats[i].is_nice {
            nice = "VERY"
            bty = "beautiful"
        } else {
            nice = "NOT"
            bty = "hideous"
        }
        fmt.Printf("%s is %d years old. %s is a %s %s cat who is %s nice.\n", my_cats[i].name, my_cats[i].age, pn, bty, my_cats[i].color, nice)
        fmt.Printf("Is is %s's birthday??\n", my_cats[i].name)
        my_cats[i].increaseAge()
    }
}
