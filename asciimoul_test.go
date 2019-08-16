package asciimoul

import (
	"fmt"
)

func ExampleHead() {
	fmt.Println(Head)
	// Output:
	//    ++
	//    ++++
	//     ++++
	//   ++++++++++
	//  +++       |
	//  ++         |
	//  +  -==   ==|
	// (   <*>   <*>
	//  |           |
	//  |         __|
	//  |      +++
	//   \      =+
	//    \      +
	//     \++++++
	//       ++++
}

func ExampleReverseHead() {
	fmt.Println(ReverseHead())
	// Output:
	//          ++
	//        ++++
	//       ++++
	//   ++++++++++
	//   |       +++
	//  |         ++
	//  |==   ==-  +
	//  <*>   <*>   )
	// |           |
	// |__         |
	//    +++      |
	//    +=      /
	//    +      /
	//    ++++++/
	//     ++++
}
