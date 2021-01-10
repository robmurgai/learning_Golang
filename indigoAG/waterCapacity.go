// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

// Example:
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

// looking at the heights, any change in elevation becomes an option to collect water.
// elevation on both sides are needed, so the embankments starts at non zero height
//  And every time I see a non zero value - I can add

// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// 0 ignore.
// 1 add the height.
// 0 and I can compare to the last embankment, if less add a volume - 1, which is the last embankment.
// 2, increase, but the height of water is 1,  min (of the two embankments)

package main
