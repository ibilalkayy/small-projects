static WORDS: &str = include_str!("../assets/words.json");

pub fn get_words() -> String {
    WORDS.to_string()
}
