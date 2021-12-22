//..............................................................................................................
// Сокращенный вариант для выбора элемен(та/тов) на странице. По умолчанию выбираеться один элемент.
function E(CssSelector, All = false) {
    if (All == false) {
        // Возвращает обьект тега html
        return document.querySelector(CssSelector)
    } else {
        // Возвращает масив обьектов тега html
        return document.querySelectorAll(CssSelector)
    }
}
//................................................................................................................


//.................................................................................................................

// Один из лучших костылей которые я придумал на 21.05.2020
// Данный костыль позволяет добавлять события и функции на элементы созданные на странице
// динамически. Поскольку часто js отказываеться стандартными способами вешать событие на
// динамически созданный элемент. Функциия добавляет элементу событие ввиде атрибута тега элемента html.
// Пример использования:
// AddFunc("button", function () {
// ...............................
// }) 
// В обьявлении функции важно не делать лишних пробелов
// обязательно обьявлять функцию стандартным способом как в примере. 
function AddFunc(CssSelector, functionStandartViev, subevent = 'onclick') {
    let obj = {};
    // Добавляем функцию в маcсив.
    obj.func = functionStandartViev;
    // Приобразуем код функции в строку.
    obj.func = obj.func + '';
    let newstr = '';
    // Проходимся по строке функции и добавляем название.
    for (let i = 0; i < obj.func.length; i++) {
        newstr += obj.func[i];
        if (i == 8) {
            newstr += '_FncN';
        };
        if (i == obj.func.length - 1) {
            newstr += '_FncN();';
        };
    };
    // Добавляем элементу атрибут события с отредактированной строкой функции.
    E(`${CssSelector}`).setAttribute(subevent, `${newstr}`);
};

//..............................................................................................................

// Функция для краcивой замены картинки на текст.

var Img_replacement_text = (CssSelector, WrapperSelector, n = 0, displacement = "left", displacementPx = 100) => {
        E(CssSelector).addEventListener('mouseover', () => {
            E(WrapperSelector + ' p', true)[n].style.opacity = 1
            E(WrapperSelector + ' img', true)[n].style.opacity = 0
            if (displacement == "left")
                E(WrapperSelector + ' p', true)[n].style.marginLeft = 0 + 'px'
        })

        E(CssSelector).addEventListener('mouseout', () => {
            E(WrapperSelector + ' img', true)[n].style.opacity = 1
            E(WrapperSelector + ' p', true)[n].style.opacity = 0
            if (displacement == "left")
                E(WrapperSelector + ' p', true)[n].style.marginLeft = (-1 * displacementPx) + "px"
        })
    }
    //..............................................................................................................

// Функкция для смены картинки при наведении.
function imgRecolection(cssSelector, srcX, srcY) {
    E(cssSelector).addEventListener("mouseover", () => {
        E(cssSelector + " img").src = srcY;
    });
    E(cssSelector).addEventListener("mouseout", () => {
        E(cssSelector + " img").src = srcX;
    });
}