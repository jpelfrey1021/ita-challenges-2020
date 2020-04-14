# Python – Week 2 Challenges
For these challenges, you'll find starter files marked with the name of the challenge itself. 

## 1. OOP (Making Candies)
### Context
This problem is meant to give you practice with Objects/Classes including constructing them, attributes, and methods.

### Task
You will create a Candy Class that will have a constructor method that takes in the name and cost of the candy. Will also have a total_cost method to calculate the total cost for x amount of candies. Lastly, you will make a method inside the Candy class that will overwrite the default string method.

### To-do
* Create a Class called `Candy`
* The constructor method for the Class (`init` method) will take a name, and a cost as parameters and assign them as attributes of the class
* Make a class method called total_cost that takes in a number as a parameter. Use this number to print the following formatted text:
```py
"4 Chocolates at 0.50 will cost 2.00"
```
  * 4 is the number that is passed into the method
  * Then the name of the candy with an "s" at the end
  * Then "at" and the cost that the class has
  * lastly, "will cost" and the number passed in multiplied by the cost

* Write a magic _ `Str` _ method that returns the following string:
```py
"A Chocolate costs $1"
```
  * "Chocolate" is the name of the candy
  * 1 is the cost of a candy.

## 2. Error Handling (Houston, We Have a Problem)
### Context
This problem is meant to give you practice with error handling.

### Task
You will create a function that will take in two strings and return the concatenation of both of them, but what if someone tries to pass in something that does not work. You must catch the error and print an error message.

### To-do
* Create a function called concatenate that takes in two parameters
* Inside the function you need to concatenate the two parameters passed with a space between them
```py
concatenate ("candy", "cane") ==> "candy cane"
```
* If something is passed in that will cause an issue when concatenating the two strings, print the error message: `"Houston, we have a problem"`

The function will need to always return a string back.
* If there was no error, return the concatenated string
* Otherwise, print the error message out and return an empty string
```py
concatenate ("candy", 1.50) ==> print ("Houston, we have a problem") & return ""
```

## 3. Generator (Whenever)
### Context
This problem is meant to give you practice with the datetime module and the generators.

### Task
For this problem you be working to create a random time generator. You will make a generator that generates a random hour and minute.

### To-do
* Import the `datetime` module and `randint` function from the `random` module
* Create a function called `random_time` that takes in a parameter for how many times it will generate
* Inside the `random_time` function, create a generator that will generate a random time up to the number that was passed in
* Use the `randomint()` function to generate random hours and minutes for the times

#### Example
```py
# This should generate up to two random times. 
random_time(2)
```

## 4. Everything Else (The Final Countdown)
### Context
These last challenges will have several little challenges that are important for a python developer, but we do not want to overwhelm you with problems so we will go through these last few problems quickly.

### Task
To go through the challenges and complete them using the instructed way. There are other ways to solve some of these problems, but please refrain from doing so as the intention is to get you to practice using these methods.

### To-do
1. List comprehension - `sales_tax()`:
  * `sales_tax()` should take in a list of prices (floats & ints) as a parameter
  * Use list comprehension to create a list of prices that include the sales tax for the prices
  * Sales tax for the purpose of this problem will be 6%
2. Filter function - `under_twenty()`:
  * `under_twenty()` should take in a list of numbers as a parameter
  * Use the filter function to return a list of numbers under twenty
  * **HINT:** Use lambda function to reduce code
3. Map function - `dec_to_binary()`:
  * `dec_to_binary()` should take in a list of numbers as a parameter
  * Use the map function to return a list of numbers that represent the binary representations of the original numbers
  * Ex. `[1,4,5] => ['0b1', '0b100', '0b101']`
  * **HINT:** `bin(x)` is the binary function
4. `*args` - `average()`:
  * `average()` should be able to take in any number of variables as parameters
  * All the parameters will be numbers
  * Inside the function you need to find the average for the parameters passed in
  * Ex. `average(1,2,3,4,5) => 3.0`