import { i18n } from '../plugins/vue-i18n';

export const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export const rules = {
    username: { required: value => !!value || i18n.global.t('message.usernameRequired') },
    userEmail: { required: value => !!value || i18n.global.t('message.emailRequired'), email: value => regex.test(value) || i18n.global.t('message.validEmail') },
    vcode: { required: value => !!value || i18n.global.t('message.vCodeRequired'), minLength: value => value.length === 4 || i18n.global.t('message.vCodeLength') },
    password: { required: value => !!value || i18n.global.t('message.passwordRequired'), minLength: value => value.length >= 8 || i18n.global.t('message.passwordLength') },
    confirmPassword: { required: value => !!value || i18n.global.t('message.confirmPwdRequired'), minLength: value => value.length >= 8 || i18n.global.t('message.confirmPwdLength') },
};