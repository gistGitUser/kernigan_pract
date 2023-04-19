package printf

import "fmt"

func printfDescription() {
	a := 10
	fmt.Printf("Число в десятичном виде %d\n", a)
	fmt.Printf("Число в двоичном виде %b\n", a)
	fmt.Printf("Число в 16 виде %x\n", a)
	fmt.Printf("Число в 8 виде %o\n", a)
	var d = 1.0 / 17.0
	fmt.Printf("1/17 \n")
	fmt.Printf("Число в float32 %f\n", d)
	fmt.Printf("Число в float64 виде %g\n", d)
	fmt.Printf("Число с перенесенной степенью %e\n", d)

	fmt.Printf("boolean %t \n", false)
	fmt.Printf("symbol %c\n", 'а')
	fmt.Printf("print string in qoutes %q\n", "example string")
	fmt.Printf("print type  %T\n", d)
	fmt.Printf("percent symb %%\n")

}
