function filter() {
    const urlParams = new URLSearchParams(window.location.search);
    const tag = urlParams.get('tag');

    if (!tag) {
        return
    }

    let filteredAnything = false;

    document.querySelectorAll(".toc tbody tr").forEach(row => {
        const hasTag = row.querySelector(`.tags a[href="/posts/?tag=${tag}"]`);
        if (!hasTag) {
            row.classList.add("hide");

            filteredAnything = true;
        } else {
            row.classList.remove("hide")
        }
    });

    if (filteredAnything) {
        document.querySelectorAll(".toc tbody").forEach(tbody => {
            const row = document.createElement("tr")
            const cell = document.createElement("td")
            cell.innerHTML = `Filtered on tag <code>${tag}</code> (<a href="./">clear filter</a>)`
            cell.setAttribute("colspan", "9")
            cell.setAttribute("style", "text-align: center; font-size: 0.9rem")

            row.appendChild(cell)
            tbody.prepend(row)
        });
    }
}

document.addEventListener("DOMContentLoaded", filter);