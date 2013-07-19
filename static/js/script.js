// $('a.btn').click(function(){
//     $(this).addClass("active");
// });

// var width = $('div.col1').width();

// $('header div').each(function(index) {
//     $(this).css('height', width);
// });

$(window).bind('scroll',function(e){
    parallaxScroll();
});

function parallaxScroll(){
    var scrolled = $(window).scrollTop();
    $('#parallax1').css('top',(100+Math.sin(scrolled*0.01)*100)+'px');
    $('#parallax1').css('left',(500-Math.cos(scrolled*0.01)*100)+'px');
    $('#parallax2').css('left',(100-(scrolled*1))+'px');
    $('#parallax3').css('top',(300-(scrolled*-0.75))+'px');
}
