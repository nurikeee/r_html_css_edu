const inputBox = document.getElementById("input-box")
const todoList = document.getElementById("todo-list")

todoList.addEventListener("click", function (e) {
    if (e.target.tagName === 'LI') {
        e.target.classList.toggle('checked');
    } else if (e.target.tagName === 'IMG') {
        e.target.parentElement.parentElement.remove();
    }
    save();
}, false)

document.addEventListener("DOMContentLoaded", restore)


function addTask(e) {
    // console.log("PRESSED!!!")

    if (inputBox.value === '') {

    } else {
        let newListItem = document.createElement("li");
        newListItem.innerHTML = inputBox.value;
        todoList.appendChild(newListItem);

        let deleteIcon = document.createElement("span")
        let deleteImage = document.createElement("img")
        deleteImage.setAttribute("src", "./images/cross.png")

        deleteIcon.appendChild(deleteImage)

        newListItem.appendChild(deleteIcon)

    }
    inputBox.value = "";
    save();
}

function save() {
    localStorage.setItem("todo-data", todoList.innerHTML)
}

function restore() {
    todoList.innerHTML = localStorage.getItem("todo-data")
}

// restore();

