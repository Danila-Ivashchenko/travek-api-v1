function Points() {

    let country = document.getElementById('country').value;
    let point = document.getElementById('you');
    let flag = false;
    if (country == 'Russia') {
        point.style.top = "480px";
        point.style.left = "1050px";
        point.style.display = "block";
        flag = true;
    }
    else if (country == 'China') {
        point.style.top = "540px";
        point.style.left = "1050px";
        point.style.display = "block";
        flag = true;
    }
    else if (country == 'USA') {
        point.style.top = "475px";
        point.style.left = "650px";
        point.style.display = "block";
        flag = true;
    }
    else if (country == 'UK') {
        point.style.top = "495px";
        point.style.left = "882px";
        point.style.display = "block";
        flag = true;
    }
    else if (country == 'Japan') {
        point.style.top = "535px";
        point.style.left = "1110px";
        point.style.display = "block";
        flag = true;
    }
    else if (country == 'Canada') {
        point.style.top = "500px";
        point.style.left = "710px";
        point.style.display = "block";
        flag = true;
    }
    else {
        alert("Select a country from the list");
    }

    if (flag == true) {
        document.getElementById("main__form__box").style.display = "none";
        document.getElementById("main__form__box__second").style.display = "flex";
        document.getElementById("back_btn").style.display = "block";
		load_tags()
    }
}