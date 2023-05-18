function view(id) {

	// console.log(jsonData)
	let host = "http://84.201.139.249:8080"
	let url = "/country?id=" + id
	
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("GET", host + url, true)
	request.setRequestHeader('Content-Type', 'application/json');

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			console.log(obj)
			viewCountry(obj["Data"]["country"])
		}
	});
	 
	request.send();
}

function viewC(id) {

	// console.log(jsonData)
	let host = "http://84.201.139.249:8080"
	let url = "/country?id=" + id
	
	const request = new XMLHttpRequest();
	request.responseType =	"json";
	request.open("GET", host + url, true)
	request.setRequestHeader('Content-Type', 'application/json');

	request.addEventListener("readystatechange", () => {
		if (request.readyState === 4 && request.status === 200) {
			let obj = request.response;
			console.log(obj)
			viewCCountry(obj["Data"]["country"])
		}
	});
	 
	request.send();
}

function disable() {
	if (document.getElementById("main__form__box__third") != null)
    	document.getElementById("main__form__box__third").style.display = "none";
	if (document.getElementById("map") != null)
    	document.getElementById("map").style.display = "none";
	if (document.getElementById("back_btn") != null)
   		document.getElementById("back_btn").style.display = "none";
	if (document.getElementById("main__form__box__fourth") != null)
    	document.getElementById("main__form__box__fourth").style.display = "block";
	if (document.getElementById("sliderRU") != null)
		document.getElementById("sliderRU").style.display = "block";
}

function viewCountry(data) {
	disable()
    let text = document.getElementById("desc");
    let name = document.getElementById("cnt_name");

	console.log(data)
	text.innerText = data["description"]
	name.innerText = data["name"]

	document.getElementById("slideone").src = "./images/" + data["name"] + "/2.jpg";
    document.getElementById("slidetwo").src = "./images/" + data["name"] + "/1.jpg";
}


function viewCCountry(data) {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("countries_box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");

    text.innerHTML = data["description"]
    name.innerHTML = data["name"]

    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }

	document.getElementById("slideone").src = "./images/" + data["name"] + "/2.jpg";
	document.getElementById("slidetwo").src = "./images/" + data["name"] + "/1.jpg";
}

function lastback() {
    document.getElementById("main__form__box__third").style.display = "block";
    document.getElementById("main__form__box__fourth").style.display = "none";
    document.getElementById("back_btn").style.display = "block";
    document.getElementById("map").style.display = "block";
    document.getElementById("sliderRU").style.display = "none";
    // let ele = document.getElementsByClassName('dot');
    // for (var i = 0; i < ele.length; i++) {
    //     ele[i].style.display = "block";
    // }
}

function viewJP() {
    document.getElementById("main__form__box__third").style.display = "none";
    document.getElementById("map").style.display = "none";
    document.getElementById("back_btn").style.display = "none";
    document.getElementById("main__form__box__fourth").style.display = "block";
    let text = document.getElementById("desc");
    let name = document.getElementById("cnt_name");

    text.innerHTML = "Japan, located in East Asia, is known for its unique blend of ancient traditions and modern innovations. Its archipelago consists of four main islands and numerous smaller ones, all surrounded by beautiful seas. Japan is renowned for its rich cultural heritage, including traditional arts such as calligraphy and flower arrangement, as well as its cuisine, which includes sushi, ramen, and many other delicious dishes.";
    name.innerHTML = "Japan";

    document.getElementById("sliderRU").style.display = "block";

    document.getElementById("slideone").src = "https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663137765_58-mykaleidoscope-ru-p-gorod-tokio-pinterest-65.jpg";
    document.getElementById("slidetwo").src = "https://sportishka.com/uploads/posts/2022-04/1650631589_27-sportishka-com-p-osaka-gorod-v-yaponii-krasivo-foto-28.jpg";

    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }
}

function viewCH() {
    document.getElementById("main__form__box__third").style.display = "none";
    document.getElementById("map").style.display = "none";
    document.getElementById("back_btn").style.display = "none";
    document.getElementById("main__form__box__fourth").style.display = "block";
    let text = document.getElementById("desc");
    let name = document.getElementById("cnt_name");

    text.innerHTML = "China, located in East Asia, is the world's most populous country with a rich history and culture that dates back thousands of years. It is known for its iconic landmarks such as the Great Wall of China, the Forbidden City, and the Terracotta Army. China is also famous for its cuisine, which includes dishes such as Peking duck, dumplings, and hot pot.";
    name.innerHTML = "China";

    document.getElementById("slideone").src = "https://look.com.ua/pic/201901/1600x1200/look.com.ua-319924.jpg";
    document.getElementById("slidetwo").src = "https://mobimg.b-cdn.net/v3/fetch/f5/f5cb4646e8a838c558289d49de4095f1.jpeg?w=1470&r=0.5625";

    document.getElementById("sliderRU").style.display = "block";
    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }
}

function viewCA() {
    document.getElementById("main__form__box__third").style.display = "none";
    document.getElementById("map").style.display = "none";
    document.getElementById("back_btn").style.display = "none";
    document.getElementById("main__form__box__fourth").style.display = "block";
    let text = document.getElementById("desc");
    let name = document.getElementById("cnt_name");

    text.innerHTML = "Canada is a North American country known for its vast landscapes, diverse culture, and friendly people. It is the second-largest country in the world by land area and has a population of over 37 million people. Canada is known for its natural beauty, including the Rocky Mountains, Niagara Falls, and the Northern Lights. The country is also famous for its multiculturalism, with a mix of Indigenous, French, British.";
    name.innerHTML = "Canada";

    document.getElementById("slideone").src = "https://media.cntraveler.com/photos/59555722b8636a1b8cd9a129/4:3/w_1600%2Cc_limit/GettyImages-488116889.jpg";
    document.getElementById("slidetwo").src = "https://futuresabroad.com/wp-content/uploads/2018/02/Canada-toronto-1426205.jpg";

    document.getElementById("sliderRU").style.display = "block";
    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }
}
function viewUSA() {
    document.getElementById("main__form__box__third").style.display = "none";
    document.getElementById("map").style.display = "none";
    document.getElementById("back_btn").style.display = "none";
    document.getElementById("main__form__box__fourth").style.display = "block";
    let text = document.getElementById("desc");
    let name = document.getElementById("cnt_name");

    text.innerHTML = "The United States of America is a country located in North America, bordered by Canada to the north and Mexico to the south. It is known for its diverse geography, ranging from the beaches of Florida to the mountains of Colorado and the deserts of Arizona. The country is also home to many iconic landmarks, such as the Statue of Liberty, the Golden Gate Bridge, and the Grand Canyon.";
    name.innerHTML = "USA";

    document.getElementById("slideone").src = "https://www.ysgsolar.com/sites/default/files/styles/panopoly_image_original/public/blog/pexels-life-of-pix-7613.jpg?itok=u2CqO-us";
    document.getElementById("slidetwo").src = "https://vsegda-pomnim.com/uploads/posts/2022-04/1649326699_19-vsegda-pomnim-com-p-maiami-plyazh-foto-19.jpg";

    document.getElementById("sliderRU").style.display = "block";
    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }
}

// countriesPAGE
function lastback_cnt() {
    document.getElementById("desc_box").style.display = "none";
    document.getElementById("countries_box").style.display = "flex";
    document.getElementById("sliderRU").style.display = "none";
}

function viewRU_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");

    text.innerHTML = "Russia (Russian Federation) is the largest state in the world, occupying 1/6 of the entire land area of the Earth.This is a country with a long history, rich culture and generous nature.Some monuments of architecture, history and culture on the territory of the Russian Federation are included in the list of UNESCO World Heritage Sites.";
    name.innerHTML = "Russia";

    document.getElementById("slideone").src = "https://ic.pics.livejournal.com/storm100/57290853/6706082/6706082_original.jpg";
    document.getElementById("slidetwo").src = "https://kartinkin.net/pics/uploads/posts/2022-08/1660202241_6-kartinkin-net-p-isaakievskaya-ploshchad-sankt-peterburg-kr-6.jpg";
}
function viewJP_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Japan, located in East Asia, is known for its unique blend of ancient traditions and modern innovations. Its archipelago consists of four main islands and numerous smaller ones, all surrounded by beautiful seas. Japan is renowned for its rich cultural heritage, including traditional arts such as calligraphy and flower arrangement, as well as its cuisine, which includes sushi, ramen, and many other delicious dishes.";
    name.innerHTML = "Japan";

    document.getElementById("slideone").src = "https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663137765_58-mykaleidoscope-ru-p-gorod-tokio-pinterest-65.jpg";
    document.getElementById("slidetwo").src = "https://sportishka.com/uploads/posts/2022-04/1650631589_27-sportishka-com-p-osaka-gorod-v-yaponii-krasivo-foto-28.jpg";

}
function viewCH_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "China, located in East Asia, is the world's most populous country with a rich history and culture that dates back thousands of years. It is known for its iconic landmarks such as the Great Wall of China, the Forbidden City, and the Terracotta Army. China is also famous for its cuisine, which includes dishes such as Peking duck, dumplings, and hot pot.";
    name.innerHTML = "China";

    document.getElementById("slideone").src = "https://look.com.ua/pic/201901/1600x1200/look.com.ua-319924.jpg";
    document.getElementById("slidetwo").src = "https://mobimg.b-cdn.net/v3/fetch/f5/f5cb4646e8a838c558289d49de4095f1.jpeg?w=1470&r=0.5625";

}
function viewCA_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Canada is a North American country known for its vast landscapes, diverse culture, and friendly people. It is the second-largest country in the world by land area and has a population of over 37 million people. Canada is known for its natural beauty, including the Rocky Mountains, Niagara Falls, and the Northern Lights. The country is also famous for its multiculturalism, with a mix of Indigenous, French, British.";
    name.innerHTML = "Canada";

    document.getElementById("slideone").src = "https://media.cntraveler.com/photos/59555722b8636a1b8cd9a129/4:3/w_1600%2Cc_limit/GettyImages-488116889.jpg";
    document.getElementById("slidetwo").src = "https://futuresabroad.com/wp-content/uploads/2018/02/Canada-toronto-1426205.jpg";

}
function viewUSA_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "The United States of America is a country located in North America, bordered by Canada to the north and Mexico to the south. It is known for its diverse geography, ranging from the beaches of Florida to the mountains of Colorado and the deserts of Arizona. The country is also home to many iconic landmarks, such as the Statue of Liberty, the Golden Gate Bridge, and the Grand Canyon.";
    name.innerHTML = "USA";

    document.getElementById("slideone").src = "https://www.ysgsolar.com/sites/default/files/styles/panopoly_image_original/public/blog/pexels-life-of-pix-7613.jpg?itok=u2CqO-us";
    document.getElementById("slidetwo").src = "https://vsegda-pomnim.com/uploads/posts/2022-04/1649326699_19-vsegda-pomnim-com-p-maiami-plyazh-foto-19.jpg";

}
function viewAV_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Australia is a country located in the southern hemisphere, surrounded by the Indian and Pacific oceans. It is the world's sixth-largest country by total area and has a population of approximately 25 million people. The capital city is Canberra, while the largest city is Sydney. Australia is known for its diverse landscapes, including the Great Barrier Reef, the Outback, and the beautiful beaches.";
    name.innerHTML = "Australia";

    document.getElementById("slideone").src = "https://blog.pokke.in/wp-content/uploads/2017/09/shutterstock_523437463-min.jpg";
    document.getElementById("slidetwo").src = "https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663445557_62-mykaleidoscope-ru-p-velikii-melburn-dostoprimechatelnosti-dost-68.jpg";

}
function viewGR_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Greece is a country located in southeastern Europe, known for its rich history and culture. It is home to many ancient ruins and landmarks, including the Acropolis in Athens and the Parthenon temple. Greece is also famous for its beautiful islands, such as Santorini and Mykonos, which attract millions of tourists every year. The country is known for its delicious cuisine, including dishes like moussaka and souvlaki.";
    name.innerHTML = "Greece";

    document.getElementById("slideone").src = "https://img.imageboss.me/greca/width/1920/5dd6a737d7b68.jpeg";
    document.getElementById("slidetwo").src = "https://gulaytour.ru/wp-content/uploads/2018/02/iSDJR5I1W.jpg";

}
function viewGO_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Georgia is a country located in the Caucasus region of Eurasia, bordered by Russia to the north, Turkey and Armenia to the south, and Azerbaijan to the east. It is known for its stunning mountain landscapes, ancient cities, and unique culture. Georgia has a rich history, with evidence of human habitation dating back to the Paleolithic era.";
    name.innerHTML = "Georgia";

    document.getElementById("slideone").src = "https://sportishka.com/uploads/posts/2022-04/thumbs/1650490897_22-sportishka-com-p-gruziya-tbilisi-dostoprimechatelnosti-kras-22.jpg";
    document.getElementById("slidetwo").src = "https://uniglobeevents.com/content/shutterstock594149126-20180619021739.jpg";

}
function viewDM_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Denmark is a country located in Northern Europe, bordered by Germany to the south. It is known for its beautiful coastline, historic castles, and modern architecture. The capital city, Copenhagen, is a vibrant and cosmopolitan city with a rich cultural heritage. Denmark is also famous for its design, fashion, and culinary scene.";
    name.innerHTML = "Denmark";

    document.getElementById("slideone").src = "https://s1.1zoom.ru/b5357/603/Denmark_Copenhagen_Houses_Motorboat_Canal_564924_1600x1200.jpg";
    document.getElementById("slidetwo").src = "https://www.scandichotels.com/imagevault/publishedmedia/ps5wimz1gn6rypa7ecoo/Aarhus_town_canal_summer.jpg";

}
function viewAB_cnt() {
    document.getElementById("desc_box").style.display = "block";
    document.getElementById("main__form__box").style.display = "none";
    document.getElementById("sliderRU").style.display = "block";
    let text = document.getElementById("desc_cnt");
    let name = document.getElementById("cnt_name_cnt");


    text.innerHTML = "Abkhazia is a disputed region located in the South Caucasus, bordered by Russia to the north and Georgia to the south. It has a population of approximately 240,000 people and its capital is Sukhumi. Abkhazia declared independence from Georgia in 1999, but its sovereignty is only recognized by a few countries, including Russia, Nicaragua, and Venezuela.";
    name.innerHTML = "Abkhazia";

    document.getElementById("slideone").src = "https://mirplaneta.com/wp-content/uploads/2022/08/Naberezhnaya-Suhuma.jpg";
    document.getElementById("slidetwo").src = "https://vsegda-pomnim.com/uploads/posts/2022-04/1649111641_34-vsegda-pomnim-com-p-gagra-priroda-foto-36.jpg";

}