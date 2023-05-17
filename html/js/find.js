function Find() {
	data = extractData()
	sendData(data)
    document.getElementById("main__form__box__third").style.display = "block";
    document.getElementById("main__form__box__second").style.display = "none";
    document.getElementById('you').style.display = "none";
    // let ele = document.getElementsByClassName('dot');
    // for (var i = 0; i < ele.length; i++) {
    //     ele[i].style.display = "block";
    // }
}

function getCountryId(name){
	const counties = document.querySelectorAll(".country_id_name")
	for (let i = 0; i < counties.length; i++){
		if (name == counties[i].value){
			return parseInt(counties[i].getAttribute("country_id"))
		}
	}
	return -1
}


function extractData() {
	let data = {}
	let name = document.querySelector("#country").value
	data["sourse_country"] = getCountryId(name)

	let tags = []
	let all_tags = document.querySelectorAll(".tag_value")
	for (let i = 0; i < all_tags.length; i++){
		if (all_tags[i].checked) {
			tags.push(parseInt(all_tags[i].value))
		}
	}
	const form = document.querySelector(".second_form")
	// data["language"] = [form.querySelector("#language").value]
	// data["transports"] = []
	if (form.querySelector("#transports").value != ''){
		data["transports"] = [form.querySelector("#transports").value]
	} else {
		data["transports"] = []
	}
	if (form.querySelector("#max_time").value != '') {
		data["max_time"] = parseInt(form.querySelector("#max_time").value)
	} else {
		data["max_time"] = 0
	}
	if (document.querySelector('input[name="internal_passport"]:checked') != null) {
		data["internal_passport"] = parseInt(document.querySelector('input[name="internal_passport"]:checked').value) == 1;
	}
	if (document.querySelector('input[name="poss_rad"]:checked') != null) {
		data["possibility_to_stay_forever"] = parseInt(document.querySelector('input[name="poss_rad"]:checked').value) == 1;
	}
	data["tags"] = tags
	console.log(data)
	return data
}

function showData(data){
	let count_result = document.querySelector("#result_countries_len")
	let countries = data["Data"]["counties"]
	count_result.innerText = countries.length
	if (data["Data"]["counties"].length == 1) {
		count_result.innerText += " country"
	} else {
		count_result.innerText += " countries"
	}
	let countr__box = document.querySelector(".countr__box")
	countr__box.replaceChildren()
	
	for (let i = 0; i < countries.length; i++){
		let box = document.createElement("div")
		let name = countries[i]["name"]
	
		box.classList = "countr__box__item"
		id = countries[i]["id"]
		box.innerText = (i + 1) + ". " + name
		
		let but = document.createElement("input")
		but.setAttribute("type", `button`)
		but.setAttribute("onclick", `view(${id})`)
		but.setAttribute("value", `view`)

		box.appendChild(but)
		countr__box.appendChild(box)
	}
}

function sendData(data) {
	var jsonData= JSON.stringify(data);

	// console.log(jsonData)
	let host = "http://84.201.139.249:8080"
	let url = "/find"
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("POST", host + url, true)
	request.setRequestHeader('Content-Type', 'application/json');

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			console.log(obj)
			showData(obj)
		}
	});
	 
	request.send(jsonData);
}