# Python – Week 1 Challenges
For these challenges, you'll find starter files marked with the name of the challenge itself. There will also be a project which you will find the in the designated `python-week-1/project-files` folder with instructions for completion.

## 1. Variable Assignment (Cost Per Candy)
### Context
This problem is a simple variable assignment problem to get you used to the Qualified system and help you practice some basic variable assignment with Python.

### Task & Objectives
For task will be to assign a number to the candy variable and a decimal number to the cost variable. Then using these variables, find the cost_per_candy.

### To-do
* Assign an integer to the variable called "num_of_candy"
* Assign a float number to the variable called "total_cost"
* Assign to the variable called "cost_per_candy" the "total_cost" divided by the "num_of_candy"

## 2. Strings (Encrypted Message)
### Context
This problem will have you work with strings to help you better understand them.

### Task
Your task is to decrypt an encrypted a message. The variable "secrete" holds the encrypted message. You must decrypt this message and assign it to "message"

### To-do
Don't worry we know the key to decrypting the message. All you need to do is:
* reverse the string 
* get every other letter from the string
* remember to assign the decrypted message to the variable "message"
* **HINT:** A step size of -1 will reverse a string.

## 3. Lists (Cheese is Not a Pet)
### Context
In this problem you will be working to add and delete items from a list.

### Task
We have a list of pets, but we accidentally added "cheese" to the list.
Your task is remove **cheese** from the list and add **"cat"** to the end of the list.

### To-do
* Remove "cheese" from the pets list
* Add "cat" to the end of the pets list

## 4. Strings + Lists (Counting Words)
### Context
This problem is meant to give you practice using build in methods in Python.

### Task
Your task is to find out how many words are in the string **"escalator_msg"**, which we have put one  of our RV sayings in. Assign the total word count to the variable named "word_count".

### To-do
* Take the string and split it into individual words
* Count how many words there are in that list of words

## 5. Strings + Lists, Pt. 2 (Unique Words)
### Context
This problem is very similar to the previous one and will also help you work on using the build in methods in Python.

### Task
Now that we know how many words there are in the message, we want to know how many different words there are in this phrase.

### To-do
* Make sure that capitalization does not affect the outcome
* Count how many different words there are
**HINT:** Choosing the right data structure is crucial since we are looking for unique items, think of data structure that does not allow duplicates

## 6. If Else & Format Printing (Can I Afford It?)
### Context
This problem is meant to give you practice with control flow, logical operators, and formatted printing.

### Task
I am currently on a tight budget, so anything above $20 is out of my budget. So, I need a tool that will help me decide whether or not I should buy something. Anything above $20 is not currently within my budget. Anything under $20 is within my budget, but I would like to know how much I would have left over if I did buy the product.

For the purpose of this problem you will be programming within a function. This is just for testing purposes. All you need to know is:
* "product" is the string name of the product I am thinking of buying.
* and "price" is the price of the product.
* For the code to be in the function just indent four spaces under the function name. (DO NOT USE TABS)

### To-do
* Determine if the price is within my budget which is $20 or less (including $20)
  * If I can afford it. Use format printing to print the following message
  * "You can afford `${product}`, you will spend `${price}` and have `${left over}` left over"
  * Since this is money make, sure that 2 decimal places are displayed.
  * Ex. If you were to have something cost $1, make sure it prints $1.00.
  * left over amount is $20 - the price of the product I am thinking of buying.
* If I cannot afford it Print the following message:
  * "You cannot afford a {product} at the time, it is ${amount over 20} over your budget."
  * The second formatted variable will need to also have two decimal places and be the amount more than the $20 I have, to buy the product.

## 7. For Loops (Getting Even)
### Context
This problem will help you practice working with for loops.

### Task
For this problem you have a list of integers called "my_numbers" and an empty list called "even_numbers". You will go through the numbers list and copy the even numbers and add them to the other list.

### To-do
Create a for loop inside the function that will add the `numbers` from `my_numbers` list to the even_numbers list if they are even. To do that, you will need to:
1. Iterate through the my_numbers list
2. If one of the numbers is an even number, add it to the even_numbers list

Do not change my_numbers list at all, you should just get the values from this list

#### Expected
```py
# Before
my_numbers = [63, 24, 7, 99, 17, 82, 45, 21, 97, 78]
even_numbers = []
```

```py
# Given the function
even_nums(my_numbers, even_numbers):
  # Your code inside of even_nums
```

To work inside the function, just make sure that your code has an indentation of 4 spaces at the start of each line.

```py
# After
my_numbers = [63, 24, 7, 99, 17, 82, 45, 21, 97, 78]
even_numbers = [24, 82, 78]
```

## 8. Scope (Getting New Candy)
### Content
For this problem you be working with global variables, block scope to be able to access variables outside of a function, and default values in a function.

### Task
You will create a global variable called candy and assign it a value. You will then change the global variable from inside the function without passing it in to the function.

### To-do
* Create a **global variable** called `candy` and assign it the string `Chocolate`
* Create a **function** called `new_candy` with the parameter `candy_name`
* Make `candy_name` default to `Candy cane` if nothing is passed in
* Capitalize the string `candy_name`, and assign it to the global variable `candy`

For the purpose of this problem, we will be using the global key word, but note that this can also be done with returning the new value and reassigning it.

#### Scope Example
```py
foo = 10

def bar(number):
  # This should be the global foo variable, not a new one. 
  foo = number

bar(13)
# foo here should be 13 not 10.
```

In this example, `foo` is the global variable.