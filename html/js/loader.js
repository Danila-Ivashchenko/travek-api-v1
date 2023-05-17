load_cityname()

function load_tags() {
	console.log("loaded")
	let host = "http://api:8080"
	let url = "/tags"
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("GET", host + url, true)

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			console.log(obj)
			create_tags(obj["Data"]["tags"])
		}
	});
	request.send()
}

function create_tags(data) {
	console.log(data)
	let tags_list = document.querySelector("#tags_list")
	tags_list.replaceChildren()
	for (let i = 0; i < data.length; i++){
		let li = document.createElement("li")
		li.className = "tag"
		let label = document.createElement("label")
		label.innerText = data[i]["name"]
		label.className = "tag_label"
		label.setAttribute("for", "tag_" + data[i]["id"])
		let checkbox = document.createElement("input")
		checkbox.setAttribute("type", "checkbox")
		checkbox.setAttribute("value", data[i]["id"])
		checkbox.setAttribute("id", "tag_" + data[i]["id"])
		checkbox.className = "tag_value"
		li.appendChild(checkbox)
		li.appendChild(label)
		// li.appendChild(checkbox)
		tags_list.appendChild(li)
	}
}

function load_cityname() {
	let host = "http://api:8080"
	let url = "/country/all"
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("GET", host + url, true)

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			console.log(obj)
			console.log(obj)
			create_countries_list(obj)
		}
	});
	request.send()
}

function create_countries_list(data){
	console.log(data)
	if (!data["success"]){
		console.log(data["exaption"])
		return
	}
	let countries_list = document.querySelector("#cityname")
	countries_list.replaceChildren()
	for (let i = 0; i < data["Data"]["countries"].length; i++){
		let option = document.createElement("option")
		option.setAttribute("value", data["Data"]["countries"][i]["name"])
		option.setAttribute("country_id", data["Data"]["countries"][i]["id"])
		option.className = "country_id_name"
		countries_list.appendChild(option)
	}
}