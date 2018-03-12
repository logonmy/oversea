$(function () {

    $('#captchaImg').on('click', function () {
        var src = $(this).attr('src');
        var p = src.indexOf('?');
        if (p >= 0) {
            src = src.substr(0, p);
        }
        $(this).attr('src', src + "?reload=" + (new Date()).getTime() + '&_xsrf=' + $('meta[name=_xsrf]').attr('content'));
    });


    $('#registerForm').bootstrapValidator({
        fields: {
            password: {
                message: '密码无效',
                validators: {
                    notEmpty: {
                        message: '密码不能为空'
                    },
                    stringLength: {
                        min: 6,
                        max: 30,
                        message: '密码长度必须在6到30之间'
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '密码只能由字母、数字、点和下划线'
                    }
                }
            },
            repassword: {
                message: '密码无效',
                validators: {
                    notEmpty: {
                        message: '确认密码不能为空'
                    },
                    stringLength: {
                        min: 6,
                        max: 30,
                        message: '密码长度必须在6到30之间'
                    },
                    identical: {
                        field: 'password',
                        message: '两次密码不一致'
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '密码只能由字母、数字、点和下划线'
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
                toastr.success("您已成功注册");
                // self.location = "/home/home/index";//"/home/wish/platform";
            }
        }, 'json').complete(function () {
            $('#btnRegister').removeAttr('disabled')
        });
    });
});