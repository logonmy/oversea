$(function () {
    // 修改密码页面-提交密码修改
    $('#changePasswordForm').bootstrapValidator({
        fields: {
            oldpassword: {
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
            newpassword: {
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
                    },
                    // callback: {
                    //     message: '新旧密码不能一致',
                    //     callback:function(value, validator,$field){
                    //         if (value === $('input[name="oldpassword"]').val()) {
                    //             return false;
                    //         }
                    //         return true;
                    //     }
                    // },
                    different: {
                        field: 'oldpassword',
                        message: '新旧密码不能一致'
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
                        field: 'newpassword',
                        message: '两次密码不一致'
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '密码只能由字母、数字、点和下划线'
                    }
                }
            }
        }
    }).on('success.form.bv', function (e) {
        e.preventDefault();
        var $form = $(e.target);
        $.post($form.attr('action'), $form.serialize(), function (result) {
            if (result.errorCode != 0) {
                toastr.error(result.errorMsg);
                return false;
            } else {
                toastr.success("密码修改成功,请重新登陆");
                self.location = "/loginout";
            }
        }, 'json').complete(function () {
            $('.changePwd').removeAttr('disabled')
        });
    });


    // 修改个人信息页面-提交个人信息修改
    $('#userForm').bootstrapValidator({
        fields: {
            real_name: {
                message: '真实姓名无效',
                validators: {
                    notEmpty: {
                        message: '真实姓名不能为空'
                    },
                    stringLength: {
                        min: 2,
                        max: 30,
                        message: '真实姓名长度必须在2到20之间'
                    }
                }
            },
            email: {
                message: '邮箱地址无效',
                validators: {
                    notEmpty: {
                        message: '邮箱地址不能为空'
                    },
                    emailAddress: {     //　　邮箱格式校验
                        message: '邮箱地址'
                    }
                }
            },

        }
    }).on('success.form.bv', function (e) {
        e.preventDefault();
        var $form = $(e.target);
        $.post($form.attr('action'), $form.serialize(), function (result) {
            if (result.errorCode != 0) {
                toastr.error(result.errorMsg);
                return false;
            } else {
                toastr.success("信息修改成功,请重新登陆");
            }
        }, 'json').complete(function () {
            $('.changeInfo').removeAttr('disabled')
        });
    });
});