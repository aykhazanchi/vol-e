package lib

import "fmt"

func PrintAsciiArt() {
	asciiArt := `
    .-"""""""-.
  .'           '.
/  .-.     .-.  \     	,-----------.
|  /   \\   /   \\ |   ,| Hej!      |
| |\\_.  | |    /| |  , | Jag heter |
|\\|  | /| | |\\ |/|    | Vol-E!    |
| \'---' | | \'---'|    '-----------'
|       | |       |
|       | |       |
 \\     /   \\    /
  \'---'     \'---'
`
	fmt.Println(asciiArt)
}
