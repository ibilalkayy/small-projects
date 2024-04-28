const reviews = [
    {
        id: 1,
        name: "Susan Smith",
        job: "Web developer",
        img: "https://res.cloudinary.com/djywxphqy/image/upload/v1696570177/tj7qhk0jkzdfvpgv1ccg.jpg",
        text: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla, voluptatem consequuntur ipsum architecto unde eius magnam fuga in. Quasi, ad!",
    },
    {
        id: 2,
        name: "John Smith",
        job: "Web designer",
        img: "https://res.cloudinary.com/djywxphqy/image/upload/v1696570177/lgcouvteh3mz2kjs6pqf.jpg",
        text: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla, voluptatem consequuntur ipsum architecto unde eius magnam fuga in. Quasi, ad!",
    },
    {
        id: 3,
        name: "Andrew Smith",
        job: "Android Developer",
        img: "https://res.cloudinary.com/djywxphqy/image/upload/v1696570178/vzf0cxd4nkq9jux0al3r.jpg",
        text: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla, voluptatem consequuntur ipsum architecto unde eius magnam fuga in. Quasi, ad!",
    },
]

const img = document.getElementById("person-img");
const author = document.getElementById("author");
const job = document.getElementById("job");
const info = document.getElementById("info");

const prevBtn = document.querySelector(".prev-btn");
const nextBtn = document.querySelector(".next-btn");

// set starting item
let currentItem = 0;

// load initial item
window.addEventListener("DOMContentLoaded", function() {
    showPerson(currentItem)
});

function showPerson(person) {
    const item = reviews[currentItem];
    img.src = item.img;
    author.textContent = item.name;
    job.textContent = item.job;
    info.textContent = item.text;
}

nextBtn.addEventListener("click", function(){
    currentItem++;
    if (currentItem > reviews.length - 1) {
        currentItem = 0;
    }
    showPerson(currentItem)
});

prevBtn.addEventListener("click", function(){
    currentItem--;
    if (currentItem < 0) {
        currentItem = reviews.length -1;
    }
    showPerson(currentItem)
})