$(function () {
    $(window).scroll(function(event){
        var top = $("#demo2").height();
        var scrolla=$(window).scrollTop();
        var cha=parseInt(top)-parseInt( scrolla);
        if(cha<=0)
        {
            $("#demo").addClass("navbar-fixed-top");
        } else {
            $("#demo").removeClass("navbar-fixed-top");
        }

    });

    $('.navbar-toggle').fancynav();

});