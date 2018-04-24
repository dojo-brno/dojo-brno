package main

import "testing"
import "fmt"

var itemsDay0 = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func Test_day1(t *testing.T) {
	var numberOfDays int

	fmt.Println("Starting status of items[0]:", items[2], "\n")
	//fmt.Println(items)
	for numberOfDays = 1; numberOfDays < 20; numberOfDays++ {
		var offSet int
		if items[2].sellIn > 0 {
			offSet = 1
		} else {
			offSet = 2
		}
		var expectedQuality = items[2].quality - offSet
		if expectedQuality < 0 {
			expectedQuality = 0
		}
		GildedRose()

		// test of the item called: "+5 Dexterity Vest"
		if itemsDay0[2].sellIn-numberOfDays != items[2].sellIn {
			t.Error()
			fmt.Println("********** ", items[2])
			fmt.Println("Error in sellIn value of ", items[2].name, " after ", numberOfDays, " days")
		}

		if expectedQuality != items[2].quality {
			t.Error()
			fmt.Println("********** ", items[2])
			fmt.Println("Error in quality value of ", items[2].name, " after ", numberOfDays, " days")
		}
	}

}
