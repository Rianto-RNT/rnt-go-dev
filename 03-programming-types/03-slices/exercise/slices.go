//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Bagian string

func tampilkanGaris(garis []Bagian) {
	for i := 0; i < len(garis); i++ {
		bagian := garis[i]
		fmt.Println(bagian)
	}
}

func main() {
	gabunganBaris := make([]Bagian, 3)

	gabunganBaris[0] = "Pipa"
	gabunganBaris[1] = "Baut"
	gabunganBaris[2] = "Besi Plat"

	fmt.Println("3 Material barang:")
	tampilkanGaris(gabunganBaris)

	gabunganBaris = append(gabunganBaris, "Alat Pembersih", "Ban/roda")
	fmt.Println("\nTambah dua barang/sukucadang lagi:")
	tampilkanGaris(gabunganBaris)

	gabunganBaris = gabunganBaris[3:]
	fmt.Println("\nSliced atau barang yang di tambahkan adalah:")
	tampilkanGaris(gabunganBaris)

}
