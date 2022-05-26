$(function ($) {
    // Smooth scroll to anchor
    $('body.home a[href*="#"]:not([href="#"])').click(function () {
        if (location.pathname.replace(/^\//, '') === this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
            if (target.length) {
                $('html,body').animate({
                    scrollTop: target.offset().top
                }, 1000);
                return false;
            }
        }
    });

    // gheading links
    $('.docs-wrapper').find('a[name]').each(function () {
        var anchor = $('<a href="#' + this.name + '"/>');
        $(this).parent().next('strong').wrapInner(anchor);
    });

    // Fixes FOUC for the search box
    $('.search.invisible').removeClass('invisible');

    // collapse and expand for the sidebar
    var toggles = document.querySelectorAll('.sidebar strong'),
        togglesList = document.querySelectorAll('.sidebar strong + ul');

    for (var i = 0; i < toggles.length; i++) {
        toggles[i].addEventListener('click', expandItem);
        toggles[i].addEventListener('keydown', expandItemKeyboard);
        toggles[i].setAttribute('tabindex', '0');
    }

    // Via https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API/Using_the_Web_Storage_API#Testing_for_availability
    function storageAvailable(type) {
        try {
            var storage = window[type],
                x = '__storage_test__';
            storage.setItem(x, x);
            storage.removeItem(x);
            return true;
        } catch (e) {
            return e instanceof DOMException && (
                    // everything except Firefox
                    e.code === 22 ||
                    // Firefox
                    e.code === 1014 ||
                    // test name field too, because code might not be present
                    // everything except Firefox
                    e.name === 'QuotaExceededError' ||
                    // Firefox
                    e.name === 'NS_ERROR_DOM_QUOTA_REACHED') &&
                // acknowledge QuotaExceededError only if there's something already stored
                storage.length !== 0;
        }
    }

    var docCollapsed = true;

    function expandDocs(e) {
        for (var i = 0; i < toggles.length; i++) {
            if (docCollapsed) {
                toggles[i].classList.add('is-active')
            } else {
                toggles[i].classList.remove('is-active')
            }
        }

        docCollapsed = !docCollapsed;
        document.getElementById('doc-expand').text = (docCollapsed ? '全部展开' : '全部收缩');

        if (storageAvailable('localStorage')) {
            localStorage.setItem('laravel_docCollapsed', docCollapsed);
        }
        if (e) {
            e.preventDefault();
        }
    }

    if (document.getElementById('doc-expand')) {
        if (storageAvailable('localStorage')) {
            if (localStorage.getItem('laravel_docCollapsed') === null) {
                localStorage.setItem('laravel_docCollapsed', true)
            } else {
                localStorage.getItem('laravel_docCollapsed') == 'false' ? expandDocs() : null
            }
        }

        document.getElementById('doc-expand').addEventListener('click', expandDocs);
    }

    if ($('.sidebar ul').length) {
        var current = $('.sidebar ul').find('li a[href="' + window.location.pathname + '"]');

        if (current.length) {
            current.parent().css('font-weight', 'bold');
            current.css('color', '#e74430');
            if (docCollapsed) {
                current.closest('ul').prev().toggleClass('is-active');
            }
        }
    }

    function expandItem(e) {
        var elem = e.target;

        if (elem.classList.contains('is-active')) {
            elem.classList.remove('is-active');
        } else {
            clearItems();
            elem.classList.add('is-active');
        }
    }

    function expandItemKeyboard(e) {
        var elem = e.target;

        if ([13, 37, 39].includes(e.keyCode)) {
            clearItems();
        }

        if (e.keyCode === 13) {
            elem.classList.toggle('is-active');
        }

        if (e.keyCode === 39) {
            elem.classList.add('is-active');
        }

        if (e.keyCode === 37) {
            elem.classList.remove('is-active');
        }
    }

    function clearItems() {
        for (var i = 0; i < toggles.length; i++) {
            toggles[i].classList.remove('is-active');
        }
    }
});
