import random
from words import words

def get_random_word():
    word = random.choice(words)
    while " " in word:
        word = random.choice(words)
    return word

def hangman():
    word = get_random_word()
    print(word)
    print("Guess the word: ")
    for _ in range(len(word)):
        print("_", end=" ")

    specific = input("\n\nEnter a letter to guess: ")
    if specific in word:
        print("Found")
    else:
        print("Not found")

hangman()