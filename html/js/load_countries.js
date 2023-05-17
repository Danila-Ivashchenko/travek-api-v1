loadCountries()

function showCountries(data) {
	let countries_box = document.querySelector("#countries_box")
	countries_box.replaceChildren()
	for (let i = 0; i < data.length; i++){
		let country_box = document.createElement("div")
		country_box.className = "county_view"
		country_box.innerText = (i + 1) + '. ' + data[i]["name"]
		let but = document.createElement("input")
		but.setAttribute("type", "button")
		but.setAttribute("value", "view")
		let id = data[i]["id"]
		but.setAttribute("onclick", `viewC(${id})`)
		country_box.appendChild(but)
		countries_box.appendChild(country_box)
	}
}

function loadCountries() {
	let host = "http://84.201.139.249:8080"
	let url = "/country/all"
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("GET", host + url, true)

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			showCountries(obj["Data"]["countries"])
		}
	});
	request.send()
}