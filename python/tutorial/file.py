# Variables 
first_variable = "Hello World"

print(first_variable)

age = 20
age = 30

is_online = False

print(is_online)

# Short porgram
name = "John Smith"
age = "20"
patient = "new patient"

print(name + " is " + age + " years old and he is a " + patient)

# # Input
name = input("Enter your name: ")
print("Hello, " + name)

birth_year = input("Enter your birth year: ")
age = 2023 - int(birth_year)
print(age)

# Exercise
first_num = input("First number: ")
second_num = input("Second number: ")

sum = float(first_num) + float(second_num)
print("Sum: " + str(sum))

# Strings
course = "Python for Beginners"

# Capitalize
print(course.capitalize())

# Find the word index
print("Here is the index of 'y' in this sentence: " + str(course.find('y')))

# Replace it with another text
print(course.replace("for", "4"))

# In operator
print("Python" in course)


# Arithmetic operator

# Divided
print(10 / 3)

# Don't get values after dot
print(10 // 3)

# Remainder
print(10 % 3)

# 10 to the power of 3
print(10 ** 3)

# Augment assigned operator
x = 10
x = x + 3
# OR
x += 3
x *= 3

print(x)

# Division and multiplication has highest priority
x = 10 / 2 * 3
print(x)

# If-statement

temperature = 15

if temperature > 30:
    print("It's a hot day")
    print("Drink plenty of water")
elif temperature > 20:
    print("It's a nice day")
elif temperature > 10:
    print("It's a bit cold")
else:
    print("It's cold")
print("Done")

# Exercise

weight = int(input("Weight: "))
unit = input("K(gs) or L(bs): ")

if unit.upper == "K":
    converted = weight / 0.45
    print("Weight in Lbs: " + str(converted))
else:
    converted = weight * 0.45
    print("Weight in Kgs: " + str(converted))

# while loop
i = 1
while i <= 10:
    print(i * "*")
    i += 1

# Lists
names = ["John", "Mary", "Smith", "Carlos", "Sam"]
names[0] = "Bilal"
print(names[0])

# Insert Method

num = [1, 2, 3, 4, 5]
num.insert(0, -1)
print(num)

num.remove(1)
print(num)

num.clear()
print(num)

print(1 in num)

print(len(num))

# for loop

numbers = [1, 2, 3, 4, 5]
for item in numbers:
    print(item)

i = 0
while i < len(numbers):
    print(numbers[i])
    i += 1

# Range functions
number = range(5)
for num in number:
    print(num)

# 2 steps
numb = range(5, 10, 2)
for n in numb:
    print(n)

# just use range
for i in range(4):
    print(i)

# Tuples - It can'be changed
first = (1, 2, 3, 2)
result = first.count(3)
result = first.index(3, 0, 3)
print(result)