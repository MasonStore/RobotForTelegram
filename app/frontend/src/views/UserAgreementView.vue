<script setup lang="ts">

import {useRouter} from "vue-router";
import {Quit} from "@go/views/AgreementView";
import LocalStorage from "@/stores/ls";
import {useI18n} from "vue-i18n";
import {BrowserOpenURL} from "../../wailsjs/runtime";

const router = useRouter();
const agree = () => {
  LocalStorage.agree()
  router.push("/");
}
const reject = () => {
  LocalStorage.disagree()
  //应用退出
  Quit();
}

const {t} = useI18n();

const contactUs = () => {
  BrowserOpenURL("https://t.me/RobotForTelegramSupport")
}
</script>

<template>
  <div class="page">
    <div id="bd" class="max-w-2xl mx-auto p-8">
      <h1 class="text-3xl font-semibold mb-4" v-for="i in 0">&nbsp;</h1>
      <h1 class="text-3xl font-semibold mb-4">{{ t('用户协议') }}</h1>
      <p class="text-gray-400 mb-6">{{ t('用户协议开头') }}</p>
      <h2 class="text-xl font-semibold mb-4">1. {{ t('产品介绍') }}</h2>
      <p class="text-gray-400 mb-6">{{ t('产品介绍内容') }}</p>
      <h2 class="text-xl font-semibold mb-4">2. {{ t('功能说明') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('功能1') }}</li>
        <li>{{ t('功能2') }}</li>
        <li>{{ t('功能3') }}</li>
        <li>{{ t('功能4') }}</li>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">3. {{ t('免责声明') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('免责1') }}</li>
        <li>{{ t('免责2') }}</li>
        <li>{{ t('免责3') }}</li>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">4. {{ t('知识产权') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('产权1') }}</li>
      </ul>
      <h2>&nbsp;</h2>
      <h2 class="text-xl font-semibold mb-4">5. {{ t('其他条款') }}</h2>
      <ul class="list-disc pl-5 space-y-2">
        <li>{{ t('其他1') }}</li>
        <li>{{ t('其他2') }}</li>
      </ul>
      <h2>&nbsp;</h2>

      <div class="bg-gray-700 p-4 rounded-md mb-4">
        <p class="mb-2">{{ t('用户协议结尾') }}</p>
        <p class="mb-2">{{ t('如有任何疑问或建议，请') }}
          <a class="text-blue-400"
             @click="contactUs"
          >{{t('联系')}}</a>
          {{ t('我们')}}
        </p>
        <p class="mb-2">{{ t('用户协议致谢') }}</p>
      </div>

      <van-row justify="center">
        <van-col class="content-center" span="12">
          <van-button color="#6bd35f" @click="agree">{{ t('我已阅读并同意') }}</van-button>
        </van-col>
        <van-col class="content-center" span="12">
          <van-button color="#eb4d3d" @click="reject">{{ t('拒绝并退出') }}</van-button>
        </van-col>
      </van-row>
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
</style>