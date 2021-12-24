let qur = (x) => {
    return document.querySelector(x)
}
let qurAll = (x) => {
    return document.querySelectorAll(x)
}



var lime = (qure, n) => {
    qur(qure).addEventListener('mouseover', () => {
        qurAll('.wrapp-links p')[n].style.opacity = 1
        qurAll('.wrapp-links img')[n].style.opacity = 0
    })

    qur(qure).addEventListener('mouseout', () => {
        qurAll('.wrapp-links img')[n].style.opacity = 1
        qurAll('.wrapp-links p')[n].style.opacity = 0
    })
}



lime('.o1', 0)
lime('.o2', 1)
lime('.o3', 2)
lime('.o4', 3)
    // lime('.o5', 4)


function imgRecolection(cssSelector, srcX, srcY) {
    qur(cssSelector).addEventListener("mouseover", () => {
        qur(cssSelector + " img").src = srcY;
    });
    qur(cssSelector).addEventListener("mouseout", () => {
        qur(cssSelector + " img").src = srcX;
    });
}
imgRecolection('.vk', "../static/img/icons8-вконтакте-50.png", "../static/img/icons8-вконтакте-50-in.png");
imgRecolection('.facebook', "../static/img/icons8-facebook-100.png", "../static/img/icons8-facebook-100-in.png");
imgRecolection('.twiter', "../static/img/icons8-twitter-в-квадрате-50.png", "../static/img/icons8-twitter-в-квадрате-50-in.png");
imgRecolection('.instagram', "../static/img/instagram-new.png", "../static/img/instagram-new-in.png")

$('html').on('mousemove', function(e) {
    let bubble = $('<div class="bubble"></div>');
    bubble.css({ 'left': e.clientX - 50, 'top': e.clientY - 50 });
    $('body').append(bubble);
    setInterval(function() { bubble.remove() }, 1000)
});