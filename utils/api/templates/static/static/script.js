const form = document.querySelector("form")
const loader = document.getElementById("loader")
const input = document.getElementById("textInput")
const downloadPNG = document.getElementById("downloadPNG")

form.addEventListener("submit", () => {

    loader.style.display = "block"

    downloadPNG.href = "/download/png?text=" + input.value
})

function copyASCII() {

    const ascii = document.getElementById("ascii").innerText

    navigator.clipboard.writeText(ascii)

    alert("ASCII copied!")
}

document.getElementById("darkToggle").onclick = () => {

    document.body.classList.toggle("dark")
}