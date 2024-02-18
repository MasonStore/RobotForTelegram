class LocalStorage {
    public static KEY_FOR_AGREEMENT: string = 'agreement_v1';
    public static KEY_FOR_LOCALE: string = 'locale';

    public static ifAgree() {
        return localStorage.getItem(LocalStorage.KEY_FOR_AGREEMENT) === 'true';
    }

    public static agree() {
        localStorage.setItem(LocalStorage.KEY_FOR_AGREEMENT, 'true');
    }

    public static disagree() {
        localStorage.setItem(LocalStorage.KEY_FOR_AGREEMENT, 'false');
    }

    public static getLocale() {
        return localStorage.getItem(LocalStorage.KEY_FOR_LOCALE);
    }

    public static setLocale(locale: string) {
        localStorage.setItem(LocalStorage.KEY_FOR_LOCALE, locale);
    }
}

export default LocalStorage;