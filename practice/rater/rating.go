package main
import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	// Front end
	// take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name: ")
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Burger shack between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// Back end
	fmt.Printf("Hello %v, \n Thanks for rating our burger shack with a %v star rating. \n\n Your rating was recorded at %v \n\n", strings.TrimSpace(name), mynumRating, time.Now().Format(time.Stamp))

	if mynumRating == 5 {
		fmt.Println("Yaay! Thanks a lot")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("Please come back to see our future improvements")
	} else if mynumRating < 3 {
		fmt.Println("Sorry that you were not happy with our service. We'll be sure to improve")
	}
}
