package asciimoul

import "strings"

const Head = `
   ++
   ++++
    ++++
  ++++++++++
 +++       |
 ++         |
 +  -==   ==|
(   <*>   <*>
 |           |
 |         __|
 |      +++
  \      =+
   \      +
    \++++++
      ++++
`
const headLongestLine = 14

func HeadWithPadding() string {
	lines := []string{}
	for _, line := range strings.Split(Head, "\n") {
		lines = append(lines, line+strings.Repeat(" ", headLongestLine-len(line)))
	}
	return strings.Join(lines, "\n")
}

func ReverseHead() string {
	lines := []string{}
	for _, line := range strings.Split(HeadWithPadding(), "\n") {
		lines = append(lines, strings.TrimRight(reverse(line), " "))
	}
	return strings.Join(lines, "\n")
}

const Body = `
      ++
      ++++
       ++++
     ++++++++++
    +++       |
    ++         |
    +  -==   ==|
   (   <*>   <*>
    |           |
    |         __|
    |      +++
     \      =+
      \      +
      |\++++++
      |  ++++      ||//
  ____|   |____   _||/__
 /     ---     \  \|  |||
/  _ _  _     / \   \ /
| / / //_//_//  |   | |
`

const Text = `
                   __
  __ _  ___  __ __/ /
 /  ' \/ _ \/ // / /
/_/_/_/\___/\_,_/_/
`
