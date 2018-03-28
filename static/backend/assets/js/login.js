$(function () {

    $('#captchaImg').on('click', function () {
        var src = $(this).attr('src');
        var p = src.indexOf('?');
        if (p >= 0) {
            src = src.substr(0, p);
        }
        $(this).attr('src', src + "?reload=" + (new Date()).getTime() + '&_xsrf=' + $('meta[name=_xsrf]').attr('content'));
    });

    $('#rememberPwd').on('click', function () {
        if ($(this).find('i').hasClass('checked')) {
            $(this).find('i').removeClass('checked');
        } else {
            $(this).find('i').addClass('checked');
        }
    });

    $('input').blur(function () {
        if ($(this).val() != '') {
            $(this).removeAttr('style');
        } else {
            $(this).css('border-color', 'red');
        }
    });

    function checkLogin() {
        if ($('input[name="username"]').val() == '') {
            $('input[name="username"]').css('border-color', 'red');
            toastr.error("账号不能为空");
            return false;
        } else if ($('input[name="password"]').val() == '') {
            $('input[name="password"]').css('border-color', 'red');
            toastr.error("密码不能为空");
            return false;
        } else if ($('input[name="captcha"]').val() == '') {
            $('input[name="captcha"]').css('border-color', 'red');
            toastr.error("验证码不能为空");
            return false;
        }
        $('input').removeAttr('style');
        return true
    }

    $('#btnLogin').on('click', function (e) {

        e.preventDefault();

        var $form =  $('#loginForm');

        if (!checkLogin()) {
            return false;
        }

        $('#btnLogin').removeAttr('disabled');
        $('#btnLogin').addClass('disabled');

        var remember = 'no';
        if ($('#rememberPwd').find('i').hasClass('checked')) {
            remember = 'yes';
        }

        $.post($form.attr('action'), $form.serialize() + "&remember=" + remember, function (result) {
            if (result.errorCode != 0) {
                toastr.error(result.errorMsg);
                $('input[name="captchaId"]').val(result.data.captchaId);
                $('#captchaImg').attr("src", result.data.captchaUrl);
                $('input[name="captcha"]').val('');
                return false;
            } else {
                toastr.success(result.errorMsg);
                self.location = "/admin/home/index";
            }
        }, 'json').complete(function () {
            $('#btnLogin').removeAttr('disabled');
            $('#btnLogin').removeClass('disabled');
        });

    });


});