const forms = document.querySelectorAll("form")
const loader = document.getElementById("loader")
const input = document.getElementById("textInput")

forms.forEach(form => {
    form.addEventListener("submit", (e) => {
        const text = input.value
        if (form.id === "logoForm") {
            document.getElementById("logoText").value = text
        }
        loader.style.display = "block"
        document.getElementById("downloadPNG").href = "/download/png?text=" + text
        document.getElementById("downloadASCII").href = "/download/ascii?text=" + text
    })
})

function copyASCII() {
    const ascii = document.getElementById("ascii").innerText
    navigator.clipboard.writeText(ascii)
    alert("QR copied!")
}

function copyLogo() {
    const logo = document.getElementById("logo").innerText
    navigator.clipboard.writeText(logo)
    alert("Logo copied!")
}

document.getElementById("darkToggle").onclick = () => {

    document.body.classList.toggle("dark")
}