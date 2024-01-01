package file

import "fmt"
import "time"

func Debug(){

	start := time.Now()
	// only runs at end of main func
	defer func(){
		fmt.Println("Time taken: ", time.Since(start), " Seconds")
	}()

	fmt.Println("\nRunning tests...")
	time.Sleep(2 * time.Second)

	fmt.Println("All tests returned successful!")

}