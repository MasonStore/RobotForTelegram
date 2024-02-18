class CommonJS {
    /**
     * 拷贝文本
     * @param text
     */
    static copyText(text) {
        let input = document.createElement('input');
        input.value = text;
        document.body.appendChild(input);
        input.select();
        document.execCommand('copy');
        document.body.removeChild(input);
    }

    /**
     * 获取base64图片的url
     * @param pic
     * @returns {string}
     */
    static getBase64URL(pic) {
        const blob = CommonJS.base64ImgToFile(pic)
        return window.URL.createObjectURL(blob)
    }

    /**
     * base64转File
     * @param dataurl
     * @param filename
     * @returns {File}
     */
    static base64ImgToFile(dataurl, filename = 'file') {
        //将base64格式分割：['data:image/png;base64','XXXX']
        const arr = dataurl.split(',')
        // .*？ 表示匹配任意字符到下一个符合条件的字符 刚好匹配到：
        // image/png
        const mime = arr[0].match(/:(.*?);/)[1]  //image/png
        //[image,png] 获取图片类型后缀
        const suffix = mime.split('/')[1] //png
        const bstr = atob(arr[1])   //atob() 方法用于解码使用 base-64 编码的字符串
        let n = bstr.length
        const u8arr = new Uint8Array(n)
        while (n--) {
            u8arr[n] = bstr.charCodeAt(n)
        }
        return new File([u8arr], `${filename}.${suffix}`, {
            type: mime
        })
    }

    /**
     * url转base64
     * @param url
     * @returns {Promise<unknown>}
     */
    static urlToBase64(url) {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest()
            xhr.open('get', url, true)
            xhr.responseType = 'blob'
            xhr.onload = function () {
                if (this.status === 200) {
                    const blob = this.response
                    const fileReader = new FileReader()
                    fileReader.onloadend = function (e) {
                        const result = e.target.result
                        resolve(result)
                    }
                    fileReader.readAsDataURL(blob)
                }
            }
            xhr.send()
        })
    }

}

//禁止右键
document.oncontextmenu = function () {
    return false;
};

//F5刷新
document.onkeydown = function (e) {
    if (e.keyCode === 116) {
        e.preventDefault();
        window.location.reload();
    }
};