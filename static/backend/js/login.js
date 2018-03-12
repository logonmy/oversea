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
            phone: {
                message: '手机号码不能为空',
                validators: {
                    notEmpty: {
                        message: '手机号码不能为空'
                    },
                    regexp: {
                        regexp: /^1[0-9]{10}$/,
                        message: '请输入正确的手机号码'
                    }
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
                    // remote: {
                    //     url: '/captcha/verify',
                    //     message:"验证码不正确",
                    //     type: "post",
                    //     dataType: 'json',
                    //     data: {
                    //         _xsrf: $('meta[name=_xsrf]').attr('content'),
                    //         captchaId: $('input[name="captchaId"]').val(),
                    //         captcha: $('input[name="captcha"]').val(),
                    //     },
                    //     success: function () {
                    //
                    //     },
                    //     delay: 2000,
                    // },
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
                self.location = "/admin/index";//"/home/wish/platform";
            }
        }, 'json').complete(function () {
            $('#btnLogin').removeAttr('disabled')
        });

    });
});