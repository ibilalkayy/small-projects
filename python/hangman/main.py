import random
from words import words

def get_random_word():
    word = random.choice(words)
    while " " in word:
        word = random.choice(words)
    return word

def hangman():
    word = get_random_word()
    word_letters = set(word)
    guessed_letters = set()
    attempts = 6

    print("Let's play hangman!")

    while len(word_letters) > 0 and attempts > 0:
        print("You have", attempts, "attempts left. Letters you've used:", ' '.join(guessed_letters))

        current_word = [letter if letter in guessed_letters else '_' for letter in word]
        print("Current word", ' '.join(current_word))

        specific = input("\nEnter your guess: ").lower()
        if specific in guessed_letters:
            print("You've already guessed the letter")
            continue

        guessed_letters.add(specific)

        if specific in word_letters:
            word_letters.remove(specific)
        else:
            attempts -= 1
            print("Letter not in the word")

    if attempts == 0:
        print("Sorry you lost. The word is", word)
    else:
        print("Congratulations, you guessed the word", word)

hangman()