import { createI18n } from "vue-i18n";

import zhHans from "./locales/zh-Hans.json";
import en from "./locales/en.json";
import LocalStorage from "@/stores/ls";

const i18n = createI18n({
  locale: LocalStorage.getLocale() || "en",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    "zh-Hans": zhHans,
    en: en,
  },
});

export default i18n;
