function beforeAfter(){
    var sliderBA = document.getElementById('before-after-slider');
    var beforeBA = document.getElementById('before-image');
    var beforeImageBA = beforeBA.getElementsByTagName('img')[0];
    var resizerBA = document.getElementById('resizer');
    var activeBA = false;

    function active(){ activeBA = true; resizerBA.classList.add('resize');}
    function inactive(){ activeBA = false; resizerBA.classList.remove('resize');}

    //Sort overflow out for Overlay Image
    var width = sliderBA.offsetWidth;
    beforeImageBA.style.width = width + 'px';

    //Adjust width of image on resize 
    window.addEventListener('resize', function () {
        var width = sliderBA.offsetWidth;
        beforeImageBA.style.width = width + 'px';
    })

    //Basic Events
    resizerBA.addEventListener('mousedown', function(){active()});
    document.body.addEventListener('mouseup',function(){inactive();});
    document.body.addEventListener('mouseleave', function (){inactive();});
    resizerBA.addEventListener('touchstart', function () {active();});
    document.body.addEventListener('touchend', function () {inactive();});
    document.body.addEventListener('touchcancel', function () {inactive();});

    //Drag & Swipe Events
    document.body.addEventListener('mousemove', function (e) {
        if (!activeBA) return;
        var x = e.pageX;
        x -= sliderBA.getBoundingClientRect().left;
        slideIt(x);
        pauseEvent(e);
    });
    document.body.addEventListener('touchmove', function (e) {
        if (!activeBA) return;
        let x;
        let i;
        for (i = 0; i < e.changedTouches.length; i++) {x = e.changedTouches[i].pageX;}
        x -= sliderBA.getBoundingClientRect().left;
        slideIt(x);
        pauseEvent(e);
    });

    function slideIt(x) {
        var transform = Math.max(0, (Math.min(x, sliderBA.offsetWidth)));
        beforeBA.style.width = transform + "px";
        resizerBA.style.left = transform - 0 + "px";
    }

    //stop divs being selected.
    function pauseEvent(e) {
        if (e.stopPropagation) e.stopPropagation();
        if (e.preventDefault) e.preventDefault();
        e.cancelBubble = true;
        e.returnValue = false;
        return false;
    }
}
beforeAfter();