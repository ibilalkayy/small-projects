function changeColor() {
    const colorNames = ["blue", "pink", "orange", "green", "yellow", "red", "white"]

    let randomIndex = Math.floor(Math.random() * colorNames.length)
    let randomColor = colorNames[randomIndex]

    document.body.style.backgroundColor = randomColor

    document.getElementById("colornames").innerHTML = randomColor
}

document.querySelector("#colornames").addEventListener("click", changeColor)