## go-easycsv

This package is an easy to use wrapper around reading from, and writing to CSV files.

### Examples

#### Reading From a CSV File

Let's assume we have a CSV file named "dogs.csv" with the following content:

```
German Shepherd,Male,3,Black and Tan
Shiba-Inu,Female,13,Black and Tan
Shepherd/Husky,Male,11,Tan
```

To read each row, we use the `Open()` function. This returns a `[][]string` of the contents you can easily iterate over.

```Go
package main

import (
    "fmt"
    "github.com/scottdware/go-easycsv"
)

func main() {
    dogs, err := easycsv.Open("dogs.csv")
    if err != nil {
        fmt.Println(err)
    }

    for _, dog := range dogs {
        // Find out how many columns we have in each row
        cols := len(dog)

        // Print the number of columns for each row, as well as the entire row's contents
        fmt.Printf("Columns: %d, Data: %+v\n", cols, dog)
    }
}
```

This will output the following:

```
Columns: 4, Data: [German Shepherd Male 3 Black and Tan]
Columns: 4, Data: [Shiba-Inu Female 13 Black and Tan]
Columns: 4, Data: [Shepherd/Husky Male 11 Tan]
```

Remember that column numbering starts at 0, so if you want to assign each column to a variable name, you can do so as follows:

```Go
    for _, dog := range dogs {
        // Assign variables for each column
        breed := dog[0]
        sex := dog[1]
        age := dog[2]
        color := dog[3]

        // Print the number of columns for each row, as well as the entire row's contents
        fmt.Printf("Breed: %s, Sex: %s, Age: %s, Color: %s\n", breed, sex, age, color)
    }
```

The above will output the following:

```
Breed: German Shepherd, Sex: Male, Age: 3, Color: Black and Tan
Breed: Shiba-Inu, Sex: Female, Age: 13, Color: Black and Tan
Breed: Shepherd/Husky, Sex: Male, Age: 11, Color: Tan
```

#### Creating and Writing to CSV Files

To create a new CSV file, you can use the `NewCSV()` function. This only takes one parameter, which is the name/path to the CSV file. You then
can use the `Write()` function to create content. Once you are finished writing to the file (buffer), you must call the `End()` function to actually write the contents
to the file and close it from further writing.

```Go
package main

import (
	"fmt"

	"github.com/scottdware/go-easycsv"
)

func main() {
	dogsCSV, err := easycsv.NewCSV("dogs.csv")
	if err != nil {
		fmt.Println(err)
	}

	dogsCSV.Write("German Shepherd,Male,3,Black and Tan\n")
	dogsCSV.Write("Shiba-Inu,Female,13,Black and Tan\n")
    dogsCSV.Write("Shepherd/Husky,Male,11,Tan\n")
    dogsCSV.End()
}
```