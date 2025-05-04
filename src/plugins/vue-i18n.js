import { createI18n } from "vue-i18n";
import ar from "../plugins/locales/ar.js";
import en from "../plugins/locales/en.js";
import es from "../plugins/locales/es.js";
import fr from "../plugins/locales/fr.js";
import it from "../plugins/locales/it.js";
import ja from "../plugins/locales/ja.js";
import pt from "../plugins/locales/pt.js";
import ru from "../plugins/locales/ru.js";
import zh_CN from "../plugins/locales/zh_CN.js";
import zh_TW from "../plugins/locales/zh_TW.js";

const messages = {
  ar,
  en,
  es,
  fr,
  it,
  ja,
  pt,
  ru,
  zh_CN,
  zh_TW
};

export const i18n = createI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages,
})