$(function () {
    // 保存管理员
    $('#userForm').bootstrapValidator({
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
            real_name: {
                message: '管理员名称不能为空',
                validators: {
                    notEmpty: {
                        message: '管理员名称不能为空'
                    },
                    stringLength: {
                        min: 2,
                        max: 10,
                        message: '密码长度必须在2到10之间'
                    },
                }
            },
            email: {
                validators: {
                    emailAddress: {
                        message: 'email地址格式错误'
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
                toastr.success(result.errorMsg);
                setTimeout(function () {
                    location.href = result.data.href
                }, 2000)
            }
        }, 'json').complete(function () {
            $('#btnAddUser').removeAttr('disabled')
        });
    });

    // 编辑管理员-保存
    $('#userEditForm').bootstrapValidator({
        fields: {
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
            email: {
                validators: {
                    emailAddress: {
                        message: 'email地址格式错误'
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
                toastr.success(result.errorMsg);
                setTimeout(function () {
                    location.href = result.data.href
                }, 2000)
            }
        }, 'json').complete(function () {
            $('#btnEditUser').removeAttr('disabled')
        });

    });
});