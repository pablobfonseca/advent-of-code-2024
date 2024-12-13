## Day 7

Each line represents a single equation. The test value appears before the colon on each line; it is your job to determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (\*).

Only three of the above equations can be made true by inserting operators:

- 190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing _ would give the test value (10 _ 19 = 190).
- 3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the right side to match the test value: 81 + 40 _ 27 and 81 _ 40 + 27 both equal 3267 (when evaluated left-to-right)!
- 292: 11 6 16 20 can be solved in exactly one way: 11 + 6 \* 16 + 20.
  The engineers just need the total calibration result, which is the sum of the test values from just the equations that could possibly be true. In the above example, the sum of the test values for the three equations listed above is 3749.

Determine which equations could possibly be true. What is their total calibration result?
