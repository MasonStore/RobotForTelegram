<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import LocalStorage from "@/stores/ls";
const {t, availableLocales: languages, locale} = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
  //存储到localStorage
  LocalStorage.setLocale(item);
};

let router = useRouter();
const onClickHelp = () => {
  router.push("/help");
}

</script>

<template>
  <van-config-provider theme="dark">
    <!-- Header -->
    <div class="header">
      <div class="nav"></div>
      <div class="menu">
        <div class="bar">
          <div class="language">
            <div
                v-for="item in languages"
                :key="item"
                :class="{ active: item === locale }"
                @click="onclickLanguageHandle(item)"
                class="lang-item"
            >
              {{ t("languages." + item) }}
            </div>
          </div>
          <van-icon name="question-o" color="#e0e0e0" class="cursor-pointer"
          @click="onClickHelp"
          />
          &nbsp;
          &nbsp;
        </div>
      </div>
    </div>
    <div class="header" v-if="false">
      <!-- navigation -->
      <div class="nav">
        <!--      <router-link to="/">{{ t("nav.home") }}</router-link>-->
        <!--      <router-link to="/about">{{ t("nav.about") }}</router-link>-->
      </div>
      <!-- Menu -->
      <div class="menu">
        <div class="nav">
          <router-link to="/">{{ t("nav.home") }}</router-link>
          <router-link to="/about">{{ t("nav.about") }}</router-link>
        </div>
        <div class="language">
          <div
              v-for="item in languages"
              :key="item"
              :class="{ active: item === locale }"
              @click="onclickLanguageHandle(item)"
              class="lang-item"
          >
            {{ t("languages." + item) }}
          </div>
        </div>
      </div>
    </div>
    <!-- Page -->
    <div class="view">
      <router-view/>
    </div>
  </van-config-provider>
</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  background-color: rgba(30, 30, 30, 0.9);
  overflow: hidden;
}

.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: rgba(50, 50, 50, 1);

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #363a3d;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #363a3d;
        color: #ffffff;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #3478f6;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        //background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #e0e0e0;
        font-size: 14px;

        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }

        &.active {
          background-color: rgba(54, 58, 61, 1);
          color: #e0e0e0;
          cursor: not-allowed;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 25px;
        line-height: 25px;
        padding: 0 5px;
        margin-left: 8px;
        //background-color: #ab7edc;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        //color: #000000;
        font-size: 14px;

        &:hover {
          //background-color: #d7a8d8;
          //color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}

.view {
  position: absolute;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}

</style>
