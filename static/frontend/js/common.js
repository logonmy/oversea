$(function () {
    $('.navbar-toggle').fancynav();

    // 关于我们tab 切换
    $('.about-tab a').on('click', function () {
        var index = $(this).index();
        $('.about-tab a').removeClass('active');
        $(this).addClass('active');
        $('.detail-content').css('display', 'none');
        $('.detail-content').eq(index).css('display', 'block');
    })

    // 移民项目
    $('.travel-type-list li').on('click', function () {
        $('.travel-type-list li').find('a').removeClass('active');
        $(this).find('a').addClass('active');
    });

});