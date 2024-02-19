<script setup lang="ts">
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {ref} from "vue";
import {SelectOneAccount} from "@go/views/AccountsView";
import anime from "animejs";

let router = useRouter();
const goHome = () => {
  SelectOneAccount(0).then((resp) => {
    if (resp.ok) {
      // 动画退出
      anime({
        // 动画
        targets: '.page',
        opacity: [1, 0],
        duration: 260,
        easing: 'easeInOutQuad'
      });
      // 400ms后跳转
      setTimeout(() => {
        router.push('/');
      }, 300);
    } else {
    }
  });
}

const back = () => {
  router.back();
}

const {t} = useI18n();

const contactUs = () => {
  BrowserOpenURL("https://t.me/RobotForTGSupport")
}

const jumpTo = (url: string) => {
  BrowserOpenURL(url)
}

const openSourceList = ref([{
  name: "telegram-bot-api",
  url: "https://github.com/go-telegram-bot-api/telegram-bot-api",
  license: "MIT License"
}, {
  name: "go-openai",
  url: "https://github.com/sashabaranov/go-openai",
  license: "Apache License 2.0"
}, {
  name: "wails",
  url: "https://github.com/wailsapp/wails",
  license: "MIT License"
}, {
  name: "sqlite",
  url: "https://github.com/go-gorm/sqlite",
  license: "MIT License"
}, {
  name: "gorm",
  url: "https://github.com/go-gorm/gorm",
  license: "MIT License"
}, {
  name: "animejs",
  url: "https://github.com/juliangarnier/anime",
  license: "MIT License"
}, {
  name: "piniajs",
  url: "https://github.com/vuejs/pinia",
  license: "MIT License"
}, {
  name: "vantjs",
  url: "https://github.com/vant-ui/vant",
  license: "MIT License"
}, {
  name: "vuejs",
  url: "https://github.com/vuejs/core",
  license: "MIT License"
}, {
  name: "vue-i18n-next",
  url: "https://github.com/intlify/vue-i18n-next",
  license: "MIT License"
}, {
  name: "vue-router",
  url: "https://github.com/vuejs/router",
  license: "MIT License"
}])
</script>

<template>
  <van-nav-bar
      :title="t('使用手册')"
      :left-text="t('返回')"
      left-arrow
      @click-left="back"
  />
  <div class="page">
    <div id="bd" class="max-w-2xl mx-auto p-8">


      <div class="bg-gray-700 p-4 rounded-md mb-4">
        <p class="mb-2">{{ t('使用手册开头') }}</p>
        <p class="mb-2">{{ t('如有任何疑问或建议，请') }}
          <a class="text-blue-400"
             @click="contactUs"
          >{{ t('联系') }}</a>
          {{ t('我们') }}
        </p>
      </div>

      <h2 class="text-xl font-semibold mb-4">1. {{ t('TG相关') }}</h2>
      <ul class="list-disc pl-5 space-y-2 color-78c6f6">
        <li>{{ t('怎么创建机器人') }}</li>
        <p class="text-l">
          {{ t('联系') }} <a class="text-blue-400" @click="jumpTo('https://t.me/BotFather')">@BotFather</a>
          {{ t('创建机器人') }}
        </p>
        <img src="@/assets/images/createbot.png" height="948" width="1464"/>
        <li>{{ t('怎么创建游戏') }}</li>
        <p class="text-l">
          {{ t('联系') }} <a class="text-blue-400" @click="jumpTo('https://t.me/BotFather')">@BotFather</a>
          {{ t('创建游戏') }}
        </p>
        <img src="../assets/images/newgame.jpg" height="2824" width="1570"/>
        <li>{{ t('怎么启用内联查询') }}</li>
        <p class="text-l">
          {{ t('联系') }} <a class="text-blue-400" @click="jumpTo('https://t.me/BotFather')">@BotFather</a>
          {{ t('启用内联查询') }}
        </p>
        <img src="@/assets/images/setinline.png" height="768" width="1496"/>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">2. {{ t('产品相关') }}</h2>
      <ul class="list-disc pl-5 space-y-2 color-78c6f6">
        <li>{{ t('如何使用键盘组') }}</li>
        <p class="text-l">
          {{ t('Configure your robot keyboard on this page.') }}
        </p>
        <img src="../assets/images/26key.png" height="1272" width="1872"/>
        <p class="text-l">
          {{ t('The effect seen on the mobile phone is like this.') }}
        </p>
        <img src="../assets/images/26keymobile.jpg" height="741" width="1179"/>
        <li>{{ t('如何使用回复') }}</li>
        <p class="text-l">
          {{ t('If your reply is static content, do not check the template option.') }}
        </p>
        <img src="../assets/images/reply.png" height="1338" width="1872"/>
        <p class="text-l">
          {{ t('If your reply is dynamic content, check the template option.') }}
        </p>
        <img src="../assets/images/tmpl.png" height="1320" width="1872"/>
        <li>{{ t('如何监听命令') }}</li>
        <p class="text-l">
          {{ t('If the reply you selected is static.') }}
        </p>
        <img src="../assets/images/command.png" height="1320" width="1872"/>
        <p class="text-l">
          {{ t('If the reply you selected is dynamic. You must set Hook URL. When the command is triggered, i will call your api and get the reply.') }}
        </p>
        <img src="../assets/images/commandtmpl.png" height="1320" width="1872"/>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">3. {{ t('OpenAI相关') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('如何获取api-key') }}:</li>
        <p class="text-l">
          <a class="text-blue-400 cursor-pointer" @click="jumpTo(' https://platform.openai.com/account/api-keys')">
            {{ t('如何获取api-key介绍.') }}</a>
        </p>
        <li>{{ t('最大Token') }}:</li>
        <p class="text-l">
          {{ t('The maximum number of tokens to generate. Requests can use up to 4096 tokens for content generation, and 2048 tokens for classification.') }}
        </p>
        <li>{{ t('随机性') }}:</li>
        <p class="text-l">
          {{ t('What is the temperature of the sampling. Higher temperature results in more random completions. Lower temperature results in more conservative completions.') }}
        </p>
        <li>{{ t('核采样') }}:</li>
        <p class="text-l">
          {{ t('An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass.') }}
        </p>
        <li>{{ t('频繁惩罚') }}:</li>
        <p class="text-l">
          {{ t('The frequency of the request. The higher the frequency, the more the request will be used.') }}
        </p>
        <li>{{ t('重复惩罚') }}:</li>
        <p class="text-l">
          {{ t('重复惩罚介绍') }}
        </p>
        <li>{{ t('上文数量') }}:</li>
        <p class="text-l">
          {{ t('上文数量介绍') }}
        </p>
        <li>{{ t('系统提示') }}:</li>
        <p class="text-l">
          {{ t('系统提示介绍') }}
        </p>
        <li>{{ t('超时时间') }}:</li>
        <p class="text-l">
          {{ t('超时时间介绍') }}
        </p>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">4. {{ t('Hook协议') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('产品调用HookAPI的请求') }}</li>
          <pre>
            <code>
{
  "form": {
    "id": 1234567890, //number
    "isBot": false, //boolean
    "firstName": "string", //string
    "lastName": "string", //string
    "userName": "string" //string
  },
  "messageId": 1234567890, //number
  "date": 1234567890, //number
  "text": "string", //string
}
            </code>
          </pre>
        <li>{{ t('HookAPI的响应(文本)') }}</li>
          <pre>
            <code>
{
  "Content": "" //string
}
            </code>
          </pre>
        <li>{{ t('HookAPI的响应(图片)') }}</li>
          <pre>
            <code>
{
  "Base64": "" //string
}
            </code>
          </pre>
        <li>{{ t('HookAPI的响应(游戏)') }}</li>
          <pre>
            <code>
{
  "GameShortName": "", //string
  "GameTitle": "", //string
  "Url": "", //string
  "ButtonText": "" //string
}
            </code>
          </pre>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">{{ t('使用的开源软件') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li v-for="app in openSourceList"><a href="#" class="text-blue-400"
                                             @click="jumpTo(app.url)"
        >{{ app.name }} </a> : {{ app.license }}
        </li>
      </ul>
      <h2 v-for="i in 3">&nbsp;</h2>
    </div>

  </div>
</template>

<style scoped lang="scss">
.page {
  //上下左右居中
  display: flex;
  justify-content: center;
  height: 100%;
  width: 100%;
  background-color: #282828;
  //滚轮
  overflow-y: auto;
  overflow-x: hidden;
  color: #b0b0b0;
  //允许用户选中文本
  user-select: text;
}

.text-l {
  font-size: 12px;
  color: #b0b0b0;
}

.color-78c6f6 {
  color: #78c6f6;
}
</style>