// We will be writing an algorithm to count the number of islands present
// in a map.

// Our map is represented by a 2D array of integers in {0, 1} where
//   0 corresponds to water
//   1 corresponds to land

// An island is what you might expect: a body of land surrounded
// completely by water (or an edge of the array).

// Two islands that touch only at the corner are considered two
// islands and are not merged.

// The challenge here is to code an algorithm that counts the number
// of islands in input array.

// *****

// Example input (see islands.png):

// 00000000
// 01001000
// 01000100
// 00000000
// 00000111
// 00000010

// Expected output: 4 (This ^ map contains 4 islands!)

// as a 2D array:
//   var arr = [6][8]int{
//     {0,0,0,0,0,0,0,0},
//     {0,1,0,0,1,0,0,0},
//     {0,1,0,0,0,1,0,0},
//     {0,0,0,0,0,0,0,0},
//     {0,0,0,0,0,1,1,1},
//     {0,0,0,0,0,0,1,0},
//

package main
