export const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export const rules = {
    username: { required: value => !!value || '用户名是必填项。' },
    userEmail: { required: value => !!value || '邮箱是必填项。', email: value => regex.test(value) || '请输入有效的邮箱地址。' },
    vcode: { required: value => !!value || '验证码是必填项。', minLength: value => value.length === 4 || '验证码长度必须为4个字符。' },
    password: { required: value => !!value || '密码是必填项。', minLength: value => value.length >= 8 || '密码长度至少为8个字符。' },
    confirmPassword: { required: value => !!value || '确认密码是必填项。', minLength: value => value.length >= 8 || '确认密码长度至少为8个字符。' },
};