package generator

import (
	"database/sql"
	"fmt"
	"leetcode/model"
)
var languages = []string{"Go", "Python3", "C", "C++", "C#", "Rust", "Java", "JavaScript", "Swift", "Kotlin", "PHP"}

var problems = []model.Problem{
	{
		QuestionNumber:26,
		Title           :"Remove Duplicates from Sorted Array",
		DifficultyLevel : "Easy",
		Description     : `
		Given an integer array nums sorted in non-decreasing order, 
		remove the duplicates in-place such that each unique element appears only once.
		The relative order of the elements should be kept the same. 
		Then return the number of unique elements in nums.
		Consider the number of unique elements of nums to be k, to get accepted, you need 
		to do the following things:

		Change the array nums such that the first k elements of nums contain 
		the unique elements in the order they were present in nums initially. 
		The remaining elements of nums are not important as well as the size of nums.
		Return k.

		Custom Judge:

		The judge will test your solution with the following code:

		int[] nums = [...]; // Input array
		int[] expectedNums = [...]; // The expected answer with correct length

		int k = removeDuplicates(nums); // Calls your implementation

		assert k == expectedNums.length;
		for (int i = 0; i < k; i++) {
			assert nums[i] == expectedNums[i];
		}

		If all assertions pass, then your solution will be accepted.`,

		Examples        :[]string{
			`Input: nums = [1,1,2]
			Output: 2, nums = [1,2,_]
			Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
			It does not matter what you leave beyond the returned k (hence they are underscores).`,
			`Input: nums = [0,0,1,1,1,2,2,3,3,4]
			Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
			Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
			It does not matter what you leave beyond the returned k (hence they are underscores).`,
		},
		Hints           :[]string{
			`In this problem, the key point to focus on is
			the input array being sorted. As far as duplicate
			elements are concerned, what is their positioning 
			in the array when the given array is sorted? Look at 
			the image below for the answer. If we know the position 
			of one of the elements, do we also know the positioning of all the duplicate elements? `,

			`We need to modify the array in-place and 
			the size of the final array would potentially 
			be smaller than the size of the input array. So, we 
			ought to use a two-pointer approach here. One, that would 
			keep track of the current element in the original array and another 
			one for just the unique elements.`,

			`Essentially, once an element is encountered, you simply
			need to bypass its duplicates and move on to the next unique element.`,
		},
		Constraints: []string{
			`1 <= nums.length <= 3 * 104`,
			`-100 <= nums[i] <= 100`,
			`nums is sorted in non-decreasing order.`,
		},
	},
	{
		QuestionNumber:  27,
		Title:           "Remove Element",
		DifficultyLevel: "Easy",
		Description:     "Given an array nums and a value val, remove all instances of that value in-place and return the new length.",
		Examples: []string{
			"Input: nums = [3,2,2,3], val = 3\nOutput: 2 with nums = [2,2]",
			"Input: nums = [0,1,2,2,3,0,4,2], val = 2\nOutput: 5 with nums = [0,1,3,0,4]",
		},
		Hints: []string{
			"The order of elements can be changed. It doesn't matter what you leave beyond the new length.",
			"Two pointers approach can be useful here.",
		},
		Constraints: []string{
			"0 <= nums.length <= 100",
			"0 <= nums[i] <= 50",
			"0 <= val <= 100",
		},
	},
	{
		QuestionNumber:  28,
		Title:           "Implement strStr()",
		DifficultyLevel: "Easy",
		Description:     "Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.",
		Examples: []string{
			"Input: haystack = 'hello', needle = 'll'\nOutput: 2",
			"Input: haystack = 'aaaaa', needle = 'bba'\nOutput: -1",
		},
		Hints: []string{
			"The naive approach is to check each substring of the haystack with the needle.",
			"Use the built-in string library functions in Go to implement this efficiently.",
		},
		Constraints: []string{
			"0 <= haystack.length, needle.length <= 5 * 10^4",
			"haystack and needle consist of only lower-case English characters.",
		},
	},
	{
		QuestionNumber:  29,
		Title:           "Divide Two Integers",
		DifficultyLevel: "Medium",
		Description:     "Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.",
		Examples: []string{
			"Input: dividend = 10, divisor = 3\nOutput: 3",
			"Input: dividend = 7, divisor = -3\nOutput: -2",
		},
		Hints: []string{
			"Think about bit manipulation (shift operations).",
			"Consider edge cases like overflow and handling of negative numbers.",
		},
		Constraints: []string{
			"-2^31 <= dividend, divisor <= 2^31 - 1",
			"divisor != 0",
		},
	},
	{
		QuestionNumber:  30,
		Title:           "Substring with Concatenation of All Words",
		DifficultyLevel: "Hard",
		Description:     "You are given a string s and an array of words words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.",
		Examples: []string{
			"Input: s = 'barfoothefoobarman', words = ['foo','bar']\nOutput: [0,9]",
			"Input: s = 'wordgoodgoodgoodbestword', words = ['word','good','best','word']\nOutput: []",
		},
		Hints: []string{
			"Use a sliding window approach combined with a hash map to efficiently find the substrings.",
			"Consider the length of words and how they can concatenate in any order.",
		},
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"words.length <= 5000",
			"1 <= words[i].length <= 30",
		},
	},
	{
		QuestionNumber:  31,
		Title:           "Next Permutation",
		DifficultyLevel: "Medium",
		Description:     "Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers. If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).",
		Examples: []string{
			"Input: nums = [1,2,3]\nOutput: [1,3,2]",
			"Input: nums = [3,2,1]\nOutput: [1,2,3]",
			"Input: nums = [1,1,5]\nOutput: [1,5,1]",
		},
		Hints: []string{
			"To generate the next permutation, you need to find the first pair of two successive numbers a[i] and a[i-1], from the right, which satisfy a[i] > a[i-1].",
			"Once the successor is found, the next step is to find the smallest number on right side of partion index which is greater than value found in previous step.",
		},
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 100",
		},
	},
	{
		QuestionNumber:  32,
		Title:           "Longest Valid Parentheses",
		DifficultyLevel: "Hard",
		Description:     "Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.",
		Examples: []string{
			"Input: s = '(()'\nOutput: 2",
			"Input: s = ')()())'\nOutput: 4",
		},
		Hints: []string{
			"Use a stack to keep track of the indices of '(' characters.",
			"Scan the string from left to right and from right to left to find the longest valid substring.",
		},
		Constraints: []string{
			"0 <= s.length <= 3 * 10^4",
			"s[i] is '(' or ')'.",
		},
	},
	{
		QuestionNumber:  33,
		Title:           "Search in Rotated Sorted Array",
		DifficultyLevel: "Medium",
		Description:     "There is an integer array nums sorted in ascending order (with distinct values). Prior to being rotated at some unknown pivot, nums was originally a sorted ascending array. You are given a target value to search. If found in the array return its index, otherwise return -1.",
		Examples: []string{
			"Input: nums = [4,5,6,7,0,1,2], target = 0\nOutput: 4",
			"Input: nums = [4,5,6,7,0,1,2], target = 3\nOutput: -1",
		},
		Hints: []string{
			"Consider using binary search for an efficient solution.",
			"Think about the conditions for determining whether to search left or right of the mid-point in the array.",
		},
		Constraints: []string{
			"1 <= nums.length <= 5000",
			"-10^4 <= nums[i], target <= 10^4",
			"All values of nums are unique.",
			"nums is guaranteed to be rotated at some pivot.",
		},
	},
	{
		QuestionNumber:  34,
		Title:           "Find First and Last Position of Element in Sorted Array",
		DifficultyLevel: "Medium",
		Description:     "Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value. If the target is not found in the array, return [-1, -1].",
		Examples: []string{
			"Input: nums = [5,7,7,8,8,10], target = 8\nOutput: [3,4]",
			"Input: nums = [5,7,7,8,8,10], target = 6\nOutput: [-1,-1]",
		},
		Hints: []string{
			"Consider using binary search to find the leftmost and rightmost positions of the target.",
			"Implement two separate binary searches for finding the starting and ending positions.",
		},
		Constraints: []string{
			"0 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
			"nums is a non-decreasing array.",
		},
	},
	{
		QuestionNumber:  35,
		Title:           "Search Insert Position",
		DifficultyLevel: "Easy",
		Description:     "Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.",
		Examples: []string{
			"Input: nums = [1,3,5,6], target = 5\nOutput: 2",
			"Input: nums = [1,3,5,6], target = 2\nOutput: 1",
			"Input: nums = [1,3,5,6], target = 7\nOutput: 4",
			"Input: nums = [1,3,5,6], target = 0\nOutput: 0",
		},
		Hints: []string{
			"Consider using binary search to find the insertion position efficiently.",
			"Handle edge cases such as the target being smaller or larger than any element in the array.",
		},
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"-10^4 <= nums[i] <= 10^4",
			"nums is sorted in ascending order.",
		},
	},
	{
		QuestionNumber:  36,
		Title:           "Valid Sudoku",
		DifficultyLevel: "Medium",
		Description:     "Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:\n\n1. Each row must contain the digits 1-9 without repetition.\n2. Each column must contain the digits 1-9 without repetition.\n3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.",
		Examples: []string{
			"Input: board = [['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: true",
			"Input: board = [['8','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: false",
		},
		Hints: []string{
			"Use sets or arrays to track the presence of numbers in rows, columns, and sub-boxes.",
			"Validate each row, column, and sub-box separately using nested loops.",
		},
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit or '.'.",
			"It is guaranteed that the Sudoku board will be valid.",
		},
	},
	{
		QuestionNumber:  37,
		Title:           "Sudoku Solver",
		DifficultyLevel: "Hard",
		Description:     "Write a program to solve a Sudoku puzzle by filling the empty cells.\n\nA sudoku solution must satisfy all of the following rules:\n\n1. Each of the digits 1-9 must occur exactly once in each row.\n2. Each of the digits 1-9 must occur exactly once in each column.\n3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.\n\nThe '.' character indicates empty cells.",
		Examples: []string{
			"Input: board = [['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: [['5','3','4','6','7','8','9','1','2'],['6','7','2','1','9','5','3','4','8'],['1','9','8','3','4','2','5','6','7'],['8','5','9','7','6','1','4','2','3'],['4','2','6','8','5','3','7','9','1'],['7','1','3','9','2','4','8','5','6'],['9','6','1','5','3','7','2','8','4'],['2','8','7','4','1','9','6','3','5'],['3','4','5','2','8','6','1','7','9']]",
		},
		Hints: []string{
			"Use backtracking to explore all possibilities.",
			"Implement functions to check the validity of placing a number in a specific cell based on row, column, and 3x3 sub-box constraints.",
		},
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit or '.'.",
			"It is guaranteed that the Sudoku puzzle will have a single unique solution.",
		},
	},
	{
		QuestionNumber:  38,
		Title:           "Count and Say",
		DifficultyLevel: "Easy",
		Description:     "The count-and-say sequence is a sequence of digit strings defined by the recursive formula:\n\n1. countAndSay(1) = '1'\n2. countAndSay(n) is the way you would 'say' the digit string from countAndSay(n-1), which is then converted into a different representation.\n\nTo determine how you 'say' a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.\n\nFor example, the saying and conversion for digit string '3322251' is '2 3 3 2 2 1' ('two 3s, two 2s, then one 1').",
		Examples: []string{
			"Input: n = 1\nOutput: '1'\nExplanation: This is the base case.",
			"Input: n = 4\nOutput: '1211'\nExplanation: countAndSay(1) = '1', countAndSay(2) = '11', countAndSay(3) = '21', countAndSay(4) = '1211'.",
		},
		Hints: []string{
			"Use iterative approach to generate the next sequence based on the previous one.",
			"Consider using two pointers to traverse and count consecutive characters in the string.",
		},
		Constraints: []string{
			"1 <= n <= 30",
		},
	},
	{
		QuestionNumber:  39,
		Title:           "Combination Sum",
		DifficultyLevel: "Medium",
		Description:     "Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.\n\nThe same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.",
		Examples: []string{
			"Input: candidates = [2,3,6,7], target = 7\nOutput: [[2,2,3],[7]]\nExplanation:\n2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times, similarly with 3.\n7 is a candidate, and 7 = 7.\nThese are the only two combinations.",
			"Input: candidates = [2,3,5], target = 8\nOutput: [[2,2,2,2],[2,3,3],[3,5]]\nExplanation:\n2, 3, and 5 are candidates, and 2 + 2 + 2 + 2 = 8, 2 + 3 + 3 = 8, and 3 + 5 = 8 are the only three combinations.",
			"Input: candidates = [2], target = 1\nOutput: []\nExplanation:\n2 is the only candidate, and 2 < 1. There are no combinations.",
			"Input: candidates = [1], target = 1\nOutput: [[1]]",
			"Input: candidates = [1], target = 2\nOutput: [[1,1]]",
		},
		Hints: []string{
			"Use backtracking to explore all possible combinations.",
			"Sort the candidates array to handle duplicates easily and improve efficiency.",
		},
		Constraints: []string{
			"1 <= candidates.length <= 30",
			"1 <= candidates[i] <= 200",
			"All elements of candidates are distinct.",
			"1 <= target <= 500",
		},
	},
	{
		QuestionNumber:  40,
		Title:           "Combination Sum II",
		DifficultyLevel: "Medium",
		Description:     "Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.\n\nEach number in candidates may only be used once in the combination.\n\nNote: The solution set must not contain duplicate combinations.",
		Examples: []string{
			"Input: candidates = [10,1,2,7,6,1,5], target = 8\nOutput: [[1,1,6],[1,2,5],[1,7],[2,6]]\nExplanation:\n1 + 1 + 6 = 8\n1 + 2 + 5 = 8\n1 + 7 = 8\n2 + 6 = 8\nThese are the only unique combinations.",
			"Input: candidates = [2,5,2,1,2], target = 5\nOutput: [[1,2,2],[5]]\nExplanation:\n1 + 2 + 2 = 5\n5 = 5\nThese are the only unique combinations.",
		},
		Hints: []string{
			"Use backtracking to explore all possible combinations.",
			"Sort the candidates array to handle duplicates easily and improve efficiency.",
		},
		Constraints: []string{
			"1 <= candidates.length <= 100",
			"1 <= candidates[i] <= 50",
			"1 <= target <= 30",
		},
	},
	{
		QuestionNumber:  41,
		Title:           "First Missing Positive",
		DifficultyLevel: "Hard",
		Description:     "Given an unsorted integer array nums, return the smallest missing positive integer.\n\nYou must implement an algorithm that runs in O(n) time and uses constant extra space.",
		Examples: []string{
			"Input: nums = [1,2,0]\nOutput: 3",
			"Input: nums = [3,4,-1,1]\nOutput: 2",
			"Input: nums = [7,8,9,11,12]\nOutput: 1",
		},
		Hints: []string{
			"Think about how to make use of the array itself to store information.",
			"Consider using a cyclic sort approach.",
		},
		Constraints: []string{
			"1 <= nums.length <= 5 * 10^5",
			"-2 * 10^9 <= nums[i] <= 2 * 10^9",
		},
	},
	{
		QuestionNumber:  42,
		Title:           "Trapping Rain Water",
		DifficultyLevel: "Hard",
		Description:     "Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.\n\nThe height at each index i represents the elevation of the terrain at that point. The width of each bar is 1.\n\nConstraints:\n\nn == height.length\n0 <= n <= 3 * 10^4\n0 <= height[i] <= 10^5",
		Examples: []string{
			"Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]\nOutput: 6\nExplanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.",
			"Input: height = [4,2,0,3,2,5]\nOutput: 9",
		},
		Hints: []string{
			"Try to solve it with two pointers.",
			"Simulate the process of water flowing from the highest point to both ends.",
		},
		Constraints: []string{
			"1 <= n <= 3 * 10^4",
			"0 <= height[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  43,
		Title:           "Multiply Strings",
		DifficultyLevel: "Medium",
		Description:     "Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.\n\nNote: You must not use any built-in BigInteger library or convert the inputs to integer directly.",
		Examples: []string{
			"Input: num1 = '2', num2 = '3'\nOutput: '6'",
			"Input: num1 = '123', num2 = '456'\nOutput: '56088'",
		},
		Hints: []string{
			"Use an array to store intermediate results of multiplication.",
			"Consider how traditional multiplication works with pen and paper.",
		},
		Constraints: []string{
			"1 <= num1.length, num2.length <= 200",
			"num1 and num2 consist of digits only.",
			"Both num1 and num2 do not contain any leading zero, except the number 0 itself.",
		},
	},
	{
		QuestionNumber:  44,
		Title:           "Wildcard Matching",
		DifficultyLevel: "Hard",
		Description:     "Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.\n\n'?' Matches any single character.\n'*' Matches any sequence of characters (including the empty sequence).\nThe matching should cover the entire input string (not partial).\n\nConstraints:\n\n0 <= s.length, p.length <= 2000\ns contains only lowercase English letters.\np contains only lowercase English letters, '?' or '*'.",
		Examples: []string{
			"Input: s = 'aa', p = 'a'\nOutput: false\nExplanation: 'a' does not match the entire string 'aa'.",
			"Input: s = 'adceb', p = '*a*b'\nOutput: true\nExplanation: The first '*' matches the empty sequence, while the second '*' matches the substring 'dce'.",
			"Input: s = 'acdcb', p = 'a*c?b'\nOutput: false\nExplanation: The matching substring is 'acdcb', which clearly does not match 'a*c?b'.",
		},
		Hints: []string{
			"Dynamic programming approach can be used to solve this problem efficiently.",
			"Consider how '?' and '*' can be handled in the matching process.",
		},
		Constraints: []string{
			"0 <= s.length, p.length <= 2000",
			"s contains only lowercase English letters.",
			"p contains only lowercase English letters, '?' or '*'.",
		},
	},
	{
		QuestionNumber:  45,
		Title:           "Jump Game II",
		DifficultyLevel: "Hard",
		Description:     "Given an array of non-negative integers nums, you are initially positioned at the first index of the array.\n\nEach element in the array represents your maximum jump length at that position.\n\nYour goal is to reach the last index in the minimum number of jumps.\n\nYou can assume that you can always reach the last index.",
		Examples: []string{
			"Input: nums = [2,3,1,1,4]\nOutput: 2\nExplanation: The minimum jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.",
			"Input: nums = [2,3,0,1,4]\nOutput: 2\nExplanation: The minimum jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 4 steps to the last index.",
		},
		Hints: []string{
			"Use a greedy approach to track the furthest point you can reach with the current number of jumps.",
			"Consider how to optimize the jumps needed using the maximum reach strategy.",
		},
		Constraints: []string{
			"1 <= nums.length <= 1000",
			"0 <= nums[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  46,
		Title:           "Permutations",
		DifficultyLevel: "Medium",
		Description:     "Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.",
		Examples: []string{
			"Input: nums = [1,2,3]\nOutput: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]",
			"Input: nums = [0,1]\nOutput: [[0,1],[1,0]]",
			"Input: nums = [1]\nOutput: [[1]]",
		},
		Hints: []string{
			"Backtracking is an efficient way to solve this problem.",
			"Think about how to swap elements to generate different permutations.",
		},
		Constraints: []string{
			"1 <= nums.length <= 6",
			"-10 <= nums[i] <= 10",
			"All the integers of nums are unique.",
		},
	},
	{
		QuestionNumber:  47,
		Title:           "Permutations II",
		DifficultyLevel: "Medium",
		Description:     "Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.",
		Examples: []string{
			"Input: nums = [1,1,2]\nOutput: [[1,1,2],[1,2,1],[2,1,1]]",
			"Input: nums = [1,2,3]\nOutput: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]",
		},
		Hints: []string{
			"Use a similar approach to generate permutations as in the previous problem, but handle duplicates carefully.",
			"Consider how to skip duplicates to generate unique permutations.",
		},
		Constraints: []string{
			"1 <= nums.length <= 8",
			"-10 <= nums[i] <= 10",
		},
	},
	{
		QuestionNumber:  48,
		Title:           "Rotate Image",
		DifficultyLevel: "Medium",
		Description:     "You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).\n\nYou have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.",
		Examples: []string{
			"Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]\nOutput: [[7,4,1],[8,5,2],[9,6,3]]",
			"Input: matrix = [[1]]\nOutput: [[1]]",
			"Input: matrix = [[1,2],[3,4]]\nOutput: [[3,1],[4,2]]",
		},
		Hints: []string{
			"To rotate the matrix in-place, consider how elements move during rotation.",
			"Think about how to transpose the matrix and then reverse each row to achieve the rotation.",
		},
		Constraints: []string{
			"matrix.length == n",
			"matrix[i].length == n",
			"1 <= n <= 20",
			"-1000 <= matrix[i][j] <= 1000",
		},
	},
	{
		QuestionNumber:  49,
		Title:           "Group Anagrams",
		DifficultyLevel: "Medium",
		Description:     "Given an array of strings strs, group the anagrams together. You can return the answer in any order.\n\nAn Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.",
		Examples: []string{
			"Input: strs = ['eat','tea','tan','ate','nat','bat']\nOutput: [['bat'],['nat','tan'],['ate','eat','tea']]",
			"Input: strs = ['']\nOutput: [['']]",
			"Input: strs = ['a']\nOutput: [['a']]",
		},
		Hints: []string{
			"Use a hash map to store groups of anagrams.",
			"Consider how to represent and compare anagrams efficiently.",
		},
		Constraints: []string{
			"1 <= strs.length <= 10^4",
			"0 <= strs[i].length <= 100",
			"strs[i] consists of lower-case English letters.",
		},
	},
	{
		QuestionNumber:  50,
		Title:           "Pow(x, n)",
		DifficultyLevel: "Medium",
		Description:     "Implement `pow(x, n)`, which calculates `x` raised to the power `n` (i.e., `x^n`).",
		Examples: []string{
			"Input: x = 2.00000, n = 10\nOutput: 1024.00000",
			"Input: x = 2.10000, n = 3\nOutput: 9.26100",
			"Input: x = 2.00000, n = -2\nOutput: 0.25000\nExplanation: 2^-2 = 1/2^2 = 1/4 = 0.25",
		},
		Hints: []string{
			"Consider how to optimize the calculation using recursion and divide-and-conquer techniques.",
			"Handle both positive and negative values of `n`.",
		},
		Constraints: []string{
			"-100.0 < x < 100.0",
			"-2^31 <= n <= 2^31-1",
			"-10^4 <= x^n <= 10^4",
		},
	},


}


// for creating all mock data
func GenerateAllMockData(db *sql.DB) {
	fmt.Println(problems)
}
