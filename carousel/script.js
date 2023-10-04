const professionals = {
    firstName: "Susan Smith",
    firstPersonProfession: "WEB DEVELOPER",
    secondName: "John Smith",
    secondPersonProfession: "WEB DEVELOPER",
    thirdName: "Andrew Smith",
    thirdPersonProfession: "WEB DESIGNER"
};

document.getElementById("first-name").innerHTML = professionals.firstName
document.getElementById("susan-profession").innerHTML = professionals.firstPersonProfession

document.getElementById("second-name").innerHTML = professionals.secondName
document.getElementById("john-profession").innerHTML = professionals.secondPersonProfession

document.getElementById("third-name").innerHTML = professionals.thirdName
document.getElementById("andrew-profession").innerHTML = professionals.thirdPersonProfession