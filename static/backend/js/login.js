$(function () {

    $('#captchaImg').on('click', function () {
        var src = $(this).attr('src');
        var p = src.indexOf('?');
        if (p >= 0) {
            src = src.substr(0, p);
        }
        $(this).attr('src', src + "?reload=" + (new Date()).getTime() + '&_xsrf=' + $('meta[name=_xsrf]').attr('content'));
    });


    $('#loginForm').bootstrapValidator({
        fields: {
            password: {
                message: '密码无效',
                validators: {
                    notEmpty: {
                        message: '密码不能为空'
                    }
                }
            },
            username: {
                message: '管理员账号不能为空',
                validators: {
                    notEmpty: {
                        message: '管理员账号不能为空'
                    },
                }
            },
            captcha: {
                message: '验证码不能为空',
                validators: {
                    notEmpty: {
                        message: '验证码不能为空'
                    },
                    stringLength: {
                        min: 6,
                        max: 6,
                        message: '验证码长度为6位',
                    },
                }
            }
        }
    }).on('success.form.bv', function (e) {
        e.preventDefault();
        var $form = $(e.target);
        $.post($form.attr('action'), $form.serialize(), function (result) {
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
            $('#btnLogin').removeAttr('disabled')
        });

    });
});