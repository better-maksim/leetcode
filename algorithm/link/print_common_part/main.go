package main

import (
	"fmt"
	"gitbuh.com/better-maksim/leetcode/pkg/link"
	"strconv"
)

func main() {
	linkOne := link.NewLink(1)
	linkOne.Add(2)
	linkOne.Add(3)
	linkOne.Add(4)

	link2 := link.NewLink(2)
	link2.Add(3)
	printCommonPart(linkOne, link2)

	link3 := linkOne.Reverse()
	link3                                                                            .Print()

}

func printCommonPart(head1 *link.LinkNode, head2 *link.LinkNode) {
	fmt.Println("Common Part:")
	for {

		if head1 == nil || head2 == nil {
			break
		}

		if head1.Value < head2.Value {
			head1 = head1.Next
		} else if head1.Value > head2.Value {
			head2 = head2.Next
		} else {
			fmt.Print(strconv.Itoa(head1.Value) + "")
			head1 = head1.Next
			head2 = head2.Next
		}

		fmt.Println()
	}
}
