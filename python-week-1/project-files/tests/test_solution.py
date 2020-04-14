# Here is where you will add the test cases used to validate the solution.
# The number and style of the test cases can significantly change how well you can evaluate your candidate's solution.
import unittest
from solution import file_to_string, string_to_list, final_grade


# Note: the class must be called Test
class Test(unittest.TestCase):
    def test_file_read(self):
        self.assertEqual(file_to_string("grades.txt"), "Quiz1 95\nQuiz2 90\nQuiz3 85\nQuiz4 100\nTest1 95\nTest2 85\nTest3 84\nTest4 80\nMidterm 70\nFinal 80")  

    def test_string_to_list(self):
        my_string = "Quiz1 100\nQuiz2 90"
        self.assertEqual(string_to_list(my_string), [("Quiz1", 100), ("Quiz2", 90)])
		
        my_grades = "Quiz1 95\nQuiz2 90\nQuiz3 85\nQuiz4 100\nTest1 95\nTest2 85\nTest3 84\nTest4 80\nMidterm 70\nFinal 80"
        grades = [('Quiz1', 95),('Quiz2', 90),('Quiz3', 85),('Quiz4', 100),('Test1', 95),('Test2', 85),('Test3', 84),('Test4', 80),('Midterm', 70),('Final', 80)]
        self.assertEqual(string_to_list(my_grades), grades)
		
    def test_final_grade(self):
        grades = [('Quiz1', 95),('Quiz2', 90),('Quiz3', 85),('Quiz4', 100),('Test1', 95),('Test2', 85),('Test3', 84),('Test4', 80),('Midterm', 70),('Final', 80)]
        self.assertEqual(final_grade(grades), 81.15)