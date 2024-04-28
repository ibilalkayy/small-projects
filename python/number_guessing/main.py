import random

user_number = input("Enter a number from 1-100: ")
computer_number = random.randint(1, 100)
print("Computer number: " + str(computer_number))

if int(user_number) == computer_number:
    print("Your number is matched with the computer number")
elif int(user_number) > computer_number:
    print("Your number is higher than a computer number")
elif int(user_number) < computer_number:
    print("Your number is less than a computer number")
else:
    print("Enter the number in range")