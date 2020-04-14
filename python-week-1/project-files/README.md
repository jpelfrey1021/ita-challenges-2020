# Getting the Grade
## Context
This problem is meant to challenge your knowledge and to think about the way you can use Python to get an answer. There is more than one way to get to the answer, so do not be afraid to experiment.

## Task & Objectives
For this weekly problem you will be reading in grades from a file, putting the data in a data structure, and then use that data structure to get the final grade of the student.

### Tip
For the purpose of this challenge you can assume that the file will always be in the same format. The only thing that could be changed is the grade number.

## To-do
The function `file_to_string(file_name)` should open up the file passed in as the parameter and return the contents of the file. The `file_to_String` function should do the following:
* Open the file passed in as a parameter
* Read the contents of the file into a string
* Close the open file
* Return the string contents of the file

In the function `string_to_list(string)`, you have the string representation of a file. You will need to create and a list of tuples that holds the graded name and the grade as an integer.

⋅⋅⋅ Ex. The list you return should have this structure:
```py
[("Quiz1", 95), ("Quiz2", 90), ("Quiz3", 85), ...]
```

The `string_to_list(string)` function should do the following:
* Split up the string to their individual pieces
* Create a tuple with the (grade name, grade) for each of the grades
* Add the tuples to a list and return it
⋅⋅* **HINT:** You can use the step size in the range function to skip items in a list.
⋅⋅* For example, if you have range(0, 10, 2) it will get every other item in the list.

In the function `final_grade(grade_list)` you will calculate and return the final grade for the student. The function `final_grade(grade_list)` should do the following:
* Find the average for Quizzes and Test
* Calculate & return the final weighted grade.
```py 
# Example: If quizzes were 25% and test 75%
(quizzes * .25) + (test * .75)
# This would give you the weighted grade.
```

Grades are weighted by the following weight:
* Quizzes - 10%
* Tests - 40%
* Midterm - 25%
* Final - 25%

For this one, you can assume that there will always be 4 quizzes and 4 test to calculate the average.