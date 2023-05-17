function Back() {
    document.getElementById("main__form__box").style.display = "flex";
    document.getElementById("main__form__box__second").style.display = "none";
    document.getElementById("back_btn").style.display = "none";
    document.getElementById("main__form__box__third").style.display = "none";
}
function SecBack() {
    document.getElementById("main__form__box__third").style.display = "none";
    document.getElementById("main__form__box__second").style.display = "flex";
    let ele = document.getElementsByClassName('dot');
    for (var i = 0; i < ele.length; i++) {
        ele[i].style.display = "none";
    }
}
