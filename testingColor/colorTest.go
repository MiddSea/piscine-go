package main

import(

// "fmt"
"github.com/fatih/color"

)

func main() {
// Print with default helper functions
color.Cyan("Prints text in cyan.")

// A newline will be appended automatically
color.Blue("Prints %s in blue.", "text")

// These are using the default foreground colors
color.Red("We have red")
color.Magenta("And many others ..")


color.Green (`Package color is an ANSI color package to output colorized or SGR defined output to the standard output. The API can be used in several way, pick one that suits you.

Use simple and default helper functions with predefined foreground colors:

color.Cyan("Prints text in cyan.")`)

// a newline will be appended automatically
color.Blue("Prints %s in blue.", "text")

// More default foreground colors..
color.Red("We have red")
color.Yellow("Yellow color too!")
color.Magenta("And many others ..")

// Hi-intensity colors
color.HiGreen("Bright green color.")
color.HiBlack("Bright black means gray..")
color.HiWhite("Shiny white color!")
color.HiGreen(`However, there are times when custom color mixes are 

required. Below are some examples to create custom color 
objects and use the print functions of each separate color object.`)

// Create a new color object
c := color.New(color.FgCyan).Add(color.Underline)
c.Println("Prints cyan text with an underline.")

// Or just add them to New()
d := color.New(color.FgCyan, color.Bold)
d.Printf("This prints bold cyan %s\n", "too!.")

// Mix up foreground and background colors, create new mixes!
red := color.New(color.FgRed)

boldRed := red.Add(color.Bold)
boldRed.Println("This will print text in bold red.")

whiteBackground := red.Add(color.BgWhite)
whiteBackground.Println("Red text with White background.")

// Use your own io.Writer output
color.New(color.FgBlue).Fprintln(myWriter, "blue color!")

blue := color.New(color.FgBlue)
blue.Fprint(myWriter, "This will print text in blue.")


}

