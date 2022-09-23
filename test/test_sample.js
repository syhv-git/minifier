clearIcon = document.querySelector(".clear-icon");
searchBar = document.querySelector(".search-input");

searchBar.addEventListener("keyup", () => {
    if(searchBar.value && clearIcon.style.visibility !== "visible"){
        clearIcon.style.visibility = "visible";
    } else if(!searchBar.value) {
        clearIcon.style.visibility = "hidden";
    }
});

clearIcon.addEventListener("click", () => {
    searchBar.value = "";
    clearIcon.style.visibility = "hidden";
})
