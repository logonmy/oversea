$(function () {

    //自定义正则表达示验证方法
    $.validator.addMethod("checkMobile", function (value, element, params) {
        var phoneReg = /^1[0-9]{10}$/;
        return this.optional(element) || (phoneReg.test(value));
    }, "*请输入正确的手机号码！");

    $.validator.addMethod("checkPwd", function (value, element, params) {
        var checkPwd = /^\w{6,16}$/g;
        return this.optional(element) || (checkPwd.test(value));
    }, "*只允许6-16位英文字母、数字或者下画线！");

    $.validator.addMethod("checkEmail", function (value, element, params) {
        var checkEmail = /^[a-z0-9]+@([a-z0-9]+\.)+[a-z]{2,4}$/i;
        return this.optional(element) || (checkEmail.test(value));
    }, "*请输入正确的邮箱！");

    // 保存管理员
    $('#userForm').validate({
        rules: {
            password: {
                required: true,
                checkPwd: true,
            },
            repassword: {
                required: true,
                checkPwd: true,
                equalTo: 'input[name="password"]'
            },
            phone: {
                required: true,
                checkMobile: true
            },
            email: {
                required: true,
                checkEmail: true
            },
            real_name: {
                required: true,
                minlength: 2,
                maxlength: 10
            },
        },
        messages: {
            real_name: {
                required: "管理员名称不能为空",
                minlength: "密码长度必须在2到10之间"
            },
            phone: {
                required: "手机号码不能为空",
                checkMobile: "请输入正确手机号码"
            },
            email: {
                required: "邮箱地址不能为空",
                checkEmail: "请输入正确邮箱地址"
            },
            password: {
                required: "请填写密码",
                checkPwd: "密码必须是6-16位英文字母、数字或者下画线"
            },
            repassword: {
                required: "请填写密码",
                checkPwd: "密码必须是6-16位英文字母、数字或者下画线"
            }
        },
        invalidHandler: function (form, validator) {  //不通过回调
            return false;
        },
        submitHandler: function (e) {
            // e.preventDefault();
            var $form = $(e);

            $.ajax({
                type: 'post',
                url: $form.attr('action'),
                data: $form.serialize(),
                dataType: "json",
                success: function (data) {
                    toastr.success(result.errorMsg);
                    setTimeout(function () {
                        location.href = result.data.href
                    }, 2000)
                },
                complete: function () {
                    alert(111);
                    $('button[type="submit"]').removeAttr('disabled')
                }
            });

            return false;

        },
    });


    // 编辑管理员-保存
    $('#userEditForm').validate({
        rules: {
            phone: {
                required: true,
                checkMobile: true
            },
            email: {
                required: true,
                checkEmail: true
            },
        },
        messages: {
            phone: {
                required: "手机号码不能为空",
                checkMobile: "请输入正确手机号码"
            },
            email: {
                required: "邮箱地址不能为空",
                checkEmail: "请输入正确邮箱地址"
            },
        },
        submitHandler: function (e) {
            var $form = $(e);
            $.ajax({
                type: 'post',
                url: $form.attr('action'),
                data: $form.serialize(),
                dataType: "json",
                success: function (data) {
                    toastr.success(result.errorMsg);
                    setTimeout(function () {
                        location.href = result.data.href
                    }, 2000)
                },
                complete: function () {
                    $('button[type="submit"]').removeAttr('disabled')
                }
            });

            return false
        }
    });
});