package main

import (
	"fmt"

	"github.com/clbanning/mxj"
)

// Sino Centre, Mong Kok. Instax Mini.
func sinoCentre() {
	fmt.Printf("\n%s\n\n", "-- sinoCentre() -- type assertions --")

	menu := []byte(menuXML) // http client.Do() returns []byte
	m, err := mxj.NewMapXml(menu)
	catch(err)

	// bm := m["breakfast_menu"]
	// food := bm[0] // invalid operation: bm[0] (type interface {} does not support indexing)

	bm := m["breakfast_menu"].(map[string]interface{})
	food := bm["food"].([]interface{})

	for _, it := range food {
		item := it.(map[string]interface{})

		name := item["name"].(string)
		price := item["price"].(string)
		description := item["description"].(string)
		calories := atoi(item["calories"].(string))

		fmt.Printf("name: %s, price: %s, description: %s, calories: %d\n\n", name, price, description, calories)
	}
}

// https://www.w3schools.com/xml/xml_examples.asp
const menuXML = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?><breakfast_menu>
    <food>
        <name>Belgian Waffles</name>
        <price>$5.95</price>
        <description>Two of our famous Belgian Waffles with plenty of real maple syrup</description>
        <calories>650</calories>
    </food>
    <food>
        <name>Strawberry Belgian Waffles</name>
        <price>$7.95</price>
        <description>Light Belgian waffles covered with strawberries and whipped cream</description>
        <calories>900</calories>
    </food>
    <food>
        <name>Berry-Berry Belgian Waffles</name>
        <price>$8.95</price>
        <description>Light Belgian waffles covered with an assortment of fresh berries and whipped cream</description>
        <calories>900</calories>
    </food>
    <food>
        <name>French Toast</name>
        <price>$4.50</price>
        <description>Thick slices made from our homemade sourdough bread</description>
        <calories>600</calories>
    </food>
    <food>
        <name>Homestyle Breakfast</name>
        <price>$6.95</price>
        <description>Two eggs, bacon or sausage, toast, and our ever-popular hash browns</description>
        <calories>950</calories>
    </food>
</breakfast_menu>`
